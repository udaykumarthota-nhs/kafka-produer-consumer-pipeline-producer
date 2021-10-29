// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/segmentio/kafka-go"
// )

// const (
// 	topicGC = "telemetrycheckGC315"
// 	// topicCPU      = "telemetrycheckCPU1"
// 	brokerAddress = "103.249.77.21:9092"
// )

// func main() {
// 	ctx := context.Background()
// 	go produceGenericCountersData(ctx)
// 	produceCPUProcessData(ctx)
// }

// func produceCPUProcessData(ctx context.Context) {

// 	i := 0
// 	ipaddresses := []string{"114.157.28.57", "18.20.53.200", "104.75.99.24", "250.151.66.138", "39.28.25.101", "61.93.243.29", "219.187.94.202", "123.166.246.199", "68.186.91.220", "83.103.17.82", "248.10.206.164", "139.7.252.212", "51.2.199.78", "64.252.36.58", "72.144.157.24", "238.147.124.14", "39.213.233.186", "187.52.248.207", "37.5.26.108", "130.174.94.107", "86.172.247.112", "145.242.221.29", "176.33.175.175", "131.216.96.201", "125.106.3.199", "234.166.216.19", "96.194.133.154", "73.56.44.210", "136.137.204.214", "140.143.46.237", "122.94.14.85", "208.238.105.191", "224.226.139.10", "208.102.52.164", "243.65.46.188", "248.28.225.99", "198.247.63.23", "2.136.9.12", "89.9.130.206", "236.127.104.1", "122.238.241.150", "116.210.37.86", "31.79.1.5", "203.89.179.152", "115.139.104.27", "15.145.183.31", "47.30.198.139", "156.20.13.50", "239.108.132.160", "183.182.12.200", "59.155.76.88", "66.227.221.51", "238.128.184.17", "76.101.59.153", "181.246.221.195", "165.103.87.201", "255.182.163.130", "190.200.248.31", "69.161.241.179", "6.64.202.16", "18.236.138.40", "94.77.209.96", "216.238.135.59", "139.211.20.108", "67.38.120.197", "194.137.14.107", "246.191.137.131", "31.11.243.208", "236.92.103.220", "100.253.45.47", "68.238.124.215", "132.134.52.64", "242.140.41.77", "231.243.84.122", "5.193.50.194", "147.20.28.251", "77.95.87.82", "250.207.121.163", "185.55.245.78", "73.22.238.17", "132.123.181.34", "248.16.190.38", "207.94.173.132", "178.169.118.101", "82.163.103.103", "224.196.116.70", "201.249.54.233", "55.76.114.51", "149.233.107.8", "34.143.51.164", "252.9.132.217", "248.171.149.105", "152.146.217.36", "70.176.214.238", "47.76.191.33", "185.56.191.162", "91.52.80.151", "122.232.54.101", "177.171.226.82", "81.89.70.129"}

// 	l := log.New(os.Stdout, "kafka writer: CPU Process ", 0)
// 	w := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topicGC,
// 		Logger:  l,
// 	})

// 	for {
// 		d, e := json.Marshal(createCPUProcessData(ipaddresses[i]))
// 		if e != nil {
// 			fmt.Errorf("could not write message " + e.Error())
// 		}
// 		err := w.WriteMessages(ctx, kafka.Message{
// 			Key:   []byte(strconv.Itoa(i)),
// 			Value: []byte(d),
// 		})

// 		if err != nil {
// 			fmt.Errorf("could not write message " + err.Error())
// 		}

// 		fmt.Println("writes:", i)
// 		i++

// 		time.Sleep(time.Second)
// 		if i == 50 {
// 			i = 1
// 		}
// 	}
// }

// func produceGenericCountersData(ctx context.Context) {
// 	i := 0
// 	ipaddresses := []string{"114.157.28.57", "18.20.53.200", "104.75.99.24", "250.151.66.138", "39.28.25.101", "61.93.243.29", "219.187.94.202", "123.166.246.199", "68.186.91.220", "83.103.17.82", "248.10.206.164", "139.7.252.212", "51.2.199.78", "64.252.36.58", "72.144.157.24", "238.147.124.14", "39.213.233.186", "187.52.248.207", "37.5.26.108", "130.174.94.107", "86.172.247.112", "145.242.221.29", "176.33.175.175", "131.216.96.201", "125.106.3.199", "234.166.216.19", "96.194.133.154", "73.56.44.210", "136.137.204.214", "140.143.46.237", "122.94.14.85", "208.238.105.191", "224.226.139.10", "208.102.52.164", "243.65.46.188", "248.28.225.99", "198.247.63.23", "2.136.9.12", "89.9.130.206", "236.127.104.1", "122.238.241.150", "116.210.37.86", "31.79.1.5", "203.89.179.152", "115.139.104.27", "15.145.183.31", "47.30.198.139", "156.20.13.50", "239.108.132.160", "183.182.12.200", "59.155.76.88", "66.227.221.51", "238.128.184.17", "76.101.59.153", "181.246.221.195", "165.103.87.201", "255.182.163.130", "190.200.248.31", "69.161.241.179", "6.64.202.16", "18.236.138.40", "94.77.209.96", "216.238.135.59", "139.211.20.108", "67.38.120.197", "194.137.14.107", "246.191.137.131", "31.11.243.208", "236.92.103.220", "100.253.45.47", "68.238.124.215", "132.134.52.64", "242.140.41.77", "231.243.84.122", "5.193.50.194", "147.20.28.251", "77.95.87.82", "250.207.121.163", "185.55.245.78", "73.22.238.17", "132.123.181.34", "248.16.190.38", "207.94.173.132", "178.169.118.101", "82.163.103.103", "224.196.116.70", "201.249.54.233", "55.76.114.51", "149.233.107.8", "34.143.51.164", "252.9.132.217", "248.171.149.105", "152.146.217.36", "70.176.214.238", "47.76.191.33", "185.56.191.162", "91.52.80.151", "122.232.54.101", "177.171.226.82", "81.89.70.129"}

