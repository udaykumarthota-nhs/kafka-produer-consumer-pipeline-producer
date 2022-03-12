// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/Kamva/octopus"
// 	"github.com/Kamva/octopus/base"
// 	"github.com/gorilla/mux"
// 	"github.com/rs/cors"
// 	"github.com/segmentio/kafka-go"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const (
// 	topicGC = "telemetry"
// 	// topicCPU      = "telemetrycheckCPU1"
// 	brokerAddress     = "103.249.77.26:9092"
// 	mapBoxAccessToken = "pk.eyJ1IjoidWRheWt1bWFyLTgzMjkiLCJhIjoiY2t2bWN0NDZmM29xMjMxcGdmcDh1bGQyMSJ9.WVNK0XbI0qTmNKzN8ALP_A"
// )

// func main() {

// 	fmt.Println("Hello World")
// 	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	// ctx := context.Background()
// 	// client, e := mongo.Connect(ctx, clientOptions)
// 	// if e != nil {
// 	// 	fmt.Println("error:", e)
// 	// }
// 	// go consumeCPUProcessData(client, ctx)
// 	// go consumeGenericCountersData(client, ctx)
// 	HandleRouter()
// 	// getCoordinates("vinukonda", "vinukonda", "andhra pradesh", "india")

// }

// func HandleRouter() {
// 	r := mux.NewRouter()

// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowCredentials: true,
// 	})
// 	handler := c.Handler(r)
// 	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(rw, "requested for index page")
// 	})

// 	r.HandleFunc("/fetch/gc/all", getAllGCInfo)
// 	r.HandleFunc("/fetch/cpu/all", getAllCPUInfo)
// 	r.HandleFunc("/fetch/gc/all/{ipaddress}/{portaddress}", getAllGCInfo)
// 	r.HandleFunc("/fetch/cpu/all/{ipaddress}/{portaddress}", getAllCPUInfo)
// 	r.HandleFunc("/fetch/ip/{ipaddress}/{portaddress}", filterByIPAddress)
// 	r.HandleFunc("/devices/add", addDevice).Methods("POST")
// 	r.HandleFunc("/fetch/devices/all", getAllDevices).Methods("GET")

// 	r.HandleFunc("/macdetails/add", addMacDetail).Methods("POST")
// 	r.HandleFunc("/fetch/macdetails/all", getAllMacDetails).Methods("GET")
// 	r.HandleFunc("/fetch/devices/{_id}", getDeviceDetailsById).Methods("GET")
// 	http.ListenAndServe(":8088", handler)
// }

// func getAllDevices(rw http.ResponseWriter, r *http.Request) {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	collectionDevices := client.Database("ConsumedDataDB1").Collection("devices")
// 	result, e := collectionDevices.Find(ctx, bson.M{})
// 	var devices []bson.M
// 	if e != nil {
// 		json.NewEncoder(rw).Encode(e)
// 	} else {

// 		err = result.All(context.Background(), &devices)
// 		// err = result.All(context.Background(), )
// 		fmt.Print(devices)
// 		json.NewEncoder(rw).Encode(devices)
// 	}
// 	defer client.Disconnect(ctx)
// }

// func getDeviceDetailsById(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["_id"]
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Print(id)
// 	collectionDevices := client.Database("ConsumedDataDB1").Collection("devices")
// 	objId, e := primitive.ObjectIDFromHex(id)
// 	if e != nil {
// 		json.NewEncoder(rw).Encode(e)
// 	}
// 	result := collectionDevices.FindOne(ctx, bson.M{"_id": objId})
// 	var device bson.M
// 	//  else {

// 	err = result.Decode(&device)
// 	// err = result.All(context.Background(), )
// 	fmt.Print(device)
// 	json.NewEncoder(rw).Encode(device)
// 	// }
// 	defer client.Disconnect(ctx)
// }

// func addDevice(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "application/json")
// 	body, err := ioutil.ReadAll(r.Body)
// 	var device Device
// 	err = json.Unmarshal(body, &device)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Print(device)

// 	collectionDevices := client.Database("ConsumedDataDB1").Collection("devices")
// 	result, e := collectionDevices.InsertOne(ctx, device)
// 	if e != nil {
// 		json.NewEncoder(rw).Encode(e)
// 	} else {
// 		json.NewEncoder(rw).Encode(result)
// 	}
// 	client.Disconnect(ctx)
// }

