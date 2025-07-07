
## go build

GOOS=linux GOARCH=amd64 go build -o getUserId

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

```text
$GOARCH
	The execution architecture (arm, amd64, etc.)
$GOOS
	The execution operating system (linux, windows, etc.)
$GOFILE
	The base name of the file.
$GOLINE
	The line number of the directive in the source file.
$GOPACKAGE
	The name of the package of the file containing the directive.
$GOROOT
	The GOROOT directory for the 'go' command that invoked the
	generator, containing the Go toolchain and standard library.
$DOLLAR
	A dollar sign.
$PATH
	The $PATH of the parent process, with $GOROOT/bin
	placed at the beginning. This causes generators
	that execute 'go' commands to use the same 'go'
	as the parent 'go generate' command.
```