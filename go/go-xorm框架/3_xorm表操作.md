[toc]

# 1 创建表
xorm中通过结构体按照一定的规则转化为mysql中的表信息。
struct与数据库表的转化规则：
示例：
```go
package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
)
// 数据库连接对象
var engine *xorm.Engine = nil
var err error = nil
func init() {...连接数据库...}

type Book struct {
    ID     uint64
    Name   string
    Author string
    Type   int8
    Price  float32
}

func main() {
    // main函数结束后关闭数据库连接
    defer engine.Close()
    /*根据Book结构体信息创建数据库表*/
    err = engine.Sync2(Book{})
    if err != nil {
        fmt.Println("create Table Book failed:", err.Error())
        return
    }

    /*也可以使用CreateTables方法创建表*/
    // err = engine.CreateTables(&Book{})

    /*也可以使用Sync方法创建表*/
    // err = engine.Sync(Book{})
}
```
生成的数据库为：
```shell
mysql> desc book;
+--------+---------------------+------+-----+---------+-------+
| Field  | Type                | Null | Key | Default | Extra |
+--------+---------------------+------+-----+---------+-------+
| i_d    | bigint(20) unsigned | YES  |     | NULL    |       |
| name   | varchar(255)        | YES  |     | NULL    |       |
| author | varchar(255)        | YES  |     | NULL    |       |
| type   | int(11)             | YES  |     | NULL    |       |
| price  | float               | YES  |     | NULL    |       |
+--------+---------------------+------+-----+---------+-------+
5 rows in set (0.00 sec)
```

**Sync和Sync2的区别**：
在xorm中，Sync和Sync2都是用于创建表的方法，它们之间的区别如下：

1. 参数类型不同
Sync方法的参数是结构体类型，而Sync2方法的参数是表名字符串和结构体类型。

2. Sync2方法支持更多的选项
Sync方法仅支持结构体中字段定义的选项，例如not null、unique等。而Sync2方法除了支持结构体中的字段选项，还支持其它一些选项，例如表选项、存储引擎选项等。

3. 兼容性问题
Sync方法是早期版本的API，经过多个版本的迭代之后，Sync2方法逐渐成为了主流API。虽然Sync方法仍然能正常工作，但在一些极端情况下，可能会发生不兼容的问题。

因此，如果在开发中需要使用xorm创建表，建议使用Sync2方法，以获得更好的兼容性和功能。

在使用 `Sync` 和 `Sync2` 方法创建表时，需要注意以下几点：

- 表名称是根据 Go 结构体中的 `TableName` 标签中定义的名称来创建的，如果没有定义 `TableName`，则会使用结构体名称来作为表名称。
- 列名称是根据 Go 结构体中的字段名称来创建的。
- 默认情况下，Go 中的 `int32`、`int64`、`uint32`、`uint64` 类型都会被映射为 MySQL 中的 `BIGINT` 类型。
- 默认情况下，Go 中的 `string` 类型会被映射为 MySQL 中的 `VARCHAR(255)` 类型。
- 如果需要映射到 MySQL 中的其他数据类型，则需要通过在字段上添加 `xorm` 标签来指定数据类型。

总之，`Sync` 和 `Sync2` 方法都是 xorm 中用来创建表的方法，但是它们的参数不同。在使用这些方法时，需要注意表名称和列名称的定义以及字段类型的映射，以确保表结构与预期一致。

## 1.1 结构体与数据库表的映射规则
结构体转化为表时，结构体名字和字段名都会按照一定的规则映射到数据库表的表名和字段名。
结构体的名称默认是驼峰命名(struct)-->小写+下划线
例如，
|结构体名称|表名|
|---------|----|
|Book|book|
|MyBook|my_book|
同时也可以通过设置自定义的表名

### 1.1.1 通过TableName方法设置表名
通过给结构体定义```func TableName() string```方法来设置表名，这个函数返回的字符换就是表的名字
  ```go
    /*给Book结构体添加TableName方法，该方法返回的字符串就用作表名*/
    func (book *Book) TableName() string {
        return "MyBook"
    }
  ```
