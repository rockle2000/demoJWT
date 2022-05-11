package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var secretKey = os.Getenv("JWT_SECRET")

type User struct {
	Id       int
	Username string
	Password string
}

type claims struct {
	Payload TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (u User) Generate(expiry int) (*Token, error) {
	data := TokenPayload{
		UserId:   u.Id,
		UserName: u.Username,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			//ExpiresAt: time.Now().Local().AddDate(0, 0, -1).Unix(),
			IssuedAt: time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (u User) Validate(myToken string) (*TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// validate the token
	if !res.Valid {
		return nil, errors.New("invalid Token")
	}

	claims, ok := res.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid Token 3")
	}

	return &claims.Payload, nil
}
