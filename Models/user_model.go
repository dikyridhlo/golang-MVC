package Models

import (
	"belajar-mvc-go/Helper"
	_ "belajar-mvc-go/Helper"
	"fmt"
	"log"
)

type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func GetAllUser() []User {
	Setup()
	rows, err := Helper.Db.Query("SELECT id , firstname , lastname, email FROM user")
	if err != nil {
		log.Fatal("Error Connection")
	}
	got := []User{}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			log.Fatal("Scan: %v", err)
		}
		got = append(got, user)
	}
	fmt.Println(got)
	return got
}
