
## go pprof 
参考官网文档: https://go.dev/blog/pprof

### pprof的作用
发现golang程序在运行中的问题：比如内存分配、gc、goroutine、thread trace 、cpu等等

当程序出现意外的oom，需要排查定位oom的原因并解决
当程序运行比预期的慢，期望让其运行得更快一些


### 得到 pprof的方式

1、通过http的方式
```go
import _ "net/http/pprof"

// http://localhost:6060/debug/pprof/
func func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
}
```

2、通过程序生成的方式
```go
// 可以通过flag的方式获取 cpuprofile memoryprofile ....
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
    f, err := os.Create(*cpuprofile)
    if err != nil {
    log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    }
}
```

### 分析 pprof

1、分析http得到pprof
```linux
go tool pprof http://localhost:6060/debug/pprof/profile   # 30-second CPU profile
go tool pprof http://localhost:6060/debug/pprof/heap      # heap profile
go tool pprof http://localhost:6060/debug/pprof/block     # goroutine blocking profile
```

2、基于go tool pprof 分析pprof文件
```linxu
# ./havlak1 是go二进制文件，获取cpuprofile文件
./havlak1 -cpuprofile=havlak1.prof

# 通过go tool pprof 分析pprof文件
$ go tool pprof havlak1 havlak1.prof
```
①、top N
> 获取前N个样例信息(shows the top N samples in the profile)

②、web 以web方式查看
> web mapaccess1 只展示指定的方法

③、list DFS  
> 查看DFS方法的调用信息，并展示了资源消耗情况(具体到哪一行)

这里可以查看heapprof中在使用list时查看实例里面的对象个数
> go tool pprof --inuse_objects havlak3 havlak3.mprof

这里表示在使用web时，只展示权重在10%以上的
> go tool pprof --nodefraction=0.1 havlak4 havlak4.prof

### 分析 pprof 命令得到的信息

cpuprofile
```text
(pprof) top10
Total: 2525 samples
     298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
     268  10.6%  22.4%     2124  84.1% main.FindLoops
     251   9.9%  32.4%      451  17.9% scanblock
     178   7.0%  39.4%      351  13.9% hash_insert
     131   5.2%  44.6%      158   6.3% sweepspan
     119   4.7%  49.3%      350  13.9% main.DFS
      96   3.8%  53.1%       98   3.9% flushptrbuf
      95   3.8%  56.9%       95   3.8% runtime.aeshash64
      95   3.8%  60.6%      101   4.0% runtime.settype_flush
      88   3.5%  64.1%      988  39.1% runtime.mallocgc
```
```text
表示有2525个样本
第一列表示方法处于运行中的样本数(不包含等待中的)
第二列表示处于运行中的func个数占总的个数的百分比
第三列是累计的占比：11.8% = 0% + 11.8%； 22.4% = 11.8% + 10.6% ...
第四列表示处于运行和等待中的func到的个数
第五列表处理运行和等待中的func占总的个数的百分比
第六列表示对应的方法
```
heapprofile
```text
$ go tool pprof havlak3 havlak3.mprof
Adjusting heap profiles for 1-in-524288 sampling rate
Welcome to pprof!  For help, type 'help'.
(pprof) top5
Total: 82.4 MB
    56.3  68.4%  68.4%     56.3  68.4% main.FindLoops
    17.6  21.3%  89.7%     17.6  21.3% main.(*CFG).CreateNode
     8.0   9.7%  99.4%     25.6  31.0% main.NewBasicBlockEdge
     0.5   0.6% 100.0%      0.5   0.6% itab
     0.0   0.0% 100.0%      0.5   0.6% fmt.init
(pprof)
```
```text
表示样本总共82.4MB
第一列表示func使用的heap
第二列表示func使用的heap占比
第三列表示累计占比
```