### 1.1.2 通过engine.Table方法进行命名
通过xorm.Engine.Table(...)进行指定命名的优先级是最高的
示例：
```go
func main() {
	defer engine.Close()
    /*将表名命名为YYY*/
    err = engine.Table("YYY").Sync2(Book{})
	if err != nil {
		fmt.Println("create Table Book failed:", err.Error())
		return
	}
}
```

### 1.1.3 使用xorm提供的映射
- 使用xorm提供的映射
xorm提供了3种映射配置，**<font color=green>names.SnakeMapper ， names.SameMapper ， names.GonicMapper</font>**

**这三种映射不仅会作用到表名中，也会作用到字段名上。**

|映射|描述|
|----|----|
|names.SnakeMapper| 支持struct为驼峰式命名，表结构为下划线命名之间的转换这个是默认的Maper; </br>例如(表名-->字段名):ID --> i_d,  Name --> name,  CreateTime --> create_time|
|names.SameMapper|支持结构体名称和对应的表名称 以及 结构体field名称与对应的表字段名称相同的命名；</br>例如(表名-->字段名):ID --> ID, CreateTime --> CreateTime|
|names.GonicMapper|和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d|
**示例**:
```go
package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
    "xorm.io/xorm/names"
)

var engine *xorm.Engine = nil
var err error = nil

func init() {
    engine, err = xorm.NewEngine("mysql", "test:zxcvbnm1997@(127.0.0.1:3306)"+
    "/test_database?charset=utf8&parseTime=True")
    if err != nil {
        fmt.Println("create database engine failed:", err.Error())
        return
    }
    if err = engine.Ping(); err != nil {
        fmt.Println("connect mysql failed:", err.Error())
        return
    }
    /*设置使用 SmameMapper映射*/
    engine.SetMapper(names.SameMapper{})
    fmt.Println("connected mysql database")
}
func main(){
    defer engine.Close()
}
```

### 1.1.4 前缀映射
通过
```go
names.NewPrefixMapper(names.SnakeMapper{}, "prefix")
```
可以创建一个在SnakeMapper的基础上命名中添加统一的前缀，当然也可以把SnakeMapper更换成SameMapper、GonicMapper或者你自定义的Mapper
**示例**:
```go
func init() {
    engine, err = xorm.NewEngine("mysql", "test:zxcvbnm1997@(127.0.0.1:3306)/test_database"+
        "?charset=utf8&parseTime=True")
    if err != nil {
        fmt.Println("create database engine failed:", err.Error())
        return
    }
    if err = engine.Ping(); err != nil {
        fmt.Println("connect mysql failed:", err.Error())
        return
    }
    /*在SameMapper的基础上添加一个前缀 `My_`*/
    prefixMapper := names.NewPrefixMapper(names.SameMapper{}, "My_")
    engine.SetMapper(prefixMapper)

    fmt.Println("connected mysql database")
}
```
### 1.1.5 后缀映射
```go
//在创建的数据库表名及其字段名上添加后缀_Same
mapper := names.NewSuffixMapper(names.SameMapper{}, "_Same")
engine.SetMapper(mapper)
```
### 1.1.6 缓存映射
通过names.NewCacheMapper(names.SnakeMapper{})可以创建一个组合了其他映射规则，起到在内存中缓存曾经映射过的命名映射。
### 1.1.7 通过标签来进行命名映射
如果所有的命名都是按照Mapper的映射来操作的，那当然是最理想的，但是如果碰到某个表名或者某个字段跟映射规则不匹配时，我们就需要使用标签来改变。
通过struct中对应的Tag中使用
`xorm:"'column_name'"`可以使得该field对应的Column名称指定名称
**示例**：
```go
type Book struct {
    ID     uint64
    /*生成表后，Name字段在表中的名字就是"FullName"*/
    Name   string `xorm:"FullName"`
    Author string
    Type   int8
    Price  float32
}
```
### 1.1.8 各个映射的优先级
当各个映射都存在时，它们的优先级为
**对于表名：**
1.  engine.Table()指定的临时表名优先级最高
2.  TableName() string其次
3. Mapper自动映射的表名优先级最后

**对于字段名字段名**
1. 结构体Tag指定字段名优先级最高
2. Mapper自动映射的字段名优先级较低

