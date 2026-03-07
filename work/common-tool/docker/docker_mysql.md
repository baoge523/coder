## docker install mysql

### docker 拉去指定mysql镜像
docker pull mysql:8.0

docker pull mysql:laster

### docker 启动mysql服务
```text
# 运行容器
docker run -d \
  --name mysql8 \
  -e MYSQL_ROOT_PASSWORD=a123456 \
  -e MYSQL_DATABASE=my_test \
  -e MYSQL_USER=a \
  -e MYSQL_PASSWORD=a123456 \
  -p 3306:3306 \
  -v mysql_data:/Users/yongzhao/tools/mysql \
  --restart=unless-stopped \
  mysql:8.0 \
  --character-set-server=utf8mb4 \
  --collation-server=utf8mb4_unicode_ci
```

### 常用命令
```text
# 查看运行中的容器
docker ps

# 查看 MySQL 日志
docker logs mysql-container

# 停止容器
docker stop mysql-container

# 启动容器
docker start mysql-container

# 删除容器（数据会保留在卷中）
docker rm mysql-container

# 删除数据和容器
docker rm -v mysql-container

# 备份数据
docker exec mysql-container sh -c 'exec mysqldump --all-databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > backup.sql
```

### mysql 使用样例
```text
# 连接到本地 MySQL
mysql -u root -p

# 连接到远程 MySQL
mysql -h 192.168.1.100 -u username -p

# 指定数据库
mysql -u username -p database_name

# 执行 SQL 文件
mysql -u username -p < file.sql
```