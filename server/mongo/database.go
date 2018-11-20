// Package mongo deals with all of mello's mongo database data
package mongo

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
)

// Database is the mgo Database instance
var Database *mgo.Database

// Config is configuration that mgo takes to connect to MongoDB
type Config struct {
	APIKey string
	DBHost string
	DBPort string
	DBName string
}

// NewDBInstance initializes the Mongo session constant
func NewDBInstance(config Config) *mgo.Database {
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
