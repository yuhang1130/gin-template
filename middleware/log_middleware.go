package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()

		// 执行时间
		costTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		entry := logger.WithFields(log.Fields{
			"status_code": statusCode,
			"cost_time":   costTime,
			"client_ip":   clientIP,
			"req_method":  reqMethod,
			"req_uri":     reqUri,
		})

		// 根据状态码输出日志
		if len(c.Errors) > 0 {
			// 请求有错误发生
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			// 请求正常处理完成
			entry.Info()
		}
	}
}
