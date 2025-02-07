# pod相关的操作

## 查看pod的版本信息
```text
快捷方式:
kg policy  # 模糊查询，拿到pod名称

kubectl get pod <pod-name> -n tce -o=jsonpath='{.spec.containers[*].image}'

kubectl get pod <pod-name> -n <namespace> -o=jsonpath='{.spec.containers[*].image}'
```