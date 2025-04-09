[toc]

# 1 正常打印日志
logrus正常打印日志提供了跟log包类似的接口
```go
package logrusTest

import (
	"github.com/sirupsen/logrus"
)

func LogrusTest() {
    // 设置日志等级未Debug级别
    logrus.SetLevel(logrus.DebugLevel)

    // time="2024-03-13T10:53:07+08:00" level=info msg="log test 888"
    logrus.Println("log test", "888")

    // time="2024-03-13T10:54:01+08:00" level=info msg="I have a apple"
    logrus.Printf("I have a %s", "apple")

    // time="2024-03-13T10:54:01+08:00" level=info msg="I have a apple"
    logrus.Infoln("logrus.Infof test")

    // time="2024-03-13T10:55:48+08:00" level=info msg="I have a pen"
    logrus.Infof("I have a %s", "pen")

    // time="2024-03-13T10:56:28+08:00" level=warning msg="warning log test"
    logrus.Warning("warning log test")

    // time="2024-03-13T10:58:22+08:00" level=warning msg="warning log test 2"
    logrus.Warningln("warning log test 2")

    // time="2024-03-13T10:59:00+08:00" level=warning msg="You have a apple"
    logrus.Warningf("You have a %s", "apple")

    // time="2024-03-13T10:59:48+08:00" level=error msg="error log test"
    logrus.Errorln("error log test")

    // time="2024-03-13T10:59:48+08:00" level=error msg="error log fff"
    logrus.Errorf("error log %s", "fff")

    // 需要logrus.SetLevel(logrus.DebugLevel)才能打印
    // time="2024-03-13T11:01:59+08:00" level=debug msg="debug log fff"
    logrus.Debugf("debug log %s", "fff")

    // 需要logrus.SetLevel(logrus.DebugLevel)才能打印
    // time="2024-03-13T11:01:59+08:00" level=debug msg="debug test"
    logrus.Debugln("debug test")

	// time="2024-03-13T11:12:21+08:00" level=fatal msg="fatal error:error"
	logrus.Fatalf("fatal error:%s", "error")

	// time="2024-03-13T11:12:58+08:00" level=fatal msg="fatal error ln"
	logrus.Fatalln("fatal error ln")

	// time="2024-03-13T11:13:35+08:00" level=panic msg="panic error:panic"
	logrus.Panicf("panic error:%s", "panic")

	// time="2024-03-13T11:14:08+08:00" level=panic msg="panic test"
	logrus.Panic("panic test")
}
```

# 2 日志等级

|日志等级|描述|
|-------|----|
|logrus.PanicLevel|panic级别，只有panic日志才能被打印|
|logrus.FatalLevelt|fatal级别，只有fatal及其之上等级的日志才能被打印|
|logrus.ErrorLevel|error级别，只有error及其之上等级的日志才能被打印|
|logrus.WarnLevel|警告日志，只有warn及其之上等级的日志才能被打印|
|logrus.InfoLevel|info日志，只有info及其之上等级的日志才能被打印|
|logrus.DebugLevel|debug日志，只有debug及其之上等级的日志才能被打印|
|logrus.TraceLevel|trace日志，只有trace及其之上等级的日志才能被打印|

以上日志级别从上到下依次递减
可以通过```logrus.SetLevel```方法来设置日志等级，在设置的等级之下的日志将不会被打印。

# 3 将日志输出到文件
如果想要将日志输出到文件可以通过```logrus.SetOutput```方法来设置日志输出的文件。
以下例子通过lumberjack包来管理日志文件

```go
package logrusTest

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogrusTest() {
    // 创建一个新的日志记录器
	log := logrus.New()

    // 创建一个lumberjack的实例，用于日志文件的轮转管理
	logFile := &lumberjack.Logger{
		Filename:   "logs/test/log.txt",
		MaxSize:    30,     // 每个日志的最大大小，单位Mb,最大100
		MaxAge:     7,     // 文件存在天数的最大值
		MaxBackups: 3,     // 要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件
		LocalTime:  true,  // 是否使用本地时间，默认是UTC时间
		Compress:   false, // 是否压缩日志文件
	}

    // 设置日志输出到文件
	log.SetOutput(logFile)

    // 测试打印日志
	idx := 1
	for ; idx < 100000; idx++ {
		log.Printf("log test %d, I have a apple, I have a pen", idx)
	}
	logFile.Close()
}
```



