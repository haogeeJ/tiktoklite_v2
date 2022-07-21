package middleware

import (
	"TikTokLite_v2/util/trace_id_log/loggers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jtolds/gls"
)

func TraceMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		traceCtx := &loggers.LoggerContext{
			traceId,
		}
		gls.EnsureGoroutineId(func(gid uint) {
			loggers.LoggerContextMgr.SetValues(gls.Values{gid: traceCtx}, func() {
				c.Next()
			})
		})
	}
}
