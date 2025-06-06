## mysql 是如何加锁的
环境版本:mysql5.7.xxx

查询、新增、删除、修改

注意：检索到的行都需要加锁
锁是锁在index上的

### 锁信息
行锁、gap lock(间隙锁)、next-key lock

加锁规则
```text
两个“原则”、两个“优化”和一个“bug”。
原则 1：加锁的基本单位是 next-key lock。next-key lock 是前开后闭区间。
原则 2：查找过程中访问到的对象才会加锁。
优化 1：索引上的等值查询，给唯一索引加锁的时候，next-key lock 退化为行锁。
优化 2：索引上的等值查询，向右遍历时且最后一个值不满足等值条件的时候，next-key lock 退化为间隙锁。
一个 bug：唯一索引上的范围查询会访问到不满足条件的第一个值为止。 (范围查询时，唯一索引和普通索引一样)
```
mysql version 8.0.35 时，已经没有一个bug的加锁规则了

#### next-key lock
next-key lock = 行锁 + gap lock

比如一个操作需要加(5,10]的next-key lock，它会先加(5,10)的gap lock，然后再加10行锁

next-key lock 是左开右闭加锁规则，比如(5,10]

#### 行锁 (RC,RR)
锁定某一行，行锁分读锁和写锁；

冲突规则：读-写，写-读，写-写


#### gap lock (RR)
gap lock (间隙锁)：只有在mysql的可重复读的隔离级别(RR)下才有;

冲突规则：gap lock 与往这gap中插入数据时才会冲突

注意点: 多个事务可以同时对一个gap加gap lock，不会冲突

gap lock的目的是为了解决mysql中的幻读问题，保证了在同一个事务中多次相同sql的查询操作，不会读取到不一样的数据(强调插入的数据)

如果在业务中觉得gap lock 锁住的区间太大了，可以考虑，将隔离级别设置成read-committed + binlog = row 模式，解决了幻读问题，解决了binlog的日志恢复出来的数据与当前数据库不一致的行为
binlog = row：
 表示sql涉及到的每行都记录一个sql到binlog中，这个虽然解决了read-committed下的幻读问题，但是会导致binlog文件大小剧增

### 操作
select ... for update 和 update 的加锁规则是一致的


## 样例
```sql
CREATE TABLE `t` ( `id` int(11) NOT NULL, `c` int(11) DEFAULT NULL, `d` int(11) DEFAULT NULL, PRIMARY KEY (`id`), KEY `c` (`c`)) ENGINE=InnoDB;
insert into t values(0,0,0),(5,5,5),(10,10,10),(15,15,15),(20,20,20),(25,25,25);
```

### 不走索引的查询(for update)，会锁住整张表
```sql
begin;
select * from t where d = 5 for update;
```
因为d没有索引，所以mysql优化器只能选择遍历主键索引，走全表扫描
对主键索引的加锁规则如下(相当于锁着了整张表)
> (负无穷,0],(0,5],(5,10],(10,15],(15,20],(20,25],(25,num]

```sql
transaction2:
update t set c=16 where id =15; -- block

transaction3:
update t set c = 16 where c = 15; -- block 这里即使是走了c索引，但更新始终需要到主键索引上面去做更新，为了保证一致性，mysql肯定是需要锁主键索引上面的行的
```

### 非唯一索引等值查询，锁例子
非主键索引时，考虑锁范围的时候，应该需要将其对应的主键考虑的锁范围中
比如下面的：
索引c上的锁是：(5,10],(10,15)，这个是没有锁住5，15的
id索引：只锁住了10
但是insert id=15，c=15时却被阻塞，原因是，插入c索引值是，发现(c=15,id=15)这个被锁住了
c索引存放的值信息：(5,5),(10,10),(15,15)
```text
begin;
select * from t where c = 10 for update;

对于索引c，上锁信息(5,10],(10,15)
对应主键索引，上锁信息 10
    
猜测：是不是只有c不在(5,15)之间，且id不是10就可以了

但事实上：
但是为什么下面的失败了:
mysql> insert into t values (11,5,11);
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
mysql> 
mysql> insert into t values (12,15,12);
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction

----
这些却又成功了
mysql> insert into t values (11,4,11);
Query OK, 1 row affected (8.03 sec)

mysql> insert into t values (12,16,12);
Query OK, 1 row affected (0.00 sec)

------- 分析
c上的索引信息：c与id的关系信息
(0,0),(5,5),(10,10),(15,15),(20,20),(25,25)

c上的锁：(5,5),(10,10),(15,15)，

再次猜测：所以当c为5，15时，id不能在5-15之间

验证如下：解决上面的问题
mysql> insert into t values (16,15,16);
Query OK, 1 row affected (0.00 sec)

mysql> 
mysql> insert into t values (4,5,4);
Query OK, 1 row affected (0.01 sec)

总结一下:
当事务执行是：
begin;
select * from t where c = 10 for update;
锁着的区间是：
c: (5,10],(10,15)
id: 10
c-id存放信息：(5,5),(10,10),(15,15)

所以：针对于insert gap lock
c不在(5,15)之间，且id不是10就可以，
同时如果c是5或者15，id不能在[5，15]之间

针对于 update: 
(5,10],(10,15)
```

