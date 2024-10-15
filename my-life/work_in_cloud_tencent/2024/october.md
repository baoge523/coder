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