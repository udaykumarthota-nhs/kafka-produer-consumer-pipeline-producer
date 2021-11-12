package models

type MacDetails struct {
	Macaddress string `json:"Macaddress"`
	Area       string `json:"Area"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
}
