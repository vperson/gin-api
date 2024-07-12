package router

import (
	"gin-api/config"
	"gin-api/server/middleware"
	v1 "gin-api/server/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	metricsOnServer(r)
	actuator := r.Group("/actuator")
	{
		actuator.GET("/health", v1.Health)
		actuator.GET("/info", v1.Info)
	}

	anonymous := r.Group("/anonymous")
	{
		anonymous.GET("/info", v1.Info)
		// 在下面增加其他匿名的接口
	}

	// 需要认证的接口
	auth := r.Group("/auth", middleware.JWTAuthMiddleware())
	{
		auth.GET("/info", v1.Info)
		// 在下面增加其他需要认证的接口
	}
	return r
}

func metricsOnServer(r *gin.Engine) {
	if !config.Get().Server.MetricsEnable {
		return
	}
	r.Use(middleware.CustomPrometheusMiddleware())
	//r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
