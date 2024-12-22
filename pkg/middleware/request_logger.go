package middleware

import (
	"time"

	"github.com/NeoJay0705/go-server-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate or extract Trace ID
		traceID := c.Request.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		start := time.Now()

		c.Next() // Process request

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Log the request details
		logger.Logger.Info("API Request",
			zap.String("trace_id", traceID),
			zap.String("user_id", "user_id"),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status_code", statusCode),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)

		logger.Logger.Error("Failure to create")

		// Add Trace ID to response headers
		c.Writer.Header().Set("X-Trace-ID", traceID)
	}
}
