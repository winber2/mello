package models

import (
	"fmt"
	"mello/server/mongo"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// User table contains the information for each user
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

var info = &mgo.CollectionInfo{
	Validator: bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"name", "email", "password"},
			"properties": bson.M{
				"name": bson.M{
					"bsonType":  "string",
					"minLength": 6,
					"maxLength": 20,
				},
				"email": bson.M{
					"bsonType":  "string",
					"minLength": 1,
					"maxLength": 256,
					"pattern":   "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$",
				},
				"password": bson.M{
					"bsonType":  "string",
					"minLength": 6,
				},
			},
		},
	},
}

var index = mgo.Index{
	Name:     "UserIndex",
	Key:      []string{"username", "email"},
	Unique:   true,
	DropDups: true,
	Sparse:   true,
}

// UserCollection defines the settings of the MongoDB Collection that Users will be saved to
var UserCollection = mongo.Collection{
	Name:           "users",
	CollectionInfo: info,
	Indexes:        []mgo.Index{index},
}

// Save will try to save a User to the database if it is valid
func (u *User) Save() error {
	now := time.Now()
	session := mongo.Database.Session.Copy()
	defer session.Close()

	c := mongo.Database.C("users").With(session)
	result, _ := FindUserByName(u.Username)

	if result != nil {
		fmt.Println("asdf")
		return mongo.CreateFormError("User", []string{"username"})
	}

	password, _ := getPasswordHash(u.Password)

	err := c.Insert(bson.M{
		"name":     u.Username,
		"email":    u.Email,
		"password": password,
		"created":  now,
		"updated":  now,
	})

	return err
}

func (u *User) Valid() error {
	return nil
}

// FindUserByName is a function to help find users in the database
func FindUserByName(name string) (*User, error) {
	var result *User
	err := mongo.Database.C("users").Find(name).One(&result)

	return result, err
}

// getPasswordHash hashes and salts a given password using bcrypt
func getPasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}
