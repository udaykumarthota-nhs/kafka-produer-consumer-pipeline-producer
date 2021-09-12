package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	topicGC = "telemetrycheckGC311"
	// topicCPU      = "telemetrycheckCPU1"
	brokerAddress = "localhost:9092"
)

func main() {

	fmt.Println("Hello World")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx := context.Background()
	client, e := mongo.Connect(ctx, clientOptions)
	if e != nil {
		fmt.Println("error:", e)
	}
	// go consumeCPUProcessData(client, ctx)
	go consumeGenericCountersData(client, ctx)
	HandleRouter()

}

func HandleRouter() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "requested for index page")
	})

	r.HandleFunc("/fetch/gc/all", getAllGCInfo)
	r.HandleFunc("/fetch/cpu/all", getAllCPUInfo)
	r.HandleFunc("/fetch/gc/all/{ipaddress}/{portaddress}", getAllGCInfo)
	r.HandleFunc("/fetch/cpu/all/{ipaddress}/{portaddress}", getAllCPUInfo)
	r.HandleFunc("/fetch/ip/{ipaddress}/{portaddress}", filterByIPAddress)
	http.ListenAndServe(":8088", handler)
}

func filterByIPAddress(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ipaddress := vars["ipaddress"]
	portAddress := vars["portaddress"]
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	} else {
		fmt.Println("Connected to mongoDB")
	}

	cpuData := getData(client, ipaddress, portAddress)

	z := OutputData{
		cpuProcessData: cpuData,
		// genericCountersData: gcData,
	}
	x, er := json.Marshal(z)
	if er != nil {
		fmt.Println(er)
	} else {
		rw.Write(x)
	}
}

type OutputData struct {
	cpuProcessData      []CPUTelemetryData
	genericCountersData []GenericCountersTelemetryData
}

func getAllGCInfo(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipaddress := vars["ipaddress"]
	portAddress := vars["portaddress"]
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	} else {
		fmt.Println("Connected to mongoDB")
	}

	fmt.Print(ipaddress, portAddress)
	output := getData1(client, ipaddress, portAddress)
	xy, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	rw.Write(xy)
}

func getAllCPUInfo(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ipaddress := vars["ipaddress"]
	portAddress := vars["portaddress"]
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	} else {
		fmt.Println("Connected to mongoDB")
	}

	output := getData(client, ipaddress, portAddress)
	xy, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	rw.Write(xy)
}

func getData(client *mongo.Client, ipaddress string, port string) (cpuData []CPUTelemetryData) {

	// collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
	collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if ipaddress != "" {
		// cursor, err := collectionGC.Find(ctx, bson.M{"source": ipaddress + ":" + port})
		// err = cursor.All(context.Background(), &gcData)
		cursor, err := collectionCPU.Find(ctx, bson.M{"source": ipaddress + ":" + port})
		err = cursor.All(context.Background(), &cpuData)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Messages")
		// cursor, err := collectionGC.Find(ctx, bson.M{})
		// err = cursor.All(context.Background(), &gcData)
		cursor, err := collectionCPU.Find(ctx, bson.M{})
		err = cursor.All(context.Background(), &cpuData)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(cpuData)
	}
	return cpuData
}

func getData1(client *mongo.Client, ipaddress string, port string) (gcData []GenericCountersTelemetryData) {

	collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
	// collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if ipaddress != "" {
		cursor, err := collectionGC.Find(ctx, bson.M{"source": ipaddress + ":" + port})
		err = cursor.All(context.Background(), &gcData)
		// cursor, err := collectionCPU.Find(ctx, bson.M{"source": ipaddress + ":" + port})
		// err = cursor.All(context.Background(), &cpuData)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Messages")
		cursor, err := collectionGC.Find(ctx, bson.M{})
		err = cursor.All(context.Background(), &gcData)
		// cursor, err := collectionCPU.Find(ctx, bson.M{})
		// err = cursor.All(context.Background(), &cpuData)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(gcData)
	}
	return gcData
}

func consumeGenericCountersData(client *mongo.Client, ctx context.Context) {

	l := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topicGC,
		Logger:  l,
	})

	collectionGC := client.Database("ConsumedDataDB1").Collection("GC")
	collectionCPU := client.Database("ConsumedDataDB1").Collection("CPU")

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}

		var u DType
		err = json.Unmarshal(msg.Value, &u)
		if err != nil {
			log.Print(err)
		}

		if u.Telemetry.EncodingPath == "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters" {
			var x GenericCountersTelemetryData
			fmt.Println("Consuming data in GC")
			err = json.Unmarshal(msg.Value, &x)
			if err != nil {
				log.Print(err)
			}
			fmt.Println(x)
			_, e := collectionGC.InsertOne(ctx, x)
			if e != nil {
				fmt.Println("error:", e)
			}
		} else {
			var x CPUTelemetryData
			err = json.Unmarshal(msg.Value, &x)
			fmt.Println("Consuming data in CPU")

			if err != nil {
				log.Print(err)
			}

			_, e := collectionCPU.InsertOne(ctx, x)
			if e != nil {
				fmt.Println("error:", e)
			}
		}
	}
}

// func consumeCPUProcessData(client *mongo.Client, ctx context.Context) {
// 	l := log.New(os.Stdout, "kafka reader: ", 0)

// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topicCPU,
// 		Logger:  l,
// 	})
// 	type Message struct {
// 		Key   string
// 		Value string
// 	}

// 	collection := client.Database("ConsumedDataDB1").Collection("CPU")

