[toc]

# 1 获取mongoDB交互框架
```shell
go get go.mongodb.org/mongo-driver/v2/mongo
```

对应的文档为 https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo

# 2 链接数据库
```go
package MongoDbTest

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func Run() {
    // 格式：mongodb://用户名:密码@主机:端口/数据库名  如果没有密码则用户名密码那一列可以不写
	uri := "mongodb://localhost:27017"

	// 链接mongoDB
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// 连接是否正常
	ctx := context.Background()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Println(err.Error())
		}
	}()

	/*获取test数据库下的collection集合*/
	col := client.Database("test").Collection("collection")
	fmt.Println(col.Name())
}
```

## 2.1 创建链接的选项
在创建mongodb的客户端链接时，会提供各个配置选线
```go
package MongoDbTest

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func Run() {
	uri := "mongodb://localhost:27017"

	// 设置mongodb的链接路径
    opt1 := 
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
}

```

# 3 插入记录到数据库
```go
/*获取test数据库下的collection集合*/
col := client.Database("test").Collection("collection")

// 插入数据
resp,err := col.InsertOne(ctx, bson.M{"name": "jake", "age": 27})
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Printf("id=%v\n", resp.InsertedID) // 获取插入数据的ID
```

# 4 查询
## 4.1 查询全部记录
```go
/*获取test数据库下的collection集合*/
col := client.Database("test").Collection("collection")

// 查询集合下的全部document
cursor, err := col.Find(ctx, bson.D{})
if err != nil {
    return
}
for cursor.Next(ctx) {
    result := struct {
        Foo string
        Bar string
    }{}
    err = cursor.Decode(&result)
    if err != nil {
        fmt.Println("decode failed", err.Error())
        continue
    }
    // 获取原始的bson bytes
    raw := cursor.Current
    fmt.Println(raw)
}

 /*除此之外还可以一次性获取全部document*/
var results []struct {
    Foo string
    Bar string
}
err= cursor.All(ctx, &results)
if err != nil {
    return
}
```

## 4.2 查询一条记录
```go
// 查询单条数据
result := struct {
    Foo string
    Bar string
}{}
filter := bson.D{{"name", "jake"}} // 查询条件，查询key为name, value为jake的记录
err = col.FindOne(ctx, filter).Decode(&result)
```

# 5 更新
## 5.1 更新单个文档
```go
// 更新单个文档
filter := bson.D{{"_id", 2158}} // 过滤条件
update := bson.D{{
		"$set", bson.D{
			{"name", "Mary Wollstonecraft Shelley"},
			{"role", "Marketing Director"},
		}}, 
		
		{"$inc", bson.D{
			{"bonus", 2000},
		}},
	} // 更新的内容

updateResult, err := col.UpdateOne(ctx, filter, update)
if err != nil {
    return
}
updateResult, err = col.UpdateByID(ctx, updateResult.UpsertedID, update)
```

## 5.2 更新多个文档
```go
// 更新多个文档
filter = bson.D{{"_id", 2158}}
newDoc = bson.D{{"<field>", "<value>"}, {"<field>", "<value>"}}
updateResult,err = collection.UpdateMany(ctx, filter, newDoc)
```

## 5.3 替换单个文档
```go
// 替换单个文档
filter = bson.D{{"_id", 2158}}
newDoc := bson.D{{"<field1>", "<value1>"}, {"<field2>", "<value2>"}}

collection := client.Database("test").Collection("collection")
replaceResult, err := collection.ReplaceOne(ctx, filter, newDoc)
if err != nil {
    return
}
fmt.Println("匹配到的文档数量：", replaceResult.MatchedCount)
fmt.Println("操作修改到的文档数量：", replaceResult.ModifiedCount)
fmt.Println("该操作更新或插入的文档数量：", replaceResult.UpsertedCount)
fmt.Println("更新或插入文档的 _id，如果没有则为 nil：", replaceResult.UpsertedID)
```

# 6 删除
使用删除操作从 MongoDB 中删除数据。删除操作由以下方法组成：
DeleteOne()，会删除与筛选器匹配的第一个文档
DeleteMany()，删除与过滤器匹配的所有文档

```go
// 匹配并删除 length（长度）大于 300 的文档
filter = bson.D{{"length", bson.D{{"$gt", 300}}}} 
opts := options.DeleteMany().SetHint(bson.D{{"_id", 1}}) // 指示此方法使用 _id 作为索引
result, err := collection.DeleteMany(context.TODO(), filter, opts) // 删除匹配到的所有文档
if err != nil {
    panic(err)
}
```

# 7 错误处理
```go
import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)
mongo.IsDuplicateKeyError(err) // 是否是唯一性校验错误
mongo.IsNetworkError(err) // 是否是网络错误
mongo.IsTimeout(err) // 是否是超时错误
```

