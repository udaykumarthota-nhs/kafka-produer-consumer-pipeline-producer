package main

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	kafkaTopic   = "telemetry"
	databaseUrl  = "mongodb://localhost:27017"
	databaseName = "ConsumedDataDB1"
)

func main() {
	HandleRouter()
}

func HandleRouter() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "requested for index page")
	})

	r.HandleFunc("/getData/{collection_name}", readDataFromCollection).Methods("GET")
	http.ListenAndServe(":8089", handler)
}

type BodyStruct struct {
	collection_name string
}

func readDataFromCollection(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	clientOptions := options.Client().ApplyURI(databaseUrl)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("error:", err)
	}
	vars := mux.Vars(r)
	collection_name := vars["collection_name"]
	dbcollection := client.Database(databaseName).Collection(collection_name)
	result, e := dbcollection.Find(ctx, bson.M{})
	if e != nil {
		fmt.Print(e)
	}
	switch getStructTypeForCollection(collection_name).(type) {
	case Devices:
		var data []Devices
		err = result.All(context.Background(), &data)
		if err != nil {
			fmt.Print(err)
		} else {
			json.NewEncoder(rw).Encode(data)
		}
		break
	case MacDetails:
		var data []MacDetails
		err = result.All(context.Background(), &data)
		if err != nil {
			fmt.Print(err)
		} else {
			json.NewEncoder(rw).Encode(data)
		}
		break
	}
	client.Disconnect(ctx)
}

type Devices struct {
	Name      string `json:"Name"`
	Ip        string `json:"Ip"`
	Port      string `json:"Port"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	IsEnabled bool   `json:"IsEnabled"`
	UseNSO    bool   `json:"UseNSO"`
}

type MacDetails struct {
	Macaddress string `json:"Macaddress"`
	Area       string `json:"Area"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
}

func getStructTypeForCollection(collection_name string) interface{} {
	switch {
	case collection_name == "devices":
		return Devices{}
	case collection_name == "macdetails":
		return MacDetails{}
	default:
		fmt.Print("something error")
		return bson.TypeNull
	}
}

func typeofstruct(collection_name string) reflect.Type {
	return reflect.TypeOf(getStructTypeForCollection(collection_name))
}
