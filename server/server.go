package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"mello/server/api"
	"mello/server/mongo"
	"mello/server/mongo/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const (
	staticDirectory string = "/static/"
	clientDirectory string = "/client/public/"
)

var (
	config  MelloConfig
	resetDB bool

	collections = []mongo.Collection{
		models.UserCollection,
	}
)

// MelloConfig determines the environment variables of the server
type MelloConfig struct {
	APIKey string `json:"apiKey"`
	DBHost string `json:"dbHost"`
	DBPort string `json:"dbPort"`
	DBName string `json:"dbName"`
}

func setupCollections() {
	for _, collection := range collections {
		err := collection.Create()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadConfiguration(path string) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("No configuration found. Make sure you have a config.json file inside the server folder.")
		log.Fatal(err)
	}

	var jsonParser = json.NewDecoder(jsonFile)
	jsonParser.Decode(&config)
}

func parseFlags() {
	flag.BoolVar(&resetDB, "R", false, "Resets the database and re-creates collections with defined schema")
	flag.Parse()
}

func main() {
	parseFlags()

	dir, err := filepath.Abs(filepath.Dir("../"))

	if err != nil {
		log.Fatal(err)
	}

	loadConfiguration("config.json")
	database := mongo.NewDBInstance(mongo.Config(config))

	// Reset the database with updated schema, we will probably need a better method later on
	if resetDB {
		fmt.Println("Resetting Database")
		database.DropDatabase()
		setupCollections()
	}

	router := mux.NewRouter()
	router.Use(ContentTypeJSONMiddleware)

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir+staticDirectory))))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir(dir+clientDirectory))))

	// Append api routes to server
	router.PathPrefix("/api").Handler(http.StripPrefix("/api", api.Router()))

	// Catch all endpoint should redirect everything to the React Application
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dir+clientDirectory+"index.html")
	})

	http.Handle("/", router)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server crashed and burned")
		log.Fatal(err)
	}
}