// func getCoordinates(area, city, state, country string) {
// 	resp, err := http.Get("https://api.mapbox.com/geocoding/v5/mapbox.places/" + area + "," + city + "," + state + "," + country + ".json?limit=1&access_token=pk.eyJ1IjoidWRheWt1bWFyLTgzMjkiLCJhIjoiY2t2bWN0NDZmM29xMjMxcGdmcDh1bGQyMSJ9.WVNK0XbI0qTmNKzN8ALP_A")
// 	if err != nil {
// 		log.Fatal("ooopsss an error occurred, please try again")
// 	}
// 	defer resp.Body.Close()
// 	fmt.Print(resp)
// }

// func getAllMacDetails(rw http.ResponseWriter, r *http.Request) {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	collectionDevices := client.Database("ConsumedDataDB1").Collection("macdetails")
// 	result, e := collectionDevices.Find(ctx, bson.M{})
// 	var macDetails []MacDetails
// 	if e != nil {
// 		json.NewEncoder(rw).Encode(e)
// 	} else {
// 		err = result.All(context.Background(), &macDetails)
// 		fmt.Print(macDetails)
// 		json.NewEncoder(rw).Encode(macDetails)
// 	}
// 	client.Disconnect(ctx)
// }

// func addMacDetail(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "application/json")
// 	body, err := ioutil.ReadAll(r.Body)
// 	var macDetail MacDetails
// 	err = json.Unmarshal(body, &macDetail)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	collectionDevices := client.Database("ConsumedDataDB1").Collection("macdetails")
// 	result, e := collectionDevices.InsertOne(ctx, macDetail)
// 	if e != nil {
// 		json.NewEncoder(rw).Encode(e)
// 	} else {
// 		json.NewEncoder(rw).Encode(result)
// 	}
// 	client.Disconnect(ctx)
// }

// func filterByIPAddress(rw http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	ipaddress := vars["ipaddress"]
// 	portAddress := vars["portaddress"]
// 	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		fmt.Println("mongo.Connect() ERROR:", err)
// 		os.Exit(1)
// 	} else {
// 		fmt.Println("Connected to mongoDB")
// 	}

// 	cpuData := getData(client, ipaddress, portAddress)

// 	z := OutputData{
// 		cpuProcessData: cpuData,
// 		// genericCountersData: gcData,
// 	}
// 	x, er := json.Marshal(z)
// 	if er != nil {
// 		fmt.Println(er)
// 	} else {
// 		rw.Write(x)
// 	}
// }

// type OutputData struct {
// 	cpuProcessData      []CPUTelemetryData
// 	genericCountersData []GenericCountersTelemetryData
// }

// func getAllGCInfo(rw http.ResponseWriter, r *http.Request) {

// 	// model := &GcModel{}
// 	// config := base.DBConfig{Driver: base.Mongo, Host: "localhost", Port: "27017", Database: "NHSDB"}
// 	// model.Initiate(&GenericCountersTelemetryData{}, config)

// 	vars := mux.Vars(r)
// 	ipaddress := vars["ipaddress"]
// 	portAddress := vars["portaddress"]
// 	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		fmt.Println("mongo.Connect() ERROR:", err)
// 		os.Exit(1)
// 	} else {
// 		fmt.Println("Connected to mongoDB")
// 	}

// 	fmt.Print(ipaddress, portAddress)
// 	output := getData1(client, ipaddress, portAddress)
// 	xy, err := json.Marshal(output)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	rw.Write(xy)
// }

// func getAllCPUInfo(rw http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	ipaddress := vars["ipaddress"]
// 	portAddress := vars["portaddress"]
// 	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		fmt.Println("mongo.Connect() ERROR:", err)
// 		os.Exit(1)
// 	} else {
// 		fmt.Println("Connected to mongoDB")
// 	}

// 	output := getData(client, ipaddress, portAddress)
// 	xy, err := json.Marshal(output)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	rw.Write(xy)
// }

// func getData(client *mongo.Client, ipaddress string, port string) (cpuData []CPUTelemetryData) {

// 	// collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
// 	collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	model := &GcModel{}
// 	config := base.DBConfig{Driver: base.Mongo, Host: "localhost", Port: "27017", Database: "NHSDB"}
// 	model.Initiate(&GenericCountersTelemetryData{}, config)

