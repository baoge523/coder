## kubectl rollout 

manage the rollout of a resource

valid resource types include:
 - deployment
 - daemonsets
 - statefulset

usage:
```bash
kubectl rollout subcommand
```
subcommand include:
 - history   查看rollout之前的版本和配置信息
 - pause     暂停指定的类型资源，在之后的修改就不会自动更新，直到使用resume为止
 - restart   重启指定类型资源
 - resume    恢复暂停的资源
 - status    查看状态
 - undo      回滚到上一个版本

### history
views previous rollout reversion and configurations

Usage:

```bash
kubectl rollout history (type name | type/name) [flags]
```

### pause

### restart

### resume

### status

### undo
