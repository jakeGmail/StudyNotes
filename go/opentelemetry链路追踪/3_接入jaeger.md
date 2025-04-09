
**初始化**

```go
package component
import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/zap"
	"time"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)
/*
源码：https://github.com/open-telemetry/opentelemetry-go
获取sdk: go get go.opentelemetry.io/otel  go.opentelemetry.io/otel/metric go.opentelemetry.io/otel/sdk go.opentelemetry.io/otel/exporters/prometheus go.opentelemetry.io/otel/exporters/jaeger
文档: https://opentelemetry.io/docs/languages/go/getting-started/
*/

// 初始化otel并以jaeger为导出器
func InitOtel(logger *zap.Logger, serviceName string, exporterType, exporterEndPoint string) (func(), error) {
	otel.SetErrorHandler((eh)(func(err error) {
		logger.Error("otel internal error", zap.Error(err))
	}))

	// 创建资源
	var err error
	ctx := context.Background()
	// 支持跨服务追踪
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{})) //

	// 初始化导出器
	var traceExporter sdkTrace.SpanExporter
	switch exporterType {
	case "jaeger":
		traceExporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(exporterEndPoint)))
		if err != nil {
			logrus.Errorf("new jaeger failed:%s", err.Error)
			return func() {}, err
		}
		// 初始化privider
		tracerProvider := sdkTrace.NewTracerProvider(
			sdkTrace.WithBatcher(traceExporter),
			sdkTrace.WithSampler(sdkTrace.AlwaysSample()), // 数据采集器
			sdkTrace.WithResource(
				resource.NewWithAttributes(
					semconv.SchemaURL,
					semconv.ServiceName(serviceName),
					attribute.String("env", "dev"),
					attribute.Int64("ID", 2)),
			), // 资源
			sdkTrace.WithSpanProcessor(sdkTrace.NewBatchSpanProcessor(traceExporter)), // 导出器
		)

		// 设置全局的provider
		otel.SetTracerProvider(tracerProvider)
		return func() {
			c := context.Background()
			tracerProvider.Shutdown(c)
		}, nil
	default:
		return func() {}, errors.New("undefined exporter type " + exporterType)
	}
}
```

**使用**
```go
id, _ := uuid.NewUUID()
ctx = trace.ContextWithSpanContext(ctx, trace.NewSpanContext(
    trace.SpanContextConfig{
        TraceID: trace.TraceID(id[:]),
    })) // 手动设置trace的trace id
    
tracer := otel.GetTracerProvider().Tracer("GracefulServe")
ctx, span := tracer.Start(ctx, "init service", trace.WithSpanKind(trace.SpanKindProducer)) // 开始一个span，并设置span的角色
span.SetAttributes(attribute.Key("result").String("success")) // 设置span的属性
defer span.End() // 结束span
```

**跨度的类型**:
OpenTelemetry 中的不同类型的跨度（Span）用于描述分布式追踪系统中参与者的角色和行为。每种类型的跨度反映了在系统中不同组件的交互方式。以下是几种常见的跨度类型及其区别：

1. **内部跨度（Internal Span）**:
   - **定义**：内部跨度用于表示应用程序内部的逻辑操作，没有特定的客户端或服务器角色。
   - **使用场景**：当你希望跟踪应用程序中的特定步骤或过程，而这些操作不涉及跨进程的请求。
   - **示例**：处理某个算法的计算，或应用中的函数执行。

2. **服务器端跨度（Server Span）**:
   - **定义**：服务器端跨度用于标识服务器接收到的请求处理逻辑的跨度。
   - **使用场景**：当服务器接收到请求并开始处理时，例如处理 HTTP 请求或 RPC 调用。
   - **示例**：一个 Web 服务处理来自客户端的 GET 请求时的跨度。

3. **客户端跨度（Client Span）**:
   - **定义**：客户端跨度用于表示客户端发起请求所需要的操作的跨程跟踪。
   - **使用场景**：当应用作为客户端向外部服务发送请求时，例如发起 HTTP 请求或调用数据库查询。
   - **示例**：应用程序向 RESTful 服务发送请求的时间段。

4. **生产者跨度（Producer Span）**:
   - **定义**：生产者跨度用于标识产生消息的角色。
   - **使用场景**：当应用程序生成并发送消息到消息队列或发布到某个中间介质时。
   - **示例**：推送消息到 Kafka、RabbitMQ 等消息队列时。

5. **消费者跨度（Consumer Span）**:
   - **定义**：消费者跨度用于标识接收消息并处理的角色。
   - **使用场景**：当服务从消息队列中接收消息并开始处理时。
   - **示例**：从 Kafka 队列中接收到消息并进行处理活动。

### 总体区别

- **角色区分**：各类跨度的根本区别在于它们在分布式系统中所扮演的角色。内部跨度不涉及跨进程通信，而客户端和服务器端跨度分别代表请求的发起者和处理者。

- **通信语境**：生产者和消费者跨度一般涉及消息传递的上下文，表示非直接同步的消息传递模型。

- **使用场景**：客户端和服务器端跨度常用于同步请求-响应交互，而生产者和消费者跨度则用于异步消息模型。

通过使用这些不同类型的跨度，开发者可以更准确地构建和分析分布式系统中的调用关系和性能瓶颈。每种跨度类型在 Trace 中承担的功能确保了系统的可观测性，从而更容易进行监控和优化。