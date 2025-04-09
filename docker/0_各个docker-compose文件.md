[toc]

# 1 minio镜像

```yaml
version: '3'
services:
  minio:
    image: minio/minio
    hostname: "minio"
    ports:
      - "9000:9000" # api 端口
      - "5000:5000" # 控制台端口
    environment:
      MINIO_ACCESS_KEY: admin    #管理后台用户名
      MINIO_SECRET_KEY: admin123# #管理后台密码，最小8个字符
    volumes:
      - ./data:/data               #映射当前目录下的data目录至容器内/data目录
      - ./config:/root/.minio/     #映射配置目录
    command: server --console-address ':5000' /data  #指定容器中的目录 /data
    privileged: true
    restart: always
```