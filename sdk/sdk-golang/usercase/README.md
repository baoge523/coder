## golang功能验证
所有学习golang的验证信息

### 切片相关的操作
golang中的切片并不是数组，它是真实的对象,它的定义如下

src/runtime/slice.go
```go
type slice struct {
	array unsafe.Pointer   // 指向内存的指针
	len   int    // 长度
	cap   int    // 容量
}
```
声明创建slice的方式
```go
var fruits []string
fruits = append(fruits,"apple")  // 在append中会初始化slice
frults[1] = "aaa"  // 会报错


var fruits = make([]string,10,100) // len = 10 cap = 100
fruits = append(fruits,"apple") // 添加索引位置为 11

fruits[1] = "aaa"  // 不会报错


```
slice通过[:] 操作，可能存在共享内存的问题
```go
var fruits = make([]string,0,10)

f1 := fruits[:5] // 此时是 0-4 len=5 cap=10, 与f2共享内存 5-9
f2 := fruits[5:] // 此时是 5-9 len=5 cap=5

```
解决方式: f1 := fruits[:5:5] 指定len和cap就没有问题了