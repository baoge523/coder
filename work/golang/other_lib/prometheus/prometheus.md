## prometheus 提供系统监控


```go
admin.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
```