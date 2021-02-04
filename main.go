package main

import (
	"belajar-mvc-go/Routes"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Application has been started on : 9000")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("site")
	log.Info(key)
	//creating instance of mux
	routes := Routes.GetRoutes()
	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
