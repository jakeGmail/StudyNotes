[toc]
# 1 Redis服务配置
<span id="lable"></span>
redis配置文件写好后可以通过redis-server后接配置文件名称来启动redis。启动的方式就很具配置文件中的内容来。
redis-server启动方式参考 [Redis服务启动方式](1_Redis介绍.md#12-启动Redis服务)

## 1.1 基本配置信息
### 1.1.1 启动方式
<table><tr><td bgcolor=Gray></br>


```shell
# 启动Redis服务的IP
bind 127.0.0.1

# Redis服务的端口号
# 如果绑定了主机的IP，就只能通过这个IP来访问。
# 如果不绑定IP，则别人访问这个redis服务的时候，可以通过别的IP来访问（只要它能标识到这台主机）。
port 6380

# 是否以守护进程的方式启动。
# no：在终端启动后，会在启动终端打日志。
# yes：在后台启动了，不会打印日志到终端。
daemonize yes/on

# redis启动后的工作目录，后面redis运行过程中产生的文件就在这里存放
dir /home/jake/working_redis/
```
</td></tr></table>

### 1.1.2 日志相关设置
<table><tr><td bgcolor=Gray></br>

```shell
# 以守护进程启动后(deamonize yes)，redis的日志会输出到$(dir)/redis_6380_log
logfile redis_6380_log

# 设置服务器的日志级别
# 日志级别在开发期设置为verbose即可。生产环境中配置notice以简化日志输出量，降低写日志的IO频度
loglevel debug|verbose|notice|warning
```
</td></tr></table>

### 1.1.3 设置数据库数量
<table><tr><td bgcolor=Gray></br>

```shell
# 设置redis数据库的个数
# 如果不设置值databases, 默认的数据库个数也是16
database 16
```
</td></tr></table>

### 1.1.4 客户端配置
<table><tr><td bgcolor=Gray></br>

```shell
# 设置同一时间最大客户端连接数，默认无限制（设置为0）。
# 当客户端连接数达到上限时，Redis服务会关闭新的连接
maxclients 0

# 客户端闲置等待最大时长（单位秒），达到最大值后关闭连接
# 关闭该功能设置为0
timeout 300
```
</td></tr></table>

### 1.1.5 包含其他配置文件
<table><tr><td bgcolor=Gray></br>

```shell
# 导入并加载指定配置文件信息。
# 用于快速创建公共配置较多的redis实例文件。先写公共配置信息，然后根据业务再include不同的配置信息。
include /home/jake/Redis/conf/redis_6380_config2.conf

# 是否开启保护模式，保护模式下只接受来自loopback接口的连接
protected-mode yes/no
```
</td></tr></table>

## 1.2 持久化相关配置
<table><tr><td bgcolor=Black></br>

```shell
# 设置本地数据库文件名，该文件保存RDB数据，默认值为dump.rdb. 通常设置为dunp-端口号.rdb
# 这个文件会放在dir配置指定的文件夹下
dbfilename  dump.rdb

# 设置存储至本地数据库shivering是否压缩数据，默认为yes,采用LZF压缩
# 经验：通常认为开启状态，如果设置为no,可以节省CPU运行时间，但会使存储文件变大(巨大)
rdbcompression yes/no

# 设置是否进行RDB文件格式校验，该校验过程再写文件的读文件均进行
# 经验：通常为开启状态，如果设置为no，可以节约读写过程的约10%时间消耗，但是存储数据有一定损坏风险，
rdbchecksum yes/no

# 在使用bgsave指令后，在后台存储过程中如果出现错误现象，是否应该停止保存操作。
# 经验: 通常默认为开启
stop-writes-on-bgsave-error yes/no
```
</td></tr></table>

## 1.3 影响逐出算法的配置
<table><tr><td bgcolor="Gray"></br>

```shell
# 设置最大可使用内存,单位字节。
# 占用物理内存的比例，默认值为0，表示不限制。生产环境中根据需求设定，通常设置在50%以上，当内存占用>=设置的这个值时就触发逐出算法(如果开启逐出算法)。
maxmemory <bytes>

# 进行数据逐出时，选取待删除数据的个数
# 选取数据时并不会全库扫描，导致严重的性能消耗，降低读写性能。因此采用随机获取数据的方式作为待检测删除数据
maxmemory-sample num

# 指定删除策略
maxmemory-policy no-eviction|volatile-lru|volatile-lfu|volatile-ttl|volatile-random|allkeys-lru|allkeys-lfu|allkeys-random
```
</td></tr></table>


 [点击查看删除策略详细](7_Redis删除策略.md)

## 1.4 定期删除配置
<table><tr><td bgcolor=Gray></br>

    ```shell
    # 定期删除策略中每秒会执行hz次serverCron()函数
    hz 10                        
    ```
</td></tr></table>
