# 1 Ubuntu安装Unity
- 添加公共登录key
  ```shell
  wget -qO - https://hub.unity3d.com/linux/keys/public | gpg --dearmor | sudo tee /usr/share/keyrings/Unity_Technologies_ApS.gpg > /dev/null
  ```
- 添加unity hub仓库
  ```shell
  sudo sh -c 'echo "deb [signed-by=/usr/share/keyrings/Unity_Technologies_ApS.gpg] https://hub.unity3d.com/linux/repos/deb stable main" > /etc/apt/sources.list.d/unityhub.list'
  ```
- 更新软件列表
  ```shell
  sudo apt update
  ```
- 获取unity hub
  ```shell
  sudo apt-get install unityhub
  ```
如果要移除unity hub,请使用
```shell
sudo apt-get remove unityhub
```

以上来源于：https://docs.unity.cn/hub/manual/InstallHub.html
--------