```sql
transaction 2:
update t set c=8 where c=15; -- block 相当于在c(5,10],(10,15中插入了8，被gap lock block 
```

```sql
transaction 3:
delete from t where c = 5; -- 将c=5的行删除掉
在执行
insert into t values (11,4,11);  -- 可以插入成功，说明mysql version 8.0.35 不会动态的更新lock
```

### 非唯一索引的范围查询
```sql
begin
select * from t where c >= 7 and c <= 15 for update;
c索引上的锁: (5,10],(10,15],(15,20]
id索引上的锁:10,15
c索引信息:(5,5),(10,10),(15,15),(20,20)
    
注意这里是范围查询，(15,20]不会退化成(15,20) 

transaction 1
update t set d = 6 where c = 20; -- block
insert into t values (6,6,6);  -- block
insert into t values (18,18,18); -- block
    
update t set d = 6 where c=5; -- ok
    
    
奇怪现象:
begin
select * from t where c >= 10 and c <= 15 for update;
c索引上的锁: (5,10],(10,15],(15,20)
```

```sql
begin
select * from t where c >=15 and c <= 20 lock in share mode;

c next-key lock (10,15],(10,20],(20,25)

transaction 1
insert into t values (14,14,14); --block
insert into t values (24,24,24); --block 
    
insert into t values (6,6,6); --ok
update t set d = 26 where id = 25; -- ok
```

### oder by c desc 为什么会影响加锁范围呢？
```sql
begin
select * from t where c >=15 and c <= 20 order by c desc  lock in share mode;

c next-key lock  (5,10],(10,15],(15,20],(20,25)
    
oder by c desc 让c上面的锁延长到了(5,10]

insert into t values (1,1,1); -- ok
insert into t values (27,27,27); -- ok
update t set d = 26 where id = 25; -- ok

insert into t values (6,6,6) ; -- block
insert into t values (24,24,24); -- block
```


### 唯一索引，等值查询例子

查询一个id存在的值
```sql
begin;
select * from t where id = 10 for update;

只会锁住id=10这行
transaction 1
insert into t values (9,9,9); -- insert ok
insert into t values (11,11,11); -- insert ok

transaction 2
update t set d=11 where c =10; -- block 因为需要拿到id=10的主键id的行，更新数据
```
查询一个id不存在的值
```sql
begin;
select * from t where id = 9 for update;

锁范围 (5,10) gap lock
    
transaction 1
update t set d = 6 where id =5; -- ok
update t set d = 11 where id =10; -- ok

insert into t values (7,7,7); -- block

```

### 唯一索引，范围查询例子
```sql
begin
select * from t where id >=7 and id < 12 for update; -- 从7开始的等值查询，但是表中没有7所以锁住了(5,10]

id next-key lock (5,10],(10,15)

transaction 1
insert into t values (6,6,6); -- block
insert into t values (14,14,14); -- block

insert into t values (4,4,4);  -- ok
insert into t values (16,16,16);  -- ok
update t set d = 16 where id =15;  -- ok
update t set d = 6 where id =5;  -- ok

```
```sql
begin
select * from t where id >= 10 and id < 14 for update; -- 从10开始的等值查询，退化成行锁

mysql version 8.0.35 , 这里不会锁住15这行了，退化成了间隙锁
id next-key lock 10,(10,15)
    
transaction 1
insert into t values (9,9,9); -- ok
update t set d = 16 where id = 15; -- ok mysql version 8.0.35

insert into t values (14,14,14); -- block

```
```sql
begin
select * from t where id >10 and id <=15 for update;

id next-key lock (10,15]
    
transaction 1
insert into t values (9,9,9); -- ok
insert into t values (16,16,16); -- ok
update t set d = 11 where id = 10; --ok
    
insert into t values (13,13,13); -- block
```

### limit 可以缩小加锁范围，建议删除时都添加上limit
