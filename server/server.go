package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mello/server/app/routes"
	"mello/server/mongo"
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
	config MelloConfig
)

// MelloConfig determines the environment variables of the server
type MelloConfig struct {
	APIKey string `json:"apiKey"`
	DBHost string `json:"dbHost"`
	DBPort string `json:"dbPort"`
	DBName string `json:"dbName"`
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

func main() {
	dir, err := filepath.Abs(filepath.Dir("../"))

	if err != nil {
		log.Fatal(err)
	}

	loadConfiguration("config.json")
	mongo.NewDBInstance(mongo.Config(config))

	router := mux.NewRouter()
	// router.Use(ContentTypeJSONMiddlewar)

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir+staticDirectory))))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir(dir+clientDirectory))))

	// Add app api routes to mux router
	routes.AppendAppRoutes(router)

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
