## kubectl scale
用于管理pod的副本的

usage:
```text
kubectl scale [--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME)
```


examples:
```text
  # Scale a replica set named 'foo' to 3
  kubectl scale --replicas=3 rs/foo
  
  # Scale a resource identified by type and name specified in "foo.yaml" to 3
  kubectl scale --replicas=3 -f foo.yaml
  
  # If the deployment named mysql's current size is 2, scale mysql to 3
  kubectl scale --current-replicas=2 --replicas=3 deployment/mysql
  
  # Scale multiple replication controllers
  kubectl scale --replicas=5 rc/example1 rc/example2 rc/example3
  
  # Scale stateful set named 'web' to 3
  kubectl scale --replicas=3 statefulset/web

```

```bash
# 表示将deployment-name管理的pod的副本改成3个
kubectl scale --replicas=3 deployment/deployment-name 

```