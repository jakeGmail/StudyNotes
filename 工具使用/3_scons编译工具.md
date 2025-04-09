[toc]
# 1 scons安装
## 1.1 ubuntu上安装scons
1. 安装python
   因为sCons是采用python脚本编写的。
   ```shell
    sudo apt-get install python3
   ```
   安装完成后在终端执行```python -V```可以查看python的版本,SCons适配的python版本为2.7.x或者3.5以后的版本.
   <table><tr><td bgcolor="black">
   <font color=white>
   jake@JK:~$ python -V</br>
    Python 3.6.9</br>
   </font>
   </td></tr></table>

2. 安装pip
  pip 是 Python 中的标准库管理器。它允许你安装和管理不属于 Python标准库 的其它软件包。
  安装命令：
   ```shell
   sudo apt-get install python3-pip
   ```

2. 安装sCons
   ```shell
   python3 -m pip install scons
   ```

# 2 简单编译
使用scons进行编译的时候需要创建SConstruct文件，在里面编写编译规则，然后在同级目录的终端执行scons命令即可进行编译。


示例：
<table><tr><td bgcolor=gray></br>

```c++
// main.cpp程序
#include<iostream>
using namespace std;
 
int main()
{
    cout << "Hello, World!" << endl;
    return 0;
}
```

```python
# SConstruct文件
Program('main.cpp')
```

</td></tr></table>

在执行scons编译后，会在编译目录下留下一些编译文件，如需要在删除这些编译文件(包括目标文件)，可以使用
  ```
  scons -c
 ```

# 3 SConstruct文件

SConstruct文件实际上就是一个Python脚本。可以在SConstruct文件中使用Python的注释.
重要的一点是SConstruct文件并不完全像一个正常的Python脚本那样工作，其工作方式更像一个Makefile，那就是在SConstruct文件中SCons函数被调用的顺序并不影响SCons你实际想编译程序和目标文件的顺序。换句话说，当你调用Program方法，你并不是告诉SCons在调用这个方法的同时马上就编译这个程序，而是告诉SCons你想编译这个程序。

<table><tr><td bgcolor=gray></br>

```python
# 这会编译出两个可执行程序
print("build main")
Program('main.cpp')
print("build hello")
Program('hello.cpp')
print("end build")
```

```shell
# 执行scons的输出（输出打印跟编写的顺序并不一致）
scons: Reading SConscript files ...
build main
build hello
end build
scons: done reading SConscript files.
scons: Building targets ...
g++ -o hello.o -c hello.cpp
g++ -o hello hello.o
g++ -o main.o -c main.cpp
g++ -o main main.o
scons: done building targets.
```
</td></tr></table>

# 4 scons语法
## 4.1 编译多个源文件
通常情况下，你需要使用多个输入源文件编译一个程序。在SCons里，只需要就多个源文件放到一个Python列表中就行了
```python
# 将main.cpp和test.cpp文件编译成mytest可执行程序
Program('mytest',['main.cpp','test.cpp'])
```

- 可以使用Glob函数来定义定义一个匹配规则来指定源文件列表，比如*,?以及[abc]等标准的shell模式
  ```python
  Program('program',Glob('*.cpp'))
  ```
- 为了更容易处理文件名长列表，SCons提供了一个Split函数，这个Split函数可以将一个用引号引起来，并且以空格或其他空白字符分隔开的字符串分割成一个文件名列表，示例如下：
  ```python
  Program('program', Split('main.cpp  file1.cpp  file2.cpp'))

  # 或者
  src_files=Split('main.cpp  file1.cpp  file2.cpp')
  Program('program', src_files)
  ```

