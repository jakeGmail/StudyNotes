[toc]

# 1 Delete请求的
```go
func DeleteTest() {
	putReq, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		log.Fatal("new http request failed", err.Error())
	}
	resp, err := http.DefaultClient.Do(putReq)
	if err != nil {
		log.Fatal("delete request failed:", err.Error())
	}
	log.Println("resp=", resp)
}
```