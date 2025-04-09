[toc]

# 1 grafana介绍
Grafana 是一个开源的分析和监控平台，用于可视化和分析时间序列数据。它支持从多种数据源获取数据，如 Prometheus、Graphite、Elasticsearch、InfluxDB 等，并将这些数据展示为图表、仪表盘等可视化形式。

### Grafana常用数据源有哪些?
常用的数据源有Graphite、MySQL、Influxdb、Prometheus、Elasticsearch、AWSCloudWatch等;商业化的数据源包括如Microsoft SQL Server、Oracle公司的Oracle数据库等。
除此之外，Grafana还有一个explore(探索)模式，在explore模式下我们可以编写查询语句进行查询（相当于查询客户端)，这样我们就可以先专注于查询迭代，直到有一个有效的查淘，然后再考虑将其放到仪表盘中。


### Grafana支持告譬功能吗?
支持多种告警方式，如Email、Telegram、钉钉等Webhook方式，但监控与告譬并非Grafana的强项。

### Grafana如何展示数据?
Grafana靠各种插件来展示数据。插件分原生(内置）插件和社区插件。原生插件包括:Graph(图形) 、Singlestat(单值状态图)、Stat(状态图)、Gauge(仪表度量图)、Bar Gauge(条状态度量图) 、Table(表格图)、Text(文本图)、Dashboard list(仪表盘列表)、Plugin list (插件列表)、AlertList(告警列表图)等，其中Stat和Bar Gauge目前在v6.x里仍还是Beta版。

### grafana插件
Grafana社区常用插件包括:Zabbix、Diagram、Imagelt、FlowCharting呼。另外，像Clock、Pie Chart也出自出Grafana Labs，但没有内置在Grafana中。
Grafana言网插件下载地址: https://grafana.com/grafana/plugins?orderBy=weight&direction=ascGrafana 
Dashboard地址: https://grafana.com/grafana/dashboards?orderBy=name&direction=asc

### 为什么要用Grafana?
因为Grafana支持接入当前各种主流的数据库，并且能将各数据库中的数据以非常灵活酷炫的图表展现出来，同时也因为是开源软件方便二次开发定制。另外，当前主流开源的监控系统诸如zabbix、prometheus、open-falcon等均能与Grafana完美结合来展示图表数据。作为一名IT运维人员，除了要及时有效地监控到系统运行状态，还需要展示各种数据趋势，快速发现问题。所以，熟练便用Grafana的各种插件也是运维人员必会技能。

# 2 grafana安装
在https://grafana.com/grafana/download?pg=graf&plcmt=deploy-box-1 下载对应系统/版本的grafana

### docker安装
使用Alpine基础镜像的轻量级Docker容器镜像。
```dockerfile
docker run -d --name=grafana -p 3000:3000 grafana/grafana-enterprise
```

Docker(Ubuntu base image)
对于那些喜欢Ubuntu基础镜像的人来说，可以选择Docker容器镜像。
```dockerfile
docker run -d --name=grafana -p 3000:3000 grafana/grafana-enterprise:11.0.0-ubuntu
```

### docker-compose运行grafana

```dockerfile
version: 'v2.26.1'

services:
  grafana:
    image: grafana/grafana:latest
    container_name: grafana
    ports:
      - "3000:3000"
    volumes:
      - grafana-storage:/var/lib/grafana
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin

volumes:
  grafana-storage:
```
这个docker-compose.yml文件包含以下内容：

- version: '3.7'：指定使用的Docker Compose文件版本。
- services：定义服务。在这里我们定义了一个grafana服务。
- image: grafana/grafana:latest：指定使用的Grafana Docker镜像。
- container_name: grafana：指定容器名称。
- ports: - "3000:3000"：将宿主机的3000端口映射到容器的3000端口，这是Grafana的默认访问端口。
- volumes: - grafana-storage:/var/lib/grafana：将宿主机上的一个持久化存储卷挂载到容器内的/var/lib/grafana目录，以持久化Grafana的数据。
environment：设置环境变量来配置Grafana的管理员用户名和密码。
