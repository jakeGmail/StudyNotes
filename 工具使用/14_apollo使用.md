[toc]

# 1 apollo第三方API接口
## 1.1 什么是开放平台？
Apollo提供了一套的Http REST接口，使第三方应用能够自己管理配置。虽然Apollo系统本身提供了Portal来管理配置，但是在有些情景下，应用需要通过程序去管理配置。

## 1.2 第三方应用接入Apollo开放平台
### 1.2.1 注册第三方应用
第三方应用负责人需要向Apollo管理员提供一些第三方应用基本信息。

基本信息如下：

第三方应用的AppId、应用名、部门
第三方应用负责人
Apollo管理员在 http://{portal_address}/open/manage.html 创建第三方应用，创建之前最好先查询此AppId是否已经创建。创建成功之后会生成一个token，如下图所示：
![](img/apollo_manager.png)
![](img/apollo_manager2.png)

### 1.2.2 给已注册的第三方应用授权
第三方应用不应该能操作任何Namespace的配置，所以需要给token绑定可以操作的Namespace。Apollo管理员在 http://{portal_address}/open/manage.html 页面给token赋权。赋权之后，第三方应用就可以通过Apollo提供的Http REST接口来管理已授权的Namespace的配置了。

###  1.2.3 第三方应用调用Apollo Open API
任何语言的第三方应用都可以调用Apollo的Open API，在调用接口时，需要设置注意以下两点：

Http Header中增加一个Authorization字段，字段值为申请的token
Http Header的Content-Type字段需要设置成application/json;charset=UTF-8

## 1.3 接口文档

### 1.3.1 URL路径参数说明

|参数名|	参数说明|
|------|-----------|
|env|	所管理的配置环境|
|appId	所管理的配置AppId|
|clusterName|	所管理的配置集群名， 一般情况下传入 default 即可。如果是特殊集群，传入相应集群的名称即可|
|namespaceName|	所管理的Namespace的名称，如果是非properties格式，需要加上后缀名，如sample.yml|

### 1.3.2 API接口列表
#### 1.3.2.1 获取App的环境，集群信息
- **URL** : http://{portal_address}/openapi/v1/apps/{appId}/envclusters
- **Method** : GET
- **Request Params** : 无
- **返回值Sample：**
```json
[
    {
        "env":"FAT",
        "clusters":[ //集群列表
            "default",
            "FAT381"
        ]
    },
    {
        "env":"UAT",
        "clusters":[
            "default"
        ]
    },
    {
        "env":"PRO",
        "clusters":[
            "default",
            "SHAOY",
            "SHAJQ"
        ]
    }
]

```

#### 1.3.2.2 获取集群下所有Namespace信息接口
- **URL** : http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces
- **Method**: GET
- **Request Params**: 无
- **返回值Sample**:
```json
[
  {
    "appId": "100003171",
    "clusterName": "default",
    "namespaceName": "application",
    "comment": "default app namespace",
    "format": "properties", //Namespace格式可能取值为：properties、xml、json、yml、yaml
    "isPublic": false, //是否为公共的Namespace
    "items": [ // Namespace下所有的配置集合
      {
        "key": "batch",
        "value": "100",
        "dataChangeCreatedBy": "song_s",
        "dataChangeLastModifiedBy": "song_s",
        "dataChangeCreatedTime": "2016-07-21T16:03:43.000+0800",
        "dataChangeLastModifiedTime": "2016-07-21T16:03:43.000+0800"
      }
    ],
    "dataChangeCreatedBy": "song_s",
    "dataChangeLastModifiedBy": "song_s",
    "dataChangeCreatedTime": "2016-07-20T14:05:58.000+0800",
    "dataChangeLastModifiedTime": "2016-07-20T14:05:58.000+0800"
  },
  {
    "appId": "100003171",
    "clusterName": "default",
    "namespaceName": "FX.apollo",
    "comment": "apollo public namespace",
    "format": "properties",
    "isPublic": true,
    "items": [
      {
        "key": "request.timeout",
        "value": "3000",
        "comment": "",
        "dataChangeCreatedBy": "song_s",
        "dataChangeLastModifiedBy": "song_s",
        "dataChangeCreatedTime": "2016-07-20T14:08:30.000+0800",
        "dataChangeLastModifiedTime": "2016-08-01T13:56:25.000+0800"
      },
      {
        "id": 1116,
        "key": "batch",
        "value": "3000",
        "comment": "",
        "dataChangeCreatedBy": "song_s",
        "dataChangeLastModifiedBy": "song_s",
        "dataChangeCreatedTime": "2016-07-28T15:13:42.000+0800",
        "dataChangeLastModifiedTime": "2016-08-01T13:51:00.000+0800"
      }
    ],
    "dataChangeCreatedBy": "song_s",
    "dataChangeLastModifiedBy": "song_s",
    "dataChangeCreatedTime": "2016-07-20T14:08:13.000+0800",
    "dataChangeLastModifiedTime": "2016-07-20T14:08:13.000+0800"
  }
]
```

