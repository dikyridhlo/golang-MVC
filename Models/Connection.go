package Models

import (
	"belajar-mvc-go/Helper"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	var err error
	Helper.Db, err = sql.Open("mysql", Helper.GetEnv("db-user")+":"+Helper.GetEnv("db-password")+"@/"+Helper.GetEnv("db-database"))
	if err != nil {
		log.Fatal(err)
	}
	return err
	// Other setup-related activities
}
