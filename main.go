package main

import (
	"DemoJWT/model"
	"DemoJWT/tokenprovider"
	"DemoJWT/tokenprovider/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func main() {
	var secretKey = os.Getenv("JWT_SECRET")
	//access token expire time - 15 minutes
	atExpireTime := 60 * 15
	//refresh token expire time - 7 days
	rtExpireTime := 60 * 60 * 24 * 7

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var loginUser model.User
		if err := c.ShouldBind(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		payload := tokenprovider.TokenPayload{
			UserId:   loginUser.Id,
			UserName: loginUser.Username,
		}
		provider := jwt.NewJWTProvider(secretKey)
		token, err := provider.Generate(payload, atExpireTime, rtExpireTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, token)
	})

	router.GET("/verify", func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenHeader := strings.Split(header, " ")
		if tokenHeader[0] != "Bearer" || len(tokenHeader) < 2 || strings.TrimSpace(tokenHeader[1]) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Authorization",
			})
			return
		}
		provider := jwt.NewJWTProvider(secretKey)
		validate, err := provider.Verify(tokenHeader[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid token",
			})
			return
		}
		c.JSON(http.StatusOK, validate)
	})

	router.Run(":8080")
}
