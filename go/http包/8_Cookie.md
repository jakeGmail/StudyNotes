[toc]

# 1 cookei介绍
cookie有两种分类
- 会话期Cookie: 只在单次会话中有效。登陆后再退出，然后再次登录就会失效，需要重新登录
- 持久性Cookie： 在登录成功一次后，在规定期间内，可以不需要登录而再次访问页面

# 设置会话期cookies

```go
import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func CookieJarTest() {
    // 创建默认cookieJar， 如果要访问其他页面也带上这个jar,就需要把jar这个变量单独保存，然后传给其他需要的方法
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Println("new cookiejar faild:", err.Error())
	}

    // 自定义客户端
	client := &http.Client{
		Jar: jar,
	}

    // 请求登录
	resp, err := client.Get("http://httpbin.org/cookies/set?name=jake&password=123")
	if err != nil {
		log.Fatal("get req failed:", err.Error())
	}

    // 打印响应消息
	_, _ = io.Copy(os.Stdout, resp.Body)
}
```

```go
func main(){
    jar,_ := cookiejar.New(nil)
    login(jar)
    userCenter(jar)
}
```

# 3 使用持久性Cookie
http中弄人提供的是一个会话级别的CookieJar, 如果需要持久性Cookie就需要自己手动实现http.CookieJar接口。
```go
type CookieJar interface {
	// SetCookies在给定URL的回复中处理cookies的接收。它可能选择保存cookie，也可能不选择保存cookie，这取决于jar的策略和实现。
	SetCookies(u *url.URL, cookies []*Cookie)

	// cookie返回为给定URL发送请求的cookie。这取决于实现是否遵守标准cookie使用限制，如RFC 6265。
	Cookies(u *url.URL) []*Cookie
}
```

也可以使用`github.com/juju/persistent-cookiejar`包