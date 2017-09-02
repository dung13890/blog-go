package config

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

type Context struct {
	MongoSession *mgo.Session
}

func connectDb() *mgo.Session {
	if session == nil {
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
	return session
}

func NewContext() *Context {
	session = connectDb()
	context := &Context{
		MongoSession: session.Copy(),
	}
	return context
}

func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(AppConfig.Database).C(name)
}

func (c *Context) Close() {
	c.MongoSession.Close()
}