## 1.2 字段标签
可用在声明结构体的时候，在每个字段的后面添加标签来对该字段进行描述，使得该结构体映射为数据库表时，对应字段也会添加标签所描述的属性。
标签各个属性描述用空格分开
### 1.2.1 设置字段名
如与其它关键字冲突，请使用单引号括起来。
```go
type Book struct {
    /*该字段映射为数据库表的字段名为id*/
    ID     uint64 `xorm:"id"`
    /*使用单引号将字段名包裹(单引号在内围，否则不生效)*/
    Name   string `xorm:"'FullName'"`
}
```
### 1.2.2 声明为主键
通过声明 `xorm:"pk"`标签来将该字段描述为主键。
**注意**： 对于int类型的Column使用pk标签后还需要是使用autoincr标签来进行声明为自增，不然在插入记录的时候，如果主键的值跟表中已有数据的值重复，则会插入失败。不指定struct种主键的值会使用默认值，头一次插入会成功，但第二次就会出现主键重复。
```go
type Book struct {
    /*声明为主键，字段名映射为id*/
    ID     uint64 `xorm:"id pk autoincr"`
    
    Name   string `xorm:"FullName""`
}
```
### 1.2.3 声明类型
```go
type Book struct {
    ID    uint64  `xorm:"id pk autoincr"`
    Name  string  `xorm:"char(64)"`
    Is    bool    `xorm:"bool"`
    Desc  string  `xorm:"varchar(255)"`
    Price float32 `xorm:"float(6,2)"`
}
```
对应的数据库表
```shell
mysql> desc book;
+-------+---------------------+------+-----+---------+----------------+
| Field | Type                | Null | Key | Default | Extra          |
+-------+---------------------+------+-----+---------+----------------+
| id    | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| name  | char(64)            | YES  |     | NULL    |                |
| is    | tinyint(1)          | YES  |     | NULL    |                |
| desc  | varchar(255)        | YES  |     | NULL    |                |
| price | float(6,2)          | YES  |     | NULL    |                |
+-------+---------------------+------+-----+---------+----------------+
5 rows in set (0.00 sec)
```
xorm类型与数据库类型对应关系如下：
|xorm|go类型|MySQL|取值范围|示例|备注|
|------|----|-----|----|---|-----|
|bit|任意|bit|-|```Value bool `xorm:"bit(1)"` ```|bit的长度最大64;</br>传入的值长度不能</br>超过声明的bit的长度,否则会插入失败|
|tinyint|bool/整形|tinyint|1个字节,范围-128~127|```Value int8 `xorm:"tinyint"` ```||
|smallint|bool/整形|smallint|2个字节,范围(-32768~32767)|```Value int16 `xorm:"smallint(1)"` ```||
|mediumint|bool/整形|mediumint|3个字节 范围(-8388608~8388607)|```Value int32 `xorm:"mediumint"` ```||
|int|bool/整形|int|4个字节 范围(-2147483648~2147483647)|```Value bool `xorm:"int"` ```||
|integer|bool/整形|integer|4个字节 范围(-2147483648~2147483647)|```Value int32  `xorm:"integer"` ```|1.int是作为对象，直接存储数值。</br>2.integer需要实例化对象，实际上是生成一个指针指向对象的地址。</br>3.在mysql中，integer的数据类型是引用数据类型，是对int的封装。|
|bigint|bool/整形/string|bigint|8个字节 范围(+-9.22*10的18次方)|```Value string `xorm:"bigint"` ```|如果go类型是string,那么值一定要是数字|
|float(m,d)|float32/float64|float(m,d)|单精度浮点型 8位精度(4字节) m总个数，d小数位|```Value float32 `xorm:"float(5,3)"`|1. 对于小数位超过的数字，会四舍五入。</br>对于(5,3)，表名数字的总位数是5，小数站3位，因此整数最多部分只能有2位，对于整数部分超过2位的(如100.21)会插入失败。</br>2.存储的精度有误差，甚至在赋值的浮点型很大时，存储却为0|
|double(m,d)|float32/float64|double(m,d)|双精度浮点型 16位精度(8字节) m总个数，d小数位|```Value float64 `xorm:"double(28,6)"` ```||
|decimal(m,d)|float32/float64/string|decimal(m,d)|参数m<65 是总个数，d<30且 d<m 是小数位。|```Value string `xorm:"decimal(45,6)"` ```| go的类型最好使用string类型，且值为数字，这样更精确|
|char|string|char|固定长度,最多255个字符|```Name string `xorm:"char(64)"` ```||
|varchar|string|varchar|固定长度,最多65535个字符|```Name string `xorm:"varchar(64)"` ```||
|tinytext|string|tinytext|可变长度，最多255个字符|```Value string `xorm:"tinytext"` ```||
|text|string|text|可变长度，最多65535个字符|```Value string `xorm:"text"` ```||
|mediumtext|string|mediumtext|可变长度，最多2的24次方-1个字符|```Value string `xorm:"mediumtext"` ```||
|longtext|string|longtext|可变长度，最多2的32次方-1个字符|```Value string `xorm:"longtext"` ```||

