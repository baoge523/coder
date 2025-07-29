
## k8s 部署使用plan

1、部署单节点的k8s
etcd
kube-apiserver
kube-controller-manager
kube-scheduler
Kube-proxy

kubelet 管理上面的这些static pod
```text
kubelet 通过如果不指定cgroup，默认使用systemd
systemctl start kubelet
systemctl stop kubelet
systemctl status kubelet
```

2、单节点的k8s部署好了后，需要运行一个pod，比如nginx pod，然后可以正常访问到nginx的欢迎页面

3、部署三个节点的k8s；成功后需要部署一个deployment管理的pod(nginx)，且能正常访问

4、网络访问比如：ingress、内部域名访问

5、k8s资源
> Pod
> Deployment  -> 
> DaemonSet
> StatefulSet
> ConfigMap

6、资源监控，k8s中的资源默认支持prometheus监控，需要配置监控k8s资源

7、自定义资源类型