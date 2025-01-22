
## yin.context 重要

```go

*gin.Context

ctx := c.Request.Context()
// 返回基础信息
param := ctx.Value(bpCtxKey)

与这个处理的区别
c.Request.Body

```

## c.Request.Body 只能从该流中获取一次，第二次获取就会为空