**关于mysql中的浮点型**：
mysql的float类型是单精度浮点类型不小心就会导致数据误差. 单精度浮点数用4字节（32bit）表示浮点数 采用IEEE754标准的计算机浮点数，在内部是用二进制表示的 如：7.22用32位二进制是表示不下的。 所以就不精确了。 mysql中float数据类型的问题总结 对于单精度浮点数Float: 当数据范围在±131072（65536×2）以内的时候，float数据精度是正确的，但是超出这个范围的数据就不稳定，没有发现有相关的参数设置建议：将float改成double或者decimal，两者的差别是double是浮点计算，decimal是定点计算，会得到更精确的数据

### 1.2.4 添加索引

**添加单索引**:
```go
type Book struct {
    ID     int32  `xorm:"pk autoincr"`
    Name   string `xorm:"char(12)"`
    Author string `xorm:"char(12) index"`
}
```
**添加联合索引**：
```go
type Book struct {
    ID     int32  `xorm:"pk autoincr"`

    /*添加联合索引，索引名为IDX_Book_my_index*/
    Name   string `xorm:"varchar(12) index(my_index)"`
    Author string `xorm:"char(12) index(my_index)"`
}
```

**创建唯一索引**:
```go
type Book struct {
	ID     int32  `xorm:"pk autoincr"`
	Name   string `xorm:"char(12) unique"`
	Author string `xorm:"char(12)"`
}
```

**创建联合唯一索引**：
```go
type Book struct {
    ID     int32  `xorm:"pk autoincr"`
    Name   string `xorm:"char(12) unique(my_unique)"`
    Author string `xorm:"char(12) unique(my_unique)"`
}
```
定义好以上结构体后，通过 ```Engine.Sync2(xxx)```创建表的时候，会自动创建索引

**CreateIndexes方法创建索引**：
也可以通过Engine.CreateIndexes(xxx)方法创建索引
```go
type Book struct {
    ID     int32  `xorm:"pk autoincr"`
    Name   string `xorm:"char(12) unique(my_unique)"`
    Author string `xorm:"char(12) unique(my_unique)"`
}

func main() {
    engine.ShowSQL(true)

    /*根据Book中的标签信息创建索引*/
    if engine.CreateIndexes(&Book{}) != nil {
        fmt.Println("create index failed:", err.Error())
    }
}
```

### 1.2.5 外键声明
在 xorm 中，可以使用 struct tag `xorm:"foreign"` 来给字段声明外键。

假设你的表 A 中需要对表 B 的字段 b_id 创建外键约束，那么在声明表 A 的结构体时，可以在字段 b_id 上添加 `foreign` tag，指定关联的表名和关联的字段名，例如：

```go
type A struct {
    Id    int64  `xorm:"pk autoincr"`
    BId   int64  `xorm:"notnull foreign(b.id)"` // 表示 A.BId 关联到 B.id
    Name  string `xorm:"varchar(50)"`
}

type B struct {
    Id    int64  `xorm:"pk autoincr"`
    Name  string `xorm:"varchar(50)"`
}
```

在这个示例中，`BId` 字段的类型为 `int64`，并且添加了 `foreign` tag，表示需要为该字段创建外键约束。`(b.id)` 表示该外键约束关联到表 B 的字段 id 上。这里的 `b` 表示要关联的表名，可以是一个已经声明的结构体的名称，或者是字符串类型的表名。

