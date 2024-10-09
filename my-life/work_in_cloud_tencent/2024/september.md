## 2024-09-09
```text
早上：排查了一下流水线问题，定位到需要先部署到环境上验证了，将tapd单走到带验证后才能mr check成功

下午:
1、梳理tce分级告警的流程

2、处理tce3.10.11环境信息
①、部署了amp和dbsql,并通过登录M11环境，查看到了部署后的消息模版
②、但是中途遇到一个问题，tce3.10.11的mysql信息中的密码是加密了的，需要使用工具解密
③、存在运营端登录不上，所以不能创建消息模版

3、处理工单问题，ModifyInstanceGroup处理，这个是barad-console的，需要查询旧的日志中心，通过request找到seqId，然后是eventId

4、和RocketMQ对接，支持维度接入
①、确认问题，即使知道拉取的哪个地域的数据，如果响应结果里面没有地域信息，需要补全地域信息
②、标签拉取的行为如何触发，在拉取产品维度后，会触发标签的拉取任务，基于上报的维度信息丰富出标签六段式去标签系统拉取标签信息
```

## 2024-09-10
```text
早上: 排查工单，用户策略里面不能查到产品侧的实例，经过前端排查是DescribeAllNamespace接口拿不到jsconfig数据，经过后端排查，调野鹤确实拿不到
     原因是：产品侧设置了白名单，通过调用yehe方法拿白名单，但是过期了导致拿不到，所以所有的用户都不能看到，解决方式，1、监控改代码 2、延期白名单
     
     
下午:
1、验证31011的支持电话通知，不是amp成功，需要看环境上面的版本是否是自己部署的版本，手动设置消息模版，并改环境上面的模版id，并验证

2、处理工单问题： 修改策略后，告警历史不生效，排查步骤参考告警历史文档

3、解决触发条件模版添加两个相同的指标时，告警间隔时间会莫名相同的问题，测试环境自测并提mr

4、修改前端性能告警、错误日志创建复合告警失败的bug，测试环境待验证

```

## 2024-09-11

```text
早上：
1、处理工单问题，调用es的接口超过了4k限制，如果修改的话，需要重启es集群，风险太大，所以考虑分包--还没有来得及处理
2、内网加白 --配置修改

下午：
1、内网加白重启各个服务
2、policy-api的功能验证和mr代码：分级告警和非分级告警在配置相同指标后，支持配置不同的告警时间
3、帮助容器同学，排查trpc-go支持polaris访问
4、前端性能告警错误日志支持复合告警自测验证，---无法验证
5、修改安全漏洞、go版本升级，直接依赖和间接依赖，间接依赖不知道怎么改版本
```

## 2024-09-12
```text
早上：
1、修改安全漏洞，升级golang依赖版本，通过go mod graph  module_name@version 查看依赖，如果已经升级到最高版本间接依赖还是不能升级，那么
可以通过replace 替换版本

下午:
1、内网白名单加白
2、发布policy-api服务并验证
3、梳理同步器相关的逻辑(policy-synchronizer)
4、排查 tag-server的错误，超时异常

```
## 2024-09-13
```text
上午：
1、内网白名单发布
2、多伦多去掉指定的appid

下午：
1、策略组绑定实例558个，导致mysql死锁问题
```


## 2024-09-14
```text
上午：
1、修复安全漏洞后，无法编译部署的问题 -- 未解决

下午：
1、排查修复安全漏洞后，无法编译部署的问题 -- 未解决
2、分级告警，恢复时无法通知指定人问题
```

## 2024-09-18
```text
上午：
1、将31011的m11环境部署好支持语言通知，并通知测试同学和产品同学验证
2、排查修改安全漏洞后的编译不通过的问题：
   vendor中总是少各种cgo的东西，但是通过go env 查看 cgo enable = 1

下午：
1、排查支持告警登录的链路，丢弃使用calcValue为支持告警登录的方式，换成在cPolicyGroup中添加一个alarmLevel字段
   ①、保存策略时，在1.0的cPolicyGroup中的alarmLevel字段保存值：serious、warn、remind 或者 null("")
   ②、因为有同步器的存在，所以会监听mysql binlog，于是修改policy-synchronizer 中对cPolicyGroup的处理添加alarmLevel字段并存放到adp的t_detect_policy的tags中
   ③、adp检查时，会将t_detect_policy中的tags里面的数据丰富到extraTags中，并透传给amp
   ④、amp接受到extraTags里面的alarmLevel后，将数据丰富到告警通知里面发送出去
   ⑤、amp写入告警历史，支持alarmLevel
   ⑥、告警历史支持alarmLevel的查询 (需要和前端联动)

```

