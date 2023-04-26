package router

import (
	"admin_api/internal/api/account"
	"net/http"

	"admin_api/middlewares"

	"github.com/gin-gonic/gin"
)

var JWT = middlewares.JWTAuth

func HealthCheck(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "pong"})
	})
}

func Account(r *gin.Engine) {
	//test
	g := r.Group("/account")

	var system account.Account
	g.POST("/signup", system.Signup)
	g.POST("/password_login", system.PasswordLogin)
	g.GET("/user_info", JWT(), system.GetUserDetail)
	g.PUT("/password", JWT(), system.UpdatePassword)
}
