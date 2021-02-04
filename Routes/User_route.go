package Routes

import (
	"belajar-mvc-go/Controller"
	"belajar-mvc-go/Helper"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)
	routes = routes.PathPrefix(Helper.GetEnv("url-api")).Subrouter()
	routes.HandleFunc("/user", Controller.Hello).Methods("GET")

	return routes
}
