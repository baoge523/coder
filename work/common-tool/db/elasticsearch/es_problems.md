## 使用es中遇到的问题

#### 对于es的index做增量的新增字段
es的index在创建后是不允许修改里面的定义信息的，包括已有的字段名称、字段类型
**但是却可以新增字段**

查看es index的定义信息如下:
```bash
curl -X GET http://ip:port/index_name/_mapping  -H "Content-Type:application/json"
```
我们可以执行执行有新字段的写入操作到该index中，然后再查看该index，就可以看到新的字段的定义会写入到mapping中
