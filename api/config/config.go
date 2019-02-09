package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Secret      string
	Server      string
	MongoDBHost string
	MongoDBUser string
	MongoDBPwd  string
	Database    string
}

var (
	AppConfig config
)

func Init() {
	AppConfig = loadConfig()
}

func loadConfig() config {
	conf := config{}
	file, err := ioutil.ReadFile("api/config/config.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	return conf
}
