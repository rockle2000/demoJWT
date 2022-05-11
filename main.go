package main

import (
	"DemoJWT/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//secretKey := os.Getenv("JWT_SECRET")
	//user := model.User{
	//	Id:       1,
	//	Username: "abc",
	//	Password: "123",
	//}
	//expireTime := 60 * 60
	//token, err := user.Generate(expireTime)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(token.Token)
	//
	////validate, err := user.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjEsInVzZXJfbmFtZSI6ImFiYyJ9LCJleHAiOjE2NTIxNTU3NjUsImlhdCI6MTY1MjI0MjE2NX0.O3JXte9rpBPsdSGuIWL6GpJ2SDIDfGMxlDrVhUEcnyc")
	////if err != nil {
	////	fmt.Println(err)
	////}
	////fmt.Printf("%+v", validate)

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var loginUser model.User
		if err := c.ShouldBind(&loginUser); err != nil {
			fmt.Println(err)
		}
		expireTime := 60 * 60
		token, err := loginUser.Generate(expireTime)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, token)
	})

	router.GET("/verify", func(c *gin.Context) {
		h := model.AuthHeader{}
		var user model.User
		if err := c.ShouldBindHeader(&h); err != nil {
			fmt.Println(err)
		}
		tokenHeader := strings.Split(h.IDToken, " ")
		if tokenHeader[0] != "Bearer" || len(tokenHeader) < 2 || strings.TrimSpace(tokenHeader[1]) == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization",
			})
			c.Abort()
			return
		}
		validate, err := user.Validate(tokenHeader[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, validate)
	})

	router.Run(":8080")
}
