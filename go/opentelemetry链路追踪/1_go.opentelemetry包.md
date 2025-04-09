[toc]

# 1 opentelemetry包介绍
`go.opentelemetry.io` 是 Go 语言中的一个模块，用于实现 OpenTelemetry 功能。OpenTelemetry 是一个开放源代码的分布式跟踪和度量标准，旨在为分布式系统提供可观测性解决方案。
是分布式链路追踪技术的标准，可以指定数据的格式、和系统指标
`go.opentelemetry.io` 模块包括了 OpenTelemetry 在 Go 语言中的实现，使开发者能够在 Go 应用程序中进行分布式跟踪和收集度量数据。

**官网**：

该模块通常包含以下子包：

1. **`go.opentelemetry.io/otel`**: 这是 OpenTelemetry 的核心包，包含了 OpenTelemetry API 和 SDK。
2. **`go.opentelemetry.io/otel/trace`**: 提供了分布式跟踪相关的接口和实现。
3. **`go.opentelemetry.io/otel/metric`**: 提供了一些用于度量数据收集的API和实现，这个模块主要用于支持应用程序对业务指标和性能指标进行采集、记录和导出
4. **`go.opentelemetry.io/otel/label`**: 用于定义标签，通常与跟踪和度量一起使用。
5. **`go.opentelemetry.io/otel/exporters`**: 包含了各种导出器，将收集到的数据导出到不同的后端系统。
6. **`go.opentelemetry.io/otel/propagation`**: 主要作用是处理上下文的传播。它提供了一组工具和接口，用于将上下文信息（如 Trace 信息）在分布式系统中的不同服务和组件之间传递。上下文传播器（Propagator）负责将上下文信息从一个请求中提取出来并注入到另一个请求中。这对于保持分布式跟踪链路的完整性至关重要。
具体来说，propagation 包提供了以下功能：

- **提取（Extract）**：从传入请求的载体中提取 Trace 上下文信息。
- **注入（Inject）**：将 Trace 上下文信息注入到传出请求的载体中。
- **上下文传播器（Propagator）**：定义如何在不同介质（如 HTTP 头、消息队列等）之间传递上下文信息。
7. **`go.opentelemetry.io/otel/bridge`**: 提供OpenCensus与OPentracing垫片，用于兼容OpenCensus与Opentracing，以便迁移到opentelemetry
8. **`go.opentelemetry.io/otel/baggage`**: 
   - 在分布式跟踪过程中传递和管理上下文数据
   - 并且可以跨进程，跨语言地传递它们
9. **`schema`**：提供了用于定义和解析OpenTelemetry数据格式地API和实现。该模块主要用于标准化和规范化OpenTelemetry数据格式，以便不同组件、系统和服务之间可以共享、交换和分析
10. **`sdk`**：提供了用于实现OpenTelemetry标准的软件开发工具报（SDK）。该模块主要用于收集、处理和导出OpenTelemetry数据，并且提供一些与性能相关的功能。 
11. **`codes`**：otlp状态码定义
12. **`internal`**: 提供了一些内部实现和工具，主要用于支持OpenTelemetry Go SDK的开发和维护

使用 OpenTelemetry 可以帮助开发者在分布式系统中更好地跟踪请求路径、识别瓶颈、并进行性能分析和故障排除。

市面上的链路追踪工具，像java实现的zipkin和go实现的jaeger。 因为它们的格式是不同的，如果要将zipkin换成jaeger，那么我们的代码和数据分析的promethus也需要进行修改。因此出现了opentracing+openCensus=OpenTelemetry, 它提供了统一了监控指标标准

**Opentelemetry能做什么**
1. 提供每种语言的lib库
2. 提供了中立的数据采集器，并支持多种方式部署
3. 生成、发送、收集、处理和导出遥测数据
4. 可以通过【欸之将数据并行发送到多个目的地
5. 提供了Opentracing和OpenCensus垫片

**traces链路**
traces链路就是请求调用链，例如
>A服务-->B服务-->C服务
A服务<--B服务<--C服务

同一个链路的所有span具有相同链路id

**span**: 一个完整的跨度是一个span

**metrics(指标/度量)**:包含一下几个类型
1. counter: 一个随时间累加的值，例如每分钟收到多少次请求
2. measure: 随时间聚合的值，例如直方图、不同范围的值的个数
3. observe: 捕获特定时间点的当前值，例如CPU使用率
