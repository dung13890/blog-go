package model

import (
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (u *UserRepo) Get(q map[string][]string) ([]User, int) {
	limit := 10
	offset := 0
	order := "id"
	sort := ""
	if value, ok := q["limit"]; ok {
		limit, _ = strconv.Atoi(value[0])
	}
	if value, ok := q["offset"]; ok {
		offset, _ = strconv.Atoi(value[0])
	}
	if value, ok := q["order"]; ok {
		order = value[0]
	}
	if value, ok := q["sort"]; ok {
		if value[0] == "asc" {
			sort = ""
		} else {
			sort = "-"
		}
	}
	users := []User{}
	count, _ := u.C.Find(nil).Count()
	iter := u.C.Find(nil).Sort(sort + order).Limit(limit).Skip(offset).Iter()
	user := User{}
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users, count
}

func (u *UserRepo) Create(user *User) error {
	user.Id = bson.NewObjectId()
	user.Date = time.Now()
	err := u.C.Insert(&user)
	return err
}

func (u *UserRepo) Delete(id string) error {
	err := u.C.RemoveId(bson.ObjectIdHex(id))
	return err
}
