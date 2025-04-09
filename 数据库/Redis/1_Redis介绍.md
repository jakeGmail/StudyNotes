[toc]
# 1.1 Redis安装
1. **_下载_**
   ```shell
   wget http://download.redis.io/releases/redis-5.0.0.tar.gz

   wget http://download.redis.io/releases/redis-6.0.0.tar.gz 

   wget http://download.redis.io/releases/redis-7.2.3.tar.gz 
   ```
   下载redis6.0.0的安装包
2. 解压
   ```shell
   tar -xvf redis-6.0.0.tar.gz
   ```
3. 安装
   ```shell
   sudo make install
   ```
4. 检测
   ```shell
   redis-server -v
   ```
   查看Redis服务的版本，如果能查看成功说明Redis安装成功。

# 1.2 启动Redis服务
## 1.2.1 默认配置启动
```shell
redis-server
```
执行`redis-server`命令后redis服务以默认方式启动，端口号6379，IP为127.0.0.1。
## 1.2.2 指定端口启动Redis服务
```shell
redis-server --port 1234 # 在指定端口1234上运行Redis服务
```
## 1.2.3 指定配置文件启动Reids服务
```shell
redis-server conf/redis-6380.conf
```
根据路径为`conf/redis-6380.conf`的配置文件来启动Redis服务

-------------------------------------------------------
# 1.3 Redis客户端启动
## 1.3.1 默认方式启动客户端
```shell
redis-cli
```
以默认方式启动Redis客户端，默认会去连接127.0.0.1:6375的Redis服务。如果对应的Redis服务没有启动就会连接失败。
## 1.3.2 指定IP:port的方式启动Reids客户端
```shell
redis-cli -h 127.0.0.1 -p 3456
```
启动Redis客户端，并请求连接IP为127.0.0.1，端口号为3456的Redis服务。