- SCons允许使用Python关键字参数来标识输出文件和输入文件。输出文件是target，输入文件是source，示例如下：
  ```python
  src_files=Split('main.cpp  file1.cpp  file2.cpp')
  Program(target='program', source=src_files)
  
  # 或者：
  src_files=Split('main.cpp  file1.cpp  file2.cpp')
  Program(source=src_files, target='program')
  ```
  如果需要用同一个SConstruct文件编译多个程序，只需要调用Program方法多次：
  ```python
  Program('foo.cpp')
  Program('bar', ['bar1.cpp', 'bar2.cpp'])
  ```

## 4.2 编译静态库文件

```python
# 编译生成 libltest.a静态库文件
Library('ltest',['test.cpp'])
```

除了使用源文件外，Library也可以使用目标文件
```python
Library('foo', ['f1.c', 'f2.o', 'f3.c', 'f4.o'])
```

你甚至可以在文件List里混用源文件和目标文件
```python
lib_srcs = Split('f1.cpp f2.o f3.c f4.0')
Library('foo', lib_srcs)
```
使用StaticLibrary显示编译静态库
```python
StaticLibrary('foo', ['f1.cpp', 'f2.cpp', 'f3.cpp'])
```

## 4.3 编译动态库：
如果想编译动态库（在POSIX系统里）或DLL文件（Windows系统），可以使用SharedLibrary：

```python
SharedLibrary('foo', ['f1.cpp', 'f2.cpp', 'f3.cpp'])
```

## 4.4 链接库文件

```python
Library("ltest",Split("test.cpp"))
# 注意在运行时要执行下"export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:."
# 临时导入下环境变量以便程序能够找到生成的库
Program(target="main", source="main.cpp", LIBS=["ltest"], LIBPATH=".")
```

# 5 节点对象
所有编译方法会返回一个节点对象列表，这些节点对象标识了那些将要被编译的目标文件。这些返回出来的节点可以作为参数传递给其他的编译方法。例如，假设我们想编译两个目标文件，这两个目标有不同的编译选项，并且最终组成一个完整的程序。这意味着对每一个目标文件调用Object编译方法，如下所示：
```python
Object('hello.cpp', CCFLAGS='-DHELLO')
Object('goodbye.cpp', CCFLAGS='-DGOODBYE')
Program(['hello.o', 'goodbye.o'])
```

这样指定字符串名字的问题就是我们的SConstruct文件不再是跨平台的了。因为在Windows里，目标文件成为了hello.obj和goodbye.obj。一个更好的解决方案就是将Object编译方法返回的目标列表赋值给变量，这些变量然后传递给Program编译方法： 
```python
hello_list = Object('hello.cpp', CCFLAGS='-DHELLO')
goodbye_list = Object('goodbye.c', CCFLAGS='-DGOODBYE')
Program(hello_list + goodbye_list)
```

## 5.1 显示创建文件和目录节点
在SCons里，表示文件的节点和表示目录的节点是有清晰区分的。SCons的File和Dir函数分别返回一个文件和目录节点：
```python
hello_c=File('hello.cpp')
Program(hello_c)
```
通常情况下，你不需要直接调用File或Dir，因为调用一个编译方法的时候，SCons会自动将字符串作为文件或目录的名字，以及将它们转换为节点对象。只有当你需要显示构造节点类型传递给编译方法或其他函数的时候，你才需要手动调用File和Dir函数。有时候，你需要引用文件系统中一个条目，同时你又不知道它是一个文件或一个目录，你可以调用Entry函数，它返回一个节点可以表示一个文件或一个目录：
```python
xyzzy=Entry('xyzzy')
```

## 5.2 将一个节点的文件名当作一个字符串
 
如果你不是想打印文件名，而是做一些其他的事情，你可以使用内置的Python的str函数。例如，你想使用Python的os.path.exists判断一个文件是否存在：
```python
import os.path
program_list=Program('hello.cpp')
program_name=str(program_list[0])
if not os.path.exists(program_name):
     print program_name, "does not exist!"
```