## 2024-09-19
```text
早上：
1、排查m6环境为什么不能告警的问题，原因是计算告警触发时间的时候返回了false导致，具体问题还有待排查
2、查看告警历史时，发现已经存在了alarmLevel，然后验证了alarmLevel没有被使用

下午：
1、编写policy-synchronize的转换代码
2、编写domain_policy的代码，支持创建策略时支持告警级别、支持告警级别的修改、支持告警级别的回显
3、编写amp的alarmLevel的相关处理

```

## 2024-09-20
```text
早上:
1、处理安全漏洞的vendor问题

下午：
1、处理安全漏洞的流水线版本问题、从1.15改成1.18 -- 未解决
2、支持告警等级的功能开始编译打包验证
```

## 2024-09-23
```text
早上：
1、处理monitor-alarm的编译问题
   domain_policy和console都依赖于monitor-alarm,选择他们依赖的版本，基于这个commit-id做字段修改并打一个tag共他们使用
   但是monitor-alarm的go.mod是go1.13的所以如果使用本地环境1.18生产的protocol会带有1.18才有的泛型，这样会导致冲突，所以我使用了1.16版本来生成protocol
   
下午：
1、处理console和domain_policy的编译问题
2、编写amp的支持告警等级的代码
3、编译amp
```
## 2024-09-24
```text
早上：
1、修改domain_policy中处理告警等级的逻辑
2、部署环境
3、修改数据库，支持alarm_policy字段

下午：
1、和前端同步告警等级的接口信息
2、处理domain_policy创建策略时报错的问题，// 问题
3、验证创建策略支持告警等级
```

## 2024-09-25
```text
早上:
1、验证policy-synchronizer同步策略问题
   ①、全量同步需要同步绑定对象的策略  SELECT groupId, viewName, ownerVin,appId, lastEditUin, projectId, groupName, isShielded, isUnionRule, alarnLevel FROM PolicyGroup WHERE groupid in (SELECT groupId FROM rApplicationPolicy WHERE appAddress in (-1,0,0,-1,50000001))
   ②、好像只全量同步了，没有增量同步，需要排查
   ③、使用了两个db:master_StormCloudConf(创建策略时保存到这个db)、StormCloudConf(这个db是同步器同步的db，创建策略后，会有一定的延时才会同步到该db)
2、验证安全漏洞修改后，存在// 的问题
3、排查告警没有携带标签的问题，绑定对象的标签存放在2.0数据的policy_tag_instance表中

下午:
1、提交安全工单的mr和修复//访问问题(ygin)的mr
2、排查amp告警失败问题，原因是告警丰富出现了问题，最终定位到是barad-api 中的barad_amsconfig_proj服务连接redis时报错了
  通过php代码发现: redis的连接是写死的，但是部署的时候，会渲染对应的配置文件
  所以我们需要等barad-api pod，然后通过find / -name 'config.*.inc.php'查找，找到文件后，就可以根据配置文件如何渲染配置排查问题
  
3、tce也有类似于云api的概念，所以需要需要在ted/barad 对应分支 修改yunapi/data/api3tcloud/monitor.json中的参数，然后编译
```

## 2024-09-26
```text
早上：
1、验证tce 云api修改参数后，部署
2、编写tce/barad 的db信息
3、编写tce/barad 消息模版信息

下午:
1、无法告警排查告警问题
2、尝试如何排查adp检查数据的问题  --- 待梳理文档，看adp是如何工作的，大数据量是如何保证正常处理的
```

## 2024-09-27
```text

早上：
1、确定了M6环境能正常写入告警历史，只不过无法查询出来，且告警历史中有告警分级
2、将消息模版db和消息模版yehe都支持告警等级了

下午：
1、排查告警历史编译报错问题
  原因是自己的protobuf有问题，需要重新安装一个,通过protoc --version 查看，感觉版本都不一样
  了解一下protobuf是什么
```


## 2024-09-29
```text
早上：
1、M6环境可以用，排查告警历史查询不出来的问题
  通过堆栈报错信息发现，是版本的问题，将grpc-alarm版本改成v0.0.2版本
2、编写告警历史查询展示的代码

下午：
1、验证修改版本后是否可以查询出告警历史
2、查询出来的告警历史不包含告警等级信息，排查代码问题
   ①、es查询出来的是否有告警等级
   ②、数据转换是否缺少对告警等级处理
3、针对修改后的db信息和消息模版信息，流水线部署后，看是否存在

```
## 2024-09-30
```text
早上：
1、学习protocol buffers

下午：
1、查看protocol buffers 的文档
2、处理tce3100的支持告警电话的出包问题
```