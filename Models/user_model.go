package Models

import (
	"belajar-mvc-go/Helper"
	_ "belajar-mvc-go/Helper"
	"fmt"
	"log"
)

type Exception struct {
	Message string `json:"message"`
}
type User struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Foto      string `json:"foto"`
}

func GetAllUser() []User {
	Setup()
	rows, err := Helper.Db.Query("SELECT id , firstname , lastname, username FROM user")
	if err != nil {
		log.Fatal("Error Connection " + err.Error())
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
func GetSpecificUser(Username string, Password string) User {
	Setup()
	var user User

	err := Helper.Db.QueryRow("SELECT id , firstname , lastname, username FROM user where username = ? AND password = ?", Username, Password).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username)

	if err != nil {
		log.Fatal(err)
		return user
	}
	return user
}
func InsertUser(DataUser User) {
	Setup()
	fmt.Println(DataUser)
	insert, err := Helper.Db.Query("INSERT INTO user (firstname , lastname , username , password , foto) VALUES ( ? , ? , ? , ? , ? )", DataUser.Firstname, DataUser.Lastname, DataUser.Username, DataUser.Password, DataUser.Foto)
	if err != nil {
		log.Fatal("Error Connection " + err.Error())
	}
	defer insert.Close()

}
func DeleteRowUser(id string) {
	Setup()
	delete, err := Helper.Db.Query("delete from user where id = ?", id)
	if err != nil {
		log.Fatal("Error Connection " + err.Error())
	}
	defer delete.Close()
}
func UpdateUser(DataUser User) {
	update, err := Helper.Db.Query("UPDATE user SET firstname = ? , lastname = ? , username = ? , password = ? , foto = ? WHERE id = ?", DataUser.Firstname, DataUser.Lastname, DataUser.Username, DataUser.Password, DataUser.Foto, DataUser.Id)
	if err != nil {
		log.Fatal("Error Connection " + err.Error())
	}
	defer update.Close()
}
