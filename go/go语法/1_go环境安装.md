[toc]
# 1 安装go环境
- 下载对应的go版本
https://golang.google.cn/doc

- 解压到/usr/local下
    ```shell
    sudo tar -C /usr/local/ -xzf go1.18.1.linux-amd64.tar.gz
    ```
- 配置环境变量
    ```shell
    # 在/etc/profile中添加
    export PATH=$PATH:/usr/local/go/bin

    # 使添加的环境变量立即生效 或者重启
    source /etc/profile

    # 查看go版本，成功则说明环境配置成功
    go version
    ```
- 更改下载源的代理，这样后面下载包会快
    ```shell
    #开启go module (go1.11版本以上)
    go env -w GO111MODULE=on
    # 更改下载源的代理
    go env -w GOPROXY=https://goproxy.io,direct

    # 阿里云镜像
    go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

    # 原始为 go env -w GOPROXY="https://proxy.golang.org,direct"
    ```

# 2 go环境变量
环境变量可以用```go env```命令来查看
**GOROOT**：   这是go环境所在位置
**GOPATH**：  将来GO项目所在的位置。GOPATH下有3个文件夹bin、src、package我们的项目源文件方在src目录下，package下存放编译好的包，bin下为可执行程序. 如果电脑上没有$GOPATH，就手动创建

# 3 更换go版本
将/usr/local/go下的文件全部删除，然后重新配置即可
```shell
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
```

# 4 go 介绍
**特性**
**语法简单** Go在自由与灵活上做了取舍，以此换来了更好的维护性和更低的学习难度。
**交叉编译** 允许跨平台编译代码。
**天然并发** Go语言对并发的支持是纯天然的，语法上只需要一个go关键字就能开启一个协程。
**垃圾回收** GC算不上很出色，但还比较靠谱。
**静态链接** 简化了部署操作，无需安装任何运行环境和诸多第三方库。
**内存分配** 可以说，除偶尔因性能问题而被迫采用对象池和自主内存管理外，基本无须参与内存管理操作。

Go语言抛弃了继承，弱化了OOP，类，元编程，泛型，Lamda表达式等这些特性，拥有不错的性能和较低的上手难度。Go语言非常适合用于云服务开发，应用服务端开发，甚至可以进行部分Linux嵌入式开发，不过由于带有垃圾回收，其性能肯定无法媲美C/C++这类系统级语言，但Go在其擅长的领域表现十分出色。虽然面世只有13年，但也已经有大量的行业采用了Go作为首选语言，总的来说这是一门非常年轻且具有活力的现代语言，值得一学。

# 5 vscode下配置go环境
- 下载go插件---关键字go
- 在$GOPATH/src下创建一个go项目并初始化
    ```shell
    $ cd $GOPATH/src/hello
    $ go mod init
    ```
    ```go mod init```会创建go.mod文件和go.sum文件，go.mod文件是记录我们依赖库以及版本号
- 以后编写完代码后执行```go mod tidy```即可，这个命令会自动下载依赖的库，也会删除多余的库

# 6 包管理
## 6.1 go.mod文件
Go.mod是Golang1.11版本新引入的官方包管理工具用于解决之前没有地方记录依赖包具体版本的问题，方便依赖包的管理。
Modules是相关Go包的集合，是源代码交换和版本控制的单元。go命令直接支持使用Modules，包括记录和解析对其他模块的依赖性。Modules替换旧的基于GOPATH的方法，来指定使用哪些源文件。
使用Go的包管理方式，依赖的第三方包被下载到了$GOPATH/pkg/mod路径下。

Modules和传统的GOPATH不同，不需要包含例如src，bin这样的子目录，一个源代码目录甚至是空目录都可以作为Modules，只要其中包含有go.mod文件。

golang 提供了 go mod命令来管理包。go mod 有以下命令：
|命令|说明|
|----|----|
|go mod download|下载依赖包|
|go mod tidy|删除不需要的依赖包、下载新的依赖包、更新go.sum|
|go mod init|在当前目录初始化mod|
|go mod why|解释为什么需要依赖|
|go clean -modcache|清楚所有go缓存，包括go get下载的库|
|go list -m -versions <包名>|查看指定包的所有版本|

## 6.2 设置使用mod来管理
1. 首先将go的版本升级为1.11以上
2. 设置GO111MODULE
```go env -w GO111MODULE=xxx```

GO111MODULE有三个值：off, on和auto（默认值）。
**GO111MODULE=off**，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
**GO111MODULE=on**，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
**GO111MODULE=auto**，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
当前目录在GOPATH/src之外且该目录包含go.mod文件
当前文件在包含go.mod文件的目录下面。

