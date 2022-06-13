package jwt

import (
	"DemoJWT/tokenprovider"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type jwtProvider struct {
	secret string
}
type jwtConfig struct {
}

func NewJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type claims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, atExpiry, rtExpiry int) (*tokenprovider.Token, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(atExpiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	accessToken, err := at.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(rtExpiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	refreshToken, err := rt.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AtExpiry:     atExpiry,
		RtExpiry:     rtExpiry,
		Created:      time.Now(),
	}, nil
}

func (j *jwtProvider) Verify(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	// validate token
	if !res.Valid {
		return nil, errors.New("invalid Token")
	}

	claims, ok := res.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid Token")
	}

	return &claims.Payload, nil
}

func (j *jwtProvider) ExtractToken(header string) (string, error) {
	tokenHeader := strings.Split(header, " ")
	if tokenHeader[0] != "Bearer" || len(tokenHeader) < 2 || strings.TrimSpace(tokenHeader[1]) == "" {
		return "", errors.New("invalid token")
	}
	return tokenHeader[1], nil
}
