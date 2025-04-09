
[toc]

# 1 docker简介
https://docker.easydoc.net/doc/81170005/cCewZWoN/lTKfePfP
**简介**：
Docker 是一个应用打包、分发、部署的工具
你也可以把它理解为一个轻量的虚拟机，它只虚拟你软件需要的运行环境，多余的一点都不要，
而普通虚拟机则是一个完整而庞大的系统，包含各种不管你要不要的软件。

**跟普通虚拟机的对比**：
|特性|	普通虚拟机|	Docker|
|----|-----------|-------|
|跨平台	|通常只能在桌面级系统运行，例如 Windows/Mac，无法在不带图形界面的服务器上运行	|支持的系统非常多，各类 windows 和 Linux 都支持|
|性能|	性能损耗大，内存占用高，因为是把整个完整系统都虚拟出来了	|性能好，只虚拟软件所需运行环境，最大化减少没用的配置|
|自动化|	需要手动安装所有东西	|一个命令就可以自动部署好所需环境|
|稳定性	|稳定性不高，不同系统差异大	|稳定性好，不同系统都一样部署方式|

**打包、分发、部署**
- <font color=red>打包</font>：就是把你软件运行所需的依赖、第三方库、软件打包到一起，变成一个安装包
- <font color=red>分发</font>：你可以把你打包好的“安装包”上传到一个镜像仓库，其他人可以非常方便的获取和安装
- <font color=red>部署</font>：拿着“安装包”就可以一个命令运行起来你的应用，自动模拟出一摸一样的运行环境，不管是在 Windows/Mac/Linux。

**Docker 部署的优势**:
常规应用开发部署方式：自己在 Windows 上开发、测试 --> 到 Linux 服务器配置运行环境部署。

>> 问题：我机器上跑都没问题，怎么到服务器就各种问题了

用 Docker 开发部署流程：自己在 Windows 上开发、测试 --> 打包为 Docker 镜像（可以理解为软件安装包） --> 各种服务器上只需要一个命令部署好

>> 优点：确保了不同机器上跑都是一致的运行环境，不会出现我机器上跑正常，你机器跑就有问题的情况。

例如 易文档，SVNBucket 的私有化部署就是用 Docker，轻松应对客户的各种服务器。

**:Docker 通常用来做什么**:
- 应用分发、部署，方便传播给他人安装。特别是开源软件和提供私有部署的应用
- 快速安装测试/学习软件，用完就丢（类似小程序），不把时间浪费在安装软件上。例如 Redis / MongoDB / ElasticSearch / ELK
- 多个版本软件共存，不污染系统，例如 Python2、Python3，Redis4.0，Redis5.0
- Windows 上体验/学习各种 Linux 系统

**重要概念：镜像、容器**：
- <font color=red>镜像</font>：可以理解为软件安装包，可以方便的进行传播和安装。
- <font color=red>容器</font>：软件安装后的状态，每个软件运行环境都是独立的、隔离的，称之为容器。


# 2 docker安装
桌面版：https://www.docker.com/products/docker-desktop
服务器版：https://docs.docker.com/engine/install/#server

## 2.1 Ubuntu上安装docker
**教程网址**:
https://docs.docker.com/desktop/install/ubuntu/

- 在以上网址下载好docker的deb包后在同级目录下执行
```shell
sudo apt-get update
sudo apt install docker.io
sudo apt install docker-compose
```

## 2.2 docker镜像加速源

|镜像加速器|	镜像加速器地址|
|---------|-----------------|
|Docker 中国官方镜像	|https://registry.docker-cn.com|
|DaoCloud 镜像站	|http://f1361db2.m.daocloud.io|
|Azure 中国镜像|	https://dockerhub.azk8s.cn|
|科大镜像站	|https://docker.mirrors.ustc.edu.cn|
|阿里云|	https://ud6340vz.mirror.aliyuncs.com|
|七牛云	|https://reg-mirror.qiniu.com|
|网易云|	https://hub-mirror.c.163.com|
|腾讯云|	https://mirror.ccs.tencentyun.com|


## 2.3 快速启动一个docker容器
**docker的镜像仓库**：https://hub.docker.com/

```shell
# 启动一个redis
sudo docker run -d -p 6379:6379 --name redis redis:latest
```


# 3 docker命令
命令参考文档
https://docs.docker.com/engine/reference/commandline/run/
## 3.1 运行容器
```shell
docker run -d -p 6380:6379 --name redis redis:latest
```
1. -d标识在后台运行
2. -p 6380:6379标识将docker中的6379端口映射到宿主机的6380端口
3. --name redis标识为容器奇异果名字
4. redis:latest表示运行最新版本的redis

