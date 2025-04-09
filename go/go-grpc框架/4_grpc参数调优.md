[toc]

# 1 grpc客户端参数配置
## 1.1 设置最大发送/接收消息大小
设置 gRPC 消息的最大接收和发送大小。这些参数可以防止因为体积太大的消息造成内存异常，但也需要根据实际需求来设置，以确保不会过早地截断有效的消息。
```go
// WithDefaultCallOptions returns a DialOption which sets the default
// CallOptions for calls over the connection.
func WithDefaultCallOptions(cos ...CallOption) DialOption

// 设置最大接收的消息的大小
func MaxCallRecvMsgSize(bytes int) CallOption

func MaxCallSendMsgSize(bytes int) CallOption
```

**示例代码**：
```go
package main

import (
    "github.com/sirupsen/logrus"
    "google.golang.org/grpc"
)

func main(){
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

    // 设置最大接收消息的长度为10Mb
    opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*10)))
    
    // 设置发送消息的最大size为10Mb
    opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024*10)))


    var conn *grpc.ClientConn
    var err error
    conn, err = grpc.Dial(grpcAddr, opts...)
    if err != nil {
        logrus.Infoln("connect grpc server failed:", err.Error())
    } else {
        logrus.Infoln("client: connect rpc")
    }
    defer conn.Close()
}
```

**注意**：
- 如果发送/接收的消息最大消息设置小了，会导致发送/接收的消息被截断。例如发送消息最大值设置小了会报错`rpc error: code = ResourceExhausted desc = trying to send message larger than max (92 vs. 2)`

## 1.2 设置传输的数据压缩格式
启用数据压缩可以减少传输数据的大小，提高网络效率，特别是在带宽受限的情况下。内置的压缩方法包括 Gzip。
```go
var opts []grpc.DialOption

opts = append(opts, grpc.WithDefaultCallOptions(grpc.CompressorCallOption{
    CompressorType: "gzip",
}))

// 或者这种方式设置压缩格式，是对上面的封装,gzip是grpc内置的
opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")))
```

**注意**：
- 要使用压缩功能，需要grpc服务器上配置对应的压缩格式。
可以使用
```go
import "google.golang.org/grpc/encoding/gzip"

func init() {
    // 数字范围是[0,9]
	_ = gzip.SetLevel(1)
}
```

**自定义压缩**：
```go
import "google.golang.org/grpc/encoding"
// 只能在init函数中执行

/// xxx 为自定义的Compressor
encoding.RegisterCompressor(xxx)
```

## 1.3 设置读取缓冲区的大小
“读取缓冲区大小”是指 HTTP/2 流在读取过程中用于暂存数据的内部缓冲区的大小(单位byte)。如果你将这个大小设置得比较大，那么可以一次读取更多的数据，减少系统调用次数，从而可能提高性能，尤其是在处理大量数据时。然而，这也会导致每个流消耗更多的内存，特别是当有很多并发的流时，可能会消耗很大的内存空间。反之，如果将缓冲区设置得太小，则可能需要更频繁的读取操作，从而影响性能。
```go
func WithReadBufferSize(s int) DialOption
```

**使用示例**：
```go
var opts []grpc.DialOption
opts = append(opts, grpc.WithReadBufferSize(16*1024))
```

**注意**：
- 如果设置的值非常小，它可能会降低 I/O 效率，并增加 CPU 使用率，因为系统需要执行更多的读取系统调用。
- 如果设置的值非常大，则会提高单个流的内存消耗，这在你有大量并发流的情况下可能会消耗可观的系统资源。
- 如果不设置，gRPC 将使用默认的缓冲区大小。
- 在很多情况下，你可能不需要调整这个值，特别是当默认的大小已经满足性能需要的情况下。

## 1.4 阻塞直到连接上服务器
阻塞直到连接上服务器，一般不建议使用
```go
func WithBlock() DialOption
```

***使用示例*:
```go
var opts []grpc.DialOption
opts = append(opts,grpc.WithBlock())
```

## 1.5 WithAuthority
允许你为创建的 gRPC 客户端连接指定 :authority header 的值。:authority header 是 HTTP/2 请求中的一个重要组成部分，对应 HTTP/1 中的 Host header，并且在 gRPC 调用中被用作服务器名称。

当 gRPC 客户端发起请求时，:authority header 通常用于以下目的：

虚拟主机名/Virtual Hosting：当服务端托管多个服务域，:authority header 可以帮助服务端判断请求应该路由到那个服务域。
安全：在 TLS 握手过程中，这个值被用于服务端证书的验证，特别是验证证书中的 Common Name (CN) 或 Subject Alternative Name (SAN)。
这个方法在配置客户端时非常有用，特别是在你需要覆盖由服务端证书指定的默认服务器名称或者使用不同的虚拟主机名时。
```go
func WithAuthority(a string) DialOption
```

