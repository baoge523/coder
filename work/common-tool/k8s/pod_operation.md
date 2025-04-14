# pod相关的操作
[kubectl-commands](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#scale)

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

## 设置指定资源的副本数
```bash
# 以deployment 资源类型为例
kubectl get deployment -n xxx | grep 'deployment-name'

# 查看deployment的配置信息
kubectl get deployment -n xxx deployment-name -o yaml 

kubectl scale -n xxx --replica=3 deployment/deployment-name

```