[toc]

# 1 opengl常见头文件介绍
- <GL/gl.h>：OpenGL所使用的函数和常量声明。

- <GL/glu.h>：GLU（OpenGL实用库）所使用的函数和常量声明。GLU库属于OpenGL标准的一部分。（以下各种库则不属于）

- <GL/glaux.h>：GLAUX（OpenGL 辅助库）所使用的函数和常量声明。这个库提供了创建窗口，处理键盘和鼠标事件，设置调色板等OpenGL本身不提供，但在编写OpenGL程序时又经常用 到的功能。目前这个库已经过时，只有比较少的编译环境中有提供，例如VC系列。在VC系列编译器中，使用这个头文件之前必须使用#include <windows.h>或者具有类似功能的头文件。

- <GL/glut.h>：GLUT（OpenGL实用工具 包）所使用的函数和常量声明。这个库的功能大致与GLAUX类似，目前许多OpenGL教程使用这个库来编写演示程序。一些编译系统可能不直接提供这个库 （例如VC系列），需要单独下载安装。这个头文件自动包含了<GL/gl.h>和<GL/glu.h>，编程时不必再次包含它 们。

- <GL/glext.h>：扩展头文件。因为微软公司对OpenGL的支持不太积极，VC系列编译器虽然有<GL /gl.h>这个头文件，但是里面只有OpenGL 1.1版本中所规定的内容，而没有OpenGL 1.2及其以后版本。对当前的计算机配置而言，几乎都支持OpenGL 1.4版本，更高的则到1.5, 2.0, 2.1，而VC无法直接使用这些功能。为了解决这一问题，就有了<GL/glext.h>头文件。这个头文件提供了高版本OpenGL所需要 的各种常数声明以及函数指针声明。

- <GL/wglext.h>：扩展头文件。与<GL/glext.h>类似，但这个头文件中只提供适用于Windows系统的各种OpenGL扩展所使用的函数和常量，不适用于其它操作系统。

"glee.h"：GLEE 开源库的头文件。它的出现是因为<GL/glext.h>虽然可以使用高版本的OpenGL函数，但是使用的形式不太方便。GLEE库则让高 版本的OpenGL函数与其它OpenGL函数在使用上同样方便。需要注意的是，这个头文件与<GL/gl.h>是冲突的，在包 含"glee.h"之前，不应该包含<GL/gl.h>。
#include <GL/glut.h>
#include "glee.h"   // 错误，因为glut.h中含有gl.h，它与glee.h冲突
                    // 但是如果把两个include顺序交换，则正确
"glos.h"：虽然这个也时常见到，但我也不知道它到底是什么，可能是与系统相关的各种功能，也可能只是自己编写的一个文件。我曾经看到一个glos.h头文件中只有一句#include <GL/glut.h>。

![](img/glHeader.png)

## 1.1 EGL， GLX/WGL/AGL 和GL之间的关系
1 egl是一个管理者的功能。包括管理所有的display ， context， surface，config。可能有很多的display ，每个display有很多的configs，这个display上可以创建很多的context和surface。如果这些工作都有app去做，那就太麻烦了。

2系统调用应该是没有3d render context这样的概念的，所以必须有某个组件或者某一层中间件来做维护context这个事情。而且针对surface而言，egl分配出来的surface有些是不能给native的api（也就是系统提供的绘图操作）去render的， EGL的config 有一个flag就是native_renderable，指定是不是能够给native的api去render（估计是tiling什么的东西不兼容）。

3 gles是不能直接处理surface的，他没有这样的接口。gl或者gles是一组只和3d render相关的操作，软件架构上属于很高的level。和具体的下层的物理对应是无关的。所以gl或者gles没有surface这样的东西。只有FBO VBO PBO之类的对象的概念。

4 egl更多的角色属于3d api和native api的中间耦合连接组件。app的资源管理是通过egl去实现，包括维护包括系统相关的surface操作同时包括3drender相关的context的维护，app的绘图操作交给gl或者gles。3d core 和native api和egl这三个组件在架构上应该是属于同一个level的东西，而不应该是上下之分。

5 glx wgl agl egl其实都是做了资源管理的功能，真正的3d render的过程就给gl或者gles去做了。

## 1.2 glut、glew、 gles、egl

glut：OpenGL Utility Toolkit，用于开发独立于窗口系统的OPENGL程序。其中打包了很多窗口操作相关的接口，包括窗口创建、显示、输入设备读取、窗口管理等；使用它可以在OPENGL开发中快速完成窗口的相关操作。不过已经在1998年停止更新与维护。

freeglut：一个完全开源的替代glut的库

glew：OpenGL Extension Wrangler Library，一个跨平台的C++扩展库，基于OpenGL图形接口。以Windows平台开发为例，Windows默认只支持OpenGL1.1版本，无法使用更高版本GL的特性和接口函数，但是我们只需要包含这个库和对应的.h文件，即可使用相关的接口和特性，十分方便。注意glewInit()的初始化调用，否则直接调用OpenGL1.1版本以上接口时会出现编译错误，即使你加载了对应的库文件。

gles：OpenGL for Embedded Systems，针对嵌入式设备的API子集，相对于OpenGL“阉割”去了一部分接口，而且在shader语法上有一定的区别

EGL：EGL是渲染API（如OpenGL ES）和原生窗口系统之间的接口。通常来说，OpenGL是一个操作GPU的API，它通过驱动向GPU发送相关指令，控制图形渲染管线与状态机的运行状态，但是当涉及到与本地窗口系统进行交互时，就需要一个中间层，且最好是与平台无关的。因此EGL被设计出来，作为OpenGL和原生窗口系统之间的桥梁。

如果只是在Window编译OPENG工程，一般使用freeglut+glew即可满足大部分需求；另外还有glfw、glaux等，目前还没有使用过。

# 2 opengl相关文档
- **opengl学习文档**: [learnopengl-cn.github.io/intro](https://learnopengl-cn.github.io/intro/)