package redis_tracing

import (
	"context"
	"fmt"
	"github.com/gistao/RedisGo-Async/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	traceLog "github.com/opentracing/opentracing-go/log"
)

type syncDo func(commandName string, args ...interface{}) (reply interface{}, err error)
type asyncDo func(commandName string, args ...interface{}) (reply redis.AsyncRet, err error)

func SyncDoAndTracing(ctx context.Context, f syncDo, commandName string, args ...interface{}) (reply interface{}, err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "redis")
	defer span.Finish()
	ext.DBType.Set(span, "redis")
	reply, err = f(commandName, args...)
	if err != nil {
		span.LogFields(traceLog.Error(err))
		return
	}
	span.LogFields(traceLog.String("redis-cmd", fmt.Sprintf("%s %v", commandName, args)))
	return
}
func AsyncDoAndTracing(ctx context.Context, f asyncDo, commandName string, args ...interface{}) (reply redis.AsyncRet, err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "redis")
	defer span.Finish()
	ext.DBType.Set(span, "redis")
	reply, err = f(commandName, args...)
	if err != nil {
		span.LogFields(traceLog.Error(err))
		return
	}
	span.LogFields(traceLog.String("redis-cmd", fmt.Sprintf("%s %v", commandName, args)))
	return
}
