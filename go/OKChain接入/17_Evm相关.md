[toc]

# 1 查看合约的代码
```go
QueryCode(contractAddrStr string) (types.QueryResCode, error)
```

**示例代码**：
```go
func GetCode() {
    // 合约地址
	addr := "0xF77596928f0823959c2caF4810834FD227244871"
	code, err := OkClient.Evm().QueryCode(addr)
	if err != nil {
		log.Fatal("query code faield: ", err.Error())
	}
	log.Println("code:", code)
}
```