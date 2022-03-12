package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type IPPOOL struct {
	Source    string    `json:"Source"`
	Telemetry Telemetry `json:"Telemetry"`
	Pools     []POOL    `json:"Pools"`
}

type Telemetry struct {
	NodeIDStr           string `json:"node_id_str"`
	SubscriptionIDStr   string `json:"subscription_id_str"`
	EncodingPath        string `json:"encoding_path"`
	CollectionID        int    `json:"collection_id"`
	CollectionStartTime int64  `json:"collection_start_time"`
	MsgTimestamp        int64  `json:"msg_timestamp"`
	CollectionEndTime   int64  `json:"collection_end_time"`
}

type POOLS struct {
	Pools []POOL `json:"pool"`
}

type POOL struct {
	PoolNo     int      `json:"pool_no"`
	PoolId     string   `json:"pool_id"`
	Available  []string `json:"available"`
	Used       []string `json:"used"`
	RangeStart string   `json:"range_start"`
	RangeEnd   string   `json:"range_end"`
}

const (
	topicGC = "pooling"
	// topicCPU      = "telemetrycheckCPU1"
	brokerAddress = "103.249.77.21:9092"
)

func main() {
	ctx := context.Background()
	// go produceIpPoolPackets(ctx)
	// produceCPUProcessData(ctx)
	produceGenericCountersData(ctx)
}

func produceIpPoolPackets() {

}

func produceGenericCountersData(ctx context.Context) {
	i := 0
	ipaddresses := []string{"114.157.28.57", "18.20.53.200", "104.75.99.24", "250.151.66.138", "39.28.25.101", "61.93.243.29", "219.187.94.202", "123.166.246.199", "68.186.91.220", "83.103.17.82", "248.10.206.164", "139.7.252.212", "51.2.199.78", "64.252.36.58", "72.144.157.24", "238.147.124.14", "39.213.233.186", "187.52.248.207", "37.5.26.108", "130.174.94.107", "86.172.247.112", "145.242.221.29", "176.33.175.175", "131.216.96.201", "125.106.3.199", "234.166.216.19", "96.194.133.154", "73.56.44.210", "136.137.204.214", "140.143.46.237", "122.94.14.85", "208.238.105.191", "224.226.139.10", "208.102.52.164", "243.65.46.188", "248.28.225.99", "198.247.63.23", "2.136.9.12", "89.9.130.206", "236.127.104.1", "122.238.241.150", "116.210.37.86", "31.79.1.5", "203.89.179.152", "115.139.104.27", "15.145.183.31", "47.30.198.139", "156.20.13.50", "239.108.132.160", "183.182.12.200", "59.155.76.88", "66.227.221.51", "238.128.184.17", "76.101.59.153", "181.246.221.195", "165.103.87.201", "255.182.163.130", "190.200.248.31", "69.161.241.179", "6.64.202.16", "18.236.138.40", "94.77.209.96", "216.238.135.59", "139.211.20.108", "67.38.120.197", "194.137.14.107", "246.191.137.131", "31.11.243.208", "236.92.103.220", "100.253.45.47", "68.238.124.215", "132.134.52.64", "242.140.41.77", "231.243.84.122", "5.193.50.194", "147.20.28.251", "77.95.87.82", "250.207.121.163", "185.55.245.78", "73.22.238.17", "132.123.181.34", "248.16.190.38", "207.94.173.132", "178.169.118.101", "82.163.103.103", "224.196.116.70", "201.249.54.233", "55.76.114.51", "149.233.107.8", "34.143.51.164", "252.9.132.217", "248.171.149.105", "152.146.217.36", "70.176.214.238", "47.76.191.33", "185.56.191.162", "91.52.80.151", "122.232.54.101", "177.171.226.82", "81.89.70.129"}

	l := log.New(os.Stdout, "kafka writer: GC ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topicGC,
		Logger:  l,
	})

	for {
		d, e := json.Marshal(createCPUProcessData(ipaddresses[i]))
		if e != nil {
			fmt.Errorf("could not write message " + e.Error())
		}
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: d,
		})

		if err != nil {
			fmt.Errorf("could not write message " + err.Error())
		}

		fmt.Println("writes:", i)
		i++

		time.Sleep(time.Second)
		if i == 50 {
			break
		}
	}
}

func createCPUProcessData(ipaddress string) (telemetryDataValue IPPOOL) {

	// pools := new([]POOLS)

	pool := POOL{
		PoolNo:     1,
		PoolId:     "3242",
		Available:  []string{"10.10.1.10", "10.10.1.9", "10.10.1.8", "10.10.1.7", "10.10.1.6", "10.10.1.5"},
		Used:       []string{"10.10.1.1", "10.10.1.2", "10.10.1.3", "10.10.1.4"},
		RangeStart: "fgsdfgdf",
		RangeEnd:   "sdfasd",
	}

	telemetry := Telemetry{
		NodeIDStr:           "uut",
		SubscriptionIDStr:   "test",
		EncodingPath:        "Cisco-IOS-XR-wdsysmon-fd-oper:system-monitoring/pools/ip-pooling",
		CollectionID:        111808,
		CollectionStartTime: 7987987908,
		MsgTimestamp:        342342342,
		CollectionEndTime:   342342342,
	}
	telemetryDataValue.Telemetry = telemetry

	telemetryDataValue.Source = ipaddress + ":32876"

	telemetryDataValue.Pools = []POOL{pool}
	fmt.Println(telemetryDataValue)
	return telemetryDataValue
}
