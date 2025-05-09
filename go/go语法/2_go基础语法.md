[toc]
# 1 语法知识补充
## 1.1 导入包
### 1.1.1 匿名导入包
只导入不调用，通常这么做是为了调用该包下的init函数。
```go
package main
import (
   "fmt"
    _ "math" // 下划线表示匿名导入
)
func main() {
   fmt.Println(1)
}
```
**注意**: 
在Go中完全禁止循环导入，不管是直接的还是间接的。例如包A导入了包B，包B也导入了包A，这是直接循环导入，包A导入了包C，包C导入了包B，包B又导入了包A，这就是间接的循环导入，存在循环导入的话将会无法通过编译。

## 1.2 运算
**注意**
Go语言中没有自增与自减运算符，它们被降级为了语句statement，并且规定了只能位于操作数的后方，所以不用再去纠结i++和++i这样的问题。
```go
a++ // 正确
++a // 错误
a-- // 正确
```
还有一点就是，它们不再具有返回值，因此```a = b++```这类语句的写法是错误的。

## 1.3 数据类型
**字符类型**
Go对于Unicode编码百分百兼容和支持。

|类型|描述|
|------|----|
|byte|等价 uint8 可以表达ANSCII字符|
|rune|等价 int32 可以表达Unicode字符|
|string|字符串即字节序列，可以转换为[]byte类型即字节切片|

**复数类型**
|类型|描述|
|------|----|
|complex128|64位实数和虚数|
|complex64|32位实数和虚数|

**nil**
源代码中的nil，可以看出nil仅仅只是一个变量。
```go
var nil Type
```
Go中的nil并不等同于其他语言的null，nil仅仅只是一些类型的零值，并且不属于任何类型，所以```nil == nil```这样的语句是无法通过编译的。

**iota**
iota是一个内置的常量标识符，通常用于表示一个常量声明中的无类型整数序数，一般都是在括号中使用。

const iota = 0 
先看几个例子，看看规律。
```go
const (
   Num = iota // 0
   Num1 // 1
   Num2 // 2
   Num3 // 3
   Num4 // 4
)
```
也可以这么写
```go
const (
   Num = iota*2 // 0
   Num1 // 2
   Num2 // 4
   Num3 // 6
   Num4 // 8
)
```
还可以
```go
const (
   Num = iota << 2*3 + 1 // 1
   Num1 // 13
   Num2 // 25
   Num3 = iota // 3
   Num4 // 4
)
```
通过上面几个例子可以发现，iota受一种序号的变化而变化，第一个常量使用iota值的表达式，根据序号值的变化会自动的赋值给后续的常量，直到用新的iota重置，这个序号其实就是代码的相对行号，是相对于当前分组的起始行号，看下面的例子
```go
const (
	Num  = iota<<2*3 + 1 // 1 第一行
	Num2 = iota<<2*3 + 1 // 13 第二行
	_ // 25 第三行
	Num3 //37 第四行
	Num4 = iota // 4 第五行
	_ // 5 第六行
	Num5 // 6 第七行
)
```
例子中使用了匿名标识符_占了一行的位置，可以看到iota的值本质上就是iota所在行相对于当前const分组的第一行的差值。而不同的const分组则相互不会影响。
## 1.4 变量赋值
短变量初始化不能使用nil，因为nil不属于任何类型，编译器无法推断其类型。
```go
name := nil // 无法通过编译
```
短变量声明也可以批量初始化
```go
name, age := "jack", 1
```
在Go中，如果想要交换两个变量的值，不需要使用指针，可以使用=直接进行交换，语法上看起来非常直观，例子如下
```go
num1, num2 := 25, 36
nam1, num2 = num2, num1
```
三个变量也是同样如此
```go
num1, num2, num3 := 25, 36, 49
nam1, num2, num3  = num3, num2, num1
```

## 1.5 输入输出
**标准**
```go
var (
   Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
   Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
   Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```
在os包下有三个外暴露的文件描述符，其类型都是*File，分别是：
Stdin - 标准输入
Stdout - 标准输出
Stderr - 标准错误
Go中的控制台输入输出都离不开它们。

**输出**
输出一句Hello 世界!，比较常用的有三种方法，第一种是调用os.Stdout
```go
os.Stdout.WriteString("Hello 世界!")
```
第二种是使用内置函数println
```go
println("Hello 世界!")
```
第三种也是最推荐的一种就是调用fmt包下的Println函数
```go
fmt.Println("Hello 世界!")
```
fmt.Println会用到反射，因此输出的内容通常更容易使人阅读，不过性能很差强人意。

**格式化**
|0	|格式化|	描述|	接收类型|
|---|-----|-----|---------|
|1	|%%|    输出百分号%	|任意类型|
|2	|%s|	输出string/[] byte值	|string,[] byte|
|3	|%q|	格式化字符串，输出的字符串两端有双引号""|	string,[] byte|
|4	|%d|    输出十进制整型值|	整型类型|
|5	|%f|	输出浮点数|	浮点类型|
|6	|%e|	输出科学计数法形式 ,也可以用于复数	|浮点类型|
|7	|%E|	与%e相同	|浮点类型|
|8	|%g|	根据实际情况判断输出%f或者%e,会去掉多余的0	|浮点类型|
|9	|%b|	输出整型的二进制表现形式	|数字类型|
|10	|%#b|	输出二进制完整的表现形式	|数字类型|
|11	|%o|	输出整型的八进制表示	|整型|
|12	|%#o|	输出整型的完整八进制表示	|整型|
|13	|%x|	输出整型的小写十六进制表示	|数字类型|
|14	|%#x|	输出整型的完整小写十六进制表示	|数字类型|
|15	|%X|	输出整型的大写十六进制表示	|数字类型|
|16	|%#X|	输出整型的完整大写十六进制表示	|数字类型|
|17	|%v|	输出值原本的形式，多用于数据结构的输出	|任意类型|
|18	|%+v|	输出结构体时将加上字段名	|任意类型|
|19	|%#v|	输出完整Go语法格式的值	|任意类型|
|20	|%t|	输出布尔值	|布尔类型|
|21	|%T|	输出值对应的Go语言类型值	|任意类型|
|22	|%c|	输出Unicode码对应的字符	|int32|
|23	|%U|	输出字符对应的Unicode码	|rune,byte|
|24	|%p|	输出指针所指向的地址	|指针类型|


**提示**
在%与格式化动词之间加上一个空格便可以达到分隔符的效果，例如
```go
func main() {
	str := "abcdefg"
	fmt.Printf("%x\n", str)
	fmt.Printf("% x\n", str)
}
```
该例输出的结果为
```
61626364656667
61 62 63 64 65 66 67
```