[toc]
# 1 服务端代码
**proto文件**：
```protobuf
syntax = "proto3";

option go_package="./myrpc/;MyRPC";

message LoginRequest{
  string uId = 1;
  string passWord = 2;
}

message LoginResponse{
  bool ok = 1;
}

service login{
  rpc Login(LoginRequest) returns(LoginResponse);
}
```
**生成代码**：
```shell
protoc --go_out=plugins=grpc:. test.proto
```
**生成的代码**：
省略的部分代码
```go
package MyRPC

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

/*根据.proto文件中定义的message LoginRequest生成的结构体*/
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId      string `protobuf:"bytes,1,opt,name=uId,proto3" json:"uId,omitempty"`
	PassWord string `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
}

func (x *LoginRequest) Reset() {
	...
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	...
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	...
}

/*获取Uid*/
func (x *LoginRequest) GetUId() string {
	if x != nil {
		return x.UId
	}
	return ""
}

/*获取PassWord*/
func (x *LoginRequest) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

/*根据.proto文件中定义的message LoginResponse生成的结构体*/
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *LoginResponse) Reset() {
	...
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	...
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

/*获取Ok的值*/
func (x *LoginResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}



/*根据.proto文件中定义的service生成的客户端的接口*/
type LoginClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

/*创建客户端*/
func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	...
}


/*根据.proto文件中定义的service生成的服务器的接口*/
type LoginServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	...
}

/*注册服务*/
func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}
```

# 2 编写服务器代码
## 2.1 实现生成代码中定义的接口
在protoc生成的代码目录下创建另一个go文件来实现接口定义。

因为我们在test.proto文件中定义了一个服务login, 因此会在生成的go文件中有一个
LoginServer的接口。这个接口中有方法Login<font color=gray>(跟test.proto文件中定义的login服务的rpc是一致的)</font>

```go
package MyRPC

import "context"

type LoginService struct {
}

var MyLoginServer LoginService

// 实现LoginServer接口
func (login LoginService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// 实现服务器业务代码
	var response LoginResponse
	if req.UId == "jake" {
		if req.PassWord == "zxcvbnm1997" {
			response.Ok = true
		}
	}
	return &response, nil
}
```
## 2.2 启动grpc服务器
```go
package main

import (
	"fmt"
	"google.golang.org/grpc"
	MyRPC "hello/myrpc"
	"net"
)

func main() {
	// 创建一个grpc Server对象
	var server *grpc.Server
	server = grpc.NewServer()

	// 将grpc对象注册到生成的service中，第二个参数是我们实现的grpc服务
	MyRPC.RegisterLoginServer(server, MyRPC.MyLoginServer)

	/* 创建一个监听器
	注意1: Listen的第二个参数是网络地址<ip:port>形式,冒号是必须的，当省略ip时，默认是127.0.0.1*/
	var listener net.Listener
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("get listener failed")
		return
	}

	// 将监听器加到grpc Server中，启动grpc服务
	err = server.Serve(listener)
}
```

# 3 实现客户端
```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	MyRPC "hello/myrpc"
)

func main() {
	// new一个证书,没有证书可能会报错，当前是无认证状态
	var transport grpc.DialOption
	transport = grpc.WithTransportCredentials(insecure.NewCredentials())

	// 根据地址和证书连接服务端grpc
	var connect *grpc.ClientConn
	connect, err := grpc.Dial("127.0.0.1:9999", transport)
	if err != nil {
		fmt.Println("connect failed")
		return
	}
	defer connect.Close()

	// 创建客户端
	var client MyRPC.LoginClient
	client = MyRPC.NewLoginClient(connect)

	// 调用定义的远程方法 Login
	var loginReq MyRPC.LoginRequest
	loginReq.UId = "jake"
	loginReq.PassWord = "zxcvbnm1997"

	loginRes, err := client.Login(context.Background(), &loginReq)
	if err != nil {
		fmt.Println("login failed :", err.Error())
		return
	}
	fmt.Printf("login response is %v\n", loginRes.Ok)
}
```


