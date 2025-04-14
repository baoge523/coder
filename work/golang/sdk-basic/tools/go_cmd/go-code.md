# golang 代码规范相关的工具

## gofmt
[gofmt](https://pkg.go.dev/cmd/gofmt)
Usage
```text
gofmt [flags] [path ...]
```
```text
-d
	Do not print reformatted sources to standard output.   不打印源文件，当存在不规范时，输出不同点到标准输出
	If a file's formatting is different than gofmt's, print diffs
	to standard output.
-e
	Print all (including spurious) errors.     打印所有的格式错误信息
-l
	Do not print reformatted sources to standard output.   不打印源文件，当存在不规范时，输出文件名到标准输出
	If a file's formatting is different from gofmt's, print its name
	to standard output.
-r rule
	Apply the rewrite rule to the source before reformatting.  重写规则用于fmt
-s
	Try to simplify code (after applying the rewrite rule, if any). 尝试简化代码
-w
	Do not print reformatted sources to standard output.   不打印源文件，当存在不规范时，尝试使用当前的gofmt版本重写它
	If a file's formatting is different from gofmt's, overwrite it
	with gofmt's version. If an error occurred during overwriting,
	the original file is restored from an automatic backup.
```

样例:
```bash
gofmt -d /proto/

gofmt -e /proto/

gofmt -l /proto/

gofmt -r '(a) -> a' -w *.go

gofmt -r 'test_a -> testA' -w sdk/sdk-golang/tools/go_cmd/go_code_style/demoStyle.go

```


## goimports
[go imports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
用于更新go project imports信息，添加未引入的，删除引入了但是没有使用的
```bash
# 安装goimports命令到$GOPATH/bin下
go install golang.org/x/tools/cmd/goimports@latest
```

goimports --help 查看帮助命令

usage：
```text
 goimports [flags] [path ...]
```
```text
-cpuprofile string
        CPU profile output
  -d    display diffs instead of rewriting files
  -e    report all errors (not just the first 10 on different lines)
  -format-only
        if true, don't fix imports and only format. In this mode, goimports is effectively gofmt, with the addition that imports are grouped into sections.
  -l    list files whose formatting differs from goimport's
  -local string
        put imports beginning with this string after 3rd-party packages; comma-separated list
  -memprofile string
        memory profile output
  -memrate int
        if > 0, sets runtime.MemProfileRate
  -srcdir dir
        choose imports as if source code is from dir. When operating on a single file, dir may instead be the complete file name.
  -trace string
        trace profile output
  -v    verbose logging
  -w    write result to (source) file instead of stdout
```


## go vet
[go vet](https://pkg.go.dev/cmd/vet@go1.23.2)
```text
Vet examines Go source code and reports suspicious constructs, 
such as Printf calls whose arguments do not align with the format string. 
Vet uses heuristics that do not guarantee all reports are genuine problems, 
but it can find errors not caught by the compilers.

-----
单词:
examine  检查、调查、考核   三单: examines
suspicious  感觉可疑的、怀疑的、不信任的
constructs  结构体
align with  与某人或者组织达成一致
heuristics  启发式、搜索
guarantee   v 确保、保证、担保   n 保证、担保、保修单
genuine adj 真正的、非伪造的、真诚的、真心的
caught  catch的过去分词 捕获
```

可用的check 列表，每个的作用可以通过 go tool vet help 查看参数作用信息
```text
appends          check for missing values after append
asmdecl          report mismatches between assembly files and Go declarations
assign           check for useless assignments
atomic           check for common mistakes using the sync/atomic package
bools            check for common mistakes involving boolean operators
buildtag         check //go:build and // +build directives
cgocall          detect some violations of the cgo pointer passing rules
composites       check for unkeyed composite literals
copylocks        check for locks erroneously passed by value
defers           report common mistakes in defer statements
directive        check Go toolchain directives such as //go:debug
errorsas         report passing non-pointer or non-error values to errors.As
framepointer     report assembly that clobbers the frame pointer before saving it
httpresponse     check for mistakes using HTTP responses
ifaceassert      detect impossible interface-to-interface type assertions
loopclosure      check references to loop variables from within nested functions
lostcancel       check cancel func returned by context.WithCancel is called
nilfunc          check for useless comparisons between functions and nil
printf           check consistency of Printf format strings and arguments
shift            check for shifts that equal or exceed the width of the integer
sigchanyzer      check for unbuffered channel of os.Signal
slog             check for invalid structured logging calls
stdmethods       check signature of methods of well-known interfaces
stringintconv    check for string(int) conversions
structtag        check that struct field tags conform to reflect.StructTag.Get
testinggoroutine report calls to (*testing.T).Fatal from goroutines started by a test
tests            check for common mistaken usages of tests and examples
timeformat       check for calls of (time.Time).Format or time.Parse with 2006-02-01
unmarshal        report passing non-pointer or non-interface values to unmarshal
unreachable      check for unreachable code
unsafeptr        check for invalid conversions of uintptr to unsafe.Pointer
unusedresult     check for unused results of calls to some functions
```

查看帮助信息
```bash
go tool vet help
```
核心的参数
> -c=n 出现问题出的上下文行数
> 
> -json 以josn格式输出
> 

样例
```bash
go vet -c=5 -json -printf ./go_code_style
```


## golint
[golangci-lint](https://golangci-lint.run/)
[golangci-lint github](https://github.com/golangci/golangci-lint)
用于检查go source code 是否存在编程规范问题

### install
macos
```bash
brew install golangci-lint
brew upgrade golangci-lint
```


执行检测
```bash
golangci-lint run
```

```text
Usage:
  golangci-lint [flags]
  golangci-lint [command]

Available Commands:
  cache       Cache control and information
  completion  Generate the autocompletion script for the specified shell
  config      Config
  help        Help
  linters     List current linters configuration
  run         Run the linters
  version     Version

Flags:
      --color string              Use color when printing; can be 'always', 'auto', or 'never' (default "auto")
  -j, --concurrency int           Concurrency (default NumCPU) (default 10)
      --cpu-profile-path string   Path to CPU profile output file
  -h, --help                      help for golangci-lint
      --mem-profile-path string   Path to memory profile output file
      --trace-path string         Path to trace output file
  -v, --verbose                   verbose output
      --version                   Print version
```

### 查询所有的linters,查看哪些可用，哪些不可用
```bash
golangci-lint linters
```
执行后的结果信息:
```text
Enabled by your configuration linters:  默认开启的检查
errcheck: errcheck is a program for checking for unchecked errors in Go code. These unchecked errors can be critical bugs in some cases [fast: false, auto-fix: false]
gosimple (megacheck): Linter for Go source code that specializes in simplifying code [fast: false, auto-fix: false]
govet (vet, vetshadow): Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string [fast: false, auto-fix: false]
ineffassign: Detects when assignments to existing variables are not used [fast: true, auto-fix: false]
staticcheck (megacheck): It's a set of rules from staticcheck. It's not the same thing as the staticcheck binary. The author of staticcheck doesn't support or approve the use of staticcheck as a library inside golangci-lint. [fast: false, auto-fix: false]
typecheck: Like the front-end of a Go compiler, parses and type-checks Go code [fast: false, auto-fix: false]
unused (megacheck): Checks Go code for unused constants, variables, functions and types [fast: false, auto-fix: false]

Disabled by your configuration linters: 默认关闭的检查
...
```


### golint的疑问

#### 既然可以通过golangci-lint linters 查看配置的规则，那怎么增加规则呢？
通过官网查看自定义linters规则 https://golangci-lint.run/contributing/new-linters/