## 6.3 go.mod在项目中的使用
1. 首先我们要在GOPATH/src 目录之外新建工程，或将老工程copy到GOPATH/src 目录之外。
PS：go.mod文件一旦创建后，它的内容将会被go toolchain全面掌控。go toolchain会在各类命令执行时，比如go get、go build、go mod等修改和维护go.mod文件。
go.mod 文件内提供了module, require、replace和exclude四个关键字，这里注意区分与上表中命令的区别，一个是管理go mod命令，一个是go mod文件内的关键字。
<font color=blue>module</font>语句指定包的名字（路径）
<font color=blue>require</font>语句指定的依赖项模块
<font color=blue>replace</font>语句可以替换依赖项模块
<font color=blue>exclude</font>语句可以忽略依赖项模块
下面是我们建立了一个hello.go的文件
    ```go
    package main
    import (
        "fmt"
    )
    func main() {
        fmt.Println("Hello, world!")
    }
    ```
2. 在当前目录下，命令行运行 go mod init + 模块名称 初始化模块
即go mod init hello
运行完之后，会在当前目录下生成一个go.mod文件，这是一个关键文件，之后的包的管理都是通过这个文件管理。
官方说明：除了go.mod之外，go命令还维护一个名为go.sum的文件，其中包含特定模块版本内容的预期加密哈希 
go命令使用go.sum文件确保这些模块的未来下载检索与第一次下载相同的位，以确保项目所依赖的模块不会出现意外更改，无论是出于恶意、意外还是其他原因。 go.mod和go.sum都应检入版本控制。 
go.sum 不需要手工维护，所以可以不用太关注。
注意：子目录里是不需要init的，所有的子目录里的依赖都会组织在根目录的go.mod文件里
接下来，让我们的项目依稀一下第三方包
如修改hello.go文件如下，按照过去的做法，要运行hello.go需要执行go get 命令 下载gorose包到 $GOPATH/src
    ```go
    package main
    
    import (
        "fmt"
        "github.com/gohouse/gorose"
    )
    
    func main() {
        fmt.Println("Hello, world!")
    }
    ```
    但是，使用了新的包管理就不在需要这样做了
    直接 go run hello.go
    稍等片刻… go 会自动查找代码中的包，下载依赖包，并且把具体的依赖关系和版本写入到go.mod和go.sum文件中。
    查看go.mod，它会变成这样：

    ```go
    module test
    require (
        github.com/gohouse/gorose v1.0.5
    )
    ```
    require 关键字是引用，后面是包，最后v1.11.1 是引用的版本号
    这样，一个使用Go包管理方式创建项目的小例子就完成了。

**问题一**：依赖的包下载到哪里了？还在GOPATH/src里吗？
不在。
使用Go的包管理方式，依赖的第三方包被下载到了$GOPATH/pkg/mod路径下。

**问题二**： 依赖包的版本是怎么控制的？
在上一个问题里，可以看到最终下载在$GOPATH/pkg/mod 下的包中最后会有一个版本号 v1.0.5，也就是说，$GOPATH/pkg/mod里可以保存相同包的不同版本。

版本是在go.mod中指定的。如果，在go.mod中没有指定，go命令会自动下载代码中的依赖的最新版本，本例就是自动下载最新的版本。如果，在go.mod用require语句指定包和版本 ，go命令会根据指定的路径和版本下载包，
指定版本时可以用latest，这样它会自动下载指定包的最新版本；

**问题三**： 可以把项目放在$GOPATH/src下吗？
可以。但是go会根据GO111MODULE的值而采取不同的处理方式，默认情况下，GO111MODULE=auto 自动模式

1.auto 自动模式下，项目在$GOPATH/src里会使用$GOPATH/src的依赖包，在$GOPATH/src外，就使用go.mod 里 require的包

2.on 开启模式，1.12后，无论在$GOPATH/src里还是在外面，都会使用go.mod 里 require的包

3.off 关闭模式，就是老规矩。

**问题四**： 依赖包中的地址失效了怎么办？比如 golang.org/x/… 下的包都无法下载怎么办？
在go快速发展的过程中，有一些依赖包地址变更了。以前的做法：

1.修改源码，用新路径替换import的地址

2.git clone 或 go get 新包后，copy到$GOPATH/src里旧的路径下

无论什么方法，都不便于维护，特别是多人协同开发时。

使用go.mod就简单了，在go.mod文件里用 replace 替换包，例如

replace golang.org/x/text => github.com/golang/text latest

这样，go会用 github.com/golang/text 替代golang.org/x/text，原理就是下载github.com/golang/text 的最新版本到 $GOPATH/pkg/mod/golang.org/x/text下。

**问题五**： init生成的go.mod的模块名称有什么用？
本例里，用 go mod init hello 生成的go.mod文件里的第一行会申明module hello

因为我们的项目已经不在$GOPATH/src里了，那么引用自己怎么办？就用模块名+路径。

