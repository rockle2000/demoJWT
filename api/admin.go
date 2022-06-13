package api

import (
	"DemoJWT/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "add admin success",
		})
	}
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		admin := service.UserRes{
			UserName: "admin",
			Role:     "admin",
		}
		c.JSON(http.StatusOK, gin.H{
			"data": admin,
		})
	}
}
