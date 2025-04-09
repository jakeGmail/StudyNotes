[toc]
# 1 gdb安装
```shell
sudo apt install gdb
```
# 2 调试前的准备
在进行调试前，需要在编译命令gcc/g++中添加-g的参数，编译出来的可执行程序才能使用gdb调试。
# 3 gdb的命令
一下命令均在gdb模式下执行
## 3.1 gdb运行程序
```shell
# 执行后程序将会一直运行，指导结束或者遇见设置的断点
run

# 或则使用简写 
r
```

##  3.2 给程序设置参数
在运行程序员前，给程序传入参数。
```shell
set args <arg1> <arg2> ...
```
示例：
<table><tr><td bgcolor=gray>
<font color=white>
jake@ubuntu:$ gcc test.cpp -g -o test<br>
jake@ubuntu:$ gdb test<br>
...<br>
<font color="87CEFA">(gdb) set args name age<br>
(gdb) r</font><br>
Starting program: /home/jake/test name age<br>
arg[0]=/home/jake/test<br>
arg[1]=name<br>
arg[2]=age<br>
[Inferior 1 (process 4350) exited normally]<br>
(gdb) 
</font>
</td></tr></table>

## 3.3 给程序设置断点
```shell
# 默认设置的是main函数所在文件的行数
break <代码行数>

# 或者使用简写
b <代码行数>

# 给指定文件的num行打断点
b <源文件名>:<num>
```
**注意**：
1. 如果设置的断点在注释上，那么断点会被设置值到注释的下一句代码上
2. 当执行到设置的断点时，断点这行的代码还未真正执行

示例
```c++
  // 测试代码
  1 #include <stdio.h>
  2 
  3 int main(int argc, char* argv[]){
  4     // print
  5     for(int i=0;i<argc;i++){
  6         printf("arg[%d]=%s\n",i,argv[i]);
  7     }
  8     return 0;
  9 }
```

<table><tr><td bgcolor=gray>
<font color=white>
<font color="87CEFA">(gdb) break 6</font><br>
Breakpoint 1 at 0x401155: file test.cpp, line 6.<br>
<font color="87CEFA">(gdb) r</font><br>
Starting program: /home/jake/Desktop/test/test <br>
Breakpoint 1, main (argc=1, argv=0x7fffffffdf48) at test.cpp:6<br>
6	        printf("arg[%d]=%s\n",i,argv[i]);<br>
(gdb) <br>
</font>
</td></tr></table>

## 3.4 执行当前行的语句
### 3.4.1 直接执行当前语句
```shell
next 

# 或者简写为 n
n
```
注意：
1. 如果当前行是函数，则**不会**进入函数内部去执行，而是直接执行完这个函数

### 3.4.2 进入当前语句内部执行
如果当前语句是函数，则会进入当前函数的内的第一条语句，但暂不执行这条语句
```shell
step

# 或者简写为 s
s
```

## 3.5 继续程序运行
继续程序运行，直到下一个断点
```shell
continue

# 或者简写 c
c
```

## 3.6 打印变量
```shell
print <变量名>

# 或者简写
p <变量名>
```

## 3.7 设置变量的值

可以改变程序中的变量值
```shell
set var <变量名>=<值>
```

## 3.8 显示源码内容
显示当前运行的代码附近的代码，这还会显示源码的行号
```shell
list 

# 或者简写 l
l
```

## 3.9 跳转到指定的行数

使程序从当前要执行的代码处，直接跳转到指定位置处继续执行后续的代码。**(中间的代码不执行)**
```shell
jump <代码行号>

# 或者简写为 j 
j
```

## 3.10 结束当前的函数
一直运行直到当前函数执行完毕，返回到函数调用处并暂停程序。
```shell
finish
```

## 3.11 中止调试
```shell
q
```


