package tracer

import (
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go/config"
)

//Init initialize open tracer and set it to global tracer
func Init() io.Closer {
	cfg, err := jaeger.FromEnv()
	if err != nil {
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return nil
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return nil
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}
