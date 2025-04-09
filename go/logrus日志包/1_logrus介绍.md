[toc]

# 1 logrus安装
```shell
go get github.com/sirupsen/logrus
```

# 2 logrus包功能介绍
`logrus` 是 Go 语言中使用广泛的结构化日志库，它比标准库 `log` 包提供了更多功能和灵活性。以下是 `logrus` 的主要功能：

1. **日志级别**：支持不同的日志级别如 `Panic`, `Fatal`, `Error`, `Warn`, `Info`, 和 `Debug`。这允许你根据情况打印不同重要性的消息。

2. **结构化日志**：允许你添加结构化数据到日志消息中，强化日志的语义。

3. **可扩展的钩子系统**：可以使用钩子（Hooks）将日志发送到各种外部存储系统，实现例如日志记录到文件、ELK（Elasticsearch, Logstash, Kibana）堆栈、Loggly、Sentry、Airbrake 或 Graylog 等。

4. **文本和JSON格式化**：支持日志输出格式的定制化，例如可以指定为 JSON 格式以便于解析和存储，或为了可读性优先使用文本格式。

5. **级别别名**：添加自定义日志级别的别名，你可以创建自己的日志级别并赋予它们特定的名字。

6. **日志旋转**：配合第三方工具如 `lumberjack` 包，可以实现日志文件的轮转管理，包括按大小切分和自动清理旧文件。

7. **线程安全**：`logrus` 是线程安全的，可以在并发环境下不用担心数据竞争。

8. **字段管理**：提供 `WithField` 和 `WithFields` 方法来包含额外的上下文信息，这有助于之后的日志分析。

9. **入口（Entry）**：你可以为特定的日志上下文创建一个日志入口，含有一连串的字段，然后基于其创建更多的日志，它们会包含这些预先定义好的字段数据。

10. **错误处理**：虽然 `logrus` 的API设计类似于标准库 `log` 包，但它还提供了一个 `WithError` 方法来记录错误，并保持与其他字段相同的处理方式。

11. **调用函数信息**：它可以自动记录出日志调用的函数信息，帮助在调试期间定位日志来源。

然后在你的 Go 程序中引入并使用它：
```go
import log "github.com/sirupsen/logrus"

func main() {
    log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")
}
```

`logrus` 是灵活和强大的，可以通过各种不同的实现被扩展以满足特殊的日志记录需求。上面列出的特性只是入门级别的介绍，更深层次的使用可以通过探索 `logrus` 的文档和源码来获得更多信息和示例。