#### 1.3.2.3 获取某个Namespace信息接口
- **URL** ： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}
- **Method** ： GET
- **Request Params** ：无
- **返回值Sample** ：

```json
{
    "appId": "100003171",
    "clusterName": "default",
    "namespaceName": "application",
    "comment": "default app namespace",
    "format": "properties", //Namespace格式可能取值为：properties、xml、json、yml、yaml
    "isPublic": false, //是否为公共的Namespace
    "items": [ // Namespace下所有的配置集合
      {
        "key": "batch",
        "value": "100",
        "dataChangeCreatedBy": "song_s",
        "dataChangeLastModifiedBy": "song_s",
        "dataChangeCreatedTime": "2016-07-21T16:03:43.000+0800",
        "dataChangeLastModifiedTime": "2016-07-21T16:03:43.000+0800"
      }
    ],
    "dataChangeCreatedBy": "song_s",
    "dataChangeLastModifiedBy": "song_s",
    "dataChangeCreatedTime": "2016-07-20T14:05:58.000+0800",
    "dataChangeLastModifiedTime": "2016-07-20T14:05:58.000+0800"
  }
```

#### 1.3.2.4 创建Namespace
可以通过此接口创建Namespace，调用此接口需要授予第三方APP，APP级别的权限。

- **URL** ： http://{portal_address} /openapi/v1/apps/{appId}/appnamespaces
- **Method** ： POST
- **Request Params** ：无
- **请求内容(Request Body, JSON格式)** ：

|参数名	|必选|	类型|	说明|
|-------|---|------|-------|
|name	|true|	String|	Namespace的名字|
|appId	|true|	String|	Namespace所属的AppId|
|format	|true|	String|	Namespace的格式，只能是以下类型： properties、xml、json、yml、yaml|
|isPublic|	true	|boolean	是否是公共文件|
|comment	|false|	String|	Namespace说明|
|dataChangeCreatedBy	|true|	String	|namespace的创建人，格式为域账号，也就是sso系统的User ID|

- **返回值 Sample** ：
```json
{
    "name": "FX.public-0420-11",
    "appId": "100003173",
    "format": "properties",
    "isPublic": true,
    "comment": "test",
    "dataChangeCreatedBy": "zhanglea",
    "dataChangeLastModifiedBy": "zhanglea",
    "dataChangeCreatedTime": "2017-04-20T18:25:49.033+0800",
    "dataChangeLastModifiedTime": "2017-04-20T18:25:49.033+0800"
}
```

- **返回值说明**：
>> 如果是properties文件，name = a p p I d 所 属 的 部 门 . {appId所属的部门}.appId所属的部门.{传入的name值} ，例如调用接口传入的name=xy-z, format=properties，应用的部门为框架（FX）,那么name=FX.xy-z

>>如果不是properties文件 name = a p p I d 所 属 的 部 门 . {appId所属的部门}.appId所属的部门.{传入的name值}.${format}，例如调用接口传入的name=xy-z, format=json，应用的部门为框架（FX）,那么name=FX.xy-z.json

#### 1.3.2.5 获取某个Namespace当前编辑人接口
Apollo在生产环境（PRO）有限制规则：每次发布只能有一个人编辑配置，且该次发布的人不能是该次发布的编辑人。 也就是说如果一个用户A修改了某个namespace的配置，那么在这个namespace发布前，只能由A修改，其它用户无法修改。同时，该用户A无法发布自己修改的配置，必须找另一个有发布权限的人操作。 这个接口就是用来获取当前namespace是否有人锁定的接口。在非生产环境（FAT、UAT），该接口始终返回没有人锁定。

- **URL** ： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/lock
- **Method** ： GET
- **Request Params** ：无
- **返回值 Sample（未锁定）** ：

```json
{
  "namespaceName": "application",
  "isLocked": false
}
```

- **返回值Sample(被锁定)** ：

```json
{
  "namespaceName": "application",
  "isLocked": true,
  "lockedBy": "song_s" //锁owner
}
```

#### 1.3.2.6 新增配置接口
- **URL** ： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items
- **Method** ： POST
- **Request Params** ：无
- **请求内容(Request Body, JSON格式)** ：

|参数名	|必选|	类型|	说明|
|-------|----|-----|-------|
|key|	true|	String|	配置的key，长度不能超过128个字符。非properties格式，key固定为content|
|value|	true|	String|	配置的value，长度不能超过20000个字符，非properties格式，value为文件全部内容|
|comment|	false|	String|	配置的备注,长度不能超过1024个字符|
|dataChangeCreatedBy|	true|	String|	item的创建人，格式为域账号，也就是sso系统的User ID|

- **Request body sample :**
```json
{
    "key":"timeout",
    "value":"3000",
    "comment":"超时时间",
    "dataChangeCreatedBy":"zhanglea"
}
```