**使用示例**：
```go
var opts []grpc.DialOption
opts = append(opts, grpc.WithAuthority("test"))
```

## 1.6 设置初始窗口大小
WithInitialWindowSize(s int32) 函数是用来创建一个 DialOption，它设置了 HTTP/2 流的初始窗口大小（initial window size）。这个窗口大小决定了对端在需要接收方确认之前可以发送的数据量。这是一种流量控制（flow control）机制，用于避免发送方向接收方发送过多数据，从而导致接收方不堪重负。

HTTP/2 协议中的流量控制依赖于窗口大小的概念。每个方向上的数据传送都有一个相应的窗口大小，它定义了何时停止发送数据以等待对方的窗口更新（即等待接收方处理已经接收的数据并准备接收更多数据）。WithInitialWindowSize 配置项允许增加或减少流的初始窗口大小，从而影响流量控制的行为。

如果你设置了一个较大的窗口大小，那么在发送方需要停下来等待接收方确认之前，它可以发送更多的数据。这在传输大量数据时可能有利于减少等待时间，提高吞吐量。反之，如果窗口大小被设置得较小，那么发送方在发送较少量的数据后就需要等待接收方的确认，这可能有助于防止接收方被大量的来自发送方的数据淹没。
```go
// 设置初始窗口的大小，单位byte
func WithInitialWindowSize(s int32) DialOption
```

**使用示例**:
```go
// 设置初始窗口的大小为1Mb
opts = append(opts, grpc.WithInitialWindowSize(1<<20))
```

## 1.7 添加拦截器
WithChainStreamInterceptor 函数是用来创建一个 DialOption，该选项用于配置 gRPC 客户端的流拦截器链。拦截器 (Interceptors) 是 gRPC 中的一个强大特性，允许用户在某个 RPC（远程过程调用）过程中插入自定义逻辑，如日志记录、身份验证、指标收集等。

流拦截器 (StreamClientInterceptor) 特别用于流式 RPC 调用，拦截流式 RPC 的客户端操作。一个流拦截器能够包装 ClientStream，而如果你有多个拦截器的话，WithChainStreamInterceptor 能帮助你按顺序链式地串联它们一起。

举例来说，若你有两个拦截器，A 和 B，并且使用 WithChainStreamInterceptor(A, B) 配置你的 gRPC 客户端，那么拦截器 A 将首先被调用，而 B 则在 A 之后被调用。当客户端执行流式 RPC 时，请求首先通过拦截器 A 的逻辑，然后是 B 的逻辑，最后才到达服务端。响应走的路线则是反过来的，先经过 B，然后是 A。
```go
func WithChainStreamInterceptor(interceptors ...StreamClientInterceptor) DialOption

type StreamClientInterceptor 
func(   ctx context.Context,
        desc *StreamDesc,
        cc *ClientConn,
        method string,
        streamer Streamer,
        opts ...CallOption) (ClientStream, error)
```

**使用示例**:
```go
package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func interceptor1(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("this is interceptor1")
	return streamer(ctx, desc, cc, method, opts...)
}

func interceptor2(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("this is interceptor2")
	return streamer(ctx, desc, cc, method, opts...)
}

func main(){
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithChainStreamInterceptor(interceptor1, interceptor2))
    // ...
}
```

**注意**:
- 使用 WithChainStreamInterceptor 创建的拦截器链应该保持无状态或在并发场景下是安全的，因为它们可能会从多个 goroutines（Go 语言的并发运行单元）同时被调用。不同的拦截器之间定义好明确的操作顺序是很重要的，因为它们执行的顺序会影响处理流程。

- 请注意，混合使用 WithChainStreamInterceptor 和单独的 WithStreamInterceptor 会导致错误，这两者不应当在同一 gRPC 客户端连接配置中一起使用。如果你在连接上多次调用 WithChainStreamInterceptor，则只有最后一次调用会生效。


# 2 grpc服务器参数配置
# 1.1 设置最大发送/接收消息大小
设置 gRPC 消息的最大接收和发送大小。这些参数可以防止因为体积太大的消息造成内存异常，但也需要根据实际需求来设置，以确保不会过早地截断有效的消息。
```go
// 设置接收消息的最大大小(单位byte)，默认4M
func MaxRecvMsgSize(m int) ServerOption

// 设置发送消息的最大大小(单位byte)，默认为math.MaxInt32
func MaxSendMsgSize(m int) ServerOption
```

