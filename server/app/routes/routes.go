package routes

import (
	"github.com/gorilla/mux"
)

//
// type RESTfulRoute struct {
// 	Name
// 	Method string
//
// }
//
// type Routes []Route

func AppendAppRoutes(r *mux.Router) {
	subrouter := r.PathPrefix("/api").Subrouter()
	AppendUserRoutes(subrouter)

}
