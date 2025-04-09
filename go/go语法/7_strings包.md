[toc]
# 1 strings.Builder
## 1.1 写入（write）方法
与 bytes.Buffer 类似，strings.Builder 也支持 4 类方法将数据写入 builder 中。
```go
func (b *Builder) Write(p []byte) (int, error)
func (b *Builder) WriteByte(c byte) error
func (b *Builder) WriteRune(r rune) (int, error)
func (b *Builder) WriteString(s string) (int, error)
```
有了它们，用户可以根据输入数据的不同类型（byte 数组，byte， rune 或者 string），选择对应的写入方法。


## 1.2 strings.Builder实现原理
strins.Builder内部实现是指针+切片，同时String()返回拼接后的字符串，它是直接把[]byte转换为string，从而避免变量拷贝。

string.Builder 通过使用一个内部的 slice 来存储数据片段。当开发者调用写入方法的时候，数据实际上是被追加（append）到了其内部的 slice 上。

## 1.3高效地使用 strings.Builder
由于strings.Builder 是通过其内部的 slice 来储存内容的。当你调用写入方法的时候，新的字节数据就被追加到 slice 上。如果达到了 slice 的容量（capacity）限制，一个新的 slice 就会被分配，然后老的 slice 上的内容会被拷贝到新的 slice 上。当 slice 长度很大时，这个操作就会很消耗资源甚至引起 内存问题。我们需要避免这一情况。

关于 slice，Go 语言提供了 make([]TypeOfSlice, length, capacity) 方法在初始化的时候预定义它的容量。这就避免了因达到最大容量而引起扩容。

strings.Builder 同样也提供了 Grow() 来支持预定义容量。当我们可以预定义我们需要使用的容量时，strings.Builder 就能避免扩容而创建新的 slice 了。
```go
func (b *Builder) Grow(n int)
```
当调用 Grow() 时，我们必须定义要扩容的字节数（n）。 Grow() 方法保证了其内部的 slice 一定能够写入 n 个字节。只有当 slice 空余空间不足以写入 n 个字节时，扩容才有可能发生。举个例子：

- builder 内部 slice 容量为 10。
- builder 内部 slice 长度为 5。
- 当我们调用 Grow(3) => 扩容操作并不会发生。因为当前的空余空间为 5，足以提供 3 个字节的写入。
- 当我们调用 Grow(7) => 扩容操作发生。因为当前的空余空间为 5，已不足以提供 7 个字节的写入。

关于上面的情形，如果这时我们调用 Grow(7)，则扩容之后的实际容量是多少？
17 还是 12?
实际上，是 27。strings.Builder 的 Grow() 方法是通过 <font color=blue>current_capacity * 2 + n （n 就是你想要扩充的容量）</font>的方式来对内部的 slice 进行扩容的。所以说最后的容量是 10*2+7 = 27。

当你预定义 strings.Builder 容量的时候还要注意一点。调用 WriteRune() 和 WriteString() 时，rune 和 string 的字符可能不止 1 个字节。因为，你懂的，UTF-8 的原因。

## 1.4 String()
和 bytes.Buffer 一样，strings.Builder 也支持使用 String() 来获取最终的字符串结果。为了节省内存分配，它通过使用指针技术将内部的 buffer bytes 转换为字符串。所以 String() 方法在转换的时候节省了时间和空间。

## 1.5 Reset()
清空字符串，将底层的存储字符串的[]byte置为nil

## 1.6 不要拷贝strings.Builder
strings.Builder 不推荐被拷贝。当你试图拷贝 strings.Builder 并写入的时候，你的程序就会崩溃。
```go
var b1 strings.Builder
b1.WriteString("ABC")
b2 := b1
b2.WriteString("DEF") 
// illegal use of non-zero Builder copied by value
```
你已经知道，strings.Builder 内部通过 slice 来保存和管理内容。slice 内部则是通过一个指针指向实际保存内容的数组。

当我们拷贝了 builder 以后，同样也拷贝了其 slice 的指针。但是它仍然指向同一个旧的数组。当你对源 builder 或者拷贝后的 builder 写入的时候，问题就产生了。另一个 builder 指向的数组内容也被改变了。这就是为什么 strings.Builder 不允许拷贝的原因。

对于一个未写入任何东西的空内容 builder 则是个例外。我们可以拷贝空内容的 builder 而不报错。
```go
var b1 strings.Builder
b2 := b1
b2.WriteString("DEF")
b1.WriteString("ABC")

// b1 = ABC, b2 = DEF
strings.Builder 会在以下方法中检测拷贝操作：

Grow(n int)
Write(p []byte)
WriteRune(r rune)
WriteString(s string)
```
所以，拷贝并使用下列这些方法是允许的：
```go
// Reset()
// Len()
// String()

var b1 strings.Builder
b1.WriteString("ABC")
b2 := b1
fmt.Println(b2.Len())    // 3
fmt.Println(b2.String()) // ABC
b2.Reset() // 需要在Reset之后再写入才不会报错
b2.WriteString("DEF")
fmt.Println(b2.String()) // DEF
``` 
## 1.7 并行支持
和 bytes.Buffer 一样，strings.Builder 也<font color=red>不支持并行的读或者写</font>。所以我们们要稍加注意。

