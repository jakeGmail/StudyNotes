
-  安装mysql服务
```shell
sudo apt install mysql-server -y
```
- 检车mysql服务状态
```shell
sudo systemctl status mysql.service

# 如果输出的Active项没有出现"active (running)"说明mysql服务没有起来，需要重新启动
service mysql restart
```
- 查看管理员的账户和密码
```shell
sudo cat /etc/mysql/debian.cnf
```

<table><tr><td bgcolor=black><font color=white>
jake@ubuntu:~$ sudo cat /etc/mysql/debian.cnf</br>
# Automatically generated for Debian scripts. DO NOT TOUCH!</br>
[client]</br>
host     = localhost</br>
user     = debian-sys-maint</br>
password = UgWYqPYGsJMg6Dnv</br>
socket   = /var/run/mysqld/mysqld.sock</br>
[mysql_upgrade]</br>
host     = localhost</br>
user     = debian-sys-maint</br>
password = UgWYqPYGsJMg6Dnv</br>
socket   = /var/run/mysqld/mysqld.sock</br>
</font></td></tr></table>
通过输出可知 管理员账号debian-sys-maint， 密码：UgWYqPYGsJMg6Dnv

- 以管理员账号登录mysql
```shell
mysql -udebian-sys-maint -pUgWYqPYGsJMg6Dnv
```
-  对数据库进行初始化,设置root用户的密码
```shell
sudo mysql_secure_installation
```
进行设置时会依次提供以下选项：
1. 是否安装验证密码插件：NO
2. 输入要为root管理员设置的数据库密码：
        再次输入root用户的新密码：（可设置跟root权限密码一样）
3. 删除匿名账户： yes
4. 是否禁止root管理员从远程登录：NO
5. 是否删除test数据库并取消对它的访问权限：YES
6. 是否刷新授权表，让初始化的设定立即生效：YES
</br>
**注意**：如果在第#2步出错，无法修改密码，可以先退出去使用命令,在mysql下修改
```shell
# 切换至root用户（su）
su

# 登录root用户的mysql
mysql

# 修改root登录的密码
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password by 'zxcvbnm1997';
```
之后继续再次执行步骤#5

- 使用新的密码进行登录
```shell
mysql -uroot -pzxcvbnm1997
```

-  配置mysql远程访问功能：
1）编辑配置文件：vim /etc/mysql/mysql.conf.d/mysqld.cnf
2）将 bind-address            = 127.0.0.1 去掉
3）重启mysql使配置文件生效sudo systemctl restart mysql.service
4）再次进入mysql:  ```sudo mysql -uroot -pzxcvbnm1997```
5）使用mysql:  ```use mysql;```
6）查看user表的host和user字段
```select host,user from user;```
可以看到root对应的host是localhost，只允许本地连接，需要将对应host改成"%"，表示允许所有ip远程连接，这里对于root账户一般可不开远程连接，一般新建账号再分配权限后开远程连接权限，我这里教程就先开远程权限了.
-  修改root用户的远程访问权限
```update user set host='%' where user='root';```
- 查看user表命令
```select host,user,plugin,authentication_string from user;```
<font color=gray>（注：host为 % 表示不限制ip，localhost表示本机使用，plugin非mysql_native_password 则需要修改密码）</font>
很明显，此时的root的plugin为auth_socket，所以需要修改密码。这里密码可以修改成原来一样的密码就行了
- 修改用户密码
这里将密码修改为跟原密码一样zxcvbnm1997
```alter user 'root'@'%' identified with mysql_native_password by 'zxcvbnm1997';```
这里有可能是会报1396的错误，这是因为 root用户已存在，解决办法是删除掉root用户，然后重新新增一个root用户。
- 删除root用用户
```delete from user where user='root';```
- 新增root用户
```create user 'root'@'%';```
- 查看
我们发现root用户的authentication_string是空的，于是我们需要设置密码
- 设置密码：
   ```alter user 'root'@'%' identified with mysql_native_password by 'zxcvbnm1997';```
- 给root用户授权命令
```grant all privileges on *.* to 'root'@'%' with grant option;```
- 刷新使配置生效
    ```flush privileges;```
- 退出mysql并重启mysql

# 在Ubuntu上安装MYSQL workbench
```shell
sudo apt update
sudo apt install mysql-workbench

# 如果报依赖错误就用
sudo apt update
sudo apt --fix-broken install
sudo apt autoremove
sudo apt install libsigc++-2.0-0c2a
sudo apt install libpangomm-1.4-1
sudo apt install libatkmm-1.6-1
sudo apt install mysql-workbench
```