例如，在项目下新建目录 utils，创建一个tools.go文件:
```go
package utils
 
import “fmt”
 
func PrintText(text string) {
    fmt.Println(text)
}
```
在根目录下的hello.go文件就可以 import “hello/utils” 引用utils
```go
package main
 
import (
	"hello/utils"
	"github.com/astaxie/beego"
)
 
func main() {
	utils.PrintText("Hi")
	beego.Run()
}
```
**问题六**：以前老项目如何用新的包管理
如果用auto模式，把项目移动到$GOPATH/src外eyJsaWNlbnNlSWQiOiJTRlhVU0E4NkZNIiwibGljZW5zZWVOYW1lIjoi5pyd6Zm956eR5oqA5aSn5a24IiwibGljZW5zZWVUeXBlIjoiQ0xBU1NST09NIiwiYXNzaWduZWVOYW1lIjoiVGFvYmFv77ya5p6B5a6i5LiT5LqrICAtLS0g6LCo6Ziy55uX5Y2W77yBIiwiYXNzaWduZWVFbWFpbCI6IktyaXN0YW5fQmxvd2VAb3V0bG9vay5jb20iLCJsaWNlbnNlUmVzdHJpY3Rpb24iOiJGb3IgZWR1Y2F0aW9uYWwgdXNlIG9ubHkiLCJjaGVja0NvbmN1cnJlbnRVc2UiOmZhbHNlLCJwcm9kdWN0cyI6W3siY29kZSI6IkdPIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJSUzAiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRNIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJDTCIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiUlNVIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJSU0MiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUEMiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRTIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJSRCIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiUkMiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IlJTRiIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjp0cnVlfSx7ImNvZGUiOiJSTSIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiSUkiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRQTiIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiREIiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRDIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJQUyIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiUlNWIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOnRydWV9LHsiY29kZSI6IldTIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJQU0kiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUENXTVAiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUlMiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiRFAiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUERCIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOnRydWV9XSwibWV0YWRhdGEiOiIwMTIwMjQwMjI2TFBBQTAwMzAwOCIsImhhc2giOiI1NDY4ODAyOS8yNTk5OTU2NTotMTQ5MzMwODg5NSIsImdyYWNlUGVyaW9kRGF5cyI6NywiYXV0b1Byb2xvbmdhdGVkIjpmYWxzZSwiaXNBdXRvUHJvbG9uZ2F0ZWQiOmZhbHNlLCJ0cmlhbCI6ZmFsc2UsImFpQWxsb3dlZCI6dHJ1ZX0
进入目录，运行 go mod init + 模块名称
go build 或者 go run 一次

# 7 goland环境配置
- 在file/setting/Go Modules下将GOPROXY配置成`go env`中的一致。
![](img/goland_1.png)




http://idea.javatiku.cn/


