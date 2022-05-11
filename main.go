package main

import (
	"DemoJWT/model"
	"fmt"
)

func main() {
	//secretKey := os.Getenv("JWT_SECRET")
	user := model.User{
		Id:       1,
		Username: "abc",
		Password: "123",
	}
	expireTime := 60 * 60
	token, err := user.Generate(expireTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token.Token)

	//validate, err := user.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjEsInVzZXJfbmFtZSI6ImFiYyJ9LCJleHAiOjE2NTIxNTU3NjUsImlhdCI6MTY1MjI0MjE2NX0.O3JXte9rpBPsdSGuIWL6GpJ2SDIDfGMxlDrVhUEcnyc")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("%+v", validate)
}
