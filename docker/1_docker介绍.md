[toc]

# 1 docker安装

## 1.1 windows安装
1. Docker Desktop 官方下载地址： https://docs.docker.com/desktop/install/windows-install/

2. 安装之后，可以打开 PowerShell 并运行以下命令检测是否运行成功： `docker run hello-world`

**方式2**
1. 通过https://docs.docker.com/desktop/install/windows-install/下载docker安装包

## 1.2 ubuntu安装

**使用官方脚本自动安装**:
```shell
curl -fsSL https://test.docker.com -o test-docker.sh
sudo sh test-docker.sh
```

**使用 Docker 仓库进行安装**：
在新主机上首次安装 Docker Engine-Community 之前，需要设置 Docker 仓库。之后，您可以从仓库安装和更新 Docker 。

设置仓库
更新 apt 包索引。
```shell
$ sudo apt-get update
```
安装 apt 依赖包，用于通过HTTPS来获取仓库:
```shell
$ sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
```
添加 Docker 的官方 GPG 密钥：
```shell
$ curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
```
9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88 通过搜索指纹的后8个字符，验证您现在是否拥有带有指纹的密钥。
```shell
$ sudo apt-key fingerprint 0EBFCD88
   
pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```
使用以下指令设置稳定版仓库
```shell
$ sudo add-apt-repository \
   "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/ \
  $(lsb_release -cs) \
  stable"
```
安装 Docker Engine-Community
更新 apt 包索引。
```shell
$ sudo apt-get update
```
安装最新版本的 Docker Engine-Community 和 containerd ，或者转到下一步安装特定版本：

```shell
$ sudo apt-get install docker-ce docker-ce-cli containerd.io
```
要安装特定版本的 Docker Engine-Community，请在仓库中列出可用版本，然后选择一种安装。列出您的仓库中可用的版本：
```shell
$ apt-cache madison docker-ce

  docker-ce | 5:18.09.1~3-0~ubuntu-xenial | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  docker-ce | 5:18.09.0~3-0~ubuntu-xenial | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  docker-ce | 18.06.1~ce~3-0~ubuntu       | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  docker-ce | 18.06.0~ce~3-0~ubuntu       | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  ...
```
使用第二列中的版本字符串安装特定版本，例如 5:18.09.1~3-0~ubuntu-xenial。
```shell
$ sudo apt-get install docker-ce=<VERSION_STRING> docker-ce-cli=<VERSION_STRING> containerd.io
```
测试 Docker 是否安装成功，输入以下指令，打印出以下信息则安装成功:

```shell
$ sudo docker run hello-world

Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
1b930d010525: Pull complete                                                                                                                                  Digest: sha256:c3b4ada4687bbaa170745b3e4dd8ac3f194ca95b2d0518b417fb47e5879d9b5f
Status: Downloaded newer image for hello-world:latest


Hello from Docker!
This message shows that your installation appears to be working correctly.


To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.


To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash


Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/


For more examples and ideas, visit:
 https://docs.docker.com/get-started/

```
使用 Shell 脚本进行安装
Docker 在 get.docker.com 和 test.docker.com 上提供了方便脚本，用于将快速安装 Docker Engine-Community 的边缘版本和测试版本。脚本的源代码在 docker-install 仓库中。 不建议在生产环境中使用这些脚本，在使用它们之前，您应该了解潜在的风险：

脚本需要运行 root 或具有 sudo 特权。因此，在运行脚本之前，应仔细检查和审核脚本。

这些脚本尝试检测 Linux 发行版和版本，并为您配置软件包管理系统。此外，脚本不允许您自定义任何安装参数。从 Docker 的角度或您自己组织的准则和标准的角度来看，这可能导致不支持的配置。

这些脚本将安装软件包管理器的所有依赖项和建议，而无需进行确认。这可能会安装大量软件包，具体取决于主机的当前配置。

该脚本未提供用于指定要安装哪个版本的 Docker 的选项，而是安装了在 edge 通道中发布的最新版本。

