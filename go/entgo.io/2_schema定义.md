[toc]


我们通过[1_介绍](1_介绍.md)的命令`go generate ./ent`生成代码后会的得到类似这样的代码
```go
package schema

import "entgo.io/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
```

其中Fields方法定义表的字段， Edges方法定义表之间的关联关系。

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// 定义表的字段名及其约束
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(). // 约束个字段必须为正数
            Range(1,140). // 约束其值在[1, 140]范围内
            Comment("主键"), // 创建表时的字段描述
        field.String("name").
            Default("unknown"). // 设置这个字段的默认值为unknown
            Unique(), // 设置唯一索引
        field.Int32("scole").
            Max(100). // 限制值的最大值
            Min(5). // 限制值的最小值
            Nillable() // 设置可以为NULL
        field.Int("number").
            Negative(). // 限制其值只能为负数
            DefaultFunc(
                func()int{
                    return 12
                }
            )   // 在创建新记录时被调用，以动态生成字段的默认值
    }
}
```

# 2 定义schema
## 2.1 定义表的字段名和类型
以下是定义表的字段名的代码示例
```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type CommercialAccount struct {
	ent.Schema
}

// Fields of the Commercial.
func (CommercialAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"), // int类型
        field.Int8("int8"), // int8类型
		field.Int16("int16"), // int16类型
		field.Int32("int32"), // int32类型
		field.Int64("scole"), // bigint类型
		field.String("name"), // varchar类型
        field.Time("create_time"),   // 时间类型
		field.Text("content"),   // text类型
		field.Enum("type").Values("t1", "t2", "t3"), // 枚举
		field.Float("floatNumber"),    // float64类型
        field.Float32("float32"), // float32类型
		field.Floats("floats"), // returns a new JSON Field with type []float
		field.Bool("isOK"),   // bool类型
		field.UUID("uuid", uuid.UUID{}).Default(func() uuid.UUID {
			return uuid.New()
		}), // uuid类型
        field.Other("amount", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,18)", // 指定该字段在mysql数据库中的类型
		}), // 自定义类型
        field.JSON("json", SS{}), // json类型
		field.Uint8("uint8"),     // uint8类型
        field.Uint16("uint16"),   // uint16类型
		field.Uint32("uint32"),   // uint32类型
		field.Uint64("uint64"),   // uint64类型
        field.Any("any").SchemaType(map[string]string{
			dialect.MySQL:    "json",  // 指定该字段在mysql数据库的类型
			dialect.Postgres: "jsonb", // 指定该字段在postgres数据库的类型
		}), // 任意类型
        
	}
}
type SS struct {
	Id   int
	Name string
}

type Decimal struct {
	decimal.Decimal
}

func (v Decimal) Value() (driver.Value, error) {
	return v.GetValue(), nil
}
```

## 2.2 设置字段的索引
```go
// Indexes 设置索引
func (CommercialAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("scole", "name").Unique(), // 设置唯一联合索引
        index.Fields("uuid"), // 设置索引
	}
}
```

## 2.3 定义字段约束
### 2.3.1 定义字段牧人为空
使用Nillable()方法定义字段默认值为NULL
```go
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("updated_at").Nillable().Comment("记录更新时间"),
    }
}
```

## 2.4 自定义表名（覆盖自动生成代码的默认行为）
```go
// Annotations 自定义表名
func (CommercialAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("commercial_account"), // 自定义表名
	}
}
```



## 2.5 设置Edge关系
### 2.5.1 一对一关系
```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// User schema.
type User struct {
    ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
    return []ent.Field{
        // ...
    }
}

// 一个用户只能有1个宠物
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("pets", Pet.Type).Unique(), // pets是表名
    }
}
```

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
    ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
    return []ent.Field{
        // ...
    }
}

// 一个宠物只能有1个用户
func (Pet) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("owner", User.Type).
            Ref("pets").
            Unique(),
    }
}
```

### 2.3.1 一对多关系
例如用户和宠物的关系是，一个用户可以有多个宠物。这是一对多
多个宠物可以只有相同的用户。这是多对一
```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// User schema.
type User struct {
    ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
    return []ent.Field{
        // ...
    }
}

// 一个用户可以有多个宠物
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("pets", Pet.Type), // pets是表名
    }
}
```

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
    ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
    return []ent.Field{
        // ...
    }
}

// 多个宠物有相同的用户
func (Pet) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("owner", User.Type).
            Ref("pets").
            Unique(),
    }
}
```
可以看到，一个User实体可以有多只宠物，但一个Pet实体只能有一个主人。
在关系定义中，pets边是O2M（一对多）关系，owner边是M2O（多对一）关系。

**注意**
- 在`Edge`方法中的如果使用了Unique()，则声明我们是要声明一对一或者一对多关系。如果From和To都声明了Unique(),则表明是一对一关系

### 2.3.2 多对多关系
可以看到，一个 Group 实体可以有多个用户，一个 User 实体可以有多个Group。
```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// Group schema.
type Group struct {
    ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        // ...
    }
}

// Edges of the group.
func (Group) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("users", User.Type),
    }
}
```

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// User schema.
type User struct {
    ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
    return []ent.Field{
        // ...
    }
}

// Edges of the user.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("groups", Group.Type).
            Ref("users"),
        // "pets" declared in the example above.
        edge.To("pets", Pet.Type),
    }
}
```