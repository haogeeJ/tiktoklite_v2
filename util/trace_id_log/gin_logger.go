package trace_id_log

import (
	"TikTokLite_v2/util/trace_id_log/loggers"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
		param.Latency,
		param.ClientIP,
		param.MethodColor(), param.Method, param.ResetColor(),
		param.Path,
		param.ErrorMessage,
	)
}
var fileLogFormatter = func(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s  %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

func NewGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(now)
		param.StatusCode = c.Writer.Status()
		param.ClientIP = c.ClientIP()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path
		param.Method = c.Request.Method
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		loggers.Logger.Infof(defaultLogFormatter(param))
		//logFormatterStr := strings.ReplaceAll(fileLogFormatter(param), "\"", "'")
		//loggers.Logger.Infof(logFormatterStr) //输出到文件的格式
	}
}
