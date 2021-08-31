package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {

	// jaeger配置
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, // 类型
			Param: 1,                       // 类型值
		}, // 采样配置
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true, // 发送到服务器时是否打印日志
			LocalAgentHostPort: "192.168.56.110:6831",
		}, // jaeger服务器配置
		ServiceName: "shop",
	}

	// 生成链路Tracer
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	// 单个span记录
	//span := tracer.StartSpan("go-grpc-web")
	//time.Sleep(time.Second * 1)
	//span.Finish()

	// 嵌套span记录
	parentSpan := tracer.StartSpan("main")
	span := tracer.StartSpan("funcA", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Second * 1)
	span.Finish()

	span2 := tracer.StartSpan("funcB", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(time.Second * 2)
	span2.Finish()

	parentSpan.Finish()

}
