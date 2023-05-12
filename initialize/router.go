package initialize

import (
	"admin_api/middlewares"
	"admin_api/router"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)

	Router := gin.New()
	Router.Use(middlewares.CORS, gin.Logger(), gin.Recovery())
	ginpprof.Wrap(Router)

	router.HealthCheck(Router)
	router.Account(Router)
	router.MarCampaign(Router)
	return Router
}
