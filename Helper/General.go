package Helper

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var Db *sql.DB
var Something string = "global variable"

func GetEnv(data string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(data)
}
