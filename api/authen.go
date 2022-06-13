package api

import (
	"DemoJWT/model"
	"DemoJWT/service"
	"DemoJWT/tokenprovider"
	"DemoJWT/tokenprovider/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Login(listUser model.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = os.Getenv("JWT_SECRET")
		//access token expire time - 1 day
		atExpireTime := 60 * 60 * 24
		//refresh token expire time - 7 days
		rtExpireTime := 60 * 60 * 24 * 7
		var loginReq service.LoginReq
		if err := c.ShouldBind(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		modelUser := model.User{
			Username: loginReq.UserName,
			Password: loginReq.Password,
		}
		user := listUser.FindUser(modelUser)
		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid username or password",
			})
			return
		}
		payload := tokenprovider.TokenPayload{
			UserId:   user.Id,
			UserName: user.Username,
			Role:     user.Role,
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
	}
}
