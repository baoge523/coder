# redis的元信息

```bash
redis-cli --help

# redis 6.0 后的ACL认证:  auth  redis:redis
redis-cli -h supervisor_agent  -p port -a auth

```


```text
1. `INFO` - 获取所有信息。
2. `INFO server` - 获取服务器信息。
3. `INFO stats` - 获取统计信息。
4. `INFO memory` - 获取内存使用情况。
5. `INFO persistence` - 获取持久化信息。
6. `INFO clients` - 获取客户端连接信息。
7. `INFO replication` - 获取复制信息。
8. `INFO cpu` - 获取 CPU 使用情况。
9. `INFO commandstats` - 获取命令统计信息。

```

## 查看内存信息
info memory