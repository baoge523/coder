## docker 安装

## docker的守护进程

当我们执行docker命令时报如下错误时，可能是docker的守护进程没有启动
```text
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
```
处理方式: 重新启动docker的守护进程
```bash
# 查看docker 守护进程状态
systemctl status docker

# 启动 docker
systemctl start docker

```

