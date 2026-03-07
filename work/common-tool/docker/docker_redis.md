
## docker install redis

1. 从docker hub中拉去redis镜像（可以指定版本）
```linux
docker pull redis

docker pull redis:7.0
```

2. 运行redis容器
```linux
docker run --name my-redis -p 6379:6379 -d redis
```
```text
-name my-redis：容器名称设为 my-redis
-p 6379:6379：将主机的 6379 端口映射到容器的 6379 端口
-d：后台运行
redis：使用的镜像名称
```


常用命令
```text
# 查看运行中的容器
docker ps

# 停止容器
docker stop my-redis

# 启动已停止的容器
docker start my-redis

# 删除容器
docker rm my-redis

# 查看日志
docker logs my-redis
```

### mac 安装redis-cli
```text
# 下载稳定版
curl -O https://download.redis.io/releases/redis-7.2.4.tar.gz

# 解压
tar xzf redis-7.2.4.tar.gz
cd redis-7.2.4/src

# 直接使用（不需要安装）
./redis-cli --version

# 或复制到合适位置
cp redis-cli ~/.local/bin/
```

### 连接redis
redis-cli -h 127.0.0.1 -p 6379 -a yourpassword