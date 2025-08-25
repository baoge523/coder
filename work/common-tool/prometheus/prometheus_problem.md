## prometheus 中的问题

参考: https://yunlzheng.gitbook.io/prometheus-book/part-iii-prometheus-shi-zhan/operator/use-operator-manage-prometheus

在共计 800 个左右的pod时，prometheus在12核独占node的总体CPU就已经长期徘徊在高位
https://github.com/prometheus/prometheus/issues/8014   issue已经解决

### 排查思路

prometheus本身也是golang程序，可以通过pprof定位到底是哪里出现了问题，然后再通过源码调试，找到其问题所在的根本原因

通过pprof发现，其中scrape 里的 reload 函数

reload函数是如何触发执行的呢？看prometheus的代码发现：manager.go中每五秒判断一次是否有reload事件发生，如果有reload就会被直接

tips：reload事件放到了一个chan中，这样存在一个问题时，如果发生了多次reload的触发事件(没有收敛)导致每五秒钟，必定有reload被执行 --需要收敛reload事件
     而reload又是一件比较消耗性能的事情，这样就导致的性能瓶颈
```text
通过 client 监听到 endpoint 的变化，然后经过一系列的管道中转(chan之间的相互通信)，最终传递给 scrape/manager.go 后触发 reload。

监听的endpoint是根据namespace级别监听的

最终发现是kube-system (namespace)下的版本变更导致的： 使用 kubectl get endpoints 命令也可以确认这一点

也正是因为这点，在 k8s 集群中部署 prometheus 的时候，如果有 watch kube-system 资源的时候，其 reload 频度就会受其影响，因为 endpoint 一直变，导致不断触发 reload

尽管看起来 kube-system 的一些 endpoint 仅仅是版本号变了而已，却依旧真的会触发 reload 链路。

这里的版本号变更的原因 --- 是因为kube-controller-manager的版本频繁变更是由于leader election的心跳导致的

```


```text
单次 reload的开销所在之地：
reload 会将所有的 servicemonitor 都重新根据本 namespace 的全量 targetGroup 来刷新建立一遍 target 信息，存在冗余的 label sort
```
解决方案：
> 优先考虑合并serviceMonitor
>

### 实践指南

```text
配置的采集 job（基于 prometheus operator 实践时就是 servicemonitor），所匹配的 target 不应重复  -- 如果重复了，在并发写的时候，可能会导致版本冲突
```