pprof用于对服务进行性能监听和分析

## 使用

```go
import (
	"net/http"
	_ "net/http/pprof"
)

// 启动性能监听服务
go func() {
	_ = http.ListenAndServe(":6060", nil)
}()
```

通过以上代码启动程序后，通过访问 `http://localhost:6060/debug/pprof/` 来获取程序的实时运行数据，比如 CPU 使用情况、内存分配、Goroutine 堆栈信息等.

## 信息查看

启动pprof后，不同的访问路径对应的监控信息不同

|访问路径|信息|
|-------|----|
|/debug/pprof/profile|CPU 使用情况分析，找出 CPU 占用高的热点函数|
|/debug/pprof/heap|内存分配情况，排查内存泄漏，优化内存使用|
|/debug/pprof/goroutine|当前所有 Goroutine 的堆栈信息，排查 Goroutine 泄漏或死锁|
|/debug/pprof/block|阻塞操作的跟踪，分析锁竞争和阻塞等待|
|/debug/pprof/mutex|锁竞争情况，优化并发性能|


## 终端命令

可以使用go自带的go tool pprof工具来采集数据

`go tool pprof <访问路径>`



|数据项|	说明|	用途|
|------|--------|------|
|/debug/pprof/profile|	CPU 使用情况分析|	找出 CPU 占用高的热点函数|
|/debug/pprof/heap|	内存分配情况|	排查内存泄漏，优化内存使用|
|/debug/pprof/goroutine|	当前所有| Goroutine 的堆栈信息	排查 Goroutine 泄漏或死锁|
|/debug/pprof/block|	阻塞操作的跟踪|	分析锁竞争和阻塞等待|
|/debug/pprof/mutex|	锁竞争情况|	优化并发性能|


