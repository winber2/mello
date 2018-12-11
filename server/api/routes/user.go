package routes

import (
	"encoding/json"
	"fmt"
	"mello/server/mongo"
	"mello/server/mongo/models"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func UserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", GetUsers).Methods("GET")
	r.HandleFunc("/", PostUser).Methods("POST")
	r.HandleFunc("/{id}/", GetUser).Methods("GET")
	return r
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var result []bson.M
	query := mongo.Database.C("users").Find(bson.M{})
	query.All(&result)
	json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var result models.User
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	objectId := bson.ObjectIdHex(vars["id"])
	err := mongo.Database.C("users").Find(bson.M{"_id": objectId}).One(&result)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(result)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user)
	err := user.Save()

	fmt.Println(err)
}
