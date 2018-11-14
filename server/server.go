package main

import (
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/inconshreveable/go-update"
)

const staticDirectory string = "/static/"
const clientDirectory string = "/client/"

func updateWithPatch(patch io.Reader) error {
	err := update.Apply(patch, update.Options{
		Patcher: update.NewBSDiffPatcher(),
	})
	if err != nil {
		// error handling
	}
	return err
}

func main() {
	dir, err := filepath.Abs(filepath.Dir("../"))

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir+staticDirectory))))
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dir+clientDirectory+"index.html")
	})

	http.Handle("/", r)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

	// updateErr := updateWithPatch(io.Reader)
	// if updateErr != nil {
	// 	panic(err)
	// }
}
