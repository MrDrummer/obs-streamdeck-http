package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type secretsType struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Identity string `json:"identity"`
	Token    string `json:"token"`
	IsHTTPS  bool   `json:"https"`
}

var secrets secretsType // Global var (nasty)

func init() {
	data, err := os.ReadFile("../secrets.json")
	if err != nil {
		log.Fatal(fmt.Errorf("error reading JSON file! %w", err))
	}

	if err := json.Unmarshal(data, &secrets); err != nil {
		log.Fatal(fmt.Errorf("error decoding JSON! %w", err))
	}
}
