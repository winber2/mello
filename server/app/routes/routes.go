package routes

import (
	"github.com/gorilla/mux"
)

func AppendAppRoutes(r *mux.Router) {
	subrouter := r.PathPrefix("/api").Subrouter()
	AppendUserRoutes(subrouter)

}
