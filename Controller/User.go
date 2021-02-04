package Controller

import (
	"belajar-mvc-go/Middleware"
	"belajar-mvc-go/Models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listuser := Models.GetAllUser()
	jsondata, err := json.Marshal(listuser)
	if err != nil {
		log.Fatal("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}
func UploadFile(w http.ResponseWriter, r *http.Request) string {
	//VALIDATION FILE
	err := r.ParseMultipartForm(5 * 1024 * 1024 * 1024)
	if err != nil {
		fmt.Println("Error ParseMultipartForm: ", err) // here it fails !!! with: "multipart: NextPart: EOF"
		return ""
	}

	//
	file, handler, err := r.FormFile("foto")
	if err != nil {
		return ""
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		fmt.Println("Error dst: ", err) // here it fails !!! with: "multipart: NextPart: EOF"
		return ""
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println("Error copy: ", err) // here it fails !!! with: "multipart: NextPart: EOF"
		return ""
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	return handler.Filename
}
func SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	if Middleware.ProtectedEndpoint(w, r) {
		filename := UploadFile(w, r)
		var user Models.User
		user.Firstname = r.FormValue("firstname")
		user.Lastname = r.FormValue("lastname")
		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.Foto = filename

		Models.InsertUser(user)

		jsondata, err := json.Marshal(user)
		if err != nil {
			log.Fatal("Error %v", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsondata)
	} else {
		json.NewEncoder(w).Encode(Models.Exception{Message: "Invalid authorization token"})
	}

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	vars := mux.Vars(r)
	id := vars["id"]
	if Middleware.ProtectedEndpoint(w, r) {
		filename := UploadFile(w, r)
		var user Models.User
		user.Firstname = r.FormValue("firstname")
		user.Lastname = r.FormValue("lastname")
		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.Foto = filename
		user.Id = id

		fmt.Println(user)
		Models.UpdateUser(user)

		jsondata, err := json.Marshal(user)
		if err != nil {
			log.Fatal("Error %v", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsondata)
	} else {
		json.NewEncoder(w).Encode(Models.Exception{Message: "Invalid authorization token"})
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if Middleware.ProtectedEndpoint(w, r) {
		vars := mux.Vars(r)
		id := vars["id"]
		Models.DeleteRowUser(id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Models.Exception{Message: "Berhasil menghapus data"})
	} else {
		json.NewEncoder(w).Encode(Models.Exception{Message: "Invalid authorization token"})
	}

}