可以试一下，通过同时给 strings.Builder 添加 1000 个字符：
```go
package main

import (
    "fmt"
    "strings"
    "sync"
)

func main() {
    var b strings.Builder
    n := 0
    var wait sync.WaitGroup
    for n < 1000 {
        wait.Add(1)
        go func() {
            b.WriteString("1")
            n++
            wait.Done()
        }()
    }
    wait.Wait()
    fmt.Println(len(b.String()))
}
```
通过运行，你会得到不同长度的结果。但它们都不到 1000。

## 1.8 io.Writer 接口
strings.Builder 通过``` Write(p []byte) (n int, err error)``` 方法实现了 io.Writer 接口。所以，我们多了很多使用它的情形：
```go
io.Copy(dst Writer, src Reader) (written int64, err error)

bufio.NewWriter(w io.Writer) *Writer

fmt.Fprint(w io.Writer, a …interface{}) (n int, err error)

func (r *http.Request) Write(w io.Writer) error
```
其他使用 io.Writer 的库

## 1.9 示例代码
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    name := "hello"
    age := "-world"
    strBuilder := strings.Builder{}
    strBuilder.WriteString(name)
    strBuilder.WriteString(age)
    fmt.Println(strBuilder.String()) // hello-world

    // 清空内容
    strBuilder.Reset()
    fmt.Println(strBuilder.String())
}
```

# 2 strings.Reader
## 2.1 Read函数
函数原型
```go
func (r *Reader) Read(b []byte) (n int, err error)
```
**说明**： 读取最多b的容量的长度的byte内容到切片b中

**参数：**
- **<font color=green>b</font>** : 读取Reader中的字符串数据到b中
**返回值：**
- **<font color=green>n</font>** : 实际读取的长度，单位byte
-  **<font color=green>err</font>** : 错误，如果Reader中没有字符或则全部字符都被都完，则返回EOF

示例代码
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // 创建一个Reader
    var strReader *strings.Reader = strings.NewReader("jake")
    if strReader == nil {
        fmt.Println("new strings reader failed")
        return
    }

    // 从Reader中读取内容
    buf := make([]byte, 2)
    size, err := strReader.Read(buf)
    if err != nil {
        fmt.Println("err:", err.Error())
        return
    }
    fmt.Printf("size=%d, buf=%s", size, buf) //size=2, buf=ja
}
```

## 2.2 Reader.Len方法
函数原型：
```go
func (r *Reader) Len() int 
```
**说明**: 获取Reader中还没有被读取的字节数
**示例代码**：
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var strReader *strings.Reader = strings.NewReader("jake")

    // 开始没有被读取的字节数为4
	fmt.Println("Len=", strReader.Len()) // 4

    // 读取2字节到buf中
    buf := make([]byte, 2)
	strReader.Read(buf)

    // 剩余没有被读取的字节数
	fmt.Println("Len=", strReader.Len()) // 2
}
```

## 2.3 Reader.Seek
函数原型：
```go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```
**说明**：修改当前读写位置，修改到whence指向位置的offset偏移量处
**参数**： 
- offser: 移动位置偏移量
- whence: 移动读写位置的相对位置，有如下几个值
  - io.SeekStart  起始位置
  - io.SeekCurrent  当前位置
  - io.SeekEnd 字符串末尾

**示例代码**:
```go
package main

import (
    "fmt"
    "io"
    "strings"
)
func main() {
    var strReader *strings.Reader = strings.NewReader("jake")

    /*读取3个字节*/
    buf := make([]byte, 3)
    _, _ = strReader.Read(buf)

    fmt.Printf("unreaded size=%d\n", strReader.Len()) // 1

    /*将读取指针移到起始位置往后1字节出*/
    strReader.Seek(1, io.SeekStart)
    fmt.Printf("unreader size=%d\n", strReader.Len()) // 3
}
```

## 2.4 ReadByte方法
函数原型：
```go
func (r *Reader) ReadByte() (byte, error)
```
**说明**: 从strings.Reader中读取一个字节的数据并返回.这个函数会影响Len()

## 2.5 ReadAt方法
函数原型:
```go
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```
示例代码：
```go
func main() {
    var reader *strings.Reader = strings.NewReader("jake")
    buf := make([]byte, 2)
    reader.ReadAt(buf, 2)
    fmt.Println(string(buf)) // ke
}
```

# 2 strings包下的其他函数
## 2.1 字符串比较
相等返回0，不等返回-1
```go
func Compare(a, b string) int
```