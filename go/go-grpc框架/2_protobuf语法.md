[toc]
# 1 protobuf语法
## 1.1 .proto文件格式
首先，Protocol Buffers的文件是以.proto结尾，看一个例子
```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  string number = 2;
}
```
- 第一行syntax = "proto3"; 表示使用proto3的语法，默认使用proto2的语法
- message声明的方式类似于结构体，是proto中的基本结构
SearchRequest中定义了三个字段，每个字段都会有名称和类型
- SearchRequest消息中的query = 1, 其中的1代表query这个字段在这个消息中的唯一编号是1.
## 1.2 字段类型
下面是.proto文件的类型与go中类型的对应关系
|.proto Type	|Go Type| |.proto Type|Go Type|
|---|---|--|--|--|
|double|	float64||fixed32	|uint32|
|float|	float32||fixed64	|uint64|
|int32	|int32||sfixed32	|int32|
|int64	|int64||sfixed64	|int64
|uint32|	uint32||bool|	bool|
|uint64	|uint64||string|	string|
|sint32	|int32||bytes|	[]byte|
|sint64	|int64|

**map的声明：**
```protobuf
syntax = "proto3";

message SearchRequest {
  // map
  map<int32, string> opt = 1;
}
```


## 1.3 字段编号
事实上，proto并不是传统的键值类型，在声明的.protoc中是不会出现具体的数据的，每一次字段的=后面跟的应该是当前message中的唯一编号，这些编号用于在二进制消息体中识别和定义这些字段。编号从1开始，1-15的编号会占用1个字节，16-2047会占用两个字节，因此尽可能的将频繁出现的字段赋予1-15的编号以节省空间，并且应该留出一些空间以留给后续可能会频繁出现的字段。

## 1.4 reserved 保留字段
reserve关键字可以声明保留字段，保留字段编号声明后，将无法再被用作其他字段的编号和名称，编译时也会发生错误。谷歌官方给出的回答是：，如果一个.proto文件在新版本中删除了一些编号，那么在未来其他用户可能会重用这些已被删除的编号，但是倘若换回旧版本的编号的话就会造成字段对应的编号不一致从而产生错误，保留字段就可以在编译期起到这么一个提醒作用，提醒你不能使用这个保留使用的字段，否则编译将会不通过。
```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  string number = 2;
  map<string, int32> config = 3;
  repeated string a = 4;

  // 如下声明后前面的字段就不能用了
  reserved "a"; //声明具体名称的字段为保留字段
  reserved 1 to 2; //声明一个编号序列为保留字段
  reserved 3,4; //声明
}
```
如此一来，此文件将不会通过编译。

## 1.5 枚举
枚举
可以声明枚举常量并将其当作字段的类型来使用，需要注意的是，枚举项的第一个元素必须是零，因为枚举项的默认值就是第一个元素。

```protobuf
syntax = "proto3";

enum Type {
  GET = 0;
  POST = 1;
  PUT = 2;
  DELETE = 3;
}

message SearchRequest {
  string query = 1;
  string number = 2;
  map<string, int32> config = 3;
  repeated string a = 4;
  Type type = 5;
}
```

当枚举项内部存在相同值的枚举项时，可以使用枚举别名
```protobuf
syntax = "proto3";

enum Type {
  option allow_alias = true; //需要开启允许使用别名的配置项
  GET = 0;
  GET_ALIAS = 0; //GET枚举项的别名
  POST = 1;
  PUT = 2;
  DELETE = 3;
}

message SearchRequest {
  string query = 1;
  string number = 2;
  map<string, int32> config = 3;
  repeated string a = 4;
  Type type = 5;
}
```

## 1.6 嵌套消息
```protobuf
message Outer {                  // Level 0
  message MiddleAA {  // Level 1
    message Inner {   // Level 2
      int64 ival = 1;
      bool  booly = 2;
    }
  }
  message MiddleBB {  // Level 1
    message Inner {   // Level 2
      int32 ival = 1;
      bool  booly = 2;
    }
  }
}
```

## 1.7 Package
您可以向. proto文件添加一个可选的包修饰符，以防止协议消息类型之间的名称冲突。
```protobuf
package foo.bar;
message Open { ... }
```
然后，您可以在定义消息类型的字段时使用包名:
```protobuf
message Foo {
  ...
  foo.bar.Open open = 1;
  ...
}
```
示例：
```protobuf
syntax  = "proto3";

// 声明了代码所在的包（对于C++来说是namespace）
package foo.bar;

enum AVM_INTO_TYPE{
  option allow_alias = true;
  UNKNOWN = 0;
  OUT_AVM = 0;
  RADAR = 1;
}

message AVM{
  string config = 21;
  AVM_INTO_TYPE type = 13;
}

message Application{
  // 使用包名
  foo.bar.AVM avms = 1;
}
```

