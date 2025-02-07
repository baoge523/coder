# 记录使用redis过程中遇到的问题

## 1、redis版本问题，导致使用set命令时报错
```go
// 使用 ttl 为 -1 报错，原因是当前的redis版本是5.3，而-1表示不过期在redis6.0之后
redisCli.Set(ctx, totalPurchaseQuotaCounterKey, totalPurchaseQuotaCounter, -1)


// Set Redis `SET key value [expiration]` command.
// Use expiration for `SETEX`-like behavior.
//
// Zero expiration means the key has no expiration time. 0过期时间表示key没有过期时间
// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
// otherwise you will receive an error: (error) ERR syntax error.
func (c cmdable) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd

```

## 2、