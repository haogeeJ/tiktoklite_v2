package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer 创建一个jaeger trace
func NewTracer(serverName, address string) (opentracing.Tracer, io.Closer, error) {

	// 生成jaegercfg
	cfg := jaegercfg.Configuration{
		ServiceName: serverName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	transport, err := jaeger.NewUDPTransport(address, 0)
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(transport)

	options := jaegercfg.Reporter(reporter)
	return cfg.NewTracer(options)
}
