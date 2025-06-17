## 排查cpu使用过高的原因

学习和总结cpu的排查经验，这个对自己定位线上linux的问题很有帮助的

### 系统的CPU使用率很高，但为啥却找不到高CPU的应用？
ab（apache bench）是一个常用的 HTTP 服务性能测试工具
```text
# 并发100个请求测试Nginx性能，总共测试1000个请求
$ ab -c 100 -n 1000 http://192.168.0.10:10000/
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, 
...
Requests per second:    87.86 [#/sec] (mean)
Time per request:       1138.229 [ms] (mean)
...
```
从 ab 的输出结果我们可以看到，Nginx 能承受的每秒平均请求数，只有 87 多一点，是不是感觉它的性能有点差呀。那么，到底是哪里出了问题呢？

持续压测: 并发请求5 ，持续600秒
```text
$ ab -c 5 -t 600 http://192.168.0.10:10000/
```

排查思路：
1、通过top命令观察cpu使用量
```text
系统的整体 CPU 使用率是比较高的：用户 CPU 使用率（us）已经到了 80%，系统 CPU 为 15.1%，而空闲 CPU （id）则只有 2.8%。
```


### 系统中出现了大量的不可中断线程和僵尸进程