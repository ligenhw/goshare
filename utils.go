package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func p(a ...interface{}) {
	fmt.Println(a...)
}

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
}

var config Configuration

func init() {
	loadconfig()
}

func loadconfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Can not open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
