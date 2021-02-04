package Routes

import (
	"belajar-mvc-go/Controller"
	"belajar-mvc-go/Helper"

	_ "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {

	routes := mux.NewRouter().StrictSlash(false)

	routes = routes.PathPrefix(Helper.GetEnv("url-api")).Subrouter()

	routes.HandleFunc("/user", Controller.GetUser).Methods("GET")
	routes.HandleFunc("/login", Controller.Login).Methods("POST")
	routes.HandleFunc("/user/insert", Controller.SaveUser).Methods("POST")
	routes.HandleFunc("/user/delete/{id}", Controller.DeleteUser).Methods("DELETE")
	routes.HandleFunc("/user/update/{id}", Controller.UpdateUser).Methods("PUT")
	return routes
}
