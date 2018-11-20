package models

import (
	"mello/server/mongo"
	"time"

	"github.com/globalsign/mgo/bson"
)

// User table contains the information for each user
type User struct {
	Name     string `db:"name" bson:"name"`
	Email    string `db:"email" bson:"email"`
	Password string `db:"password" bson:"password"`
}

func (user *User) Save() error {
	now := time.Now()
	session := mongo.Database.Session.Copy()
	defer session.Close()

	c := mongo.Database.C("user").With(session)

	err := c.Insert(bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"created":  now,
		"updated":  now,
	})

	return err
}