// 	for {
// 		msg, err := r.ReadMessage(ctx)
// 		if err != nil {
// 			fmt.Println("could not read message " + err.Error())
// 		}
// 		fmt.Println("Consuming data in CPU Process")
// 		fmt.Println("received: ", string(msg.Value))
// 		// l := TelemetryData{msg.Value}
// 		var u CPUTelemetryData
// 		err = json.Unmarshal(msg.Value, &u)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		fmt.Printf("u: %+v\n", u)

// 		_, e := collection.InsertOne(ctx, u)
// 		if e != nil {
// 			fmt.Println("error:", e)
// 		}
// 	}
// }

type Telemetry struct {
	NodeIDStr           string `json:"node_id_str"`
	SubscriptionIDStr   string `json:"subscription_id_str"`
	EncodingPath        string `json:"encoding_path"`
	CollectionID        int    `json:"collection_id"`
	CollectionStartTime int64  `json:"collection_start_time"`
	MsgTimestamp        int64  `json:"msg_timestamp"`
	CollectionEndTime   int64  `json:"collection_end_time"`
}

type CPUProcessCPUPIPELINEEDIT struct {
	ProcessCPUFifteenMinute int    `json:"process-cpu-fifteen-minute"`
	ProcessCPUFiveMinute    int    `json:"process-cpu-five-minute"`
	ProcessCPUOneMinute     int    `json:"process-cpu-one-minute"`
	ProcessID               int    `json:"process-id"`
	ProcessName             string `json:"process-name"`
}

type RequiredData struct {
	ProcessName           string `json:"process-name"`
	EncodingPath          string `json:"encoding_path"`
	CollectionID          int    `json:"collection_id"`
	NodeIDStr             string `json:"node_id_str"`
	TotalCPUFifteenMinute int    `json:"total-cpu-fifteen-minute"`
	TotalCPUFiveMinute    int    `json:"total-cpu-five-minute"`
	TotalCPUOneMinute     int    `json:"total-cpu-one-minute"`
	Source                string `json:"Source"`
}

type CPUContent struct {
	ProcessCPUPIPELINEEDIT []CPUProcessCPUPIPELINEEDIT
	TotalCPUFifteenMinute  int `json:"total-cpu-fifteen-minute"`
	TotalCPUFiveMinute     int `json:"total-cpu-five-minute"`
	TotalCPUOneMinute      int `json:"total-cpu-one-minute"`
}

type CPUKey struct {
	NodeName string `json:"node-name"`
}

type CPURow struct {
	TimeStamp int64      `json:"TimeStamp"`
	Keys      CPUKey     `json:"Keys"`
	Content   CPUContent `json:"Content"`
}

type CPUTelemetryData struct {
	Source    string    `json:"Source"`
	Telemetry Telemetry `json:"Telemetry"`
	Rows      []CPURow  `json:"Rows"`
}

type GenericCountersKey struct {
	InterfaceName string `json:"interface-name"`
}

type GenericCountersContent struct {
	Applique                       int   `json:"applique"`
	AvailabilityFlag               int   `json:"availability-flag"`
	BroadcastPacketsReceived       int   `json:"broadcast-packets-received"`
	BroadcastPacketsSent           int   `json:"broadcast-packets-sent"`
	BytesReceived                  int   `json:"bytes-received"`
	BytesSent                      int64 `json:"bytes-sent"`
	CarrierTransitions             int   `json:"carrier-transitions"`
	CrcErrors                      int   `json:"crc-errors"`
	FramingErrorsReceived          int   `json:"framing-errors-received"`
	GiantPacketsReceived           int   `json:"giant-packets-received"`
	InputAborts                    int   `json:"input-aborts"`
	InputDrops                     int   `json:"input-drops"`
	InputErrors                    int   `json:"input-errors"`
	InputIgnoredPackets            int   `json:"input-ignored-packets"`
	InputOverruns                  int   `json:"input-overruns"`
	InputQueueDrops                int   `json:"input-queue-drops"`
	LastDataTime                   int   `json:"last-data-time"`
	LastDiscontinuityTime          int   `json:"last-discontinuity-time"`
	MulticastPacketsReceived       int   `json:"multicast-packets-received"`
	MulticastPacketsSent           int   `json:"multicast-packets-sent"`
	OutputBufferFailures           int   `json:"output-buffer-failures"`
	OutputBuffersSwappedOut        int   `json:"output-buffers-swapped-out"`
	OutputDrops                    int   `json:"output-drops"`
	OutputErrors                   int   `json:"output-errors"`
	OutputQueueDrops               int   `json:"output-queue-drops"`
	OutputUnderruns                int   `json:"output-underruns"`
	PacketsReceived                int   `json:"packets-received"`
	PacketsSent                    int   `json:"packets-sent"`
	ParityPacketsReceived          int   `json:"parity-packets-received"`
	Resets                         int   `json:"resets"`
	RuntPacketsReceived            int   `json:"runt-packets-received"`
	SecondsSinceLastClearCounters  int   `json:"seconds-since-last-clear-counters"`
	SecondsSincePacketReceived     int   `json:"seconds-since-packet-received"`
	SecondsSincePacketSent         int   `json:"seconds-since-packet-sent"`
	ThrottledPacketsReceived       int   `json:"throttled-packets-received"`
	UnknownProtocolPacketsReceived int   `json:"unknown-protocol-packets-received"`
}

type GenericCountersRow struct {
	Timestamp int64                  `json:"Timestamp"`
	Keys      GenericCountersKey     `json:"Keys"`
	Content   GenericCountersContent `json:"Content"`
}

type GenericCountersTelemetryData struct {
	Source    string               `json:"Source"`
	Telemetry Telemetry            `json:"Telemetry"`
	Rows      []GenericCountersRow `json:"Rows"`
}

type DType struct {
	Source    string    `json:"Source"`
	Telemetry Telemetry `json:"Telemetry"`
}
