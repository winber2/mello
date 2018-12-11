package api

import (
	"mello/server/api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/user").Handler(http.StripPrefix("/user", routes.UserRoutes()))
	return router
}