**示例代码**:
```go
serverOpts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024*1024*10),
		grpc.MaxSendMsgSize(1024*1024*10),
	}
	
	grpcServer := grpc.NewServer(serverOpts...)
```



## 1.2 设置连接同时打开流的数量
限制每个连接可以同时打开的流的数量。增加这个值可以允许更高的并发度，但可能会增加内存和处理的负担。
```go
func MaxConcurrentStreams(n uint32) ServerOption
```

**示例代码**:
```go
// grpc服务配置参数
serverOpts := []grpc.ServerOption{
    grpc.MaxConcurrentStreams(12),
}

// 创建grpc服务
grpcServer := grpc.NewServer(serverOpts...)
```

## 1.3 连接活跃配置
调整保持连接活跃的参数，包括 ping 时间间隔、超时时间等。正确设置这些参数可以帮助检测和维持活跃的连接，同时减少不必要的连接重建成本。
```go
func KeepaliveParams(kp keepalive.ServerParameters) ServerOption

type ServerParameters struct {
    // 这个时间段指定了一个空闲连接在发送GoAway帧关闭之前应该保持的最大时间。一个连接如果在这个时间段内没有任何正在进行的RPC请求，就会被认为是空闲的。
    MaxConnectionIdle time.Duration // 默认值无限长

    // 这个时间段指明了一个连接在发送GoAway帧关闭之前可以存在的最大时间。这个限制帮助服务器避免长久维持旧的连接，促成定期重新连接以便更新路由等配置。为了避免多个连接同时关闭导致的连接风暴，随机增减最多10%的时间震荡会被应用到MaxConnectionAge的值上。默认值是无穷大。
    MaxConnectionAge time.Duration // 默认值是无线长

    // 这是在MaxConnectionAge之后添加的额外时间，在这段时间内即使连接已经超过了其最大存活时间，也可以允许活跃的RPC完成。一旦宽限期结束，连接将被强制关闭。默认值同样是无穷大。
    MaxConnectionAgeGrace time.Duration // 默认值是无限长
    
    // 经过一段时间后，如果服务器没有看到任何活动，它会ping客户端，看看传输是否仍然有效。如果设置低于1s，则使用最小值1s代替。
    Time time.Duration // 默认值2小时
    
    // 在ping通keepalive检查之后，服务器等待一段时间的超时，如果在此之后没有看到任何活动，则关闭连接。
    Timeout time.Duration // 默认20秒
}
```

**示例**：
```go
serverOpts := []grpc.ServerOption{
    grpc.KeepaliveParams(keepalive.ServerParameters{
        Time: time.Minute * 10,
        Timeout: time.Second*5,
    }),
}

grpcServer := grpc.NewServer(serverOpts...)
```

## 1.4 设置连接建立超时时间
为所有新连接设置连接建立超时(包括HTTP/2握手)。如果不设置，则默认为120秒。零或负值将导致立即超时。
```go
func ConnectionTimeout(d time.Duration) ServerOption
```

**使用示例**:
```go
serverOpts := []grpc.ServerOption{
    // 设置10秒连接超时
    grpc.ConnectionTimeout(time.Millisecond * 10000),
}
```

## 1.5 设置初始窗口的大小

为连接设置窗口大小（单位byte）。窗口大小的下限是64K，任何小于这个值的值都将被忽略。
```go
func InitialConnWindowSize(s int32) ServerOption
```

**使用示例**:
```go
grpc.InitialConnWindowSize(1024*64),
```

## 1.6 设置流的动态报头表的大小
用于设置 HTTP/2 头部压缩（HPACK）的头表大小（单位byte）。设置较大的头表可以增加压缩率和性能，但也会使用更多的内存。反之，较小的表可以节省内存，但可能减少压缩效率。
```go
func HeaderTableSize(s uint32) ServerOption
```

**使用示例**：
```go
serverOpts := []grpc.ServerOption{
    grpc.HeaderTableSize(4096),
}
```

## 1.7 在服务器接受新的连接之前或处理新的 RPC 之前执行特定的代码
InTapHandle 函数用来创建一个 ServerOption，该选项允许用户为 gRPC 服务器端注册一个 “tap” 处理函数。tap（测试访问点）功能使用户能够插入自定义的钩子（hooks）函数，在服务器接受新的连接之前或处理新的 RPC 之前执行特定的代码。

