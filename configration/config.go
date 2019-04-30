package configration

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Dsn          string
}

var Conf Config

func init() {
	loadconfig()
}

func loadconfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Can not open config file", err)
	}
	decoder := json.NewDecoder(file)
	Conf = Config{}
	err = decoder.Decode(&Conf)
	if err != nil {
		log.Fatalln("Cannot get Conf from file", err)
	}
}
