package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Users []User
