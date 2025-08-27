## 通过EAV的方式存放信息
Entity-Attribute-Value
```text
Entity：实体，代表一个业务对象，比如上面的例子里的商品。
Attribute：对象的属性，属性并不是作为实体单独的一列来进行存放，而是存储在一组单独的数据库表中。
Value：指特定属性所关联的值
```

### 传统的列模型
1、定义列属性
2、每一行表示一个对象

当需要扩展属性时，就必须添加列，删除属性时，需要删除列
```text

user 表
字段信息: id、name、birthday、sex、userName、password ...
一行数据，就表示一个对象
1、"张三"、1998-04-27、男、"zhangsan"、"aaaa"
2、"李四"、1988-04-27、男、"lisi"、"bbbb"

如果要添加一个字段，那么就得增加一个列，如果经常变动，那就比较麻烦
```

### EAV （行模式）
针对与user这张表，将其拆分成三种表，分别是:
eav_entity_user: 用于存放实体对象信息
eav_attribute_user: 用于存放属性定义信息
eav_value_user: 用于存放具体的值的信息
```text
eav_entity_user
id、entity_id、check_id、create_time、update_time

eav_attribute_user:
id、attr_name、attr_type、is_key（是否具有唯一性，业务层面约束）

eav_value_user:
id、entity_id、attr_id、value

```
比如:
```sql
CREATE TABLE `eav_value_user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `entity_id` varchar(255) DEFAULT NULL,
    `attr_id` int DEFAULT NULL,
    `value` varchar(768) DEFAULT NULL,
    `checksum` varchar(255) DEFAULT NULL,
    `version` int DEFAULT '0',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE `uniq_eid_aid` (`entity_id`, `attr_id`),
    KEY `idx_aid_sv` (`attr_id`, `value`)
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARSET = utf8mb3
```

直接的关系:
eav_attribute_user: 是该表的属性定义，及表示有哪些属性信息，可以动态的往这里面添加，可以快速支持新增属性

eav_entity_user: 实体对象，一个对象就往里面插入一条数据

eav_value_user：实体对象的属性信息，attr_id表示这个属性是什么




