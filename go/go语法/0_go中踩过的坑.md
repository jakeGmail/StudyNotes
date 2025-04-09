[toc]

# 1 Wrong test signature报错
**现象**：
在工程中添加一个包后，调用时找不到自己写的包（即使手动import了---）

**原因**：
是因为在手写的包中的go文件的后缀用了_test.go, 导致go认为这是一个单元测试包。
并且在这种命名的文件下定义函数名为`TestRun`的函数时，要求必须有参数t *testing.T，如下
```go
func TestRun(t *testing.T) {
}
```

**修复措施**：
- 将go文件进行重命名，不以"_test.go"结尾。