// 	if ipaddress != "" {
// 		// cursor, err := collectionGC.Find(ctx, bson.M{"source": ipaddress + ":" + port})
// 		// err = cursor.All(context.Background(), &gcData)
// 		cursor, err := collectionCPU.Find(ctx, bson.M{"source": ipaddress + ":" + port})

// 		err = cursor.All(context.Background(), &cpuData)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	} else {
// 		fmt.Println("Messages")
// 		// cursor, err := collectionGC.Find(ctx, bson.M{})
// 		// err = cursor.All(context.Background(), &gcData)
// 		cursor, err := collectionCPU.Find(ctx, bson.M{})
// 		err = cursor.All(context.Background(), &cpuData)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// fmt.Println(cpuData)
// 	}
// 	return cpuData
// }

// func getData1(client *mongo.Client, ipaddress string, port string) (gcData []GenericCountersTelemetryData) {

// 	collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
// 	// collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	if ipaddress != "" {
// 		cursor, err := collectionGC.Find(ctx, bson.M{"source": ipaddress + ":" + port})
// 		err = cursor.All(context.Background(), &gcData)
// 		// cursor, err := collectionCPU.Find(ctx, bson.M{"source": ipaddress + ":" + port})
// 		// err = cursor.All(context.Background(), &cpuData)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	} else {
// 		fmt.Println("Messages")
// 		cursor, err := collectionGC.Find(ctx, bson.M{})
// 		err = cursor.All(context.Background(), &gcData)
// 		// cursor, err := collectionCPU.Find(ctx, bson.M{})
// 		// err = cursor.All(context.Background(), &cpuData)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// fmt.Println(gcData)
// 	}
// 	return gcData
// }

// func consumeGenericCountersData(client *mongo.Client, ctx context.Context) {

// 	l := log.New(os.Stdout, "kafka reader: ", 0)

// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topicGC,
// 		Logger:  l,
// 	})

// 	collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
// 	// collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")

// 	// model := &GcModel{}
// 	// config := base.DBConfig{Driver: base.Mongo, Host: "localhost", Port: "27017", Database: "NHSDB"}
// 	// model.Initiate(&GenericCountersTelemetryData{}, config)

// 	for {
// 		msg, err := r.ReadMessage(ctx)
// 		if err != nil {
// 			fmt.Println("could not read message " + err.Error())
// 		} else {
// 			fmt.Println(msg)
// 		}
// 		var u DType
// 		err = json.Unmarshal(msg.Value, &u)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		fmt.Println(u)

// 		if u.Telemetry.EncodingPath == "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters" {

// 			var x GenericCountersTelemetryData
// 			fmt.Println("Consuming data in GC")
// 			err = json.Unmarshal(msg.Value, &x)
// 			if err != nil {
// 				log.Print(err)
// 			}
// 			fmt.Println(x)

// 			// gcModel := NewGCTelemetryData()
// 			// e := model.Create(&x)
// 			// gcModel.CloseClient()

// 			_, e := collectionGC.InsertOne(ctx, x)
// 			if e != nil {
// 				fmt.Println("error:", e)
// 			} else {
// 				fmt.Println("Successfully Inserted")
// 			}
// 		}
// 		// else {
// 		// 	var x CPUTelemetryData
// 		// 	err = json.Unmarshal(msg.Value, &x)
// 		// 	fmt.Println("Consuming data in CPU")
// 		// 	fmt.Println(x)

// 		// 	if err != nil {
// 		// 		log.Print(err)
// 		// 	}
// 		// 	cpuModel := NewGCTelemetryData()
// 		// 	e := cpuModel.Create(&x)
// 		// 	// _, e := collectionCPU.InsertOne(ctx, x)
// 		// 	if e != nil {
// 		// 		fmt.Println("error:", e)
// 		// 	} else {
// 		// 		fmt.Println("Successfully Inserted")
// 		// 	}
// 		// }
// 	}
// }

// func consumeCPUProcessData(client *mongo.Client, ctx context.Context) {

// 	l := log.New(os.Stdout, "kafka reader: ", 0)

// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topicGC,
// 		Logger:  l,
// 	})

// 	// model := &CpuModel{}
// 	// config := base.DBConfig{Driver: base.Mongo, Host: "localhost", Port: "27017", Database: "NHSDB"}
// 	// model.Initiate(&CPUTelemetryData{}, config)

// 	// collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
// 	collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")

