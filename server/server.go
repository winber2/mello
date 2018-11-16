package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

const staticDirectory string = "/static/"
const clientDirectory string = "/client/public/"

func main() {
	dir, err := filepath.Abs(filepath.Dir("../"))

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir+staticDirectory))))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(dir+clientDirectory))))
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dir+clientDirectory+"index.html")
	})

	http.Handle("/", r)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
