package server

import (
	"context"
	"errors"
	"fmt"
	"gin-api/config"
	"gin-api/pkg/logger"
	"gin-api/server/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func StartMetrics(ctx context.Context, cfg *config.MetricsServerConfig) {
	logger.GetSugaredLogger().Infof("metrics server enable: %v", cfg.Enable)
	if !cfg.Enable {
		return
	}

	log := logger.GetSugaredLogger().With(zap.String("app", "metrics_server"))
	gin.SetMode(cfg.Mode)
	routersInit := router.InitMetricsRouter()
	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", cfg.Port)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	// 使用goroutine在后台启动服务
	go func() {
		log.Infof("start listening and serving HTTP in %s", endPoint)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("listening service failure: %v", err)
		}
	}()

	<-ctx.Done()
	// 优雅关闭服务器，不接受新的连接，等待已有连接处理完毕
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown failure: %v", err)
		return
	}
	log.Info("the server was gracefully shut down")
}
