package Middleware

import (
	"belajar-mvc-go/Models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type JwtToken struct {
	Token string `json:"token"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) string {
	var user Models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	return tokenString
}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) bool {
	token, _ := jwt.Parse(req.Header.Get("X-AUTH"), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user Models.User
		mapstructure.Decode(claims, &user)
		return true
		// json.NewEncoder(w).Encode(Models.Exception{Message: "Valid Token"})

	} else {
		return false
		// json.NewEncoder(w).Encode(Models.Exception{Message: "Invalid authorization token"})
	}
}
