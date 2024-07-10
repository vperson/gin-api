package main

import (
	"context"
	"flag"
	"fmt"
	"gin-api/config"
	"gin-api/pkg/logger"
	"gin-api/repository/store"
	"gin-api/server"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var (
		err        error
		configPath string
	)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	flag.StringVar(&configPath, "config", "config.yaml", "config file")
	flag.StringVar(&configPath, "c", "config.yaml", "config file (shorthand)")
	help := flag.Bool("help", false, "Display this help message")
	flag.BoolVar(help, "H", false, "Display this help message (shorthand)")

	// 解析命令行参数
	flag.Parse()

	if *help {
		fmt.Println("Usage of this program:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// 初始化配置文件
	log.Println("config file:", configPath)
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	// 初始化日志
	logger.InitLogger(&cfg.Log)
	logger.GetStructuredLogger().Info("logger init is ok...")

	// 初始化数据库
	err = store.New(&cfg.DB)
	if err != nil {
		logger.GetStructuredLogger().Fatal("store init is failed", zap.Error(err))
	}

	// 运行metrics服务
	wg.Add(1)
	go func(ctx context.Context, metricsServerConfig *config.MetricsServerConfig) {
		defer wg.Done()
		server.StartMetrics(ctx, metricsServerConfig)
	}(ctx, &cfg.MetricsServer)

	// 运行http服务
	wg.Add(1)
	go func(ctx context.Context, serverConfig *config.ServerConfig) {
		defer wg.Done()

		server.Start(ctx, serverConfig)
	}(ctx, &cfg.Server)

	// 捕捉到SIGTERM和SIGINT信号时调用cancel
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

		<-signals // 等待信号
		logger.GetStructuredLogger().Info("Received termination signal, exiting gracefully...")
		cancel() // 通知退出
	}()

	// 启动一个需要清理的goroutine
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		<-ctx.Done() // 等待退出的通知
		logger.GetStructuredLogger().Info("Performing cleanup tasks...")
		// 执行必要的清理操作
		// 然后优雅地停止服务或者关闭goroutine等
	}(ctx)

	logger.GetStructuredLogger().Info("The program is running... Press Ctrl+C to send the termination signal")
	wg.Wait()
	logger.GetStructuredLogger().Info("Exit the program...")
	logger.GetStructuredLogger().Info("Program exited successfully.")
	logger.GetStructuredLogger().Info("service has exited")
}
