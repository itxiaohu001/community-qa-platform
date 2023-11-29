package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 是一个简单的日志记录中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求处理前的逻辑
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 请求处理后的逻辑
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		log.Printf("请求 %s 耗时 %v", c.Request.URL.Path, latency)
	}
}
