[toc]

# 1 构造GET请求
在http中， GET请求一般是用于获取服务器上的资源
```go
func Get(url string) (resp *Response, err error)
```

**参数**：
|参数名称|类型|描述|
|--------|---|----|
|url|string|网址的路径|

**返回值**：
|返回值名称|返回值类型|描述|
|----------|--------|----|
|resp|Response|http请求的响应内容|


**使用示例**:
```go
package Http

import (
	"io"
	"log"
	"net/http"
)

func GetTest() {
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal("GET Request failed", err.Error())
	}
	defer func() { _ = res.Body.Close() }()
	httpContent, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("read http content failed:", err.Error())
	}
	log.Println(string(httpContent))
}
```

# 2 添加查询参数
```go
func GetTest() {
	getReq, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatal("new request failed:", err.Error())
	}
    

    // GET请求添加参数，其作用类似于请求 "http://httpbin.org/get?name=jake&age=23"
	params := make(url.Values)
	params.Add("name", "jake")
	params.Add("age", "23")
	getReq.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(getReq)
	if err != nil {
		log.Fatal("do get request failed:", err.Error())
	}

    // 读取响应内容
    content, err := io.ReadAll(resp.Body)
	log.Println("get resp=", string(content))

    defer resp.Body.Close()
}
```

# 3 定制请求头
```go
func GetTest() {
    getReq, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
    if err != nil {
        log.Fatal("new request failed:", err.Error())
    }


    // 设置请求头
    getReq.Header.Add("user-agent", "jake")

    resp, err := http.DefaultClient.Do(getReq)
    if err != nil {
        log.Fatal("do get request failed:", err.Error())
    }

    // 读取响应内容
    content, err := io.ReadAll(resp.Body)
    if err != nil{
		log.Fatal("read resp body failed:",err.Error())
	}
    log.Println("get resp=", string(content))
    defer func() { resp.Body.Close() }()
}
```