## 5.3 GetBuildPath：从一个节点或字符串中获得路径
   env.GetBuildPath(file_or_list)返回一个节点或一个字符串表示的路径。它也可以接受一个节点或字符串列表，返回路径列表。如果传递单个节点，结果就和调用str(node)一样。路径可以是文件或目录，不需要一定存在：
```python
import os.path
exe_target = Program(target="mymain",source=Split("main.cpp"))
exe_path = str(exe_target[0])
if not os.path.exists(exe_path):
        print("target file not exist")

env = Environment(VAR="value")
mainFile = File("main.cpp")
print(env.GetBuildPath([mainFile,"sub/dir/$VAR"]))
```

将会打印输出如下：

```shell
jake@ubuntu:~$ scons -Q
target file not exist
['main.cpp', 'sub/dir/value'] # GetBuildPath的内容
g++ -o main.o -c main.cpp
g++ -o mymain main.o
```

# 6 依赖性
隐式依赖：$CPPPATH Construction变量
```c++
#include <iostream>
#include "hello.h"
using namespace std;
 
int main()
{
    cout << "Hello, " << VAR << endl;
    return 0;
}
```
并且，hello.h文件如下： 
```c++
#define VAR "world"
```
在这种情况下，我们希望SCons能够认识到，如果hello.h文件的内容发生改变，那么hello程序必须重新编译。我们需要修改SConstruct文件如下： 
```python
Program('hello.cpp', CPPPATH='.')  #CPPPATH告诉SCons去当前目录('.')查看那些被C源文件（.c或.h文件）包含的文件。
```
就像```$LIBPATH变量```，$CPPPATH也可能是一个目录列表，或者一个被系统特定路径分隔符分隔的字符串。

# 7 环境
## 7.1 外部环境
外部环境指的是在用户运行SCons的时候，用户**环境变量**的集合。这些变量在SConscript文件中通过Python的os.environ字典可以获得。你想使用外部环境的SConscript文件需要增加一个import os语句。
```python
import os
env_value = os.environ
print(env_value)
```

运行结果
```shell
jake@ubuntu:~$ scons -Q
environ({'SHELL': '/bin/bash',
 'XDG_CONFIG_DIRS': '/etc/xdg/xdg-ubuntu-wayland:/etc/xdg',
 'GNOME_SHELL_SESSION_MODE': 'ubuntu',
  'DESKTOP_SESSION': 'ubuntu-wayland', 
 'PWD': '/home/jake/Desktop/MyCode/scons_study', 
 'LOGNAME': 'jake', 
 'XDG_SESSION_DESKTOP': 'ubuntu-wayland', 
 'XDG_SESSION_TYPE': 'wayland', 
 'HOME': '/home/jake', 
 'XDG_CURRENT_DESKTOP': 'ubuntu:GNOME', 
 'WAYLAND_DISPLAY': 'wayland-0', 
 'PATH': '/home/jake/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin'
 ···
 ··· })
scons: `.' is up to date.
```

## 7.2 构造环境
一个构造环境是在一个SConscript文件中创建的一个唯一的对象，这个对象包含了一些值可以影响SCons编译一个目标的时候做什么动作，以及决定从那一个源中编译出目标文件。SCons一个强大的功能就是可以创建多个构造环境，包括从一个存在的构造环境中克隆一个新的自定义的构造环境。

### 7.2.1 创建一个构造环境
创建一个构造环境：Environment函数
默认情况下，SCons基于你系统中工具的一个变量集合来初始化每一个新的构造环境。当你初始化一个构造环境时，你可以设置环境的构造变量来控制一个是如何编译的。例如：
```python
import os
env=Environment(CC='gcc', CCFLAGS='-O2')
env.Program('foo.c') # 使用创建的环境变量来编译程序
# 或者
env=Environment(CXX='/usr/local/bin/g++', CXXFLAGS='-02')
env.Program('foo.cpp')
```
### 7.2.2 从一个构造环境中获取值
你可以使用访问Python字典的方法获取单个的构造变量： 
<table><tr><td bgcolor=gray></br>

```python
# SConstruct文件内容
import os
env = Environment()
print("env=",env)
# 如果没有CXX环境变量会报错
print("CXX=", env["CXX"])
```

```shell
# 运行输出
jake@ubuntu:~$ scons -Q
env= <SCons.Script.SConscript.SConsEnvironment object at 0x7fe3afbb8070>
CXX= g++
scons: `.' is up to date.
```
</td></tr></table>