// 	for {
// 		msg, err := r.ReadMessage(ctx)
// 		if err != nil {
// 			fmt.Println("could not read message " + err.Error())
// 		} else {
// 			fmt.Println(msg)
// 		}
// 		var u DType
// 		err = json.Unmarshal(msg.Value, &u)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		fmt.Println(u)

// 		if u.Telemetry.EncodingPath != "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters" {

// 			var x CPUTelemetryData
// 			err = json.Unmarshal(msg.Value, &x)
// 			fmt.Println("Consuming data in CPU")
// 			fmt.Println(x)

// 			if err != nil {
// 				log.Print(err)
// 			}
// 			// cpuModel := NewGCTelemetryData()
// 			// e := model.Create(&x)
// 			// cpuModel.CloseClient()
// 			_, e := collectionCPU.InsertOne(ctx, x)
// 			if e != nil {
// 				fmt.Println("error:", e)
// 			} else {
// 				fmt.Println("Successfully Inserted")
// 			}
// 		}
// 	}

// }

// type Telemetry struct {
// 	NodeIDStr           string `json:"node_id_str"`
// 	SubscriptionIDStr   string `json:"subscription_id_str"`
// 	EncodingPath        string `json:"encoding_path"`
// 	CollectionID        int    `json:"collection_id"`
// 	CollectionStartTime int64  `json:"collection_start_time"`
// 	MsgTimestamp        int64  `json:"msg_timestamp"`
// 	CollectionEndTime   int64  `json:"collection_end_time"`
// }

// type CPUProcessCPUPIPELINEEDIT struct {
// 	ProcessCPUFifteenMinute int    `json:"process-cpu-fifteen-minute"`
// 	ProcessCPUFiveMinute    int    `json:"process-cpu-five-minute"`
// 	ProcessCPUOneMinute     int    `json:"process-cpu-one-minute"`
// 	ProcessID               int    `json:"process-id"`
// 	ProcessName             string `json:"process-name"`
// }

// type RequiredData struct {
// 	ProcessName           string `json:"process-name"`
// 	EncodingPath          string `json:"encoding_path"`
// 	CollectionID          int    `json:"collection_id"`
// 	NodeIDStr             string `json:"node_id_str"`
// 	TotalCPUFifteenMinute int    `json:"total-cpu-fifteen-minute"`
// 	TotalCPUFiveMinute    int    `json:"total-cpu-five-minute"`
// 	TotalCPUOneMinute     int    `json:"total-cpu-one-minute"`
// 	Source                string `json:"Source"`
// }

// type CPUContent struct {
// 	ProcessCPUPIPELINEEDIT []CPUProcessCPUPIPELINEEDIT
// 	TotalCPUFifteenMinute  int `json:"total-cpu-fifteen-minute"`
// 	TotalCPUFiveMinute     int `json:"total-cpu-five-minute"`
// 	TotalCPUOneMinute      int `json:"total-cpu-one-minute"`
// }

// type CPUKey struct {
// 	NodeName string `json:"node-name"`
// }

// type CPURow struct {
// 	TimeStamp int64      `json:"TimeStamp"`
// 	Keys      CPUKey     `json:"Keys"`
// 	Content   CPUContent `json:"Content"`
// }

// type GenericCountersKey struct {
// 	InterfaceName string `json:"interface-name"`
// }

