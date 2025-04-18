# golang 中的error处理
[go1.23 error文档](https://pkg.go.dev/errors@go1.23.3#section-documentation)

error用于表示golang中的错误

the method list for create error object：
- errors.new("xxx")     // text error
- fmt.Errorf("%s","ss") // text error
- fmt.Errorf("%w",err)  // warp error

join error:
> errors.Join(err1,err2) return err

check error:
> errors.Is(a,b) bool // depth-first traversal to find b is parent of a ?
> 
> errors.As(a,err_type_point_addr) bool // depth-first traversal

usage:
```go

func TestAs() {
    var pathError *PathError
	check := errors.As(err,&pathError)
}

```

unwrap error:
- Unwarp(err) err
- Unwarp(err) []err

for example:
```go
 wrapErr := fmt.Errorf("this is one err= %w,this is another one err=%w",err1,err2)
 errList := errors.Unwrap(wrapErr)
```
