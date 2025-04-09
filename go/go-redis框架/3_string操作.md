[toc]

# 1 string操作命令
## 1.1 set
```go
// go-redis/V8版本
func (c cmdable) Set(ctx context.Context, key string/, value interface{}, expiration time.Duration) *StatusCmd

// go-redis版本
func (c *cmdable) Set(key string, value interface{}, expiration time.Duration) *StatusCmd
```
**参数：**
<font color=blue>ctx</font> :
<font color=blue>key</font> : 需要设置的key
<font color=blue>value</font> : key的值，可以是字符串或者数字
<font color=blue>expiration</font>:  设置过期时间，单位us;  0表示永久
**返回值：**
<font color=blue>StatusCmd</font>：redis命令的执行状态，StatusCmd.Val()用于获取返回结果字符串

**示例**：
```go
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

...

func main(){
    defer redisClient.Close()
    ctx := context.Background()
    var status *redis.StatusCmd
    /*设置name为99，时效性为10秒*/
    status = redisClient.Set(ctx, "name", 99, 10*time.Second)
    
    str, err := status.Result()
    if err != nil {
        fmt.Println("cmd set failed:", err.Error())
    } else {
        /*status.Result()返回的第一个参数就是status.Val()*/
        fmt.Println("str=", str) // OK
        fmt.Println("value=", status.Val()) // OK
    }
}
```

## 1.2 get
```go
// go-redis/V8版本
func (c cmdable) Get(ctx context.Context, key string) *StringCmd

// go-redis版本
func (c *cmdable) Get(key string) *StringCmd
```

**参数：**
<font color=blue>ctx</font> :
<font color=blue>key</font> : string类型的键

**返回值**：
<font color=blue>*StringCmd</font> : 参见[4 redis.StringCmd结构体](#4-redisstringcmd结构体)

**示例**：
```go
func TestRun() {
    ctx := context.Background()
    strCmd := redisClient.Get(ctx, "name")
    if strCmd.Err() != nil {
        fmt.Println("get cmd exec failed:", strCmd.Err().Error())
    } else {
        // 返回name的值
        fmt.Println("get key=", strCmd.Val())
    }
}
```

## 1.3 MSet
```go
// go-redis/V8版本
// 同时设置多个string类型的key
func (c cmdable) MSet(ctx context.Context, values ...interface{}) *StatusCmd


// go-redis版本
func (c *cmdable) MSet(pairs ...interface{}) *StatusCmd
```
**参数**:
<font color=blue>ctx</font> :
<font color=blue>values</font> : 键值对的列表、map或者键-值-键-值的参数，其中值对应的参数可以是数字

**示例1**：
```go
func TestRun() {
    ctx := context.Background()
    str, err := redisClient.MSet(ctx, "name", "jake", "age", 26.3).Result()

    // 第二参数也可以是接口列表
    //str, err := redisClient.MSet(ctx, []interface{}{"name", "jake", "age", 22}).Result()
    if err != nil {
        fmt.Println("mset cmd exec failed:", err.Error())
    } else {
        fmt.Println("mset smd:", str)
    }
}
```

**示例2**：
map作为第二参数
```go
func TestRun() {
    ctx := context.Background()
    cmdMaps := make(map[string]interface{}, 1)
    cmdMaps["name"] = "jake"
    cmdMaps["age"] = 22
    str, err := redisClient.MSet(ctx, cmdMaps).Result()
    if err != nil {
        fmt.Println("mset cmd exec failed:", err.Error())
    } else {
        fmt.Println("mset smd:", str)
    }
}
```

## 1.4 Incr
将string类型的key中的数字+1
```go
// go-redis/V8版本
func (c cmdable) Incr(ctx context.Context, key string) *IntCmd

// go-redis版本
func (c *cmdable) Incr(key string) *IntCmd
```

**示例**：
```go
func TestRun() {
    ctx := context.Background()
    intVal, err := redisClient.Incr(ctx, "age").Result()
    if err != nil {
        fmt.Println("incr cmd failed:", err.Error())
    } else {
        // 
        fmt.Println("incr ok:", intVal)
    }
}
```


# 2 redis.baseCmd的方法
```go
// 返回执行命令的名字，返回的值包括`set`
func (cmd *baseCmd) Name() string

/*返回执行的redis命令*/
func (cmd *baseCmd) Args() []interface{}

/*设置redis命令执行的错误信息*/
func (cmd *baseCmd) SetErr(e error)

/*获取redis命令执行的错误信息*/
func (cmd *baseCmd) Err() error

func (cmd *baseCmd) SetFirstKeyPos(keyPos int8)

```

# 3 redis.StatusCmd的方法
redis.StatusCmd继承自redis.baseCmd，因此redis.baseCmd的方法redis.StatusCmd都有。
```go
/* 返回redis命令的执行结果 
string: 返回对应redis命令的输出结果
error:  错误信息*/
func (cmd *StatusCmd) Result() (string, error)

func (cmd *StatusCmd) Val() string
```

# 4 redis.StringCmd结构体
redis.StringCmd继承自redis.baseCmd.
```go
type StringCmd struct {
    baseCmd
    val string
}
```

redis.StringCmd的方法
```go
// 获取redis命令的输出值
func (cmd *StringCmd) Val() string

// 设置redis命令的输出值
func (cmd *StringCmd) SetVal(val string)

func (cmd *StringCmd) Result() (string, error)
```

# 5 redis.IntCmd

```go
// 将redis的命令的返回值转化为string
func (cmd *IntCmd) String() string


```