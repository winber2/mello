package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const staticDirectory string = "/static/"
const clientDirectory string = "/client/public/"

// Configuration determines the environment variables of the server
type Configuration struct {
	APIKey string `json:"apiKey"`
}

func loadConfiguration() {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println("No configuration found. Make sure you have a config.json file inside the server folder.")
		log.Fatal(err)
	}

	var config Configuration
	var jsonParser = json.NewDecoder(jsonFile)
	jsonParser.Decode(&config)

	defer jsonFile.Close()
	os.Setenv("API_KEY", config.APIKey)
}

func main() {
	dir, err := filepath.Abs(filepath.Dir("../"))

	if err != nil {
		log.Fatal(err)
	}

	loadConfiguration()

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir+staticDirectory))))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(dir+clientDirectory))))
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dir+clientDirectory+"index.html")
	})

	http.Handle("/", r)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server crashed and burned")
		log.Fatal(err)
	}
}
