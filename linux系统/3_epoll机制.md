[toc]
# 1 epoll背景
在linux的网络编程中，很长的时间都在使用select来做事件触发。在linux新的内核中，有了一种替换它的机制，就是epoll。
相比于select，epoll最大的好处在于它不会随着监听fd数目的增长而降低效率。因为在内核中的select实现中，它是采用轮询来处理的，轮询的fd数目越多，自然耗时越多。并且，在linux/posix_types.h头文件有这样的声明：
\#define __FD_SETSIZE    1024
表示select最多同时监听1024个fd，当然，可以通过修改头文件再重编译内核来扩大这个数目，但这似乎并不治本。

epoll机制是linux下的一个多路复用机制，多路是指多个业务方（句柄）并发下来的 IO ；复用是指复用这一个后台处理程序.

对此linux内核提供了3中多路复用工具： select、poll、epoll。通过历史的改进select-->poll-->epoll. 其中epoll效率最高
**select**:

epoll的接口非常简单，一共就三个函数
```c++
// 创建epoll池
int epoll_create(int size);

// 向epoll句柄中添加/修改/删除 监听内容
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);

// 等待监听的时间发生
int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
```

# 2  epoll API
使用epoll API需要包含有文件 ```#include <sys/epoll.h>```

## 2.1 eopll_create
**函数原型:**
```c++
// 创建epoll池
int epoll_create(int size);
```
**参数：**
- size: 设置epoll池监听的fd的大致数量，但在linux内核2.6.8及其之后，这个参数就没有任何意义了，有内核来管理监听fd的数量，但**还是需要大于0, 以兼容新的epoll代码在旧的内核版本的运行**。

**返回值：**
返回创建的epoll示例的句柄(文件描述符)---在不用的时候需要调用close关闭，不然可能导致fd消耗完。 
创建失败就返回0并且设置errno.

## 2.2 epoll_ctl
**函数原型**:
```c
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
```
**参数**:
- **<font color=green>epfd：</font>** epoll句柄，由```epoll_create```创建
- **<font color=green>op：</font>** epoll操作选项，用于epoll的增加、修改、删除监听的文件描述符。
  **EPOLL_CTL_ADD** : 在文件描述符epfd引用的epoll实例上注册目标文件描述符fd，并将事件event与fd链接的内部文件相关联
- **<font color=green>fd：</font>** 要进行操作的文件描述符
- **<font color=green>event：</font>** 描述要对fd文件描述符进行监听的事件。
epoll_event的定义如下：
```c
typedef union epoll_data {
    void        *ptr;
    int          fd;
    uint32_t     u32;
    uint64_t     u64;
} epoll_data_t;

struct epoll_event {
    uint32_t     events;      /* Epoll events */
    epoll_data_t data;        /* User data variable */
};
```
其中epoll_event的成员events定义监听的事件：
**EPOLLIN**： 监听写事件
**EPOLLOUT**: 监听读事件
**EPOLLRDHUP**: (Linux 2.6.17版本支持)
**返回值**
成功返回0，失败返回-1并设置errno

## 2.3 epoll_wait
**函数原型**
```c
int epoll_wait(int epfd, struct epoll_event *events,int maxevents, int timeout);
```
**描述**
等待监听的事件发生，如果

**参数**：
- **<font color=green>epfd:</font>**  epoll池文件描述符
- **<font color=green>event:</font>** 监听到的事件信息。
- **<font color=green>maxevents:</font>** 
- **<font color=green>timeout:</font>** 超时时间，阻塞超过这个时间就会返回。单位毫秒

# 3 epoll底层原理
## 3.1 添加到epoll池的fd的存储结构
添加到epoll池中的fd是以红黑树存储的。
## 3.2 epoll高效及时同时fd事件的原理
Linux 设计成一切皆是文件的架构，这个不是说说而已，而是随处可见。实现一个文件系统的时候，就要实现这个文件调用，这个结构体用 struct file_operations 来表示。如下：
```c++
struct file_operations {
ssize_t (*read) (struct file *, char __user *, size_t, loff_t *);
ssize_t (*write) (struct file *, const char __user *, size_t, loff_t *);
__poll_t (*poll) (struct file *, struct poll_table_struct *);
int (*open) (struct inode *, struct file *);
int (*fsync) (struct file *, loff_t, loff_t, int datasync);
// ....
};
```
其中的```poll```方法，是定制监听事件的实现。通过```poll```让上层能直接告诉底层，我这个 fd 一旦读写就绪了，请底层硬件（比如网卡）回调的时候自动把这个 fd 相关的结构体放到指定队列中，并且唤醒操作系统。
举个例子：网卡收发包其实走的异步流程，操作系统把数据丢到一个指定地点，网卡不断的从这个指定地点掏数据处理。请求响应通过中断回调来处理，中断一般拆分成两部分：硬中断和软中断。poll 函数就是把这个软中断回来的路上再加点料，只要读写事件触发的时候，就会立马通知到上层，采用这种事件通知的形式就能把浪费的时间窗就完全消失了。

因此**这个 poll 事件回调机制则是 epoll 池高效最核心原理**
也就是说，只有实现了poll这个回调的fd才能使用epoll机制-----epoll仅仅支持pipe、socket、driver、等文件的操作，不支持普通文件操作。如果对应的fd没有实现那么在调用```epoll_ctl```时会调用失败并设置errno(Operation not permittedepoll end)

epoll 之所以做到了高效，最关键的三点：

1. 内部管理 fd 使用了高效的红黑树结构管理，做到了增删改之后性能的优化和平衡；

2. epoll 池添加 fd 的时候，调用 file_operations->poll ，把这个 fd 就绪之后的回调路径安排好。通过事件通知的形式，做到最高效的运行；

3. epoll 池核心的两个数据结构：红黑树和就绪列表。红黑树是为了应对用户的增删改需求，就绪列表是 fd 事件就绪之后放置的特殊地点，epoll 池只需要遍历这个就绪链表，就能给用户返回所有已经就绪的 fd 数组；

