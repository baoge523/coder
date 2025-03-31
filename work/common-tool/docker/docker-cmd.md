## docker 命令
[docker_cli_command](https://docs.docker.com/reference/cli/docker/)
### docker的基础命令

查看docker镜像
```bash
docker images 
docker images | grep 'xxx'
```

启动docker镜像
```bash
docker run -d nginx:alpine    // -d 表示后台运行

docker run -i -t nginx:alpine /bin/bash  // 启动并打开控制台
```

查看docker镜像进程
```bash
docker ps
```

进入一个已经运行的docker程序
```bash
docker exec -i -t docker_image_id  /bin/bash
```






### docker 系统性命令