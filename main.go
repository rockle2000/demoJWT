package main

import (
	"DemoJWT/api"
	"DemoJWT/middlewares"
	"DemoJWT/model"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	listUser := model.User{}.InitData()
	cb, err := casbin.NewEnforcer("rbac_model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("unable to create Casbin enforcer: %v", err)
	}
	allSubjects := cb.GetAllSubjects()
	fmt.Println("all subject: ", allSubjects)
	allNamedRoles := cb.GetAllNamedRoles("g")
	fmt.Println("all named: ", allNamedRoles)
	policy := cb.GetPolicy()
	fmt.Println("all policy: ", policy)
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
	//router.GET("/verify", func(c *gin.Context) {
	//	header := c.GetHeader("Authorization")
	//	tokenHeader := strings.Split(header, " ")
	//	if tokenHeader[0] != "Bearer" || len(tokenHeader) < 2 || strings.TrimSpace(tokenHeader[1]) == "" {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": "Invalid Authorization",
	//		})
	//		return
	//	}
	//	provider := jwt.NewJWTProvider(secretKey)
	//	validate, err := provider.Verify(tokenHeader[1])
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": "Invalid token",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, validate)
	//})

	router.Run(":8080")
}
