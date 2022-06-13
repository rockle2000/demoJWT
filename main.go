package main

import (
	"DemoJWT/api"
	"DemoJWT/middlewares"
	"DemoJWT/model"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Init user test data
	listUser := model.User{}.InitData()
	cb, err := casbin.NewEnforcer("rbac_model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("unable to create Casbin enforcer: %v", err)
	}
	router := gin.Default()
	router.POST("/login", api.Login(listUser))
	admin := router.Group("/admin").Use(middlewares.Authorization(cb))
	{
		admin.POST("", api.AddAdmin())
		admin.GET("", api.GetAdmin())
	}
	user := router.Group("/users").Use(middlewares.Authorization(cb))
	{
		user.GET("", api.GetUser())
		user.POST("", api.AddUser())
	}

	router.Run(":8080")
}