实际上生成的env是一个带有方法的对象，可以通过其中的Dictionary方法来获取环境变量的字典
```python
import os
env = Environment(TEST="mytest",CC="g++")
dic = env.Dictionary()
for key in ["TEST","CC"]:
	print(key, ":", dic[key])
```

### 7.2.3 默认的构造环境
默认的构造环境：DefaultEnvironment函数
你可以控制默认构造环境的设置，使用DefaultEnvironment函数：
```python
DefaultEnvironment(CC='/usr/local/bin/gcc')
```
这样配置以后，所有Program或者Object的调用都将使用/usr/local/bin/gcc编译目标文件。注意到DefaultEnvironment返回初始化了的默认构造环境对象，这个对象可以像其他构造环境一样被操作。所以如下的代码和上面的例子是等价的： 
```python
env=DefaultEnvironment()
env['CC']='/usr/local/bin/gcc'
```
### 7.2.4 多个构造环境
构造环境的真正优势是你可以创建你所需要的许多不同的构造环境，每一个构造环境对应了一种不同的方式去编译软件的一部分或其他文件。比如，如果我们需要用-O2编译一个程序，编译另一个用-g，我们可以如下做：
```python
opt=Environment(CCFLAGS='-O2')
dbg=Environment(CCFLAGS='-g')
opt.Program('foo','foo.cpp')
dbg.Program('bar','bar.cpp')
```

### 7.2.5 拷贝构造环境：Clone方法 
有时候你想多于一个构造环境对于一个或多个变量共享相同的值。当你创建每一个构造环境的时候，不是重复设置所有共用的变量，你可以使用Clone方法创建一个构造环境的拷贝。Environment调用创建一个构造环境，Clone方法通过构造变量赋值，重载拷贝构造环境的值。例如，假设我们想使用gcc创建一个程序的三个版本，一个优化版，一个调试版，一个其他版本。我们可以创建一个基础构造环境设置$CC为gcc，然后创建两个拷贝：
```python
env=Environment(CC='gcc')
opt=env.Clone(CCFLAGS='-O2')
dbg=env.Clone(CCFLAGS='-g')
env.Program('foo','foo.cpp')
o=opt.Object('foo-opt','foo.cpp')
opt.Program(o)
d=dbg.Object('foo-dbg','foo.cpp')
dbg.Program(d)
```

### 7.2.6 替换值：Replace方法
你可以使用Replace方法替换已经存在的构造变量：
```python
env=Environment(CCFLAGS='-DDEFINE1');
env.Replace(CCFLAGS='-DDEFINE2');
env.Program('foo.cpp')
```

### 7.2.7 在没有定义的时候设置值：SetDefault方法
有时候一个构造变量应该被设置为一个值仅仅在构造环境没有定义这个变量的情况下。你可以使用SetDefault方法，这有点类似于Python字典的set_default方法：
```python
# 如果SPECIAL_FLAG已经有定义了则这里的设置不会生效
env.SetDefault(SPECIAL_FLAG='-extra-option')
```

### 7.2.8 控制目标文件的路径：env.Install方法 
```python
import os
env = Environment()
env.Program(target="mymain", source="main.cpp")

#表示要将mymain和main.o放到./bin目录下
env.Install("./bin",["main.o","mymain"]) 
```

### 7.2.9 执行SConstruct脚本文件
在我们使用SConstruct进行编译的时候，可能会涉及到对子文件的SConstruct文件进行执行。







