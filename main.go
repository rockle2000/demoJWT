package main

import (
	"DemoJWT/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var loginUser model.User
		if err := c.ShouldBind(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		expireTime := 60 * 60
		token, err := loginUser.Generate(expireTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, token)
	})

	router.GET("/verify", func(c *gin.Context) {
		//h := model.AuthHeader{}
		var user model.User
		header := c.GetHeader("Authorization")
		//if err := c.ShouldBindHeader(&h); err != nil {
		//	fmt.Println(err)
		//}
		tokenHeader := strings.Split(header, " ")
		if tokenHeader[0] != "Bearer" || len(tokenHeader) < 2 || strings.TrimSpace(tokenHeader[1]) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Authorization",
			})
			return
		}
		validate, err := user.Validate(tokenHeader[1])
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
