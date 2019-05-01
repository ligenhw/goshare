package configration

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"runtime"
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

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

func loadconfig() {
	currentPath := getCurrentPath()

	file, err := os.Open(currentPath + "/config.json")
	if err != nil {
		log.Fatalln("Can not open config file", err)
	}
	decoder := json.NewDecoder(file)
	Conf = Config{}
	err = decoder.Decode(&Conf)
	if err != nil {
		log.Fatalln("Cannot get Conf from file", err)
	}
	log.Println("load config : ", Conf)

}
