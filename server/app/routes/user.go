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
	r.HandleFunc("/user", GetUsers).Methods("GET")
	r.HandleFunc("/user", PostUser).Methods("POST")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var result []bson.M
	query := mongo.Database.C("users").Find(bson.M{})
	query.All(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)
	req.Save()
}
