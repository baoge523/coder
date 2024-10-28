## 2024-10-08
```text
早上:
1、排查消息模版固化部署后，不生效的原因，--原因是自己改错地方了
2、处理tce 告警等级31011环境的问题

下午:
1、验证rum创建复合告警问题
2、处理工单问题
3、tce支持告警等级的mr信息
```

## 2024-10-09
```text
早上:
1、排查tce31011的amp问题，不支持电话语音告警的原因 -- 未果
2、排查tce3100的告警问题，原因是告警丰富问题，真实原因是barad-api中/use/local/services/conf下面没有redis_ckv_barad.conf文件

下午：
1、排查告警列表查询问题，原因是告警历史更新时将alarmLevel字段给更新为空串了，解决方案在amp处理的入口处，将alarmLevel丰富到params中
2、处理白名单问题，并发布
3、处理工单问题
4、tce3100告警信息支持告警等级，邮件支持，短信的持续触发支持，但是短信的触发存在问题
```

## 2024-10-10
```text
早上：
1、发布内网白名单和标签多选白名单
2、tce3100告警历史页面支持告警等级多选
3、排查工单问题

下午：
1、排查prometheus告警收到不到的问题
2、排查告警丰富问题
3、tce3100告警历史页面支持告警等级多选、涉及的模版有：tce/barad的yunapi、monitor-alarm、barad-api-go(console)、consumption 
```

## 2024-10-11
```text
早上：
1、修复message-quota-tool组件的安全漏洞问题、并流水线验证  --- mr合并出现问题，导致tapd单不支持合并代码
2、和测试对tce31011支持语音告警的需求
3、验证tce3100告警等级问题
4、rum创建复合告警的bug提供mr

下午:
1、处理rum、apm、cat不需要支持复合告警的问题
2、处理工单问题
3、处理qce/tdmq pulsar接入维度服务的模型创建、source config创建、pipeline的创建、并验证接受数据
```

## 2024-10-12
```text
早上:
1、处理安全漏洞的mr问题
2、讨论优化基于cmdb的默认策略的的代码优化问题
3、json序列化和反序列protocol buffer结构数据可能存在的问题

下午：
1、处理31011的amp代码冲突问题，并合并代码，注意git的集中状态:工作区(working tree)、暂存区(stage)、版本库
2、梳理默认策略的处理逻辑
3、排查工单问题
```

## 2024-10-14
```text
早上:
1、发布多标签的白名单
2、tce31011的ted/barad的mr有问题，排查原因；
   ①、为什么改了代码会导致所有的组件都会监听改动 --- tce31011极光上的监听文件的路径存在问题
   ②、yunapi_barad 在ted/barad 的.ted/applications/下不存在yunapi_barad，导致mr check时找不到接入的项目
        改变方式： 在ted/barad的.ted/applications下添加yunapi_barad，并配置相关信息，比如chart.yaml、values.yaml、template\applications.yaml
                 在极光的yunapi-barad项目tce3.10.11然后编辑版本，修改配置架构

下午:
1、学习gofmt
2、查看golang编程规范
3、处理tce31011新增电话告警的mr问题
4、处理tdmq的pulsar的默认策略问题
```

## 2024-10-15
```text
早上:
1、tce3.10.11中的mr检查到多个组件，但是组件中没有包含指定修改组件的对应的tce3.10.11版本，而是tad3.10.11版本
   而且有时候修改了极光的对应组件的监听文件后，重新打开mr没有生效
2、安全工单问题，当时自己为了修复安全工单，帮自己私有仓库做了代码提交，这就导致远端私有仓库和远端公有仓库的提交信息不一致，所以需要回滚自己的提交
  通过 git reset --hard commit_id （这个commit_id是需要回滚的信息的前一个commit_id）
  git push -f 

  git pull upstream master --rebase
  git push -f 

下午:
1、处理工单问题，告警触发条件模版中有事件信息，但是策略列表却展示不出来
   排查定位原因：首次使用事件创建触发条件模版时使用的viewname是dts，但是查询策略使用的viewname是dts_replication，导致查询不出来
2、tce3.10.0支持告警等级的测试需求会议
3、mr组件问题
```