## 1.8 Any
当新旧协议数据进行交换时，双方的字段版本不一致，一些新字段无法识别，即被称为未知字段。在proto3刚刚推出时，序列化输出时总是会抛弃未知字段，不过3.5以后又重新保留了未知字段，为了能够兼容低版本的proto文件。

Anymessage 类型允许您将消息作为嵌入类型使用，而不需要它们的. proto 定义。
```protobuf
syntax  = "proto3";

// 声明包名
package foo.bar;

// 要使用任意数据类型的声明需要inport此包
import "google/protobuf/any.proto";

message Person{
  string name =1;
  // 任意类型的字段
  repeated google.protobuf.Any data = 2;
}
```

## 1.9 OneOf
这里的官方文档给出的解释实在是太繁琐了，其实就是表示一个字段在通信时可能会有多种不同的类型，但最终只可能会有一个类型被使用（联想switch），并且oneof 内部不允许出现repeated修饰的字段。
```protobuf
message Stock {
    // Stock-specific data
}

message Currency {
    // Currency-specific data
}

message ChangeNotification {
    // 字段编号不能重复
    int32 id = 1;
    oneof instrument {
        Stock stock = 2;
        Currency s = 3;
    }
}
```

## 1.10 Service
service关键字可以定义一个RPC服务，并且可以使用已定义的消息类型。
```protobuf
syntax = "proto3"; // 声明了protobuf的版本

option go_package="./myrpc/;MyPRC";

package fixbug; // 声明了代码所在的包（对于C++来说是namespace）

//定义下面的选项，表示生成service服务类和rpc方法描述，默认不生成
option cc_generic_services = true;

message ResultCode//封装一下失败类
{
  int32 errcode = 1;//表示第1字段
  bytes errmsg = 2;//表示第2字段
}

// 定义登录请求消息类型  name   pwd
message LoginRequest
{
  bytes name = 1;//表示第1字段
  bytes pwd = 2;//表示第2字段
}

// 定义登录响应消息类型
message LoginResponse
{
  ResultCode result = 1;//表示第1字段
  bool success = 2;//表示第2字段
}

//在protobuf里面怎么定义描述rpc方法的类型 - service
service UserServiceRpc
{
  rpc Login(LoginRequest) returns(LoginResponse);
}
```

## 1.11 Options
option就是修改当前文件一些处理的配置，前文中已经出现过了一次别名option，以下官网给出的常用配置项。

```protobuf
option optimize_for = CODE_SIZE;
```
代码生成配置，总共分为三种：

**<font color=green>SPEED</font>**: 将生成用于序列化、解析和对消息类型执行其他常见操作的代码。这段代码经过了高度优化，这个阶段解析最快，但是空间占用最大。

**<font color=green>CODE_SIZE</font>**: 将生成最小的类，并依赖于共享的、基于反射的代码来实现序列化、解析和各种其他操作。因此，生成的代码将比使用 SPEED 时小得多，但是操作会更慢。

**<font color=green>LITE_RUNTIME</font>**: 将生成仅依赖于“ lite”运行时库的类(libProtobuf-lite 而不是 libProtobuf)。Lite 运行时比完整库小得多(数量级更小) ，但省略了某些特性，比如描述符和反射。
```protobuf
option go_package = "dir;filename";
```
代码生成时，指定的生成路径，以及文件名。

## 1.12 生成代码
```shell
protoc --proto_path=<IMPORT_PATH> --go_out=<DST_DIR> <.proto文件路径>
```
IMPORT_PATH：指定解析import指令时要去寻找依赖的目录。
如果DST_DIR以zip结尾的话，会自动将其打包为.zip的压缩文件。

注意：
1. 在.proto文件中需要添加
    ```protobuf
    // 指定生成的=代码存放在./rpc路径下，包名为MyRPC
    option go_package="./rpc/;MyPRC";
    ```
    不然使用```protoc --go_out=<代码输出路径> <.proto文件>```命令会失败

示例：
```protobuf
syntax = "proto3";

option go_package="./myrpc/;MyPRC";

message LoginRequest{
  fixed64 uId = 1;
  string passWord = 2;
}

message LoginResponse{
  bool ok = 1;
}

service login{
  rpc NormalLogin(LoginRequest) returns(LoginResponse);
  rpc MasterLogin(LoginRequest) returns(LoginResponse);
}
```
命令：
```shell
protoc --go_out=./rpc test.proto

// 或者
protoc --go_out=plugins=grpc:./rpc test.proto
```

生成的代码存放在执行protoc命令的"./rpc/myrpc/"目录下. 也就是说通过命令protoc的--go_out参数指定的存放代码的大路径，而在这个大路径下，通过.proto文件中的```option go_package="./myrpc/;MyPRC";```指定在大路径下的代码存放位置。因此代码最终存放的位置在protoc指定的路径和.proto文件指定路径的拼接。

**注意**:
- 如果通过plugins=grpc生成代码后有对应的包找不到，使用```go mod tidy``下载依赖包