通过这种方式，你可以方便地在 xorm 中声明外键约束，简化了 SQL 语句的编写和维护。 **需要注意的是，xorm 只提供了外键的声明功能，实际的外键约束需要在数据库中使用 SQL 语句创建。** 因此，在使用 xorm 的同时，仍然需要对 SQL 语句的创建和执行有一定的了解。

# 2 添加数据
向数据库中擦汗如数据使用```engine.Insert(...)```方法, Insert方法的参数可以是**一个或多个Struct、一个或多个struct指针、一个或多个Struct的Slice、一个或多个Struct的切片指针**。
如果传入的是Slice并且当数据库支持批量插入时，Insert会使用批量插入的方式进行插入。
```sql
func main() {
    defer engine.Close()
    engine.ShowSQL(true)
    user := []User{
        {Name: "tom", Age: 33},
        {Name: "rr", Age: 44},
    }
    number, err := engine.Insert(&user)
    if err != nil {
        fmt.Println("insert record failed:", err.Error())
    } else {
        fmt.Printf("insert %d record\n", number)
    }
}
```

# 3 删除数据
## 3.1 删除条件
**通过Id函数**：
## 3.2 软删除
执行软删除时，并不会真正的删除，如果要使用软删除，需要在创建表对应的Struct的时候指定deleted字段，例如
对于结构体
```sql
type User struct{
    ID int64 `xorm:"pk INT UNSIGNED autoincr"`
    Name string `xorm:"char(32) NOT NULL"`
    Age int `xorm:"tinyint NOT NULL"`
    DeletedTime time.Time `xorm:"deleted"`
};
```
如果执行Delete删除函数：
```go
func main() {
    defer engine.Close()
    engine.ShowSQL(true)
    /*会将user的字段作为查找条件*/
    user := User{Name: "tom", Age: 33}
    number, err := engine.Delete(user)
    if err != nil {
        fmt.Println("insert record failed:", err.Error())
    } else {
        fmt.Printf("insert %d record\n", number)
    }
}
```
会执行一下SQL语句更新DeleteTime字段，而不是删除该记录
```sql
UPDATE `User` SET `DeletedTime` = ? WHERE `Name`=? AND `Age`=? AND (`DeletedTime`=? OR `DeletedTime` IS NULL)
```
这样```Delete```函数执行完成后，user对应的记录还是存在于数据库中，只是DeleteTime字段被标记成了删除的时间。

## 3.3 硬删除
那么如果记录已经被标记为删除后，要真正的获得该条记录或者真正的删除该条记录，需要启用```Unscoped```，如下所示
```sql
func main() {
    defer engine.Close()
    engine.ShowSQL(true)
    user := User{Name: "tom", Age: 33}
    /*直接删除，而不是设置DeletedTime字段*/
    number, err := engine.Unscoped().Delete(user)
    if err != nil {
        fmt.Println("insert record failed:", err.Error())
    } else {
        fmt.Printf("insert %d record\n", number)
    }
}
```

**注意**：
如果是对于没有声明```xorm:"deleted"```标签的，会直接删除对应的记录。
即使在开始声明了xorm:"deleted"标签，如果后续的某次执行中把这个标签从结构体声明中删除了，那么本次的```Delete```方法会***直接从数据库中删除对应的记录**。

## 3.4 修改记录
更新记录使用 engine.Update(...)方法
```go
type User struct {
    ID          int64     `xorm:"pk INT UNSIGNED autoincr"`
    Name        string    `xorm:"char(32) NOT NULL"`
    Age         int8      `xorm:"tinyint NOT NULL"`
    DeletedTime time.Time `xorm:"deleted"`
}


user := User{Name: "jake", Age: 9}
number, err := engine.Where("ID=6").Update(user)

/*执行的SQL语句*/
UPDATE `User` SET `Name` = ?, `Age` = ? WHERE (ID=6) AND (`DeletedTime`=? OR `DeletedTime` IS NULL)
```
Update的参数可以是结构体或者结构体指针。
Update方法会自动根据结构体中的主键的值来生成条件，此外如果结构体中有标记为`xorm:"version"`的字段，那么也同时会加上version的判断条件。
**注意：**
- 不能只单独使用Update()方法，这会更新整张表
```go
effected, err := engine.Update(&user) // 会将整张表更新成user的内容，不要这样用
```
- Update参数不能是结构体切片或结构体切片指针。

