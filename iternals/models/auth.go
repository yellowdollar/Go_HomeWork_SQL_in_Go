package models

import "github.com/dgrijalva/jwt-go"

type AuthUser struct {
	Id       int
	Login    string
	Password string
}

type JWTcols struct {
	UserId    int    `json: userId`
	UserLogin string `json: userLogin`
	jwt.StandardClaims
}
