package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

type config struct {
	Server      string
	MongoDBHost string
	MongoDBUser string
	MongoDBPwd  string
	Database    string
}

var (
	AppConfig config
	session   *mgo.Session
)

func Init() {
	AppConfig = loadConfig()
}

func loadConfig() config {
	conf := config{}
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	return conf
}

func createSession() *mgo.Session {
	value, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.MongoDBUser,
		Password: AppConfig.MongoDBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[ConnectDB]: %s\n", err)
	}
	return value
}

func ConnectDb() {
	if session == nil {
		session = createSession()
	}
	return session
}