# 4 查询
## 4.1 查询单条数据
查询单条数据可以使用xorm.Engine下的`Get`方法：
```go
func (engine *Engine) Get(beans ...interface{}) (bool, error)
```
**参数:***
- <font color=blue>beans</font>: 数据库表所对应的结构体的**指针**，查询到的数据信息会存放改这里。如果传入的是结构体类型会报错。

**返回值**
- <font color=blue>bool: </font> true: 查询到记录
  false: 没有查询到信息
- <font color=blue>error: </font>如果查询遇到错误，返回失败的错误；如果查询成功则返回nil。没有查询到对应记录且没有报错也会返回nil

**示例**
```go
var b bool
var err error
user := User{ID: 1, Name: "jake"}
b, err = engine.Get(user)
```

如果传入`Get`方法中的结构体指针对应的结构体没有初始化，可以通过`Id`方法来指定主键的值作为查询条件,如下
```go
var b bool
var err error
user := User{}
b, err = engine.ID(1).Get(&user)
```

**注意**：
- 通过Get方法查询时，会根据参数user中的非空字段作为查询条件来进行查询，如上例子会以ID=1 AND Name="jake"作为查询条件。

## 4.2 使用Where方法来指定查询条件

```go
var b bool
var err error
user := User{ID: 1, Name: "jake"}
/*会以personId=1 AND ID=1 AND Name="jake"一起作为查询条件。
这里的personId=1是指实际数据库中的personId字段的值，
ID=1 和 Name="jake"是指User结构体中ID和Name字段对应数据库表中的字段的值*/
b, err := engine.Where("personId=1").Get(&user)
```

`Where`方法还可以使用？作为占位符.
```go
user := User{}
id := 1
name := "jake"
b, err := engine.Where("personId=? AND personName=?", id, name).Get(&user)
```

## 4.3 查询多条数据
查询多条数据使用**Find**方法，Find方法的第一个参数只能为**slice的指针**或**Map指针**。第二个参数可选，为查询的条件struct的指针或struct。
```go
func (session *Session) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error
```

1. 传入Slice用于返回数据:
```go
user := []User{}
/*这回查询所有信息，相当于`select * from <table>`*/
err := engine.Find(&user)

user1 := make([]User, 0)
err = engine.Where("personId > ?", 1).Find(&user1)

user := make([]*User, 0)
err := engine.Find(&user, User{ID: 2})


/*如果只选择单个字段，也可使用非结构体的Slice*/
ints := make([]int, 0)
// 查找Person表下的全部personAge字段的值
err := engine.Table("Person").Cols("personAge").Find(&ints)
```

2. 传入Map用户返回数据，map必须为map[int64]Userinfo的形式，map的key为id，因此对于复合主键无法使用这种方式。
```go
users := make(map[int64]User)
err := engine.Find(&users)

users := make(map[int64]*User)
err := engine.Find(&users)
```

## 4.4 统计记录的个数
统计数据使用Count方法，Count方法的参数为struct或者struct的指针并且成为查询条件。
```go
func (engine *Engine) Count(bean ...interface{}) (int64, error)
```
**参数**
<font color=blue>bean</font>: 查询条件的Struct(指针)。
**返回值**
- <font color=blue>int64</font>: 满足条件的记录的个数
- <font color=blue>error</font>: 报错信息

**示例**
```go
// 计算ID=1的记录有多少条
user := User{ID: 1}
num, err := engine.Count(user)
```

## 4.5 Exist系列方法
```go
func (engine *Engine) Exist(bean ...interface{}) (bool, error)
```
**参数**
<font color=blue>bean</font>: 结构体指针，如果传入结构体会报错。
**返回值**
- <font color=blue>bool</font>: true:存在； false:部存在
- <font color=blue>error</font>: 没有报错则返回nil


**示例1：查看是否有对应的记录**：
```go
user := User{ID: 11}
var isExist bool
isExist, err := engine.Exist(&user)
```

**示例2：查看对应的表是否存在**：
```go
// 检查Person表是否存在，如果不存在err返回不为nil
isExist, err := engine.Table("Person").Exist()
```

