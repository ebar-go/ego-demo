package data

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id    int
	Email string
}

type UserClaims struct {
	jwt.StandardClaims
	User User `json:"user"`
}
