[toc]

# 1 jwt介绍
JWT 全名为 JSON Web Token，是在网络应用间传递信息的一种基于 JSON 的开放标准。它可以通过数字签名保证传输的数据不被篡改或损坏，从而达到在身份验证和信息交流方面的安全性目标。

JWT 由三个部分组成：头部、载荷和签名。头部中通常指定 JWT 使用的加密算法。载荷是 JWT 的主体部分，通常包含有关用户或应用程序的信息。签名用来验证数据的完整性，确保信息没有被篡改或损坏。

在应用程序中，当一个用户成功登录后，服务器可以生成一个 JWT 并将其返回给客户端。客户端可以将此令牌存储起来，并在之后的请求中将其发送给服务器。服务器可以验证 JWT 的签名，以确保用户已登录，并可以从 JWT 中提取出有关用户或会话的信息，以便在请求中进行身份验证和进行其他操作。

JWT 具有以下优点：

- 可以跨不同的编程语言、平台和系统进行使用。
- 不需要在服务器端存储会话状态，从而降低了服务器的负担和提高了可伸缩性。
- JWT 是一种开放标准，由于其开放性，能够很好地适应不断变化的需求和环境。

但是，使用 JWT 也需要注意安全问题。例如，一旦将 JWT 发送给客户端，为了防止它被窃取或窃听，需要采取合适的措施，如密钥保护或加密传输。此外，在生成 JWT 时需要避免任何敏感信息，如密码等。

# 2 jwt格式
JWT (JSON Web Token) 由三个部分组成，这三部分分别是 Header（头部）、Payload（载荷）和 Signature（签名），它们之间由两个点 (`.`) 分隔。其格式如下：

```
aaaaa.bbbbb.ccccc
```

每个部分具体含义如下：

1. **Header（头部）**
   
   Header 通常由两部分组成：token 类型 (`typ`) 和所使用的哈希算法 (`alg`)。这部分的信息被编码为 Base64Url，以告诉验证服务器 JWT 使用的签名算法。

   例如：

   ```json
   {
     "alg": "HS256",
     "typ": "JWT"
   }
   ```

   然后 Header 会被 Base64Url 编码生成 JWT 的第一部分。

2. **Payload（载荷）**
   
   Payload 部分包含了所谓的 claims（声明），这些声明是关于实体（通常是用户）和其他数据的声明。Claims 可以包含多个预定义的字段（例如：`sub`（主题），`iss`（发行人），`exp`（过期时间）等）和自定义的字段。

   例如：

   ```json
   {
     "sub": "1234567890",
     "name": "John Doe",
     "admin": true
   }
   ```

   类似 Header，Payload 也会被 Base64Url 编码为 JWT 的第二部分。

3. **Signature（签名）**
   
   JWT 的最后一部分是签名，它由 Header 和 Payload 编码后的字符串通过一个密钥（对于 HMAC 算法）或一对 RSA/ECDSA 私钥公钥使用 Header 中指定的算法进行签名生成。

   例如，如果使用 HMAC SHA-256 算法，Signature 可以这样生成：

   ```plaintext
   HMACSHA256(
     base64UrlEncode(header) + "." +
     base64UrlEncode(payload),
     secret)
   ```

   最终，Signature 被编码为 Base64Url，构成 JWT 的第三部分。

完整的 JWT 看起来类似于下面这样：

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
```

这里的 `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9` 是编码的 Header，`eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9` 是编码的 Payload，而 `TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ` 是 Signature。

验证时，服务器会用相同的算法和密钥在接收到的 Header 和 Payload 上运行签名算法，以确保签名部分匹配且数据未被篡改。

# 3 jwt包
在go语言中，生成和解析jwt的包可以用
```shell
go get github.com/golang-jwt/jwt/v5
```

**使用示例**:
```go
package JWT

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// 用于给jwt签名的key
var signKey = []byte("test")

