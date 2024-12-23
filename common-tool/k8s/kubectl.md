# kubectl
[官方文档kubectl](https://kubernetes.io/docs/reference/kubectl/)

kubectl 的命令格式
```bash
kubectl get deploy tcloud-barad-alarm-amp -n tce -o yaml > 1.yaml
```


获取指定名称空间下所有的pod
```bash
kubectl get pods -n namespace_name 

# for example
kubectl get pods -n tce

# 检索指定的pod名称
kubectl get pods -n namespace_name | grep pod_name

# for example
kubectl get pods -n tce | grep policy

```

查看pod启动的详情信息:
```bash
kubectl describe pod pod_name -n namesapce_name

# such as:
kubectl describe pod tcloud-barad-alarm-policy-8467f7d74b-9hcnx -n tce
```

查看pod容器日志
```bash
kubectl logs pod_name -n namespace_name

# such as
kubectl logs tcloud-barad-alarm-policy-8467f7d74b-9hcnx -n tce
```

