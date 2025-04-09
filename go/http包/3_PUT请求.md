[top]

# 1 PUT情趣

```go
func PutTest() {
    // 床架PUT请求
	putReq, err := http.NewRequest("PUT", "http://httpbin.org/put", nil)
	if err != nil {
		log.Fatal("new http request failed", err.Error())
	}

    // 执行PUT请求
	resp, err := http.DefaultClient.Do(putReq)
	if err != nil {
		log.Fatal("put request failed:", err.Error())
	}
	log.Println("resp=", resp)
}
```