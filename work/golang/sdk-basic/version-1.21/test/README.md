# 了解如何使用测试用例
在 Go 中，测试用例通常使用 `testing` 包来编写和运行。以下是如何编写和运行测试用例的步骤：
1. **创建测试文件**：测试文件以 `_test.go` 结尾。例如，如果你的代码在 `example.go` 中，测试文件应命名为 `example_test.go`。
2. **编写测试函数**：测试函数必须以 `Test` 开头，并接受一个指向 `testing.T` 的指针
```go
package main
import (
"testing"
)
func TestMyFunction(t *testing.T) {
    result := myFunction()
    expected := "Hello from myFunction!"
    if result != expected {
        t.Errorf("expected %s, got %s", expected, result)
    }
}
   ```
3. **运行测试**：在终端中，使用以下命令运行测试：
   ```bash
   go test
   ```
如果你想查看详细的输出，可以添加 `-v` 标志：
   ```bash
   go test -v
   ```
4. **指定测试文件**：如果只想运行特定的测试文件，可以使用：
   ```bash
   go test -run TestMyFunction
   ```