## 3.2 停止/启动指定 id 的容器
```shell
docker stop/start <容器ID>
```

## 3.3 查看有哪些doker容器
docker  ps -a
docker  ps  # 查看正在运行的docker容器吧

## 3.4 进入容器的终端(-i表示交互式， -t表示分配一个伪终端)
docker exec -it <容器id或容器名> [bash,sh]

**示例**：
```shell
docker exec -it 124eac100 bash
```

## 3.5 退出容器终端
exit

## 3.6 重启容器
docker restart <容器id>

## 3.7 编译为镜像
docker build -t <容器名>:<版本号> -f <Dockerfile的路径>  <Dockerfile所在目录>

```shell
docker build -t wallet_admin -f ./Dockerfile .
```


**docker build的参数**:

- `-t`: 设置镜像名字和版本号
- `--build-arg DOCKER_BUILDKIT=1` 启用DockerKit。BuildKit 是 Docker 的下一代构建引擎，它提供了更多高级功能和优化，包括对缓存挂载的支持。包含像--mount选项

## 3.8 其他docker命令
```shell
# 查看本地镜像列表
docker images

# 查看volume 列表
docker volume ls

# 查看网络列表
docker network ls

# 删除指定 id 的镜像
docker rmi image-id

# 删除指定 id 的容器
docker rm container-id
```

# 4 docker-compose命令
如果项目依赖更多的第三方软件，我们需要管理的容器就更加多，每个都要单独配置运行，指定网络。
这节，我们使用 docker-compose 把项目的多个服务集合到一起，一键运行。

```shell
# 终端启动doker
docker-compose up

# 后台启动
docker-compose up -d

# 停止容器并移除容器、网络、卷和镜像
docker-compose down

# 查看运行状态
docker-compose ps

# 重启
docker-compose restart

# 重启单个服务
docker-compose restart service-name

# 停止运行
docker-compose stop

# 进入容器命令行
docker-compose exec service-name sh

# 查看容器运行log
docker-compose logs [service-name]
```

# 5 Dockerfile语法

**编写Dockerfile**:
```dockerfile
# 以node11为基础镜像
FROM node:11

# 维护者信息
MAINTAINER easydoc.net

# 复制代码，将宿主机的当前目录的内容复制到docker 容器中的/app目录下，如果/app目录不存在会创建
ADD . /app

# 设置容器启动后的默认运行目录
WORKDIR /app

# 运行命令，安装依赖
# RUN 命令可以有多个，但是可以用 && 连接多个命令来减少层级。
# 例如 RUN npm install && cd /app && mkdir logs
RUN npm install --registry=https://registry.npm.taobao.org

# CMD 指令只能一个，是容器启动后执行的命令，算是程序的入口。
# 如果还需要运行其他命令可以用 && 连接，也可以写成一个shell脚本去执行。
# 例如 CMD cd /app && ./start.sh
CMD node app.js
```

**Build为镜像（安装包）和运行
编译 `docker build -t test:v1 .`

-t 设置镜像名字和版本号
命令参考：https://docs.docker.com/engine/reference/commandline/build/

运行 `docker run -p 8080:8080 --name test-hello test:v1`

-p 映射容器内端口到宿主机
--name 容器名字
-d 后台运行
命令参考文档：https://docs.docker.com/engine/reference/run/

## 5.1 FROM字段
Dockerfile中FROM字段用于构建容器基础镜像,在Dockerfile中必须存在，基础镜像是新 Docker 镜像的起点，提供了操作系统层面的支持以及一些基础的工具和库。
其**基本语法**为 `FROM <image>[:<tag>]`

**示例**：
```Dockerfile
# 以golang:1.21基于 Alpine Linux 的版本为基础镜像，并该阶段命名为builder
#这种写法允许 Dockerfile 在后续的阶段使用 COPY --from=builder ... 从这个名为 builder 的阶段复制构建产物或其他文件。这是优化 Docker 镜像构建流程的一种非常有效的方式。通过这种方法，可以在初期阶段使用较大的基础镜像进行编译和构建任务，而后续阶段则可以使用更加精简的镜像运行应用程序，从而减小最终产品的镜像大小，加快部署速度并提高运行时效率。
FROM golang:1.21-alpine as builder
```

## 5.2 MAINTAINER字段
这个字段用于描述维护者信息
**基本语法**: `MAINTAINER <维护者名字>`

