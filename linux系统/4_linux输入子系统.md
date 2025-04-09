# 1 介绍
linux下一切皆文件，在处理输入事件(键盘、触摸、鼠标)时，对应的事件也会输入到一个文件中，我们只需要读取对应文件的内容就可以获取到输入事件。

# 2 获取输入事件
我们可以读取指定的设备节点的内容来获取输入事件，例如`/dev/input/event2`节点。
可以通过系统调用read来读取，读取一次的内容存入`input_event`结构体中即可。

# 2.1 input_event结构体
使用input_event结构体和读取需要包含一下头文件
```c++
#include<linux/input.h>
#include<fcntl.h>
#include<unistd.h>
```
其中input_event结构体定义如下
```c++
struct timeval {
  __kernel_time_t tv_sec;  // long类型
  __kernel_suseconds_t tv_usec;  // long类型
};

struct input_event {
  struct timeval time; // 事件发生的时间
  __u16 type;  // 时间类型
  __u16 code;  // 时间码
  __s32 value;  // 值
};
```
根据type的不同，code和value字段的含义也有所不同。
## 2.2 type的值
### 2.2.1 type = EV_KEY
当type值为EV_KEY时，表示按键事件。
**code**: 此时code表示按键的代码，0 ~ 127为键盘按键代码，0x110 ~ 0x116为鼠标按键代码， 其中0x110(BTN_ LEFT)为鼠标左键, 0x111(BTN_RIGHT)为鼠标右键, 0x112(BTN_ MIDDLE)为鼠标中键。
**value**: value字段1表示按键按下，0表示抬起，2表示长按

### 2.2.2 type = EV_ABS
当type = EV_ABS时，表示该事件为触摸事件。
**code**：此时code字段表示触摸的方式，
ABS_MT_PRESSURE为触摸按压；
ABS_MT_POSITION_X为触摸的x坐标；
ABS_MT_POSITION_Y为触摸的y坐标；
ABS_MT_SLOT为触摸的点位(多指触摸)，一般在每个触摸事件开始时，会最先读到该类型。

**value**：当type字段为ABS_MT_PRESSURE时，value的值为1表示触摸按压，0表示触摸抬起。
当type字段为ABS_MT_POSITION_X/ABS_MT_POSITION_Y时，value字段为触摸的x/y坐标，以屏幕左上角为原点。

### 2.2.3 type = EV_STN
表示一个完整事件结束。当读到这个事件后，前面读到的事件就可以组成一个完整的事件用于逻辑处理了。