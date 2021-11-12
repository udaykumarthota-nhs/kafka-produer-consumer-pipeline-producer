package models

type Devices struct {
	Name      string `json:"Name"`
	Ip        string `json:"Ip"`
	Port      string `json:"Port"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	IsEnabled bool   `json:"IsEnabled"`
	UseNSO    bool   `json:"UseNSO"`
}