## 5.3 RUN字段
在 Dockerfile 中，RUN 指令用于在构建镜像时执行命令。这些命令在当前镜像层上执行，并且它们的结果会被用于创建一个新的镜像层。最终构建出的镜像将包含所有 RUN 指令执行后的结果，这使得 RUN 成为在 Docker 镜像构建过程中安装软件包、运行脚本或执行其他配置任务的核心指令

**示例1**
```Dockerfile
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o wallet_admin main.go
```

在 Dockerfile 中使用 `RUN --mount=type=cache,target="/root/.cache/go-build" go build -o wallet_admin main.go` 这条语句是利用 Docker 的构建缓存功能，特别是在构建 Go 程序时优化构建过程和时间。下面详细解析这条命令的各个部分和它们的作用：

1. **`RUN`**: 这是 Dockerfile 的指令，用于在构建镜像的过程中执行命令。

2. **`--mount=type=cache,target="/root/.cache/go-build"`**:
   - `--mount`: 这个选项用于在执行 `RUN` 命令时挂载指定类型的存储。在这个场景中，它是用来挂载一个缓存。
   - `type=cache`: 这指定挂载类型为缓存，这意味着这个挂载的目的是用来存储构建过程中可以重用的数据，以加快后续构建的速度。
   - `target="/root/.cache/go-build"`: 这定义了缓存挂载的目标位置，即容器内部的路径。对于 Go 程序，`/root/.cache/go-build` 是 Go 编译器存储临时编译文件的默认位置。通过缓存这些文件，可以在多次构建间共享编译缓存，从而减少重复编译相同代码的时间。

3. **`go build -o wallet_admin main.go`**:
   - `go build`: 这是 Go 语言的编译命令，用于编译 Go 程序。
   - `-o wallet_admin`: 这个参数指定编译输出的文件名。在这里，编译后生成的可执行文件将命名为 `wallet_admin`。
   - `main.go`: 这是 Go 程序的入口文件，`go build` 将从这个文件开始编译整个应用。

综上所述，这条 `RUN` 命令在 Docker 构建过程中，通过缓存 Go 的编译缓存来优化和加速构建过程。这不仅节省了重复编译的时间，也使得整个 Docker 镜像构建过程更为高效。这种方法尤其适用于有大量依赖或大型项目的情况，其中编译过程可能非常耗时。通过缓存编译产物，可显著减少从源代码到可运行镜像的构建时间。

**示例2**:
```Dockerfile
RUN apk add --no-cache make gcc musl-dev linux-headers
```

# 6 目录挂载

**现存问题**：
使用 Docker 运行后，我们改了项目代码不会立刻生效，需要重新build和run，很是麻烦。
容器里面产生的数据，例如 log 文件，数据库备份文件，容器删除后就丢失了。

**几种挂载方式**
- bind mount。 直接把宿主机目录映射到容器内，适合挂代码目录和配置文件。可挂到多个容器上。可以通过使用-v参数
``` -v <宿主机目录>:<容器内部的目录>```

```shell
# 运行容器test的v1版本，并将容器的8080端口映射到宿主机的9090端口，给运行的容器起名为testName. 并将宿主机的D:/code目录(绝对路径)挂载到容器的/app目录下， 使用-d指定在后台运行
# 通过这种方式挂载的后，宿主机挂载目录下的修改和容器内的修改会同步
docker run -p 9090:8080 -name testName -v D:/code:/app -d test:v1
```

- volume 由容器创建和管理，创建在宿主机，所以删除容器不会丢失，官方推荐，更高效，Linux 文件系统，适合存储数据库数据。可挂到多个容器上
- tmpfs mount 适合存储临时文件，存宿主机内存中。不可多容器共享。

# 7 多容器通信
项目往往都不是独立运行的，需要数据库、缓存这些东西配合运作。
如果要想多容器之间互通，从 Web 容器访问 Redis 容器，我们只需要把他们放到同个网络中就可以了。
参考文档：https://docs.docker.com/engine/reference/commandline/network/

**演示**：
1. 创建一个名为test-net的网络：
```docker network create test-net```

2. 运行 Grpc 在 test-net 网络中，别名Grpc
```docker run -d --name market --network test-net --network-alias Grpc-alias market:latest```

3. 修改代码中访问redis的地址为网络别名
```go
// 使用网络别名来作为url
connect, err = grpc.Dial("Grpc-alias:1080", transport, grpc.WithDialer(dialer))
if err != nil {
    log.Fatal("connect wallet-chain-node rpc server failed:", err.Error())
}
```

4. 运行 test 项目，使用同个网络
```shell
docker run -p 8080:8080 --name test -v D:/test:/app --network test-net -d test:v1
```