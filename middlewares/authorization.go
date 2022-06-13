package middlewares

import (
	"DemoJWT/tokenprovider/jwt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Authorization(cb *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = os.Getenv("JWT_SECRET")
		header := c.GetHeader("Authorization")
		// eg: /admin
		obj := c.Request.URL.RequestURI()
		// eg: GET
		action := c.Request.Method
		provider := jwt.NewJWTProvider(secretKey)
		token, err := provider.ExtractToken(header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}
		payload, err := provider.Verify(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Invalid token",
			})
			return
		}
		if res, err := cb.Enforce(payload.UserName, obj, action); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		} else if !res {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you don't have permission to access this resource",
			})
			return
		}
		c.Next()
	}
}
