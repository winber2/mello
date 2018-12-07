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

func AppendUserRoutes(r *mux.Router) {
	r.HandleFunc("/user/", GetUsers).Methods("GET")
	r.HandleFunc("/user/", PostUser).Methods("POST")
	r.HandleFunc("/user/{id}/", GetUser).Methods("GET")
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
	user.Save()
}
