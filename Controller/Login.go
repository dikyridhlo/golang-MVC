package Controller

import (
	"belajar-mvc-go/Middleware"
	"belajar-mvc-go/Models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user Models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")
	if user.Username == "" {
		json.NewEncoder(w).Encode(Models.Exception{Message: "Username tidak boleh kosong"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		json.NewEncoder(w).Encode(Models.Exception{Message: "Password tidak boleh kosong"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)

	listuser := Models.GetSpecificUser(user.Username, user.Password)
	if user.Username != "" {
		listuser.Token = Middleware.CreateTokenEndpoint(w, r)
	}

	jsondata, err := json.Marshal(listuser)
	if err != nil {
		log.Fatal("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}
