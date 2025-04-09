# Ubuntu虚拟机系统安装
一般来讲如果用x11作为与本地窗口通信的协议只需要安装openGL ES环境即可，如果要使用wayland协议与本地窗口系统进行通信，还需要安装glfw环境 <font color=gray>（主要是有libwayland-bin libwayland-dev libxext-dev libxrandr-dev libxrender-dev x11proto-randr-dev x11proto-xext-dev）</font>

<font color=gray>我自己的ubuntu环境安装的是openGL ES环境 + libwayland-bin libwayland-dev</font>
## 1.1 安装openGL ES环境
安装命令
```shell
sudo apt-get install libgles2-mesa-dev
```

安装这个环境可以支持
```c++
#include <EGL/egl.h>
#include <EGL/eglext.h>
#include <EGL/eglplatform.h>

#include <GLES/gl.h>
#include <GLES/glext.h>
#include <GLES/glplatform.h>

#include <GLES2/gl2.h>
#include <GLES2/gl2ext.h>
#include <GLES2/gl2platform.h>

#include <GLES3/...>
...
```

这个命令还会同时安装
libegl-dev libgl-dev libgles-dev libgles1 libglvnd-dev libglx-dev
  libopengl-dev libpthread-stubs0-dev libx11-6 libx11-dev libxau-dev
  libxcb1-dev libxdmcp-dev x11proto-core-dev x11proto-dev xorg-sgml-doctools
  xtrans-dev

## 1.2 安装glfw环境
安装命令,可以选择不安装这个，这样就不能用wayland协议与本地窗口系统进行通信，而是使用x11协议
```shell
sudo apt-get install libglfw3-dev
```

安装这个dev时同时还会安装依赖
ibegl1-mesa-dev libglfw3 libvulkan-dev libwayland-bin libwayland-dev
libxext-dev libxrandr-dev libxrender-dev x11proto-randr-dev x11proto-xext-dev

可以支持的头文件
```c++
#include <wayland-client.h>
#include <wayland-server.h>
#include <wayland-egl.h>
#include <wayland-cursor.h>
...
```

glfw.h

一个轻量级的，开源的，跨平台的library。支持OpenGL及OpenGL ES，用来管理窗口，读取输入，处理事件等。因为OpenGL没有窗口管理的功能，所以很多热心的人写了工具来支持这些功能，比如早期的glut，现在的freeglut等。那么GLFW有何优势呢？glut太老了，最后一个版本还是90年代的。freeglut完全兼容glut，算是glut的代替品，功能齐全，但是bug太多。稳定性也不好（不是我说的啊），GLFW应运而生。

## 1.3 安装glew环境

不同的显卡公司，也会发布一些只有自家显卡才支 持的扩展函数，你要想用这数涵数，不得不去寻找最新的glext.h,有了GLEW扩展库，你就再也不用为找不到函数的接口而烦恼，因为GLEW能自动识别你的平台所支持的全部OpenGL高级扩展函数。也就是说，只要包含一个glew.h头文件，你就能使用gl,glu,glext,wgl,glx的全部函数。

## 1.4 安装png解析库

```shell
sudo apt-get install libpng-dev
```

## 1.5 安装freetype字体解析库

```shell
sudo apt install libfreetype6-dev
```


