[otel](https://github.com/open-telemetry/opentelemetry-go) 

所以这里目前的做法依然是提供一套直接使用的api接口:
标签的类型引用 [otel/attribute](https://pkg.go.dev/go.opentelemetry.io/otel/attribute) 
在某些语义上可以直接复用 [otel/semconv](https://pkg.go.dev/go.opentelemetry.io/otel/semconv) 


设计理念：
