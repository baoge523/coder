### 哪些场景会导致出现panic
  - 访问数组索引下标越界
  - 向已经close的chan写入数据、重复close chan、close nil chan
  - 使用nil 对象(指针 接口 map 切片)
  - 类型转换时，如果没有指定第二bool参数，转换失败会出现panic
  - 被除数为0
  - sync中的waitGroup和Mutex(不可重入)；他们的内部都是通过计数值来的，如果小于0了，就会panic
  - 反射包reflect


### 常用的工具类
go sdk 中常用的工具类，帮忙我们快速处理
#### strconv 
https://pkg.go.dev/strconv
search: std/strconv 表示查询标准库中的strconv

Package strconv implements conversions to and from string representations of basic data types.
> 该package实现了基础类型转换成string 和 string转换成基础类型

转换类型的func
```go
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)

b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)

s := "2147483647" // biggest int32
i64, err := strconv.ParseInt(s, 10, 32)
...
i := int32(i64)

s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)


q := strconv.Quote("Hello, 世界")
q := strconv.QuoteToASCII("Hello, 世界")
```


### 内存拷贝，比如数组拷贝