如果已使用其他机制将 Docker 安装在主机上，请不要使用便捷脚本。

本示例使用 get.docker.com 上的脚本在 Linux 上安装最新版本的Docker Engine-Community。要安装最新的测试版本，请改用 test.docker.com。在下面的每个命令，取代每次出现 get 用 test。
```shell
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sudo sh get-docker.sh
```
如果要使用 Docker 作为非 root 用户，则应考虑使用类似以下方式将用户添加到 docker 组：
```shell
$ sudo usermod -aG docker your-user
```
卸载 docker
删除安装包：
```shell
sudo apt-get purge docker-ce
```
删除镜像、容器、配置文件等内容：
```shell
sudo rm -rf /var/lib/docker
```

# 2 docker中的概念
## 2.1 镜像（Image）
镜像是一种轻量级、可执行的独立软件包，它包含运行某个软件所需要的全部内容，我们把应用程序和配置依赖打包好形成一个可交付的运行环境（包含代码、运行时需要的库、环境变量和配置文件），这个打包好的运行环境就是imgage镜像文件

**镜像分层**：
docker镜像层都是只读的，容器层是可写的。当容器启动时，一个新的可写层被加载到镜像的顶部，这一层通常被称作“容器层”，“容器层”之下的都叫做“镜像层”

**联合文件系统**：
联合文件系统（Union File System）是一种文件系统服务，它可以将多个不同的文件系统挂载到同一个虚拟文件系统下，并对外表现为单一一致的文件系统。这种机制通过创建层（layers）来实现，每一层都可以存放文件和目录，而上层的文件会覆盖下层相同路径和文件名的文件。这种覆盖只是视觉上的覆盖，实际物理文件并未修改，从而实现了非破坏性编辑。

联合文件系统在Docker中得到了广泛应用，尤其适合于环境虚拟化和容器化技术。在Docker的应用中，每一个image或者说是容器镜像都可以视作是一系列的只读层，当构建或者启动容器的时候，Docker会在最上层添加一个可写层。这意味着：

- **基础层保持不变**：所有基础的操作系统文件和依赖库都存放在这些只读层上，它们被多个容器共享，这样可以减少硬盘占用和提高容器启动的速度。
- **个性化设置和数据保存在可写层**：当容器运行时，对文件系统的任何修改都仅发生在顶层的可写层上。这样即使容器被销毁，这些修改也不会影响到基础层。

联合文件系统的这种设计使得Docker镜像非常适合用于快速部署、轻量化以及高效的环境隔离。通过利用已有的只读层，可以极大地提高资源复用率并减少存储空间的消耗。

## 2.2 容器
镜像（Image）和容器（Container）的关系，就像是面向对象程序设计中的类和实例一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等。
容器是独立运行的一个或一组应用，是镜像运行时的实体。

## 2.3 仓库
仓库可看成一个代码控制中心，用来保存镜像。

## 2.4 镜像，容器，仓库的关系
- Docker 使用客户端-服务器 (C/S) 架构模式，使用远程API来管理和创建Docker容器。
- Docker 容器通过 Docker 镜像来创建。

## 2.5 Docker Registry
Docker 仓库用来保存镜像，可以理解为代码控制中的代码仓库。

Docker Hub(https://hub.docker.com) 提供了庞大的镜像集合供使用。

一个 Docker Registry 中可以包含多个仓库（Repository）；每个仓库可以包含多个标签（Tag）；每个标签对应一个镜像。

通常，一个仓库会包含同一个软件不同版本的镜像，而标签就常用于对应该软件的各个版本。我们可以通过 <仓库名>:<标签> 的格式来指定具体是这个软件哪个版本的镜像。如果不给出标签，将以 latest 作为默认标签。

## 2.6 Docker Machine
Docker Machine是一个简化Docker安装的命令行工具，通过一个简单的命令行即可在相应的平台上安装Docker，比如VirtualBox、 Digital Ocean、Microsoft Azure。