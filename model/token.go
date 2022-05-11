package model

import "time"

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}
