package utils

import (
	"encoding/json"
	"log"
	"os"
)

var GlobalObject *GlobalObj

type GlobalObj struct {
	IPVersion     string `json:"ip_version"`
	IP            string `json:"ip"`
	Port          string `json:"port"`
	MaxPacketSize uint32 `json:"max_packet_size"`
}

func (g *GlobalObj) Reload() {
	file, err := os.Open("config/GlobalObj.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&GlobalObject)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		IPVersion:     "tcp4",
		IP:            "127.0.0.1",
		Port:          "8089",
		MaxPacketSize: 4096,
	}
	GlobalObject.Reload()
}
