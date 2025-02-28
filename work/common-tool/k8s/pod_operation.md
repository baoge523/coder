# pod相关的操作

## 查看pod的版本信息
```text
快捷方式:
kg policy  # 模糊查询，拿到pod名称

kubectl get pod <pod-name> -n tce -o jsonpath='{.spec.containers[*].image}'

kubectl get pod <pod-name> -n <namespace> -o jsonpath='{.spec.containers[*].image}'
```

## 查看pod 下的所有containers的名称
```bash
kubectl get pod alarm-domain-policy-5ddb45cff9-zmhm8 -n tce -o jsonpath='{.spec.containers[*].name}'
```