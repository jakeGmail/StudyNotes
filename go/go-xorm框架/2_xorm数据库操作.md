[toc]
# 1 连接数据库
```go
package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
)

var engine *xorm.Engine = nil
var err error

/*执行main函数前会先执行init函数，因此连接数据库的操作放在init函数中*/
func init() {
    /*创建数据库连接对象
    连接主机127.0.0.1:3306的数据库test_database, 
    登录账号和密码为test、zxcvbnm1997。
    字符格式uft8*/
    engine, err = xorm.NewEngine("mysql", "test:zxcvbnm1997@(127.0.0.1:3306)/test_database"+
        "?charset=utf8&parseTime=True")
    if err != nil {
        fmt.Println("create database engine failed:", err.Error())
        return
    }

    /*ping数据库，如果成功说明连接ok*/
    if err = engine.Ping(); err != nil {
        fmt.Println("connect mysql failed:", err.Error())
        return
    }
    fmt.Println("connected mysql database!")
}
func main() {
    defer engine.Close()
}
```



