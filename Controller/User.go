package Controller

import (
	"belajar-mvc-go/Models"
	"encoding/json"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listuser := Models.GetAllUser()
	jsondata, err := json.Marshal(listuser)
	if err != nil {
		log.Fatal("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}