## 4.6 对查询到个每条记录执行操作
```go
func ShowName(idx int, bean interface{}) error {
	/*这里转化的类型是main函数中声明user的指针类型*/
    user := bean.(*User)
	fmt.Printf("the %d person's name is%s\n", idx, user.Name)
	return nil
}

func main() {
	engine.ShowSQL(true)
    /*查询到的每一条记录都会赋值给user然后传入ShowName函数中执行，
    当Iterate函数执行完后从，将user设置成传入Iterate之前的状态。
    如果user的字段不为空，还会将不为空的字段作为查询条件*/
	user := User{}
	err := engine.Where("personAge > ?", 10).Iterate(&user, ShowName)
	if err != nil {
		fmt.Println("query failed with:", err.Error())
	}
}
```

## 4.7 join方法
```go
func (engine *Engine) Join(joinOperator string, tablename interface{}, condition interface{}, args ...interface{}) *Session
```
**参数**
- <font color=blue>joinOperator</font>:连接类型，当前支持INNER, LEFT OUTER, CROSS中的一个值
- <font color=blue>tablename</font>:string类型的表名，表对应的结构体指针或者为两个值的[]string，表示表名和别名
- <font color=blue>args</font>:连接条件

**返回值**
*Session ： Session，说明Join后面还可继续操作

**示例**
```go
/*数据库中没有UserBook对应的表，但有User和Book对应的表*/ 
type UserBook struct {
    User `xorm:"extends"`
    Book `xorm:"extends"`
}

func main() {
    engine.ShowSQL(true)
    userBook := []UserBook{}
    /*将Book和Person表中满足Where条件的记录拼接起来放到UserBook结构体中*/
    err := engine.Table("Person").Join("INNER", "Book", "Book.author=Person.personName").Find(&userBook)

    // 也可以使用Sql来代替
    //结构体中extends标记对应的结构顺序应和最终生成SQL中对应的表出现的顺序相同。
    //err := engine.SQL("select Person.*,Book.* FROM Person,Book WHERE Person.personName=Book.author").Find(&userBook)
    if err != nil {
        fmt.Print("error:", err.Error())
    } else {
        fmt.Println("user=", userBook)
    }
}
```

当然，如果表名字太长，我们可以使用别名：
```go
engine.Table("user").Alias("u").
	Join("INNER", []string{"group", "g"}, "g.id = u.group_id").
	Join("INNER", "type", "type.id = u.type_id").
	Where("u.name like ?", "%"+name+"%").Find(&users, &User{Name:name})
```

**拓展**
Join方法后面还可继续接Join来联合多表一起查询

## 4.8 Rows方法 
Rows方法和Iterate方法类似，提供逐条执行查询到的记录的方法，不过Rows更加灵活好用。

```go
user := new(User)
rows, err := engine.Where("id >?", 1).Rows(user)
if err != nil {
}
defer rows.Close()
for rows.Next() {
    err = rows.Scan(user)
    //...对user记录进行操作
}
```

**注意**
- Rows的参数只能是**结构体指针**

## 4.9 Sum系列方法
 
求和数据可以使用Sum, SumInt, Sums 和 SumsInt 四个方法，Sums系列方法的参数为struct的指针并且成为查询条件。
```go
Sum 求某个字段的和，返回float64
type SumStruct struct {
    Id int64
    Money int
    Rate float32
}

ss := new(SumStruct)
total, err := engine.Where("id >?", 1).Sum(ss, "money")
fmt.Printf("money is %d", int(total))
```
SumInt 求某个字段的和，返回int64
```go
type SumStruct struct {
    Id int64
    Money int
    Rate float32
}

ss := new(SumStruct)
total, err := engine.Where("id >?", 1).SumInt(ss, "money")
fmt.Printf("money is %d", total)
```

Sums 求某几个字段的和， 返回float64的Slice
```go
ss := new(SumStruct)
totals, err := engine.Where("id >?", 1).Sums(ss, "money", "rate")

fmt.Printf("money is %d, rate is %.2f", int(total[0]), total[1])
SumsInt 求某几个字段的和， 返回int64的Slice
ss := new(SumStruct)
totals, err := engine.Where("id >?", 1).SumsInt(ss, "money")

fmt.Printf("money is %d", total[0])
```