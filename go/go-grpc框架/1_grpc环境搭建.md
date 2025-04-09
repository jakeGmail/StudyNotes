# 1 安装protobuf
1. 安装通用的编译器
地址： https://github.com/protocolbuffers/protobuf/releases
下载对应系统的protoc
`git clone https://github.com/protocolbuffers/protobuf/releases/latest`
将protoc执行程序cp到环境PATH中（/usr/bin）下，方便在终端调用
安装好protoc后在终端输入
    ```shell
    protoc --version
    ```
    如果输出protoc版本信息说明安装成功
2. 安装go专用的protoc的生成器
    ```shell
    go install github.com/golang/protobuf/protoc-gen-go@latest
    ```
    安装后会在GOPATH目录下生成可执行文件，protobuf的编译器插件protoc-gen-go
    执行protoc命令会自动调用插件
    <font color=gray>(注意在使用在使用该命令时，如果不成功，就在包含go.mod的路径下执行)</front>

    ```shell
    # 安装最新版本的protoc
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    ```
    安装好的protoc包存放在$GO_PATH/pkg/mod/google.golang.org/protobuf
3. 设置环境变量
```sudo /etc/profile```
然后添加下面的命令进去，用于寻找protoc-gen-go程序，生成go代码
```shell
export GOPATH=$HOME/go 
export PATH=$PATH:$GOPATH/bin
```
