package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Mail     string `json:"mail"`
}

type Card struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type Sequence struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	Card1ID   int `json:"card_1_id"`
	Card2ID   int `json:"card_2_id"`
	CreatedAt int `json:"created_at"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
