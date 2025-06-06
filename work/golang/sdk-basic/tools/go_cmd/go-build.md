
## go build

GOOS=linux go build -o getUserId

### 构建约束，在使用go build时，通过指定tag，只有包含标签的文件才会被加载和编译
aaa.go
```go
//go:build aaa
package aaa1

import "aa1"
```
bbb.go
```go
//go:build bbb
package bbb1

import "bb1"
```

执行 go build 命令时，指定标签 其中 aaa 和 bbb 表示标签
```bash
go build aaa -o aaa_bin

go build bbb -o bbb_bin
```

### golang的交叉编译
[go cmd](https://pkg.go.dev/cmd/go)