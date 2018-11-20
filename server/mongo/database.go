// Package mongo deals with all of mello's mongo database data
package mongo

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
)

// Database is the singleton instance of the mgo session
var Database *mgo.Database

// MongoConfig is configuration that mgo takes to connect to MongoDB
type MongoConfig struct {
	APIKey string
	DBHost string
	DBPort string
	DBName string
}

// Initializes the Mongo session constant
func NewDBInstance(config MongoConfig) *mgo.Database {
	if Database != nil {
		return Database
	}

	url := config.DBHost + ":" + config.DBPort
	fmt.Printf("Connecting to MongoDB: %v ...\n", url)
	session, err := mgo.Dial(url)

	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	fmt.Println("Mongo successfully connected")
	Database = session.DB(config.DBName)
	return Database
}
