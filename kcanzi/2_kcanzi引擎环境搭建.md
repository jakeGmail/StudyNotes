[toc]
# 1. Ubuntu搭建Kanzi环境
## 1.1 环境预备
在Ubuntu下搭建kanzi环境前需要做以下前提配置：
gcc、python、scons环境
### 1.1.1 gcc环境安装
使用scons编译kanzi项目的时候需要依赖python环境和gcc环境（推荐gcc-4.8.5， gcc-7版本编译出来会有问题）
- 安装gcc-4.8
  ```shell
  sudo apt-get install g++-4.8

  # 如果安装失败,可能需要修改源
  sudo vim /etc/apt/sources.list
  # 在最后一行添加
  deb http://dk.archive.ubuntu.com/ubuntu/ trusty main universe

  # 然后更新源
  sudo apt-get update
  # 再次安装g++4.8
  sudo apt-get install g++-4.8
  ```
如果系统中有其他版本的gcc版本，则需要切换gcc版本至4.8
- 将gcc版本切换到gcc 4.8
  1. **查看系统中存在的gcc版本**
    ```shell
    ls /usr/bin | grep gcc
    ```
    <table><tr><td bgcolor=gray>
    <font color=white>
        jake@JK:~$ ls /usr/bin/ | grep gcc</br>
        c89-gcc</br>
        c99-gcc</br>
        gcc</br>
        gcc-4.8</br>
        gcc-7</br>
        ...
    </font>
    </td></tr></table>

   2. **先设置存在的gcc版本**
    ```shell
    # gcc-4.8后面的数字代表优先级
    # 
    sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.8 20 --slave /usr/bin/g++ g++ /usr/bin/g++-4.8
    sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 10 --slave /usr/bin/g++ g++ /usr/bin/g++-7
    ```
    像这样设置后后续就不需要再次设置了，除非需要添加新的gcc版本
   3. **切换gcc版本**
    ```shell
    # 执行后会让你进行选择，只需要选择对应的数字就可以将gcc版本切换到你需要的版本
    sudo update-alternatives --config gcc
    ```
    
### 1.1.2 scons工具安装
- scons环境安装
  scons是用python进行编译的工具，后续在Ubuntu上进行编译kanzi工程时也是用这个工具进行编译。
  该工具的安装见[scons环境安装](../工具使用/3_scons编译工具.md)
  **注意**:对于一些SConstruct编写的python编译脚本，需要注意是python2.x版本还是python3.x版本，有些语法不同可能会导致编译报错。这里的示例代码中使用的是python2.x版本。
#### 1.1.2.1 python2.x安装scons
1. 安装python2.7
2. 安装python2.7的pip
  ```shell
  curl -o get-pip.py https://bootstrap.pypa.io/pip/2.7/get-pip.py
  sudo python2 get-pip.py
  ```
3. 安装scons
  ```shell
  python2 -m pip install scons
  ```

在使用scons进行编译的时候，scons会使用默认的python环境来进行编译。如果系统中有其他的python版本（例如python3），可以通过切换默认python版本至python2.x

**切换python的版本**:
1. 查看系统中已经存在的python版本
   ```shell
   ls /usr/bin | grep python
   ```
    <table><tr><td bgcolor=gray>
    <font color=white>
    jake@JK:~$ ls /usr/bin/ | grep python</br>
    dh_python2</br>
    python</br>
    python2</br>
    python2.7</br>
    python2-pbr</br>
    python3</br>
    python3.6</br>
    python3.6m</br>
    python3m</br>
    </font>
    </table></tr></td>
2. **设置已经存在的python版本**
   ```shell
   # 最后的数字是优先级
   sudo update-alternatives --install /usr/bin/python python /usr/bin/python2.7 1
   sudo update-alternatives --install /usr/bin/python python /usr/bin/python3.6 2
   ```
   设置过一次后，后续就不需要再重复设置了
3. **切换默认python版本**
   ```shell
   # 执行后根据提示选择python版本
   sudo update-alternatives --config python 
   ```

## 1.2 配置kanzi
1. 将Kanzi Engine文件夹放到ubuntu中。
2. 设置系统寻找库文件的路径
   - 临时设置
     ```shell
     # kanzi库文件路径在<Kanzi Engine路径>Engine/configs/platforms/<各个平台的文件夹名称>
     export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:<kanzi引擎库文件路径>
     ```
   - 永久设置
