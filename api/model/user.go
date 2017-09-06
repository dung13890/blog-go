package model

import (
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type User struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Email string        `json:"email,omitempty"`
	Date  time.Time     `json:"date, omitempty"`
}

type Query struct {
	limit  int `json:"limit,omitempty"`
	offset int `json:offset,omitempty`
}

type UserRepo struct {
	C *mgo.Collection
}

func (u *UserRepo) Get(q map[string][]string) []User {
	limit := 10
	offset := 0
	if value, ok := q["limit"]; ok {
		limit, _ = strconv.Atoi(value[0])
	}
	if value, ok := q["offset"]; ok {
		offset, _ = strconv.Atoi(value[0])
	}
	users := []User{}
	iter := u.C.Find(nil).Limit(limit).Skip(offset).Iter()
	user := User{}
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users
}

func (u *UserRepo) Create(user *User) error {
	user.Id = bson.NewObjectId()
	user.Date = time.Now()
	err := u.C.Insert(&user)
	return err
}
