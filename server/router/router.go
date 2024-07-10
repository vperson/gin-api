package router

import (
	v1 "gin-api/server/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	actuator := r.Group("/actuator")
	{
		actuator.GET("/health", v1.Health)
		actuator.GET("/info", v1.Info)
	}
	return r
}
