
# 1 post请求简单示例
```go
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

**使用示例**:
```go
func PostTest() {
	resp, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		log.Fatal("POST Request failed", err.Error())
	}
	defer func() { _ = resp.Body.Close() }()
	httpContent, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read http content failed:", err.Error())
	}
	log.Println(string(httpContent))
}
```

# 2 post form数据
```go
func PostFormTest() {
    postData := make(url.Values)
    postData.Add("name", "jake")
    postData.Add("age", "23")
    data := postData.Encode()

    resp, err := http.Post("http://httpbin.org/post",
        "application/x-www-form-urlencoded",
        strings.NewReader(data))
    if err != nil {
        log.Fatal("post failed:", err.Error())
    }

    buf, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("read resp body failed:", err.Error())
    }
    defer resp.Body.Close()
    log.Println("buf=", string(buf))
}
```

# 3 post json数据
```go
func PostJsonTest() {
	data, err := json.Marshal(Person{Name: "jake", Age: 23})
	if err != nil {
		log.Fatal("marshal failed:", err.Error())
	}

	resp, err := http.Post(
		"http://httpbin.org/post",
		"application/json",
		strings.NewReader(string(data)),
	)

	if err != nil {
		log.Fatal("post failed:", err.Error())
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read resp data failed:", err.Error())
	}
	defer resp.Body.Close()
	log.Println("resp.Body=", string(buf))
}
```

# 4 post文件
```go
import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PostFileTest() {
    // 定于post的body
	body := &bytes.Buffer{}

    // 创建multipart字段，用于保存post的文件数据
	multipartWriter := multipart.NewWriter(body)
	// 设置file的字段和文件名
    uploadFileWriter, _ := multipartWriter.CreateFormFile("file1", "download.txt")

    // 读取文件内容到body中
	file, _ := os.Open("download.txt")
	defer file.Close()
	_, _ = io.Copy(uploadFileWriter, file)

	// multipart/form-data; boundary=0be9801253b4e336bdb2ed7bd94672813d9324180a980155b226ce04d614
	log.Println("content-type=", multipartWriter.FormDataContentType())
	
    // 需要线关闭后再上传
	_ = multipartWriter.Close()

    // 发送post请求
	resp, err := http.Post(
		"http://httpbin.org/post",

        // 需要现生成content-type，因为数据的边界是现生成的
        // multipart/form-data; boundary=0be9801253b4e336bdb2ed7bd94672813d9324180a980155b226ce04d614
		multipartWriter.FormDataContentType(),
		body,
	)
	if err != nil {
		log.Fatal("marshal failed:", err.Error())
	}

    // 获取响应体
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read resp data failed:", err.Error())
	}
	defer resp.Body.Close()
	log.Println("resp.Body=", string(buf))
}
```


# 5 Post提交的本质
post是通过request body来提交请求的，相对于get请求，限制比较少。
get提交大小有限制。