## prometheus
https://yunlzheng.gitbook.io/prometheus-book/parti-prometheus-ji-chu/quickstart

Prometheus会将所有采集到的样本数据以时间序列（time-series）的方式保存在内存数据库中，并且定时保存到硬盘上

### docker install
```bash
docker run -p 9090:9090 -d -v /etc/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

# 修复容器时间和主机时间不一致的情况，手动挂载时间
docker run -p 9090:9090 -d  -v /etc/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml -v /etc/localtime:/etc/localtime prom/prometheus
```

### prometheus 的实例和任务
通过在prometheus.yml指定绑定的数据采集任务，即从node exporter暴露的服务中获取监控指标数据。
```yaml
scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
  - job_name: 'node'
    static_configs:
      - targets: ['localhost:9100']
```
当我们需要采集不同的监控指标(例如：主机、MySQL、Nginx)时，我们只需要运行相应的监控采集程序，并且让Prometheus Server知道这些Exporter实例的访问地址。
在Prometheus中，每一个暴露监控样本数据的HTTP服务称为一个实例。例如在当前主机上运行的node exporter可以被称为一个实例(Instance)。

实例：采集指标的exporter
任务：管理多个采集相同指标的实例

### prometheus 的组件及其相关性
```text

数据采集层        存储层                  表示层
Exporter       prometheus-server       prom-ui/grafana
```
数据采集层：
提供了多种Exporter供用户使用，当然用户也可以自定义Exporter，只要满足prom的数据结构即可

存储层:
prometheus-server 本身就是一个时序存储器，可以将数据按照时间存在到磁盘中

表示层：
prometheus 提供了内置的ui页面供用户使用，grafana完整的集成了prometheus的功能，也可以使用强大的grafana

### 指标
指标格式：
```text
<metric name>{<label name>=<label value>, ...}
```
标签(label)反映了当前样本的特征维度，通过这些维度Prometheus可以对样本数据进行过滤，聚合等

### prometheus 内置的统计方式
Counter

Gauge  /ɡeɪdʒ/ 

histogram  /ˈhɪstəɡræm/  直方图; prometheus 的histogram 是累积直方图
参考: https://linuxczar.net/blog/2016/12/31/prometheus-histograms/

summary   /ˈsʌməri/  总结、概要

### PromQL
prometheus 是基于指标名称(metric name) 和一组标签(label set)组成的一条时序数据
```text
这里表示的意思是，该指标有哪些维度信息（即标签）
```
然后PromQL就是将指定的指标名称的label set拿来做**过滤，聚合，统计**从而产生新的计算后的一条时间序列

匹配方式：
 - 全值匹配  
 - 正则匹配

全值匹配：
```text
http_requests_total 等价于 http_requests_total{}

http_requests_total{code="200"}
http_requests_total{code!="200"}
```

正则匹配:
```text
http_requests_total{code=~"200|201"}
http_requests_total{code!~"200|201"}
```

瞬时向量和瞬时向量表达式
```text
http_requests_total{}
http_requests_total{code="200"}
```

范围向量和范围向量表达式
```text
http_requests_total{}[5m]
http_requests_total{code="200"}[5m]
```

支持的时间：
```text
s - 秒
m - 分钟
h - 小时
d - 天
w - 周
y - 年
```

向量偏移：
五分钟前的样本数据，瞬时向量
```text
http_requests_total{} offset 5m
```
昨天一天的样本数据，范围向量
```text
http_requests_total{}[1d] offset 1d
```


可以使用聚合的方式，将指标的通过不同的label进行聚合
```text
# 查询系统所有http请求的总量
sum(http_request_total)

# 按照mode计算主机CPU的平均使用时间
avg(node_cpu) by (mode)

# 按照主机查询各个主机的CPU使用率
sum(sum(irate(node_cpu{mode!='idle'}[5m]))  / sum(irate(node_cpu[5m]))) by (instance)
```

合法的PromQL表达式
```text
所有的PromQL表达式都必须至少包含一个指标名称(例如http_request_total)，或者一个不会匹配到空字符串的标签过滤器(例如{code="200"})。

http_request_total # 合法
http_request_total{} # 合法
{method="get"} # 合法
```