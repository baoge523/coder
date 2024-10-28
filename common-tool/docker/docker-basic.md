## docker 基础技术 linux namespace
[参考博客](https://coolshell.cn/articles/17010.html)

docker如何实现镜像环境的隔离，以达到同时在一个宿主机上运行多个docker镜像容器服务
其底层功臣就是linux namespace

通过linux namespace来达到各种隔离比如UTC、PID、文件系统等

[linux namespace的文档信息](https://lwn.net/Articles/531114/)

| 分类	|系统调用参数	|相关内核版本|
|--------|------|----------|
|Mount namespaces	|CLONE_NEWNS	|Linux 2.4.19|
|UTS namespaces	|CLONE_NEWUTS	|Linux 2.6.19|
|IPC namespaces	|CLONE_NEWIPC	|Linux 2.6.19|
|PID namespaces	|CLONE_NEWPID	|Linux 2.6.24|
|Network namespaces	|CLONE_NEWNET	|始于Linux 2.6.24 完成于 Linux 2.6.29|
|User namespaces	|CLONE_NEWUSER	|始于 Linux 2.6.23 完成于 Linux 3.8)|


### 依赖的系统调用
- clone() – 实现线程的系统调用，用来创建一个新的进程，并可以通过设计上述参数达到隔离。
- unshare() – 使某进程脱离某个namespace
- setns() – 把某进程加入到某个namespace

### UTS namespace
Unix Time-sharing System Namespace 主要作用是隔离系统标识和主机名
通过 UTS namespace 不同的进程可以拥有独自的主机名和域名，对于容器化环境十分重要

系统标识: 

### IPC namespace  
Inter-Process Communication namespace 进程间通信命名空间

主要作用: 隔离不同进程间的通信资源，通过 IPC namespace 进程可以在同一台机器中拥有自己的独立的信号量、消息队列、共享内容

主要用途: 对容器化十分重要
- 资源隔离
- 多租户管理

**IPC:** 进程间通信
> 用于允许不同进程间相互交互数据和信息
> 提供了多种机制、帮助进程协调和共享资源

常见的IPC资源有: 这些资源都是拥有linux系统中协调进程通信
- 管道 pipe
- 命名管道 fifo
- 消息队列
- 共享内存
- 信号量



### PID namespace
作用: 提供进程的隔离，在子进程容器中，创建PID=1的进程

这使得在不同的 PID Namespace 中，进程可以独立运行，而不会相互干扰

### Mount namespace
作用: 文件系统的隔离，使得不同的进程拥有独立的挂载点文件视图

具体功能:
- 文件系统隔离
- 灵活的文件系统管理
- 支持容器化


例子:
```text
top ps 命令都会去读/proc下的文件信息，但是在子进程容器中，不应该读取到父进程(其他进程的信息)；
所以需要提供文件系统的隔离，将/proc挂载到子经常容器中

比如:
创建一个子进程容器的根目录 /rootfs
将/proc目录挂载到/rootfs/proc下: mount -t proc proc /rootfs/proc
其中参数信息:
  -t proc 表示指定文件系统类型为 proc
  proc    表示源、代表/proc文件系统
  /rootfs/proc  表示目标挂载点
  
其中需要确保/rootfs/proc已经存在，如果不存在，需要先创建目录: mkdir -p /rootfs/proc

```

注意：
 通过clone系统调用创建的子进程指定使用了Mount namespace，如果不指定 "挂载"操作，其子进程和宿主机的文件系统一样
```text
Mount Namespace 修改的，是容器进程对文件系统“挂载点”的认知。
但是，这也就意味着，只有在“挂载”这个操作发生之后，进程的视图才会被改变。
而在此之前，新创建的容器会直接继承宿主机的各个挂载点


这就是 Mount Namespace 跟其他 Namespace 的使用略有不同的地方：它对容器进程视图的改变，一定是伴随着挂载操作（mount）才能生效。


```


### Network namespace
作用: 提供网络环境的隔离，使得不同的网络命名空间拥有独立的网络设备、路由表、ip、port等网络信息。这种隔离运行在同一个物理机器上运行多个网络环境

常见用途:
- 容器化
- 测试和开发
- 安全性
- 多租户环境

### User namespace
作用: 提供用户和组ID的隔离，允许不同的进程在独立的用户空间运行

```text
简而言之: 允许进程(容器)拥有自己的用户空间，容器内的用户ID(UID)、组ID(GID)都是对宿主机的映射，同时容器间具有隔离性、且对宿主机的访问受限(只允许访问自己容器内)

“允许不同的进程在独立的用户空间中运行”指的是在 Linux User Namespace 中，每个用户命名空间都可以有自己的用户 ID（UID）和组 ID（GID）映射。具体来说：
1. **UID/GID 映射**：在一个用户命名空间内，用户和组的标识符（UID 和 GID）可以与主机系统的标识符不同。这意味着容器中的用户（如 UID 1000）可以在主机上对应到另一个用户（如 UID 65534），从而实现隔离。
2. **权限隔离**：即使容器内的用户是特权用户（例如 UID 0），在主机上它们可能仍然是非特权用户。这种方式使得容器内的进程可以执行某些操作，而不会对主机系统造成安全风险。
3. **独立性**：每个用户命名空间内的进程对其他命名空间的用户和组没有可见性。这样，不同容器或用户之间的资源和权限不会相互影响。
通过这种机制，Linux 用户命名空间提升了安全性，使得可以在同一台物理机上安全地运行多个应用或容器，而不必担心它们会干扰或访问彼此的资源。
```