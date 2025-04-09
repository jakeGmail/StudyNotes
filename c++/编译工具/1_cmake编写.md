[toc]
# 0 安装cmake环境
## 0.1 windows安装cmake环境
https://cmake.org/download/

# 1 cmake基本组件
## 1.1 设定cmake的版本(必须)
```cmake
# 后面的版本数字可改变
cmake_minimum_required(VERSION 3.10)
```

## 1.2 设定c++的版本(必须)
```cmake
# 设置使用c++ 11标准编译
set(CMAKE_CXX_STANDARD 11)
```

## 1.3. 设定工程名称（必须）
设置编译项目的名称，这个名称就是编译后最终生成的文件（可执行程序或者库文件）
```cmake
project(<工程名>)
```

## 1.4 添加寻找头文件的路径
这里添加的头文件路径影响代码中的#include
```cmake
# 一次添加一个头文件路径
include_directories(<头文件路径1>)
include_directories(<头文件路径2>)
...
```


## 1.5 添加寻找库文件的路径
这里添加的库文件路径影响`target_link_libraries`关键字的库文件
```cmake
# 一次添加一个库文件路径
link_directories(<库文件路径1>)
link_directorie(<库文件路径2>)
...
```

## 1.6 添加源文件(必须)
**生成可执行文件**：
```cmake
# 第一个参数是通过```project```关键字指定的工程名，后面的参数是源文件(.cpp)名称与本CMakeLists.txt的相对路径
add_executable(
    <工程名>
    <源文件1>
    <源文件2>
    <源文件3>
    ...)
```

**生成库文件**：
```cmake
add_library(
    <工程名>
    <源文件1>
    <源文件2>
    <源文件3>
    ...)
```

## 1.7 链接
```cmake
# 这里的库名可以是去掉lib前缀和.so后缀的名称
target_link_libraries(
    <工程名>
    <库名1>
    <库名2>
    <库名3>
    ...
)
```

## 1.8 添加宏定义
```cmake
add_definitions(-D XXX)
```
