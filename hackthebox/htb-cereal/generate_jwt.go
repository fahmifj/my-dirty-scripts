package main

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type CurrentUser struct {
	UserId   string `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}

func CreateToken(secret string) string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().AddDate(0, 0, 7).UTC().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return t
}

func main() {

	cu := CurrentUser{
		UserId:   "1",
		Username: "iamf",
		Token:    CreateToken("secretlhfIH&FY*#oysuflkhskjfhefesf"),
	}

	currentUser, _ := json.Marshal(cu)
	fmt.Printf("%s", currentUser)
}