SFXUSA86FM-eyJsaWNlbnNlSWQiOiJTRlhVU0E4NkZNIiwibGljZW5zZWVOYW1lIjoi5pyd6Zm956eR5oqA5aSn5a24IiwibGljZW5zZWVUeXBlIjoiQ0xBU1NST09NIiwiYXNzaWduZWVOYW1lIjoiVGFvYmFv77ya5p6B5a6i5LiT5LqrICAtLS0g6LCo6Ziy55uX5Y2W77yBIiwiYXNzaWduZWVFbWFpbCI6IktyaXN0YW5fQmxvd2VAb3V0bG9vay5jb20iLCJsaWNlbnNlUmVzdHJpY3Rpb24iOiJGb3IgZWR1Y2F0aW9uYWwgdXNlIG9ubHkiLCJjaGVja0NvbmN1cnJlbnRVc2UiOmZhbHNlLCJwcm9kdWN0cyI6W3siY29kZSI6IkdPIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJSUzAiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRNIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJDTCIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiUlNVIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJSU0MiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUEMiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRTIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJSRCIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiUkMiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IlJTRiIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjp0cnVlfSx7ImNvZGUiOiJSTSIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiSUkiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRQTiIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiREIiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6ZmFsc2V9LHsiY29kZSI6IkRDIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJQUyIsInBhaWRVcFRvIjoiMjAyNS0wMi0xOSIsImV4dGVuZGVkIjpmYWxzZX0seyJjb2RlIjoiUlNWIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOnRydWV9LHsiY29kZSI6IldTIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOmZhbHNlfSx7ImNvZGUiOiJQU0kiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUENXTVAiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUlMiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiRFAiLCJwYWlkVXBUbyI6IjIwMjUtMDItMTkiLCJleHRlbmRlZCI6dHJ1ZX0seyJjb2RlIjoiUERCIiwicGFpZFVwVG8iOiIyMDI1LTAyLTE5IiwiZXh0ZW5kZWQiOnRydWV9XSwibWV0YWRhdGEiOiIwMTIwMjQwMjI2TFBBQTAwMzAwOCIsImhhc2giOiI1NDY4ODAyOS8yNTk5OTU2NTotMTQ5MzMwODg5NSIsImdyYWNlUGVyaW9kRGF5cyI6NywiYXV0b1Byb2xvbmdhdGVkIjpmYWxzZSwiaXNBdXRvUHJvbG9uZ2F0ZWQiOmZhbHNlLCJ0cmlhbCI6ZmFsc2UsImFpQWxsb3dlZCI6dHJ1ZX0=-JDVXZeZnNxn5sMQEXZ2TOZlrMOVI37CPE25JugHcDUdJPc75u4D+IEwoFl1GRB8GKrIhSwJa6OhgHpyXyMqLXtroe/p+qWo6kLi86iTuXpK+E4UQPQP9X9cZTxgupD4py7/Pps4qeuwiWIsbESoDDxRsuivhh1xka8lfJHoPDMwdV7DNjRFUUFpJrDr7KYp5zGRFU9hIUfh8YzZ0lQTAzboQyUwMoTRRiUOM5hs/2/RG6VA1gPaeqRaE6v0nphHTZ6By3Zvs5tj9qh6iW07jtXTxXk0MDzNrQpMh2MUvPB0dikKjDMxgUKFGEiDKvFilZJ+y0ErfdFekBn+mfInr0Q==-MIIETDCCAjSgAwIBAgIBDzANBgkqhkiG9w0BAQsFADAYMRYwFAYDVQQDDA1KZXRQcm9maWxlIENBMB4XDTIyMTAxMDE2MDU0NFoXDTI0MTAxMTE2MDU0NFowHzEdMBsGA1UEAwwUcHJvZDJ5LWZyb20tMjAyMjEwMTAwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC/W3uCpU5M2y48rUR/3fFR6y4xj1nOm3rIuGp2brELVGzdgK2BezjnDXpAxVDw5657hBkAUMoyByiDs2MgmVi9IcqdAwpk988/Daaajq9xuU1of59jH9eQ9c3BmsEtdA4boN3VpenYKATwmpKYkJKVc07ZKoXL6kSyZuF7Jq7HoQZcclChbF75QJPGbri3cw9vDk/e46kuzfwpGftvl6+vKibpInO6Dv0ocwImDbOutyZC7E+BwpEm1TJZW4XovMBegHhWC04cJvpH1u98xoR94ichw0jKhdppywARe43rGU96163RckIuFmFDQKZV9SMUrwpQFu4Z2D5yTNqnlLRfAgMBAAGjgZkwgZYwCQYDVR0TBAIwADAdBgNVHQ4EFgQU5FZqQ4gnVc+inIeZF+o3ID+VhcEwSAYDVR0jBEEwP4AUo562SGdCEjZBvW3gubSgUouX8bOhHKQaMBgxFjAUBgNVBAMMDUpldFByb2ZpbGUgQ0GCCQDSbLGDsoN54TATBgNVHSUEDDAKBggrBgEFBQcDATALBgNVHQ8EBAMCBaAwDQYJKoZIhvcNAQELBQADggIBANLG1anEKid4W87vQkqWaQTkRtFKJ2GFtBeMhvLhIyM6Cg3FdQnMZr0qr9mlV0w289pf/+M14J7S7SgsfwxMJvFbw9gZlwHvhBl24N349GuthshGO9P9eKmNPgyTJzTtw6FedXrrHV99nC7spaY84e+DqfHGYOzMJDrg8xHDYLLHk5Q2z5TlrztXMbtLhjPKrc2+ZajFFshgE5eowfkutSYxeX8uA5czFNT1ZxmDwX1KIelbqhh6XkMQFJui8v8Eo396/sN3RAQSfvBd7Syhch2vlaMP4FAB11AlMKO2x/1hoKiHBU3oU3OKRTfoUTfy1uH3T+t03k1Qkr0dqgHLxiv6QU5WrarR9tx/dapqbsSmrYapmJ7S5+ghc4FTWxXJB1cjJRh3X+gwJIHjOVW+5ZVqXTG2s2Jwi2daDt6XYeigxgL2SlQpeL5kvXNCcuSJurJVcRZFYUkzVv85XfDauqGxYqaehPcK2TzmcXOUWPfxQxLJd2TrqSiO+mseqqkNTb3ZDiYS/ZqdQoGYIUwJqXo+EDgqlmuWUhkWwCkyo4rtTZeAj+nP00v3n8JmXtO30Fip+lxpfsVR3tO1hk4Vi2kmVjXyRkW2G7D7WAVt+91ahFoSeRWlKyb4KcvGvwUaa43fWLem2hyI4di2pZdr3fcYJ3xvL5ejL3m14bKsfoOv