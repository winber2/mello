package routes

import (
	"encoding/json"
	"fmt"
	"mello/server/mongo"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func AppendUserRoutes(r *mux.Router) {
	r.HandleFunc("/user", GetUsers)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var result []bson.M
	query := mongo.Database.C("test").Find(bson.M{})
	query.All(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func PostUser(w http.ResponseWriter, r *http.Request) {

}
