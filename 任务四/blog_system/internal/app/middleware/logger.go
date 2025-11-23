// internal/app/middleware/logger.go
package middleware

import (
	"time"
	"blog_system/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 日志中间件
func GinLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        // 处理请求
        c.Next()
        
        // 记录日志
        latency := time.Since(start)
        
        logger.Logger.Info("HTTP请求",
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.String("path", c.Request.URL.Path),
            zap.String("query", c.Request.URL.RawQuery),
            zap.String("ip", c.ClientIP()),
            zap.String("user-agent", c.Request.UserAgent()),
            zap.Duration("latency", latency),
            zap.String("request_id", c.GetString("request_id")),
        )
    }
}

// 错误日志中间件
func ErrorLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        // 检查是否有错误
        if len(c.Errors) > 0 {
            for _, ginErr := range c.Errors {
                logger.Logger.Error("请求处理错误",
                    zap.String("path", c.Request.URL.Path),
                    zap.String("method", c.Request.Method),
                    zap.Any("error", ginErr.Error()),
                )
            }
        }
    }
}