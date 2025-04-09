# 1 get-apt
  **作用**：主要用于自动从互联网的软件仓库中搜索、安装、升级、卸载软件或操作系统
  **参数**：
  - ```apt-get update``` 更新源文件，并不会做任何安装升级操作. 这里的源文件是记录了安装包与其对应的下载地址的对应关系。
  - ```apt-get upgrade``` 升级所有已安装的包
  - ```apt-get install <packagename>``` 安装指定包名的包。
    A、下载的软件的存放位置：/var/cache/apt/archives

    B、安装后软件的默认位置：/usr/share

    C、可执行文件位置：/usr/bin

    D、配置文件位置：/etc

    E、lib文件位置：/usr/lib
  - ```sudo apt-get source <packagename>``` 下载该包名的源码
  - ```apt-get remove <packagename>``` 删除指定的包
  - ```apt-get remove -- purge <packagename> ``` 删除包，包括删除配置文件等
  - ```apt-get clean``` 清除无用的包
  - ```apt-get check``` 检查是否有损坏的依赖

# 2 nm
**作用**:
它显示指定文件中的符号信息（函数名、变量名），文件可以是对象文件、可执行文件或对象文件库。
**格式**：
```shell
nm [参数] <文件名>
```
**参数**：
|--选项--|--说明--|
|-------|-------|
|-a|显示所有符号|
|-D|显示动态库符号，这个选项只对动态库有意义|
|-g|只显示外部符号|
|-l|对于每一个符号，使用debug信息找到文件名和行号|
|--define-only|只显示定义的符号|