具体来说，ServerInHandle 类型定义的 tap 函数接收一个 tap.Info 类型的参数，里面包含关于新的 RPC 调用的信息，用户可以利用这个信息实现例如日志记录、度量收集、鉴权、限流或其他自定义逻辑。
```go
func InTapHandle(h tap.ServerInHandle) ServerOption

type ServerInHandle func(ctx context.Context, info *Info) (context.Context, error)
```

**使用示例**:
```go
package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/tap"
)

func tabHandle(ctx context.Context, info *tap.Info) (context.Context, error) {
	// 在服务器接受新的连接之前或处理新的 RPC 之前执行特定的代码
	return ctx, nil
}

func main(){
    serverOpts := []grpc.ServerOption{
        grpc.InTapHandle(tabHandle),
    }
    grpcServer := grpc.NewServer(serverOpts...)
    // ...
}
```

## 1.8 限制客户端可以发送的最大头部列表大小

```go
func MaxHeaderListSize(s uint32) ServerOption
```

在 `google.golang.org/grpc` 包中，`MaxHeaderListSize` 函数用于创建一个 `ServerOption`，这个选项限制客户端可以发送的最大头部列表大小，当设置在 gRPC 服务器上时。这主要是一个为了安全和性能考虑的配置，它可以防止客户端通过发送巨大的头部列表来尝试耗尽服务器资源。

HTTP/2 协议规定了一个 `SETTINGS_MAX_HEADER_LIST_SIZE` 参数，用来告诉 peer（对等端点，可能是客户端或服务器）它愿意接收的头部列表的最大尺寸。这不是一个强制的限制，但是发送方尊重这个限制可以帮助防止遭受某些类型的拒绝服务攻击（DoS 攻击）。

具体来说，`MaxHeaderListSize` 配置的大小是用字节为单位的：

- `s uint32`：指定服务器愿意接受的最大头部列表大小。

这是一个设置该选项的示例：

```go
server := grpc.NewServer(grpc.MaxHeaderListSize(16 * 1024))
```

在这个例子中，服务器被配置为只接受最大16KB大小的头部列表。如果客户端尝试发送一个超过这个大小的头部列表，服务器可以选择以PROTOCOL_ERROR作为HTTP/2错误码拒绝该连接或流。

这个选项的设置要根据应用的具体需求来确定：较大的限制可以增加灵活性，允许头部包含更多的数据，但是增加了资源耗尽的风险；较小的限制则减小了这种风险，但是可能不足以满足包含很多头部信息的应用。通常默认值就足够大，可以适用大多数情况，只有在特定环境中需要进行微调。

## 1.9 设置处理流的协程数量
用于设置应该用于处理传入流的工作协程的数量。将此值设置为零(默认值)将禁用worker并为每个流生成一个新的例程。

```go
func NumStreamWorkers(numServerWorkers uint32) ServerOption
```

**使用示例**:
```go
serverOpts := []grpc.ServerOption{
    grpc.NumStreamWorkers(128),
}
```

## 1.10 配置 gRPC 服务器的读缓冲区大小
于配置 gRPC 服务器的读缓冲区大小。此设置将影响服务器底层 TCP 连接的读操作，它确定了在一次 read 系统调用中最多可以读取多少字节的数据。默认值为 32KB。

具体来说，ReadBufferSize 的作用包括：

- 增加或减少服务器处理入站消息时的缓冲区大小。
- 当设置为零或负数时，禁用连接的读缓冲区，使数据帧（framer）可以直接访问底层的连接（conn），这有助于某些用例减少内存消耗和潜在的缓冲延迟。
- 对于每个连接来说，较大的读缓冲区可能改进了网络吞吐率，因为它减少了每个消息可能产生的 read 调用的数量。这是因为它允许以更大的块来读取数据。
```go
func ReadBufferSize(s int) ServerOption
```

**使用示例**:
```go
serverOpts := []grpc.ServerOption{
    grpc.ReadBufferSize(32768),
}
```

## 1.11 设置写缓冲区大小
WriteBufferSize决定在对网络执行写操作之前可以批处理多少数据。该缓冲区对应的内存分配将是其大小的两倍，以保持低系统调用。这个缓冲区的默认值是32KB。0或负值将禁用写缓冲区，这样每次写都将在底层连接上进行。
```go
func WriteBufferSize(s int) ServerOption
```

**使用示例**:
```go
serverOpts := []grpc.ServerOption{
    grpc.ReadBufferSize(32768),
}
```

## 1.12 设置是否允许重用每个连接的传输写缓冲区。
SharedWriteBuffer允许重用每个连接的传输写缓冲区。
如果该选项设置为true，每个连接都会在刷新数据后释放缓冲区。

```go
func SharedWriteBuffer(val bool) ServerOption
```