// 	l := log.New(os.Stdout, "kafka writer: GC ", 0)
// 	w := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topicGC,
// 		Logger:  l,
// 	})

// 	for {
// 		d, e := json.Marshal(createGenericCountersData(ipaddresses[i]))
// 		if e != nil {
// 			fmt.Errorf("could not write message " + e.Error())
// 		}
// 		err := w.WriteMessages(ctx, kafka.Message{
// 			Key:   []byte(strconv.Itoa(i)),
// 			Value: d,
// 		})

// 		if err != nil {
// 			fmt.Errorf("could not write message " + err.Error())
// 		}

// 		fmt.Println("writes:", i)
// 		i++

// 		time.Sleep(time.Second)
// 		if i == 50 {
// 			break
// 		}
// 	}
// }

// func createCPUProcessData(ipaddress string) (telemetryDataValue CPUTelemetryData) {

// 	row := new(CPURow)
// 	row.Keys.NodeName = "GigabitEthernet0/0/0/1"

// 	processCpuPipeleine := CPUProcessCPUPIPELINEEDIT{
// 		ProcessCPUFifteenMinute: 54,
// 		ProcessCPUFiveMinute:    32,
// 		ProcessCPUOneMinute:     322,
// 		ProcessID:               234,
// 		ProcessName:             "dfasd",
// 	}

// 	content := CPUContent{
// 		TotalCPUFifteenMinute:  66,
// 		TotalCPUFiveMinute:     456,
// 		TotalCPUOneMinute:      4524,
// 		ProcessCPUPIPELINEEDIT: []CPUProcessCPUPIPELINEEDIT{processCpuPipeleine},
// 	}
// 	row.Content = content
// 	row.TimeStamp = 42342342

// 	telemetry := Telemetry{
// 		NodeIDStr:           "uut",
// 		SubscriptionIDStr:   "test",
// 		EncodingPath:        "Cisco-IOS-XR-wdsysmon-fd-oper:system-monitoring/cpu-utilization",
// 		CollectionID:        111808,
// 		CollectionStartTime: 7987987908,
// 		MsgTimestamp:        342342342,
// 		CollectionEndTime:   342342342,
// 	}
// 	telemetryDataValue.Telemetry = telemetry

// 	telemetryDataValue.Source = ipaddress + ":32876"

// 	telemetryDataValue.Rows = []CPURow{*row}
// 	fmt.Println(telemetryDataValue)
// 	return telemetryDataValue
// }

// func createGenericCountersData(ipaddress string) (telemetryDataValue GenericCountersTelemetryData) {

// 	row := new(GenericCountersRow)
// 	row.Keys.InterfaceName = "GigabitEthernet0/0/0/1"

// 	content := GenericCountersContent{
// 		Applique:                       12,
// 		AvailabilityFlag:               12,
// 		BroadcastPacketsReceived:       12,
// 		BroadcastPacketsSent:           12,
// 		BytesReceived:                  12,
// 		BytesSent:                      12,
// 		CarrierTransitions:             12,
// 		CrcErrors:                      12,
// 		FramingErrorsReceived:          12,
// 		GiantPacketsReceived:           12,
// 		InputAborts:                    12,
// 		InputDrops:                     12,
// 		InputErrors:                    12,
// 		InputIgnoredPackets:            12,
// 		InputOverruns:                  12,
// 		InputQueueDrops:                12,
// 		LastDataTime:                   12,
// 		LastDiscontinuityTime:          12,
// 		MulticastPacketsReceived:       12,
// 		MulticastPacketsSent:           12,
// 		OutputBufferFailures:           12,
// 		OutputBuffersSwappedOut:        12,
// 		OutputDrops:                    12,
// 		OutputErrors:                   12,
// 		OutputQueueDrops:               12,
// 		OutputUnderruns:                12,
// 		PacketsReceived:                12,
// 		PacketsSent:                    12,
// 		ParityPacketsReceived:          12,
// 		Resets:                         12,
// 		RuntPacketsReceived:            12,
// 		SecondsSinceLastClearCounters:  12,
// 		SecondsSincePacketReceived:     12,
// 		SecondsSincePacketSent:         12,
// 		ThrottledPacketsReceived:       12,
// 		UnknownProtocolPacketsReceived: 12,
// 	}
// 	row.Content = content
// 	row.Timestamp = 42342342

// 	telemetry := Telemetry{
// 		NodeIDStr:           "uut",
// 		SubscriptionIDStr:   "test",
// 		EncodingPath:        "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters",
// 		CollectionID:        111808,
// 		CollectionStartTime: 7987987908,
// 		MsgTimestamp:        342342342,
// 		CollectionEndTime:   342342342,
// 	}
// 	telemetryDataValue.Telemetry = telemetry

// 	telemetryDataValue.Source = ipaddress + ":32876"

// 	telemetryDataValue.Rows = []GenericCountersRow{*row}

// 	return telemetryDataValue
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

// type CPUTelemetryData struct {
// 	Source    string    `json:"Source"`
// 	Telemetry Telemetry `json:"Telemetry"`
// 	Rows      []CPURow  `json:"Rows"`
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

// type GenericCountersTelemetryData struct {
// 	Source    string               `json:"Source"`
// 	Telemetry Telemetry            `json:"Telemetry"`
// 	Rows      []GenericCountersRow `json:"Rows"`
// }

// type DType struct {
// 	Source    string    `json:"Source"`
// 	Telemetry Telemetry `json:"Telemetry"`
// }