// type GenericCountersContent struct {
// 	Applique                       int   `json:"applique"`
// 	AvailabilityFlag               int   `json:"availability-flag"`
// 	BroadcastPacketsReceived       int   `json:"broadcast-packets-received"`
// 	BroadcastPacketsSent           int   `json:"broadcast-packets-sent"`
// 	BytesReceived                  int   `json:"bytes-received"`
// 	BytesSent                      int64 `json:"bytes-sent"`
// 	CarrierTransitions             int   `json:"carrier-transitions"`
// 	CrcErrors                      int   `json:"crc-errors"`
// 	FramingErrorsReceived          int   `json:"framing-errors-received"`
// 	GiantPacketsReceived           int   `json:"giant-packets-received"`
// 	InputAborts                    int   `json:"input-aborts"`
// 	InputDrops                     int   `json:"input-drops"`
// 	InputErrors                    int   `json:"input-errors"`
// 	InputIgnoredPackets            int   `json:"input-ignored-packets"`
// 	InputOverruns                  int   `json:"input-overruns"`
// 	InputQueueDrops                int   `json:"input-queue-drops"`
// 	LastDataTime                   int   `json:"last-data-time"`
// 	LastDiscontinuityTime          int   `json:"last-discontinuity-time"`
// 	MulticastPacketsReceived       int   `json:"multicast-packets-received"`
// 	MulticastPacketsSent           int   `json:"multicast-packets-sent"`
// 	OutputBufferFailures           int   `json:"output-buffer-failures"`
// 	OutputBuffersSwappedOut        int   `json:"output-buffers-swapped-out"`
// 	OutputDrops                    int   `json:"output-drops"`
// 	OutputErrors                   int   `json:"output-errors"`
// 	OutputQueueDrops               int   `json:"output-queue-drops"`
// 	OutputUnderruns                int   `json:"output-underruns"`
// 	PacketsReceived                int   `json:"packets-received"`
// 	PacketsSent                    int   `json:"packets-sent"`
// 	ParityPacketsReceived          int   `json:"parity-packets-received"`
// 	Resets                         int   `json:"resets"`
// 	RuntPacketsReceived            int   `json:"runt-packets-received"`
// 	SecondsSinceLastClearCounters  int   `json:"seconds-since-last-clear-counters"`
// 	SecondsSincePacketReceived     int   `json:"seconds-since-packet-received"`
// 	SecondsSincePacketSent         int   `json:"seconds-since-packet-sent"`
// 	ThrottledPacketsReceived       int   `json:"throttled-packets-received"`
// 	UnknownProtocolPacketsReceived int   `json:"unknown-protocol-packets-received"`
// }

// type GenericCountersRow struct {
// 	Timestamp int64                  `json:"Timestamp"`
// 	Keys      GenericCountersKey     `json:"Keys"`
// 	Content   GenericCountersContent `json:"Content"`
// }
// type DType struct {
// 	Source    string    `json:"Source"`
// 	Telemetry Telemetry `json:"Telemetry"`
// }

// type GenericCountersTelemetryData struct {
// 	octopus.MongoScheme
// 	Id        string               `json:"_id"`
// 	Source    string               `json:"Source"`
// 	Telemetry Telemetry            `json:"Telemetry"`
// 	Rows      []GenericCountersRow `json:"Rows"`
// }

// type CPUTelemetryData struct {
// 	octopus.MongoScheme
// 	Id        string    `json:"_id"`
// 	Source    string    `json:"Source"`
// 	Telemetry Telemetry `json:"Telemetry"`
// 	Rows      []CPURow  `json:"Rows"`
// }

// type CPUTelemetryDataModel struct {
// 	octopus.Model
// }

// func NewCPUTelemetryData() *CPUTelemetryDataModel {
// 	model := &CPUTelemetryDataModel{}
// 	config := base.DBConfig{Driver: base.Mongo, Host: "localhost", Port: "27017", Database: "CPUTeleData"}
// 	model.Initiate(&CPUTelemetryData{}, config)

// 	return model
// }

// func NewGCTelemetryData() *GCTelemetryDataModel {
// 	model1 := &GCTelemetryDataModel{}
// 	config := base.DBConfig{Driver: base.Mongo, Host: "localhost", Port: "27017", Database: "GCTeleData"}
// 	model1.Initiate(&GenericCountersTelemetryData{}, config)

// 	return model1
// }

// type GCTelemetryDataModel struct {
// 	octopus.Model
// }

// func (u GenericCountersTelemetryData) GetID() interface{} {
// 	return u.Id
// }
// func (u CPUTelemetryData) GetID() interface{} {
// 	return u.Id
// }

// type GcModel struct {
// 	octopus.Model
// }
// type CpuModel struct {
// 	octopus.Model
// }

// type Device struct {
// 	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Name        string             `json:"Name"`
// 	Ip          string             `json:"Ip"`
// 	Port        string             `json:"Port"`
// 	Username    string             `json:"Username"`
// 	Password    string             `json:"Password"`
// 	IsEnabled   bool               `json:"IsEnabled"`
// 	UseNSO      bool               `json:"UseNSO"`
// 	MacDetail   MacDetails         `json:"MacDetails"`
// 	Coordinates Coordinates        `json:"Coordinates"`
// }

// type MacDetails struct {
// 	// Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Macaddress string `json:"Macaddress"`
// 	Area       string `json:"Area"`
// 	City       string `json:"City"`
// 	State      string `json:"State"`
// 	Country    string `json:"Country"`
// }

// type Coordinates struct {
// 	Latitude  json.Number `json:"Latitude"`
// 	Longitude json.Number `json:"Longitude"`
// }
