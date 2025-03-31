## 记录一些优雅的代码风格

### 对一些属性做打印隐藏
需求：对敏感信息输出到控制台，或者日志文件中做隐蔽
分析：打印输出一般都是fmt.print 或者的json的序列化操作
找到print和json.Marshal的拓展即可

print: 实现了Format对象方法的，在print该对象时，就会按照Format的格式输出
```go
// Formatter is implemented by any value that has a Format method.
// The implementation controls how State and rune are interpreted,
// and may call Sprint() or Fprint(f) etc. to generate its output.
type Formatter interface {
	Format(f State, verb rune)
}
```
json.Marshal：实现了MarshalJSON对象方法，在json.Marshal会调用该对象的MarshalJSON()
```go
// Marshaler is the interface implemented by types that
// can marshal themselves into valid JSON.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
```
