package models

import (
	"fmt"
	"mello/server/mongo"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/thoas/go-funk"
	"golang.org/x/crypto/bcrypt"
)

const (
	userCollection = "users"
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
			"required": []string{"username", "email", "password"},
			"properties": bson.M{
				"username": bson.M{
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
	Name:           userCollection,
	CollectionInfo: info,
	Indexes:        []mgo.Index{index},
}

// getKeyErrors is a helper that gets the key/form errors in user validation
func getKeyErrors(u *User, results []User) error {
	var keys = make(map[string]bool)
	for _, user := range results {
		if u.Email == user.Email {
			keys["email"] = true
		}
		if u.Username == user.Username {
			keys["username"] = true
		}
	}

	var keyErrors = funk.Keys(keys).([]string)
	if len(keyErrors) > 0 {
		return mongo.CreateFormError("User", keyErrors)
	}

	return nil
}

// getPasswordHash hashes and salts a given password using bcrypt
func getPasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

// Save will try to save a User to the database if it is valid
func (u *User) Save() error {
	now := time.Now()
	session := mongo.Database.Session.Copy()
	c := mongo.Database.C(userCollection).With(session)
	defer session.Close()

	if err := u.Valid(); err != nil {
		fmt.Println(err)
		return err
	}

	password, _ := getPasswordHash(u.Password)
	err := c.Insert(bson.M{
		"username": u.Username,
		"email":    u.Email,
		"password": password,
		"created":  now,
		"updated":  now,
	})

	return err
}

func (u *User) Valid() error {
	var results []User
	if err := mongo.Database.C(userCollection).Find(bson.M{
		"$or": []bson.M{
			bson.M{"username": u.Username},
			bson.M{"email": u.Email},
		},
	}).All(&results); err != nil {
		return err
	}
	fmt.Println(results)
	if err := getKeyErrors(u, results); err != nil {
		return err
	}

	return nil
}

// FindUserByName is a function to help find users in the database
func FindUserByName(name string) (*User, error) {
	var result *User
	err := mongo.Database.C(userCollection).Find(name).One(&result)

	return result, err
}
