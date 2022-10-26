package routers

import (
	"github.com/gin-gonic/gin"
	"short-link-go/middleware"
	v1 "short-link-go/routers/api/v1"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.GinLogger())
	engine.Use(middleware.GinRecovery(true))
	engine.Use(gin.Recovery())
	routerGroup := engine.Group("/v1")
	{
		routerGroup.POST("/getShort", v1.GetShort)

	}
	engine.GET("/:path", v1.Short)
	engine.GET("/default", v1.DefaultUrl)
	return engine
}
