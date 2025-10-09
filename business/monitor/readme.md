
## tencent cloud monitor

旧架构： -- 基于流式的告警，流式存储数据
```text
存在的问题：

1、计算任务繁重，每个产品的指标的一分钟、五分钟、十五分钟、1小时 都是通过flink任务计算出来的 -- 计算资源的浪费
2、链路扩充告警维度困难，不支持原基础上新增维度，只能新建namespace+viewname
    原因：当产品侧在控制层配置好指标的维度后，会基于这些维度信息生成表（表字段）信息，该表用来存放维度信息，其目的是用于维度的补充 -- 类cmdb服务
         当产品侧需要新增维度时，表（表字段）不支持动态添加，所以不支持新增维度，需要新增一个viewname（产品上报视图），新增的又必须包含之前的维度（之前的+新增的）
         又生成一张表（表字段）【同时还需要产品侧提供接口把数据写入到表中】，又由于网关侧为了快速处理维度的丰富，索引表数据都是加载到内存中的，重复的维度就会导致内存的使用量增多
3、即便后面有了cmdb，但是业务需要拓展维度时，只能业务各自去查询cmdb里面的数据（告警丰富、告警标签、告警收敛 等都需要查询）

4、架构上面存在很多问题

5、服务上面存在职责不单一的问题
6、日志层面输出格式对查询、排查问题不友好
7、自监控层面：很多服务都缺少自监控
```



新架构： -- 基于时序数据 + openTelemetry collector 【receiver、processor、extender】对 measurement(多维度、多指标)进行加工处理
```text
解决旧架构的问题：
1、旧架构时流式告警，没有复用时序数据，计算任务会占用较大的计算资源  -- 通过时序存储，1分钟、5分钟粒度的聚合数据，后面的比如15分钟，30分钟，1小时，1天，都可以通过前者聚合出来

2、旧架构新增维度比较困难，即便有了维度服务，也只能后置各个业务功能时才能去拿维度   -- 在process中，统一补全业务需要用到的维度 （可插拔的方式）
    这样存在的问题是：
        比如： 告警丰富需要去维度服务拿维度  -- 至少一次RPC调用
              告警标签填充需要去维度服务拿标签  -- 一次RPC调用
              告警收敛（收敛规则需要的维度）需要去维度服务中拿维度  -- 一次RPC调用
              
3、提供类influxDB的AST用来解析influxQL，使用的是influxDB的抽象语法树

4、架构职责清晰，功能模块单一；
```




## 可观测平台
理论： logger、trace、metric

架构：
可观测+告警

数据收集与存储：
  时序数据库： influxDB、prometheus
```text
两者的优缺点：

influxDB

prometheus

    
```

victoriaMetrics是如何解决高基数存储与查询的问题
```text
1、存储引擎优化：
    使用更高效的压缩算法、相同数据比 Prometheus 节省更多空间
    对时序数据（包括高基数标签）进行专项压缩优化
    支持后台自动合并和数据压缩
 
2、查询时基数控制
    限制查询范围
    支持 LIMIT 和 OFFSET
    
3、架构优势
    VMCluster 的分片机制
        数据自动分片到多个节点
        查询负载分散，降低单点压力
        支持水平扩展应对高基数场景

```

prometheus 是如何存放数据的，底层的数据结构是什么
```text


```

prometheus 中的 alertManager是如何与prometheus的server联动的？
```text

```



### influxDB的特性 - 多维度，多指标

在 InfluxDB 中，fields 和 tags 是在数据写入时动态定义的，不需要预先创建 schema。这是 InfluxDB 的一个显著特点
 - 首次写入时自动定义
 - 无预定义 schema
 - 后续写入可以扩展（fields、tags）

当写入数据时，如果上报的数据中没有某个field、tag，influxDB不会有默认值，其表示null； -- 如果需要填充默认值的（客户端上报时处理、服务端写入前处理、查询时转换）

多维度(tags)： 用于指定维度信息，该维度信息是用来描述被监控的对象的
多指标(fields)：这些指标用来表示被监控对象的内部状态的

```text
measurement: server_metrics
tags: 
  host=server01
  region=us-west
  environment=production
  app=web_service
time                cpu_usage memory_used disk_io   request_count response_time
----                --------- ----------- -------   ------------- -------------
2023-08-01T00:00:00Z 45.2      5.7GB      1200iops  1254          42ms
2023-08-01T00:01:00Z 47.8      5.8GB      1350iops  1421          45ms
2023-08-01T00:02:00Z 52.1      6.1GB      1500iops  1567          48ms
```
```text
server_metrics,host=server01,region=us-west,environment=production,app=web_service cpu_usage=45.2,memory_used="5.7GB",disk_io=1200iops,request_count=1254,response_time="42ms" 1690848000000000000
server_metrics,host=server01,region=us-west,environment=production,app=web_service cpu_usage=47.8,memory_used="5.8GB",disk_io=1350iops,request_count=1421,response_time="45ms" 1690848060000000000
server_metrics,host=server01,region=us-west,environment=production,app=web_service cpu_usage=52.1,memory_used="6.1GB",disk_io=1500iops,request_count=1567,response_time="48ms" 1690848120000000000
```


### prometheus - 多维度，单指标


```text

```



### APM
APM：Application Performance Monitoring 的简称，即应用性能监控。

[Dapper](https://static.googleusercontent.com/media/research.google.com/zh-CN//archive/papers/dapper-2010-1.pdf)

SkyWalking


skyWalking goAgent
```text
通过混合编译的方式，在编译的时候，动态的将增强的监控信息插入到go应用程序中

在go build之前需要go get agent，然后将main方法中通过import 导致agent
go build -toolexec 'agent path' -a

大概原理是： 基于toolchain实现
  基于AST解析、操作代码：找到需要增强的地方：net/http、gin、gorm、kafka等等
  文件拷贝和生成 ： 找到需要增强的地方后，生成增强的代理，一般包含，对入参，出参的处理、defer的结束处理
  代理命令执行
```