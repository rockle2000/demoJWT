package tokenprovider

import "time"

type Provider interface {
	Generate(data TokenPayload, atExpiry, rtExpiry int) (*Token, error)
	Verify(token string) (*TokenPayload, error)
}

type Token struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Created      time.Time `json:"created"`
	AtExpiry     int       `json:"atExpiry"`
	RtExpiry     int       `json:"rtExpiry"`
}

type TokenPayload struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

//type AuthHeader struct {
//	IDToken string `header:"Authorization"`
//}
