package router

import (
	"admin_api/internal/api/account"
	"admin_api/internal/api/financial"
	"admin_api/internal/api/marketing"
	"admin_api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maocatooo/thin/gin_handler"
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

	var financial financial.Financial
	g.PUT("/amount_create", JWT(), financial.CreateUserAmount)
	g.GET("/user_amount", JWT(), financial.GetUserAmount)
	//	g.GET("/user_trade", JWT(), financial.DescribeUserTrade)
}

func MarCampaign(r *gin.Engine) {
	g := r.Group("/mar_campaigns", JWT())
	var mcApi marketing.MarketingCampaignApi
	g.POST("", gin_handler.JSON(mcApi.Create))
	g.GET("", gin_handler.Query(mcApi.List))

	g.GET("/:id", gin_handler.Query(mcApi.Detail))
	g.PUT("/:id/state", gin_handler.JSON(mcApi.UpdateState))

	var sendRecord marketing.SendRecord
	g.POST("/send_record/upload", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "upload",
		})
	})
	g.GET("/record", sendRecord.DescribeSendRecord)
}
