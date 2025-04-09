# 1 逐出算法
当新的数据进入redis时，如果内存不足应该怎么办呢？
- Redis使用内存存储数据，在执行每一个命令前，会调用<font color=red>freeMemoryIfNeeded()</font>检测内存是否充足。如果内存不满足新加入数据的最低存储要求， redis要临时删除一些数据为当前指令清理存储空间。清理数据的策略称为逐出算法。在进行数据删除前会按照一定策略选取一些待删除的数据样本，判断它们是否需要被逐出，选取的个数由配置文件中的maxmemory-sample指定。
**注意**：逐出数据的过程不是100%能够清理出足够的可使用的内存空间，如果不成功则反复执行。当对所有数据尝试完毕后，如果不能达到内存清理的要求，将出现错误信息。
    ```
    (error) OOM command not allowed when used memory >'maxmemory'
    ```
## 1.1 逐出算法的过程
1. 当需要逐出一些数据时，会从数据库中挑选若干条数据用于检测，看是否需要将其临时删除。
2. 检测这些待删除数据需要使用一定的策略来检测是否需要删除。检测的策略包含：
   - 对检测易失性数据（可能会过期的数据集server.db[i].expires）
    (1) <font color=blue>volatile-lru</font>: 挑选最久没有使用的数据淘汰（least recentlu used）
    (2) <font color=blue>volatile-lfu</font>：挑选使用次数最少的数据淘汰(least frequently used)
    (3) <font color=blue>volatile-ttl</font>: 挑选要过期的数据淘汰，通过ttl查看过期时间，最近要过期的淘汰掉。
    (4) <font color=blue>volatile-random</font>: 任意选择数据淘汰
    - 对全库数据
    (1) <font color=green>allkeys-lru</font>: 挑选最久没有使用的数据淘汰
    (2) <font color=green>allkeys-lfu</font>: 挑选使用次数最少的数据淘汰
    (3) <font color=green>allkeys-random</font>: 任意选择数据淘汰
    - 放弃数据逐出策略
    <font color=red>no-enviction</font>: 禁止逐出数据（redis4.0默认策略），当内存满时会引发错误（Out Of Memory）</br>

这些策略配置在配置文件中memory-policy属性配置,详见[影响逐出算法的配置](#2-影响逐出算法的配置)



# 2 影响逐出算法的配置
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

# 3 逐出策略配置依据
使用```info stats```命令输出监控信息. 在输出的信息中keyspace_hits属性代表执行获取key的命令命中次数，keyspace_misses代表查询丢失次数。可以通过这两值的运行期间的值来调整策略配置。