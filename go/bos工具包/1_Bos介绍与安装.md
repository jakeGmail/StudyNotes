[toc]
# 1 介绍
Bos是百度Oss存储服务对象。

# 2 各个语言的安装包
https://cloud.baidu.com/doc/BOS/s/Tjwvyrw7a

**go**: 

```shell
go get github.com/baidubce/bce-sdk-go
```

**SDK目录结构**:

bce-sdk-go
|--auth                   //BCE签名和权限认证
|--bce                    //BCE公用基础组件
|--http                   //BCE的http通信模块
|--services               //BCE相关服务目录
|--|--bos                 //BOS服务目录
|--|--|--bos_client.go    //BOS客户端入口
|--|--|--api              //BOS相关API目录
|--|--|--bucket.go     //BOS的Bucket相关API实现
|--|--|--object.go     //BOS的Object相关API实现
|--|--|--multipart.go  //BOS的Multipart相关API实现
|--|--|--module.go     //BOS相关API的数据模型
|--|--|--util.go       //BOS相关API实现使用的工具
|--|--sts                 //STS服务目录
|--util                   //BCE公用的工具实现