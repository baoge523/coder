# fmt 的格式化
[fmt官方文档](https://pkg.go.dev/fmt@go1.20)


### 注意：📢 %w 只能在fmt.Errorf中使用，不能在任何其他地方使用


### 使用%v时，不同类型的对应的默认占位符

普通类型
```text
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

复杂类型
```text
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]
```

精度问题处理
```text
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
```

### 常见的对应类型占位符
```text
1. **%s**: 字符串
2. **%d**: 十进制整数
3. **%f**: 浮点数
4. **%t**: 布尔值
5. **%v**: 值的默认格式
6. **%+v**: 值的详细格式（包含字段名，通常用于结构体）
7. **%#v**: 值的 Go 语法表示（用于调试）
8. **%x**: 十六进制表示
9. **%p**: 指针地址
```