## 2024-10-16
```text
早上:
1、tce31011的mr检查问题，并修复，但是有一个固化消息模版的mr需要通过极光页面的方式接入
2、参考golang中的使用redis的方式
   使用github.com/go-redis/redis/v8的方式，其中使用的基于接口的方式操作redis
   

下午:
1、默认策略绑定服务支持appid换uin的方式，并将其缓存 --- 引入redis的使用
2、redis使用github.com/go-redis/redis/v8来完成缓存
3、重构默认策略绑定服务的main方法，并解决循环依赖的问题
4、提交需要极光合并的mr，但是出现了问题，应该是激光版本的问题  ---  待完成
```

## 2024-10-17
```text
早上:
1、添加默认策略服务的redis、accountClient配置
2、编译默认策略服务的测试包、解决编译失败问题
3、测试tce3100的电话语音告警等级功能 
4、处理tce3100的固化新增告警等级db字段的功能  

下午:
1、排查tce3100只支持告警等级的电话语音模版
2、tce3100短信无法发送，原因是短信额度用完了   TOTAL_QUOTA#1255000575#202410
   tce默认使用redis存放短信额度，但是m6 master ip 上没有redis-cli导致redis登录不上，无法修改
3、准备看看go-redis的文档，学习一下
```

## 2024-10-18
```text
早上：
1、看go-redis中文文档，管道事务、发布订阅、配置项  （客户端的信息）
2、了解了redis 6.0后支持ACL认证信息后，才同时需要用户名密码；在redis 6.0之前只需要密码认证，所以在认证时需要先确认redis的版本信息
3、TLS (传输层安全 Transport Layer Security)是一种加密协议，用于在计算机网络中提供安全的通信; 
   - 数据加密
   - 身份验证
   - 数据完整性

下午:
1、看GORM的中文文档，学习如何创建、批量创建等其他细节的操作
2、看GORM的查询部分文档
```

## 2024-10-21
```text
早上：
1、查看gorm的context文档信息


下午：
1、周会
2、梳理支持tce31011的amp动态时区问题
3、和测试同学对齐告警等级的测试评审

```

## 2024-10-22
```text
早上:
体检

下午:
1、处理部署sql持久化alarm_level的问题 --未处理，看报错信息应该是数据库升级flyway的问题
2、处理告警等级测试中遇到的问题
3、学习docker的发展历史
4、理解容器编排技术：容器编排是指自动化管理和协调多个容器的部署、扩展和运行的过程。
    - 自动化管理
    - 负载均衡
    - 故障恢复
    - 服务发现
5、tce3100多次修改后导致产生多条件不同alarmId的告警历史，而在恢复时，只会恢复最新的一条告警历史
  尝试的解决方案：
     ①、恢复所有的告警历史，但是由于alarmId变了，导致alertId也跟着改变了，redis中存放的就只有最近的alarmId，所以无法实现
     ②、将其他的告警历史失效，讨论后，存在无法失效的情况，原因是amp和adp同时处理，但往往adp因为链路长，会晚到，又会将失效状态改成告警中的状态
```

## 2024-10-23
```text
早上:
1、tce3100持久化db的sql执行升级 flyway，该表在各种的db里面 flyway_schema_history
2、处理tce3100的adp的相关逻辑

下午:
1、解决tce3100查询告警内容的问题，原因是-amp下的update_template.py脚步没有执行
2、排查执行amp失效告警历史的问题，-- 排查到时因为resultText写入不进去es中导致的，定义信息有，写入的参数有
```

## 2024-10-24
```text
早上:
1、学习docker容器的基础信息，比如docker是如何基于linux namespace来做的进程容器隔离
2、排查tce问题

下午:
1、排查tce3100告警失效的问题，查询resultText查询不出来，怀疑没有写入 --- 但是经过排查后发现是查询语句将resultText给忽略了
2、学习linux中的namespace的限制
   UTS namespace  提供hostname、系统标识的隔离
   IPC namespace 提供进程间通信隔离，进程拥有独立的通信信息：管道、信号量、消息队列、共享内存
   PID namespace  提供进程的隔离，支持进程容器中的PID为1
   Mount namespace 提供文件系统的隔离，支持各个进程拥有独立的文件系统
   Network namespace 提供网络环境的隔离，比如网络设备、路由表、host、port；支持一个物理主机上拥有多个网络环境
   User namespace  提供用户、组ID的隔离
```

## 2024-10-25
```text
早上：
1、排查告警失效问题
2、查看amp中的代码处理告警失效的问题

下午：
1、修复告警失效问题，并总结原因
2、修复修改策略名称报错的问题
3、处理tce的其他的问题
```