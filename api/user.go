package api

import (
	"DemoJWT/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "add user success",
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := service.UserRes{
			UserName: "user",
			Role:     "user",
		}
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
