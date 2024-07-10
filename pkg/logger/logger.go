package logger

import (
	"gin-api/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var sugarLogger *zap.SugaredLogger
var structLogger *zap.Logger
var once sync.Once

// 定义 Prometheus 指标
var (
	logLevelCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "log_level_total",
			Help: "Total number of logs by level and logger type",
		},
		[]string{"level", "logger_type"},
	)
)

// InitLogger 初始化日志
func InitLogger(cfg *config.LogConfig) {
	once.Do(func() {
		var zapConfig zap.Config
		if cfg.Format == "json" {
			zapConfig = zap.NewProductionConfig()
		} else {
			zapConfig = zap.NewDevelopmentConfig()
			zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		switch cfg.Level {
		case "debug":
			zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		case "info":
			zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case "warn":
			zapConfig.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		case "error":
			zapConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		default:
			zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		}

		var err error
		baseLogger, err := zapConfig.Build(zap.Hooks(prometheusHook("structured")))
		if err != nil {
			panic(err)
		}
		structLogger = baseLogger
		sugarLogger = baseLogger.WithOptions(zap.Hooks(prometheusHook("sugared"))).Sugar()
	})
}

// prometheusHook 是一个 zap 的钩子函数工厂，用于记录 Prometheus 指标
func prometheusHook(loggerType string) func(zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		logLevelCounter.WithLabelValues(entry.Level.String(), loggerType).Inc()
		return nil
	}
}

// GetSugaredLogger 获取全局 SugaredLogger 实例
func GetSugaredLogger() *zap.SugaredLogger {
	if sugarLogger == nil {
		panic("Logger is not initialized")
	}
	return sugarLogger
}

// GetStructuredLogger 获取全局结构化 Logger 实例
func GetStructuredLogger() *zap.Logger {
	if structLogger == nil {
		panic("Logger is not initialized")
	}
	return structLogger
}

// Sync flushes any buffered log entries
func Sync() {
	if structLogger != nil {
		_ = structLogger.Sync()
	}
}