func JwtTest() {
    // 根据选择的加密算法创建token对象
    hs256Token := jwt.New(jwt.SigningMethodHS256)

    // 给jwt添加claim(声明)，即消息体
    claims, ok := hs256Token.Claims.(jwt.MapClaims)
    if !ok {
        fmt.Println("type to jwt.MapClaims failed")
        return
    }
    claims["name"] = "jake"
    claims["age"] = 23
    claims["exp"] = time.Now().Unix() + int64(time.Minute*5)

    // 对jwt内容进行签名
    encodeStr, err := hs256Token.SignedString(signKey)
    if err != nil {
        log.Fatal("sign failed:", err.Error())
    }
    fmt.Println("encodeStr=", encodeStr)


    // 接下来是解析jwt
    parsedToken, err := jwt.Parse(encodeStr, func(token *jwt.Token) (interface{}, error) {
        return signKey, nil
    })
    if err != nil {
        log.Fatal("parse failed:", err.Error())
    }

    // 获取jwt的消息体内容
    if c, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
        for key, value := range c {
            fmt.Printf("%s:%v\n", key, value)
        }
    }
}
```

# 4 jwt格式
在`github.com/golang-jwt/jwt/v5`中，声明claims时有几个特定的字段

## 4.1 过期时间
声明token的过期时间的字段为`exp`.
可以通过以下方法获取
```go
func (m MapClaims) GetExpirationTime() (*NumericDate, error)

// NumericDate是对time.Time的封装
```

## 4.2 token被签发的时间
token被签发的时间的对应字段为`iat`
可以通过以下方法获取
```go
func (m MapClaims) GetIssuedAt() (*NumericDate, error)
// NumericDate是对time.Time的封装
```

## 4.3 JWT的受众
在 JWT (JSON Web Token) 标准中，aud 是一个预定义的声明，代表 “Audience”，意思是这个JWT的受众。aud 声明用来指定这个令牌预期会送往的一方或多方，这些受众通常是API或者某些资源的标识符。

aud 的值可以是一个字符串或者字符串数组。当它是一个字符串时，它通常是一个单一的受众标识符，例如一个特定的服务器的标识。如果它是一个字符串数组，那么JWT可以同时针对多个受众。

aud 声明的存在是为了增强安全性，确保JWT不会意外地被用于非目标接收者的系统中。在验证JWT时，接收者应验证aud声明是否匹配或包含其自身的标识符，如果不匹配，则必须拒绝这个JWT。
可以通过以下方法获取:
```go
func (m MapClaims) GetAudience() (ClaimStrings, error)

type ClaimStrings []string
```

## 4.4 jwt的发布者
jwt的发布者在claim中通过`iss`来指定。
可以通过以下方法来获取：
```go
func (m MapClaims) GetIssuer() (string, error)
```

iss 声明的值是一个用来指示令牌发行方的字符串。它可以是一个URI、一个领域名，或者任何能唯一标识令牌颁发者身份的字符串。

## 4.5 “Not Before” 时间
在 JSON Web Token (JWT) 的标准字段中，nbf 是一个预定义的声明，代表 “Not Before” 时间。nbf 声明的值是一个数字日期，之前的时间内，JWT 是不被接受的。这个数字通常是一个以秒为单位的 UNIX 时间戳，它表示一个特定的日期和时间。

如果 JWT 的 nbf 声明中指定了一个时间，那么在这个时间之前，JWT是无效的，接收方应该拒绝处理这个令牌。这等于是说，“这个令牌将在 nbf 指定的时间开始之前，是不可使用的。”

在go中获取这个信息的方法如下:
```go
func (m MapClaims) GetNotBefore() (*NumericDate, error)
```

## 4.6 令牌主题
在 JSON Web Token (JWT) 的标准中，sub 是一个预定义的声明，它代表 “Subject”，即令牌的主题。sub 声明的值是用于唯一标识令牌的主题的字符串，通常是指用户的一个标识符。

sub 声明是JWT中用来表明这个令牌是代表谁的。例如，它可能包含用户的ID、用户名或其他一些唯一识别用户身份的信息。JWT的主题可以是访问的终端用户，也可以是某个系统中代表的一个实体。
sub 声明是一个可选字段，但它是一个用于确保JWT的有效应用的好实践。使用sub声明，JWT的接收者可以具体了解这个令牌是为谁颁发的，并根据这些信息做出授权决定。

这个信息可以通过一下方法获取：
```go
func (m MapClaims) GetSubject() (string, error)
```