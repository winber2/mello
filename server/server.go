package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"mello/server/abclientstate"
	"mello/server/app/routes"
	"mello/server/mongo"
	"mello/server/mongo/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/volatiletech/authboss"
	"github.com/volatiletech/authboss/defaults"
)

const (
	staticDirectory   string = "/static/"
	clientDirectory   string = "/client/public/"
	sessionCookieName string = "mello_ab"
)

// MelloConfig determines the environment variables of the server
type MelloConfig struct {
	APIKey string `json:"apiKey"`
	DBHost string `json:"dbHost"`
	DBPort string `json:"dbPort"`
	DBName string `json:"dbName"`
}

func loadConfiguration(path string) MelloConfig {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println("No configuration found. Make sure you have a config.json file inside the server folder.")
		log.Fatal(err)
	}

	var config MelloConfig
	var jsonParser = json.NewDecoder(jsonFile)
	jsonParser.Decode(&config)

	return config
}

func createAuthRouter() http.Handler {
	ab := authboss.New()

	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore := abclientstate.NewCookieStorer(cookieStoreKey, nil)
	sessionStore := abclientstate.NewSessionStorer(sessionCookieName, sessionStoreKey, nil)

	ab.Config.Storage.Server = &models.UserStorer{}
	ab.Config.Storage.SessionState = sessionStore
	ab.Config.Storage.CookieState = cookieStore

	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = "http://localhost:8000/"

	// This is using the renderer from: github.com/volatiletech/authboss
	// ab.Config.Core.ViewRenderer = abrenderer.New("/auth")
	// Probably want a MailRenderer here too.

	// Set up defaults for basically everything besides the ViewRenderer/MailRenderer in the HTTP stack
	defaults.SetCore(&ab.Config, true, true)
	ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}

	if err := ab.Init(); err != nil {
		panic(err)
	}

	// Mount the router to a path (this should be the same as the Mount path above)
	// mux in this example is a chi router, but it could be anything that can route to
	// the Core.Router.
	// router.PathPrefix("/auth").Handler(ab.Config.Core.Router)

	return ab.Config.Core.Router

}

func main() {
	dir, err := filepath.Abs(filepath.Dir("../"))

	if err != nil {
		log.Fatal(err)
	}

	config := loadConfiguration("config.json")
	mongo.NewDBInstance(mongo.Config(config))

	// Create Mello and AuthBoss routers
	router := mux.NewRouter()
	abRouter := createAuthRouter()

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir+staticDirectory))))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir(dir+clientDirectory))))
	router.PathPrefix("/auth").Handler(http.StripPrefix("/auth", abRouter))

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
