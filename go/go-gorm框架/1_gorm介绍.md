[toc]

文档
https://gorm.io/zh_CN/docs/hooks.html

# 1 安装gorm
```go
go get gorm.io/gorm
go get gorm.io/gorm
```

# 2 连接数据库
```go
package Gorm

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var db *gorm.DB
var err error

func init() {
    // 数据库地址：
    // <用户名>:<密码>@<连接方式>(<IP地址和端口号>)/<数据库名称>?<数据库配置>
    dsn := "root:zxcvbnm1997@tcp(127.0.0.1:3306)" +
        "/test?charset=utf8mb4&parseTime=True"
    conf := &gorm.Config{}

    // 连接数据库
    db, err = gorm.Open(mysql.Open(dsn), conf)
    if err != nil {
        log.Fatal("connect mysql failed:", err.Error())
    } else {
        log.Println("gorm connected database")
    }
}
```

# 2 配置数据库
也可以通过以下方式连接数据库
```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var db *gorm.DB
var err error

func init() {
    dsn := "root:zxcvbnm1997@tcp(127.0.0.1:3306)" +
        "/test?charset=utf8mb4&parseTime=True"
    conf := &gorm.Config{}

    // 以指定的配置方式连接数据库
    db, err = gorm.Open(mysql.New(mysql.Config{
        DSN: dsn,
    }), conf)
    if err != nil {
        log.Fatal("connect mysql failed:", err.Error())
    } else {
        log.Println("gorm connected database")
    }
}
```

其中mysql.Config的定义如下
```go
type Config struct {
	DriverName                string
	ServerVersion             string
	DSN                       string
	Conn                      gorm.ConnPool
	SkipInitializeWithVersion bool
	DefaultStringSize         uint
	DefaultDatetimePrecision  *int
	DisableDatetimePrecision  bool
	DontSupportRenameIndex    bool
	DontSupportRenameColumn   bool
	DontSupportForShareClause bool
}
```
**字段含义**:
|字段名|类型|含义|
|------|----|----|
|DriverName|string|用于设置 MySQL 驱动程序的名称。通常情况下这是不需要更改的，因为默认值已经是 “mysql”，表示使用 MySQL 驱动程序。|
|ServerVersion|string|数据库服务器版本，这可以被用来启用或禁用特定于某个版本的功能或行为。|
|DSN|string|包含了数据库连接信息的字符串。它是传给驱动程序的、用于建立数据库连接的配置信息。这个字符串通常包含如主机名、用户名、密码、数据库名等信息。|
|Conn|gorm.ConnPool|这是一个 gorm.ConnPool 接口，可以接受一个已存在的数据库连接池（比如 *sql.DB 的实例）。如果你已经有一个数据库连接和管理方式，可以直接使用此字段。|
|SkipInitializeWithVersion|bool| 设为 true 时，将跳过与数据库版本相关的初始化步骤，比如检查特定版本的功能是否可用。|
|DefaultStringSize|uint|在创建字符串字段时没有指定长度的情况下，所使用的默认长度。在 MySQL 中，默认通常是 255。|
|DefaultDatetimePrecision|*int|默认的 datetime 精度，用于设置时间日期类型字段的精度。它是一个指向 int 类型的指针，可以为 nil，如果为 nil 则使用数据库的默认设置。|
|DisableDatetimePrecision|bool|若设置为 true，则会禁用 datetime 类型的精度设置，对应所有 datetime 类型字段将不会设置精度。|
|DontSupportRenameIndex|bool|若设置为 true，则在迁移时 GORM 将不会使用 RENAME INDEX 语法，而是使用更通用的先删除旧索引然后创建新索引的方式来重命名索引。这对于一些老版本的 MySQL 或一些定制版本可能是必要的。|
|DontSupportRenameColumn|bool|若设置为 true，则在迁移时 GORM 将不会使用 RENAME COLUMN 语法来重命名列。对于不支持该语法的 MySQL 版本，你可能需要启用这一选项。|
|DontSupportForShareClause|bool|若设置为 true，则会禁用 SELECT … LOCK IN SHARE MODE 语句的支持。这通常是为了兼容不支持该语法的 MySQL 版本。|
