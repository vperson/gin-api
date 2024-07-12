package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"code", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func CustomPrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		statusCode := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.WithLabelValues(statusCode, c.Request.Method).Inc()
	}
}
