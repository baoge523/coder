## collector
https://github.com/open-telemetry/opentelemetry-collector

https://opentelemetry.io/docs/collector/



```go
import (
    promCommConfig "github.com/prometheus/common/config"
    promConfig "github.com/prometheus/prometheus/config"
    "github.com/prometheus/prometheus/discovery"
    "github.com/prometheus/prometheus/model/labels"
    "github.com/prometheus/prometheus/scrape"
)

// 这两个是作用是什么，需要参考opentelemetry-collector
    scrapeManager    *scrape.Manager
	discoveryManager *discovery.Manager
```

### open-telemetry 官网
https://opentelemetry.io/docs/languages/go/getting-started/