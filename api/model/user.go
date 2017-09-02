package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Email string        `json:"email,omitempty"`
	Date  time.Time     `json:"date, omitempty"`
}

type UserRepo struct {
	C *mgo.Collection
}

func (u *UserRepo) GetAll() []User {
	users := []User{}
	iter := u.C.Find(nil).Iter()
	user := User{}
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users
}