- **返回值Sample ：**
```json
  {
    "key": "timeout",
    "value": "3000",
    "comment": "超时时间",
    "dataChangeCreatedBy": "zhanglea",
    "dataChangeLastModifiedBy": "zhanglea",
    "dataChangeCreatedTime": "2016-08-11T12:06:41.818+0800",
    "dataChangeLastModifiedTime": "2016-08-11T12:06:41.818+0800"
}
```

#### 1.3.2.7 修改配置接口
- **URL** ： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key}
- **Method** ： PUT
- **Request** Params ：无
- **请求内容(Request Body, JSON格式) ：**

|参数名|	必选|	类型|	说明|
|------|-------|-------|-------|
|key|	true|	String|	配置的key，需和url中的key值一致。非properties格式，key固定为content|
|value|	true|	String|	配置的value，长度不能超过20000个字符，非properties格式，value为文件全部内容|
|comment|	false|	String|	配置的备注,长度不能超过1024个字符|
|dataChangeLastModifiedBy|	true|	String|	item的修改人，格式为域账号，也就是sso系统的User ID|

- **Request body sample :**
```json
{
    "key":"timeout",
    "value":"3000",
    "comment":"超时时间",
    "dataChangeLastModifiedBy":"zhanglea"
}
```

- **返回值** ：无

#### 1.3.2.8 删除配置接口
- **URL** ： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/items/{key}?operator={operator}
- **Method：** DELETE
- **Request Params ：**
|参数名|	必选|	类型|	说明|
|------|--------|------|-------|
|key|	true|	String|	配置的key。非properties格式，key固定为content|
|operator|	true|	String|	删除配置的操作者，域账号|

- **返回值**：无

#### 1.3.2.9 发布配置接口
- **URL** ： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/releases
- **Method** ： POST
- **Request Params** ：无
- **Request Body** ：

|参数名|	必选|	类型|	说明|
|------|-------|-------|-------|
|releaseTitle|	true|	String|	此次发布的标题，长度不能超过64个字符|
|releaseComment|	false|	String|	发布的备注，长度不能超过256个字符|
|releasedBy|	true|	String|	发布人，域账号，注意：如果ApolloConfigDB.ServerConfig中的namespace.lock.switch设置为true的话（默认是false），那么该环境不允许发布人和编辑人为同一人。所以如果编辑人是zhanglea，发布人就不能再是zhanglea。|

- **Request Body example ：**

```json
{
    "releaseTitle":"2016-08-11",
    "releaseComment":"修改timeout值",
    "releasedBy":"zhanglea"
}
```

- **返回值Sample ：**

```json
{
    "appId": "test-0620-01",
    "clusterName": "test",
    "namespaceName": "application",
    "name": "2016-08-11",
    "configurations": {
        "timeout": "3000",
    },
    "comment": "修改timeout值",
    "dataChangeCreatedBy": "zhanglea",
    "dataChangeLastModifiedBy": "zhanglea",
    "dataChangeCreatedTime": "2016-08-11T14:03:46.232+0800",
    "dataChangeLastModifiedTime": "2016-08-11T14:03:46.235+0800"
}
```

#### 1.3.2.10 获取某个Namespace当前生效的已发布配置接口
- **URL**： http://{portal_address}/openapi/v1/envs/{env}/apps/{appId}/clusters/{clusterName}/namespaces/{namespaceName}/releases/latest
- **Method**： GET
- **Request Params**：无
- **返回值Sample**：

```json
{
    "appId": "test-0620-01",
    "clusterName": "test",
    "namespaceName": "application",
    "name": "2016-08-11",
    "configurations": {
        "timeout": "3000",
    },
    "comment": "修改timeout值",
    "dataChangeCreatedBy": "zhanglea",
    "dataChangeLastModifiedBy": "zhanglea",
    "dataChangeCreatedTime": "2016-08-11T14:03:46.232+0800",
    "dataChangeLastModifiedTime": "2016-08-11T14:03:46.235+0800"
}
```

## 1.4 错误码说明
正常情况下，接口返回的Http状态码是200，下面列举了Apollo会返回的非200错误码说明。

### 1.4.1 400 - Bad Request
客户端传入参数的错误，如操作人不存在，namespace不存在等等，客户端需要根据提示信息检查对应的参数是否正确。

### 1.4.2 401 - Unauthorized
接口传入的token非法或者已过期，客户端需要检查token是否传入正确。

### 1.4.3 403 - Forbidden
接口要访问的资源未得到授权，比如只授权了对A应用下Namespace的管理权限，但是却尝试管理B应用下的配置。

### 1.4.3 404 - Not Found
接口要访问的资源不存在，一般是URL或URL的参数错误。

### 1.4.4 405 - Method Not Allowed
接口访问的Method不正确，比如应该使用POST的接口使用了GET访问等，客户端需要检查接口访问方式是否正确。

### 1.4.4 500 - Internal Server Error
其它类型的错误默认都会返回500，对这类错误如果应用无法根据提示信息找到原因的话，可以找Apollo研发团队一起排查问题。