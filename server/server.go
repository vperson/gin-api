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

func Start(ctx context.Context, cfg *config.ServerConfig) {
	logger.GetSugaredLogger().Infof("server enable: %v", cfg.Enable)
	log := logger.GetSugaredLogger().With(zap.String("app", "server"))

	gin.SetMode(cfg.Mode)
	routersInit := router.InitRouter()
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
		log.Infof("Start listening and serving HTTP in %s", endPoint)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("Listening service failure: %v", err)
		}
	}()

	<-ctx.Done()
	// 优雅关闭服务器，不接受新的连接，等待已有连接处理完毕
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("Server shutdown failure: %v", err)
		return
	}
	log.Info("The server was gracefully shut down")
}
