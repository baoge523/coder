## MYSQL

5.7的官方文档
https://dev.mysql.com/doc/refman/5.7/en/show.html



#### mysql中一条查询语句的执行流程

服务端

> 连接器: 处理客户端连接、用户名密码校验、权限控制
> 查询缓存: 8.0已去掉，弊大于利
> 分析器： 将sql语句做分析(抽象语法树)，检查sql里面的字段信息等
> 优化器： 比如在有多个索引的时候，选择使用更优的索引查询
> 执行器： 执行sql，与存储引擎打交道，比如：查询出一条数据，写入一条数据等

存储引擎 (只做存储数据和查询数据)

> mysaim: 表锁，没有事务、不支持外键
> innodb：行锁、支持事务、支持外键

#### 一条更新sql的执行流程

客户端  连接器  分析器  优化器  执行器 存储引擎

> 在执行器和存储引擎交互过程中，会将sql写入 redo log 和 bin log
>
> redo log有两阶段提交，来保证redo log 和binlog 的一致性
>
> redo log 是一个环形日志(写满时，需要先写入磁盘后，才能再次写入),用于mysql异常重启时恢复数据
>
> bin log 是一个日志文件，用于恢复最近的数据或者部署备库的，主从复制也需要它

更新流程： 写入 redo log(prepare) ,更新内存，写binlog，再写redo log(commit)->(提交事务的时候，再写redo log)

先写redo log的行为就是 预写日志操作



#### 事务隔离

事务就是要保证一组数据库操作，要么全部成功，要么全部失败。在 MySQL 中，事务支持是在引擎层实现的

隔离性与隔离级别

> ACID（Atomicity、Consistency、Isolation、Durability，即原子性、一致性、隔离性、持久性
>
> 多事务同时执行时，会出现脏读（dirty read）、不可重复读（non-repeatable read）、幻读（phantom read）的问题，为了解决这些问题，就有了“隔离级别”的概念。
>
> SQL 标准的事务隔离级别包括：读未提交（read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（serializable ）

事务隔离级别

- 读未提交是指，一个事务还没提交时，它做的变更就能被别的事务看到。
- 读提交是指(RC)，一个事务提交之后，它做的变更才会被其他事务看到。
- 可重复读是指(RR)，一个事务执行过程中看到的数据，总是跟这个事务在启动时看到的数据是一致的。当然在可重复读隔离级别下，未提交变更对其他事务也是不可见的。
- 串行化，顾名思义是对于同一行记录，“写”会加“写锁”，“读”会加“读锁”。当出现读写锁冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行

读提交和可重复读的比较

- 在“可重复读”隔离级别下，这个视图是在事务启动时创建的，整个事务存在期间都用这个视图。
- 在“读提交”隔离级别下，这个视图是在每个 SQL 语句开始执行的时候创建的

> “读未提交”隔离级别下直接返回记录上的最新值，没有视图概念；而“串行化”隔离级别下直接用加锁的方式来避免并行访问。

查看和设置隔离级别

> show variables like 'transaction_isolation';

事务隔离的实现

> 实际上每条记录在更新的时候都会同时记录一条回滚操作。记录上的最新值，通过回滚操作，都可以得到前一个状态的值

<img src="https://static001.geekbang.org/resource/image/d9/ee/d9c313809e5ac148fc39feff532f0fee.png?wh=1142*737" alt="img" style="zoom:50%;" />

> 同一条记录在系统中可以存在多个版本，就是数据库的多版本并发控制（MVCC）

回滚日志的删除

> 在不需要的时候才删除。也就是说，系统会判断，当没有事务再需要用到这些回滚日志时，回滚日志会被删除。
>
> 什么时候才不需要了呢？就是当系统里没有比这个回滚日志更早的 read-view 的时候

长事务的危害

> 长事务意味着系统里面会存在很老的事务视图。由于这些事务随时可能访问数据库里面的任何数据，所以这个事务提交之前，数据库里面它可能用到的回滚记录都必须保留，这就会导致大量占用存储空间。
>
> 长事务还占用锁资源，也可能拖垮整个库

事务的启动方式

> 显式启动事务语句， begin 或 start transaction。配套的提交语句是 commit，回滚语句是 rollback。
>
> set autocommit=0，这个命令会将这个线程的自动提交关掉。意味着如果你只执行一个 select 语句，这个事务就启动了，而且并不会自动提交。这个事务持续存在直到你主动执行 commit 或 rollback 语句，或者断开连接

会建议你总是使用 set autocommit=1, 通过显式语句的方式来启动事务

mysql如何查看长事务

> 在 information_schema 库的 innodb_trx 这个表中查询长事务

```sql
select * from information_schema.innodb_trx where TIME_TO_SEC(timediff(now(),trx_started))>60
```

#### 深入浅出索引(上)

索引的出现其实就是为了提高数据查询的效率，就像书的目录一样

##### 索引的常见模型

>索引的出现是为了提高查询效率，但是实现索引的方式却有很多种，所以这里也就引入了索引模型的概念。可以用于提高读写效率的数据结构很多，这里我先给你介绍三种常见、也比较简单的数据结构，它们分别是哈希表、有序数组和搜索树

- 哈希表是一种以键 - 值（key-value）存储数据的结构,等值查询快，不支持范围查询

> 哈希表这种结构适用于只有等值查询的场景

- 有序数组在等值查询和范围查询场景中的性能就都非常优秀（有序数组索引只适用于静态存储引擎）

> 用二分法就可以快速得到，这个时间复杂度是 O(log(N))。
>
> 如果仅仅看查询效率，有序数组就是最好的数据结构了。但是，在需要更新数据的时候就麻烦了，你往中间插入一个记录就必须得挪动后面所有的记录，成本太高

- 二叉搜索树 （数量大时，层级比较高，io次数多）

> 二叉搜索树的特点是：父节点左子树所有结点的值小于父节点的值，右子树所有结点的值大于父节点的值
>
> 当然为了维持 O(log(N)) 的查询复杂度，你就需要保持这棵树是平衡二叉树

- N 叉树 （一般一层和二层都在内存中，加快查询速度）

> N 叉树由于在读写上的性能优点，以及适配磁盘的访问模式，已经被广泛应用在数据库引擎中了
>
> 跳表、LSM 树等数据结构也被用于引擎设计中

**你心里要有个概念，数据库底层存储的核心就是基于这些数据模型的。每碰到一个新数据库，我们需要先关注它的数据模型，这样才能从理论上分析出这个数据库的适用场景**

在 MySQL 中，索引是在**存储引擎层**实现的，所以并没有统一的索引标准，即不同存储引擎的索引的工作方式并不一样。

##### InnoDB 的索引模型

> 在 InnoDB 中，表都是根据主键顺序以索引的形式存放的，这种存储方式的表称为索引组织表。又因为前面我们提到的，InnoDB 使用了 B+ 树索引模型，所以数据都是存储在 B+ 树中的。
>
> 每一个索引在 InnoDB 里面对应一棵 B+ 树

<img src="https://static001.geekbang.org/resource/image/dc/8d/dcda101051f28502bd5c4402b292e38d.png?wh=1142*856" alt="img" style="zoom:50%;" />

索引类型分为主键索引和非主键索引

> 主键索引的叶子节点存的是整行数据。在 InnoDB 里，主键索引也被称为聚簇索引（clustered index）
>
> 非主键索引的叶子节点内容是主键的值。在 InnoDB 里，非主键索引也被称为二级索引（secondary index）

基于主键索引和普通索引的查询有什么区别？

- 如果语句是 select * from T where ID=500，即主键查询方式，则只需要搜索 ID 这棵 B+ 树；
- 如果语句是 select * from T where k=5，即普通索引查询方式，则需要先搜索 k 索引树，得到 ID 的值为 500，再到 ID 索引树搜索一次。这个过程称为回表

**基于非主键索引的查询需要多扫描一棵索引树。因此，我们在应用中应该尽量使用主键查询。**(覆盖索引可以不用扫描主键索引，因为二级索引已经包含了要查询的内容)

##### 索引维护

B+ 树为了维护索引有序性，在插入新值的时候需要做必要的维护

> 更糟的情况是，如果 R5 所在的数据页已经满了，根据 B+ 树的算法，这时候需要申请一个新的数据页，然后挪动部分数据过去。这个过程称为页分裂
>
> 除了性能外，页分裂操作还影响数据页的利用率。原本放在一个页的数据，现在分到两个页中，整体空间利用率降低大约 50%。
>
> 当相邻两个页由于删除了数据，利用率很低之后，会将数据页做合并。合并的过程，可以认为是分裂过程的逆过程。

自增主键

> 自增主键的插入数据模式，正符合了我们前面提到的递增插入的场景。每次插入一条新记录，都是追加操作，都不涉及到挪动其他记录，也不会触发叶子节点的分裂。
>
> 而用业务逻辑的字段做主键，则往往不容易保证有序插入，这样写数据成本相对较高。

主键长度越小，普通索引的叶子节点就越小，普通索引占用的空间也就越小。

> 从性能和存储空间方面考量，自增主键往往是更合理的选择 （分布式下不合适）

#### 深入浅出索引(下)

如果我执行 select * from T where k between 3 and 5，需要执行几次树的搜索操作，会扫描多少行？

```mysql
mysql> create table T (
ID int primary key,
k int NOT NULL DEFAULT 0, 
s varchar(16) NOT NULL DEFAULT '',
index k(k))
engine=InnoDB;

insert into T values(100,1, 'aa'),(200,2,'bb'),(300,3,'cc'),(500,5,'ee'),(600,6,'ff'),(700,7,'gg');
```

<img src="https://static001.geekbang.org/resource/image/dc/8d/dcda101051f28502bd5c4402b292e38d.png?wh=1142*856" alt="img" style="zoom:50%;" />

看看这条 SQL 查询语句的执行流程

>1、在 k 索引树上找到 k=3 的记录，取得 ID = 300；
>
>2、再到 ID 索引树查到 ID=300 对应的 R3；
>
>3、在 k 索引树取下一个值 k=5，取得 ID=500；
>
>4、再回到 ID 索引树查到 ID=500 对应的 R4；
>
>5、在 k 索引树取下一个值 k=6，不满足条件，循环结束。

在这个过程中，回到主键索引树搜索的过程，我们称为回表。可以看到，这个查询过程读了 k 索引树的 3 条记录（步骤 1、3 和 5），回表了两次（步骤 2 和 4）

##### 覆盖索引

```mysql
select ID from T where k between 3 and 5
```

> 这时只需要查 ID 的值，而 ID 的值已经在 k 索引树上了，因此可以直接提供查询结果，不需要回表。也就是说，在这个查询里面，索引 k 已经“覆盖了”我们的查询需求，我们称为覆盖索引。

**由于覆盖索引可以减少树的搜索次数，显著提升查询性能，所以使用覆盖索引是一个常用的性能优化手段**

>在引擎内部使用覆盖索引在索引 k 上其实读了三个记录，R3~R5（对应的索引 k 上的记录项）
>
>但是对于 MySQL 的 Server 层来说，它就是找引擎拿到了两条记录，因此 MySQL 认为扫描行数是 2。

##### 最左前缀原则

> B+ 树这种索引结构，可以利用索引的“最左前缀”，来定位记录
>
> 这个最左前缀可以是联合索引的最左 N 个字段，也可以是字符串索引的最左 M 个字符

在建立联合索引的时候，如何安排索引内的字段顺序。

>索引的复用能力
>
>第一原则是，如果通过调整顺序，可以少维护一个索引，那么这个顺序往往就是需要优先考虑采用的
>
>第一原则就是空间: name 字段是比 age 字段大的 ，那我就建议你创建一个（name,age) 的联合索引和一个 (age) 的单字段索引。

##### 索引下推

联合索引: (name,age)

```mysql
mysql> select * from tuser where name like '张%' and age=10 and ismale=1;
```

> 你已经知道了前缀索引规则，所以这个语句在搜索索引树的时候，只能用 “张”

在 MySQL 5.6 之前，只能一个个回表。到主键索引上找出数据行，再对比字段值

<img src="https://static001.geekbang.org/resource/image/b3/ac/b32aa8b1f75611e0759e52f5915539ac.jpg?wh=1142*833" alt="img" style="zoom:50%;" />

MySQL 5.6 引入的索引下推优化（index condition pushdown)， 可以在索引遍历过程中，对索引中包含的字段先做判断，直接过滤掉不满足条件的记录，减少回表次数

<img src="https://static001.geekbang.org/resource/image/76/1b/76e385f3df5a694cc4238c7b65acfe1b.jpg?wh=1142*856" alt="img" style="zoom:50%;" />

区别是，InnoDB 在 (name,age) 索引内部就判断了 age 是否等于 10，对于不等于 10 的记录，直接判断并跳过。

重建索引

> 索引可能因为删除，或者页分裂等原因，导致数据页有空洞，重建索引的过程会创建一个新的索引，把数据按顺序插入，这样页面的利用率最高，也就是索引更紧凑、更省空间
>
> 重建索引 k 的做法是合理的，可以达到省空间的目的。

```mysql
alter table T drop primary key;
alter table T add primary key(id);
```

> 不论是删除主键还是创建主键，都会将整个表重建。所以连着执行这两个语句的话，第一个语句就白做了。
>
> 这两个语句，你可以用这个语句代替 ： alter table T engine=InnoDB

#### 全局锁和表锁

> 数据库锁设计的初衷是处理并发问题。作为多用户共享的资源，当出现并发访问的时候，数据库需要合理地控制资源的访问规则。而锁就是用来实现这些访问规则的重要数据结构。

根据加锁的范围，MySQL 里面的锁大致可以分成全局锁、表级锁和行锁三类

##### 全局锁

全局锁就是对整个数据库实例加锁

>MySQL 提供了一个加全局读锁的方法，命令是 Flush tables with read lock (FTWRL)。
>
>当你需要让整个库处于只读状态的时候，可以使用这个命令，
>
>之后其他线程的以下语句会被阻塞：数据更新语句（数据的增删改）、数据定义语句（包括建表、修改表结构、索引等）和更新类事务的提交语句

**全局锁的典型使用场景是，做全库逻辑备份**。也就是把整库每个表都 select 出来存成文本。

通过 FTWRL 确保不会有其他线程对数据库做更新，然后对整个库做备份 (在备份过程中整个库完全处于只读状态)

但是让整库都只读，听上去就很危险：

- 如果你在主库上备份，那么在备份期间都不能执行更新，业务基本上就得停摆；
- 如果你在从库上备份，那么备份期间从库不能执行主库同步过来的 binlog，会导致主从延迟;

如果直接备份数据库，在操作过程中可能会导致**数据一致性问题**：

解决方案：

1、可重复读隔离级别下开启一个事务

>官方自带的逻辑备份工具是 mysqldump。当 mysqldump 使用参数**–single-transaction** 的时候，导数据之前就会启动一个事务，来确保拿到一致性视图。而由于 MVCC 的支持，这个过程中数据是可以正常更新的。
>
>一致性读是好，但前提是引擎要支持这个隔离级别
>
>对于 MyISAM 这种不支持事务的引擎，如果备份过程中有更新，总是只能取到最新的数据，那么就破坏了备份的一致性

2、Flush tables with read lock (FTWRL) --- 系统处理只读状态

3、既然要全库只读，为什么不使用 set global readonly=true 的方式呢  --- 系统处于只读状态

>readonly 方式也可以让全库进入只读状态，但我还是会建议你用 FTWRL 方式
>
>在有些系统中，readonly 的值会被用来做其他逻辑，比如用来判断一个库是主库还是备库。因此，修改 global 变量的方式影响面更大，我不建议你使用
>
>在异常处理机制上有差异。如果执行 FTWRL 命令之后由于客户端发生异常断开，那么 MySQL 会自动释放这个全局锁，整个库回到可以正常更新的状态。而将整个库设置为 readonly 之后，如果客户端发生异常，则数据库就会一直保持 readonly 状态，这样会导致整个库长时间处于不可写状态，风险较高。

##### 表级锁

[mysql表锁官网地址](https://dev.mysql.com/doc/refman/5.7/en/lock-tables.html)

> MySQL 里面表级别的锁有两种：一种是表锁，一种是元数据锁（meta data lock，MDL)。

表锁的语法是 **lock tables … read/write**。可以用 **unlock tables** 主动释放锁，也可以在**客户端断开的时候自动释放**。需要注意，lock tables 语法除了会限制别的线程的读写外，也限定了本线程接下来的操作对象。

```mysql
LOCK TABLES
    tbl_name [[AS] alias] lock_type
    [, tbl_name [[AS] alias] lock_type] ...

lock_type: {
    READ [LOCAL]
  | [LOW_PRIORITY] WRITE
}

UNLOCK TABLES
-- 比如
lock tables `user` read;  -- 表读锁  其他查询user事务不会被阻塞
lock tables `user` write;  -- 表写锁 其他查询user事务都会被阻塞

lock tables `user` as `user_aa` write;  -- 锁住的是别名，使用时，必须通过别名去访问表，不然会报错

unlock tables -- 释放锁
```



>举个例子, 如果在某个线程 A 中执行 lock tables t1 read, t2 write; 
>
>这个语句，则其他线程写 t1、读写 t2 的语句都会被阻塞。
>
>同时，线程 A 在执行 unlock tables 之前，也只能执行读 t1、读写 t2 的操作。连写 t1 都不允许，自然也不能访问其他表

另一类表级的锁是 MDL（metadata lock) --- MySQL 5.5 版本中引入了 MDL

> MDL 不需要显式使用，在访问一个表的时候会被自动加上。
>
> MDL 的作用是，保证读写的正确性。
>
> 你可以想象一下，如果一个查询正在遍历一个表中的数据，而执行期间另一个线程对这个表结构做变更，删了一列，那么查询线程拿到的结果跟表结构对不上，肯定是不行的。

加锁时机

> 当对一个表做增删改查操作的时候，加 MDL 读锁；
>
> 当要对表做结构变更操作的时候，加 MDL 写锁

- 读锁之间不互斥，因此你可以有多个线程同时对一张表增删改查。
- 读写锁之间、写锁之间是互斥的，用来保证变更表结构操作的安全性。因此，如果有两个线程要同时给一个表加字段，其中一个要等另一个执行完才能开始执行。

虽然 MDL 锁是系统默认会加的，但却是你不能忽略的一个机制

> MDL的读锁需要在事务提交的时候才会释放
>
> MDL写锁被阻塞后，该操作会阻塞其操作后的所有获取MDL读锁、写锁的操作 --- 极端情况下导致整个库挂了

如何安全地给小表加字段

> 首先我们要解决长事务，事务不提交，就会一直占着 MDL 锁。在 MySQL 的 **information_schema 库的 innodb_trx 表中**，你可以查到当前执行中的事务。如果你要做 DDL 变更的表刚好有长事务在执行，要考虑先暂停 DDL，或者 kill 掉这个长事务

变更的表是一个热点表

> 在 alter table 语句里面设定等待时间，如果在这个指定的等待时间里面能够拿到 MDL 写锁最好，拿不到也不要阻塞后面的业务语句，先放弃。之后开发人员或者 DBA 再通过重试命令重复这个过程。
>
> MariaDB 已经合并了 AliSQL 的这个功能，所以这两个开源分支目前都支持 DDL NOWAIT/WAIT n 这个语法

```mysql
ALTER TABLE tbl_name NOWAIT add column ...
ALTER TABLE tbl_name WAIT N add column ... 
```

思考题：

```mysql
CREATE TABLE `geek` (
  `a` int(11) NOT NULL,
  `b` int(11) NOT NULL,
  `c` int(11) NOT NULL,
  `d` int(11) NOT NULL,
  PRIMARY KEY (`a`,`b`),
  KEY `c` (`c`),
  KEY `ca` (`c`,`a`),
  KEY `cb` (`c`,`b`)
) ENGINE=InnoDB;

-- 为了满足下面的查询，上面的索引是否合理
select * from geek where c=N order by a limit 1;
select * from geek where c=N order by b limit 1;
```

主键索引(a,b) 就等于 order by a,b 了， 先按a排序，再按b排序，c无序

二级索引(c) 

> 等同于  二级索引 + 主键索引
>
>  c 、 a 、 b (主键索引 a b) 

二级索引(c,a)  和二级索引c一样的效果

> 等同于  二级索引 + 主键索引
>
> c 、 a 、 b (主键索引只有b)

二级索引(c,b)  

> 等同于  二级索引 + 主键索引
>
> c 、 b 、 a (主键索引只有a)

所以二级索引(c,a) 是不需要的，浪费磁盘空间

#### 行锁功过

MySQL 的行锁是在引擎层由各个引擎自己实现的

> 减少锁冲突来提升业务并发度

##### 两阶段锁

<img src="https://static001.geekbang.org/resource/image/51/10/51f501f718e420244b0a2ec2ce858710.jpg?wh=1142*856" alt="img" style="zoom:50%;" />

> 实际上事务 B 的 update 语句会被阻塞，直到事务 A 执行 commit 之后，事务 B 才能继续执行。
>
> 事务 A 持有的两个记录的行锁，都是在 commit 的时候才释放的

如果你的事务中需要锁多个行，要把最可能造成锁冲突、最可能影响并发度的锁尽量往后放

> 最大程度地减少了事务之间的锁等待，提升了并发度

##### 死锁和死锁检测

CPU 消耗接近 100%，但整个数据库每秒就执行不到 100 个事务。这是什么原因呢？ --- 死锁和死锁检测

> 当并发系统中不同线程出现循环资源依赖，涉及的线程都在等待别的线程释放资源时，就会导致这几个线程都进入无限等待的状态，称为死锁

<img src="https://static001.geekbang.org/resource/image/4d/52/4d0eeec7b136371b79248a0aed005a52.jpg?wh=1142*856" alt="img" style="zoom:50%;" />

>  事务 A 和事务 B 在互相等待对方的资源释放，就是进入了死锁状态

当出现死锁以后，有两种策略：

- 一种策略是，直接进入等待，直到超时。这个超时时间可以通过参数 **innodb_lock_wait_timeout** 来设置。

> 在 InnoDB 中，innodb_lock_wait_timeout 的默认值是 50s，意味着如果采用第一个策略，当出现死锁以后，第一个被锁住的线程要过 50s 才会超时退出，然后其他线程才有可能继续执行
>
> 我们又不可能直接把这个时间设置成一个很小的值，比如 1s。这样当出现死锁的时候，确实很快就可以解开，但如果不是死锁，而是简单的锁等待呢？所以，超时时间设置太短的话，会出现很多误伤。

- 另一种策略是，发起死锁检测，发现死锁后，主动回滚死锁链条中的某一个事务，让其他事务得以继续执行。将参数 **innodb_deadlock_detect 设置为 on**，表示开启这个逻辑。

> nnodb_deadlock_detect 的默认值本身就是 on
>
> 每当一个事务被锁的时候，就要看看它所依赖的线程有没有被别人锁住，如此循环，最后判断是否出现了循环等待，也就是死锁
>
> 每个新来的被堵住的线程，都要判断会不会由于自己的加入导致了死锁，这是一个时间复杂度是 O(n) 的操作
>
> 假设有 1000 个并发线程要同时更新同一行，那么死锁检测操作就是 100 万这个量级的。虽然最终检测的结果是没有死锁，但是这期间要消耗大量的 CPU 资源

```mysql
show variables like 'innodb_deadlock_detect';
```

出现死锁一般建议使用死锁检测策略，但是这样会存在当有多个线程并发修改同一行时，会产生大量的检查，导致消耗大量的CPU资源

怎么解决由这种热点行更新导致的性能问题呢？

- 一种头痛医头的方法，就是如果你能确保这个业务一定不会出现死锁，可以临时把死锁检测关掉

> 但是这种操作本身带有一定的风险，因为业务设计的时候一般不会把死锁当做一个严重错误，毕竟出现死锁了，就回滚，然后通过业务重试一般就没问题了，这是业务无损的。而关掉死锁检测意味着可能会出现大量的超时，这是业务有损的。

- 另一个思路是控制并发度

> 这个并发控制要做在数据库服务端
>
> 将一行改成逻辑上的多行来减少锁冲突

**减少死锁的主要方向，就是控制访问相同资源的并发事务量**

##### 思考题:

当备库用–single-transaction 做逻辑备份的时候，如果从主库的 binlog 传来一个 DDL 语句会怎么样？

> 假设这个 DDL 是针对表 t1 的， 这里我把备份过程中几个关键的语句列出来：

```mysql
Q1:SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;   -- 设置隔离级别
Q2:START TRANSACTION  WITH CONSISTENT SNAPSHOT； -- 开启事务 立马就创建视图
/* other tables */
Q3:SAVEPOINT sp;   -- 设置一个保存点
/* 时刻 1 */
Q4:show create table `t1`; -- 拿到表结构
/* 时刻 2 */
Q5:SELECT * FROM `t1`;  -- 正式导出数据
/* 时刻 3 */
Q6:ROLLBACK TO SAVEPOINT sp; -- 回到Q3保存点
/* 时刻 4 */
/* other tables */
```

> DDL 从主库传过来的时间按照效果不同，我打了四个时刻

- 如果在 Q4 语句执行之前到达，现象：没有影响，备份拿到的是 DDL 后的表结构。
- 如果在“时刻 2”到达，则表结构被改过，Q5 执行的时候，报 Table definition has changed, please retry transaction，现象：mysqldump 终止；
- 如果在“时刻 2”和“时刻 3”之间到达，mysqldump 占着 t1 的 MDL 读锁，binlog 被阻塞，现象：主从延迟，直到 Q6 执行完成。
- 从“时刻 4”开始，mysqldump 释放了 MDL 读锁，现象：没有影响，备份拿到的是 DDL 前的表结构。

##### 思考题：

innodb行级锁是通过锁索引记录实现的。如果update的列没建索引，即使只update一条记录也会锁定整张表吗？

> 答案： 是的

测试验证

```mysql
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(64) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `love` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

insert into user (1,'aa',18,'aa'),(2,'bb',19,'bb'),(3,'cc',20,'cc');

-- 用例1
update user set age = 20 where love ='aa'; -- 事务A，先执行
update user set age = 20 where love ='cc'; -- 事务B，后执行，阻塞
-- love上没有索引导致事务A更新时，行锁升级成了表锁，从而导致事务B阻塞

-- 用例2
update user set age = 20 where name = 'aa' and love ='aa'; -- 事务A，先执行
update user set age = 20 where name = 'cc' and love ='cc'; -- 事务B，后执行，不阻塞
-- 事务A、B互不影响：name有索引，love没有索引

-- 用例3
update user set age = 20 where name = 'aa' ; -- 事务A，先执行
update user set age = 20 where id = 1; -- 事务B，后执行，阻塞
-- name有索引，id有索引，但是两条语句操作同一行，事务B被阻塞

-- 疑问 ？
-- 用例3是锁的哪个一个索引呢？
```

#### 事务到底是隔离的还是不隔离的

begin/start transaction 命令并不是一个事务的起点，在执行到它们之后的第一个操作 InnoDB 表的语句，事务才真正启动。

如果你想要马上启动一个事务，可以使用 start transaction with consistent snapshot 这个命令

> 第一种启动方式，一致性视图是在执行第一个快照读语句时创建的；
>
> 第二种启动方式，一致性视图是在执行 start transaction with consistent snapshot 时创建的。

例子：

```mysql
CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `k` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
insert into t(id, k) values(1,1),(2,2);
```

<img src="https://static001.geekbang.org/resource/image/82/d6/823acf76e53c0bdba7beab45e72e90d6.png?wh=910*509" alt="img" style="zoom: 67%;" />



事务 B 查到的 k 的值是 3，而事务 A 查到的 k 的值是 1

##### 在 MySQL 里，有两个“视图”的概念：

- 一个是 view。它是一个用查询语句定义的虚拟表，在调用的时候执行查询语句并生成结果。创建视图的语法是 create view … ，而它的查询方法与表一样。
- 另一个是 InnoDB 在实现 MVCC 时用到的一致性读视图，即 consistent read view，用于支持 RC（Read Committed，读提交）和 RR（Repeatable Read，可重复读）隔离级别的实现

##### “快照”在 MVCC 里是怎么工作的？

> 在可重复读隔离级别下，事务在启动的时候就“拍了个快照”。注意，这个快照是基于整库的
>
> InnoDB 里面每个事务有一个唯一的事务 ID，叫作 transaction id。它是在事务开始的时候向 InnoDB 的事务系统申请的，是按申请顺序严格递增的

数据表中的一行记录，其实可能有多个版本 (row)，每个版本有自己的 row trx_id

<img src="https://static001.geekbang.org/resource/image/68/ed/68d08d277a6f7926a41cc5541d3dfced.png?wh=1142*856" alt="img" style="zoom: 67%;" />

语句更新会生成 undo log（回滚日志）吗？那么，undo log 在哪呢？

> 图中的三个虚线箭头，就是 undo log；
>
> 而 V1、V2、V3 并不是物理上真实存在的，而是每次需要的时候根据当前版本和 undo log 计算出来的。比如，需要 V2 的时候，就是通过 V4 依次执行 U3、U2 算出来。

InnoDB 为每个事务构造了一个数组，用来保存这个事务启动瞬间，当前正在“活跃”的所有事务 ID。“活跃”指的就是，启动了但还没提交。

> 数组里面事务 ID 的最小值记为低水位，
>
> 当前系统里面已经创建过的事务 ID 的最大值加 1 记为高水位

<img src="https://static001.geekbang.org/resource/image/88/5e/882114aaf55861832b4270d44507695e.png?wh=1142*856" alt="img" style="zoom:50%;" />

对于当前事务的启动瞬间来说，一个数据版本的 row trx_id，有以下几种可能:

- 如果落在绿色部分，表示这个版本是已提交的事务或者是当前事务自己生成的，这个数据是可见的；
- 如果落在红色部分，表示这个版本是由将来启动的事务生成的，是肯定不可见的；
- 如果落在黄色部分，那就包括两种情况

> a.  若 row trx_id 在数组中，表示这个版本是由还没提交的事务生成的，不可见；
>
> b.  若 row trx_id 不在数组中，表示这个版本是已经提交了的事务生成的，可见

InnoDB 利用了“所有数据都有多个版本”的这个特性，实现了“秒级创建快照”的能力。

读数据都是从当前版本读起的。所以 查询语句的读数据流程是这样的：事务数组[90,100]  --> 低水位 90  高水位101

> 找到 (1,3) 的时候，判断出 row trx_id=101，比高水位大(这里准确来说是大于等于高水位)，处于红色区域，不可见；
>
> 接着，找到上一个历史版本，一看 row trx_id=102，比高水位大，处于红色区域，不可见；
>
> 再往前找，终于找到了（1,1)，它的 row trx_id=90，比低水位小，处于绿色区域，可见

简单版本的理解：（RR 和 RC都可用）

- 版本未提交，不可见；
- 版本已提交，但是是在视图创建后提交的，不可见；
- 版本已提交，而且是在视图创建前提交的，可见

##### 更新逻辑

更新数据的时候，就不能再在历史版本上更新了，否则其他事务的更新就丢失了

更新数据都是先读后写的，而这个读，只能读当前的值(最新已提交的值)，称为“当前读”（current read）。

当前读的方式：

> update 语句
>
> select 语句加锁

```mysql
select k from t where id=1 lock in share mode;  -- 共享锁、读锁
select k from t where id=1 for update;  -- 排他锁、写锁
```

用例：

<img src="https://static001.geekbang.org/resource/image/cd/6e/cda2a0d7decb61e59dddc83ac51efb6e.png?wh=906*565" alt="img" style="zoom:67%;" />

事务B会被阻塞，直到事务C`提交后，事务B才会执行

> 两阶段锁协议
>
> 事务C`持有该行的写锁并没有释放，事务B更新时需要先获取该行写锁，于是需要进行锁等待，必须等到事务 C’释放这个锁，才能继续它的当前读。

事务的可重复读的能力是怎么实现的

> 可重复读的核心就是一致性读（consistent read）；
>
> 而事务更新数据的时候，只能用当前读。如果当前的记录的行锁被其他事务占用的话，就需要进入锁等待。

读提交的逻辑和可重复读的逻辑类似，它们最主要的区别是：

- 在可重复读隔离级别下，只需要在事务开始的时候创建一致性视图，之后事务里的其他查询都共用这个一致性视图；
- 在读提交隔离级别下，每一个语句执行前都会重新算出一个新的视图

视图的计算方式：(视图计算出来后，在判断数据对一个事务是否可见)

> 用一个数组保存当前活跃的事务ID
>
> 数组中最小的事务ID为低水位
>
> 系统中最大的生成的事务ID + 1为高水位

##### 思考题

怎么删除表的前 10000 行

- 第一种，直接执行 delete from T limit 10000;
- 第二种，在一个连接中循环执行 20 次 delete from T limit 500;
- 第三种，在 20 个连接中同时执行 delete from T limit 500。

**第二种方式是相对较好的。**

> 第一种方式（即：直接执行 delete from T limit 10000）里面，单个语句占用时间长，锁的时间也比较长；而且大事务还会导致主从延迟。
>
> 第三种方式（即：在 20 个连接中同时执行 delete from T limit 500），会人为造成锁冲突

##### 思考题

现在，我要把所有“字段 c 和 id 值相等的行”的 c 值清零，但是却发现了一个“诡异”的、改不掉的情况。请你构造出这种情况，并说明其原理。

```mysql
CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
insert into t(id, c) values(1,1),(2,2),(3,3),(4,4);
```

![img](https://static001.geekbang.org/resource/image/9b/0b/9b8fe7cf88c9ba40dc12e93e36c3060b.png?wh=486*661)

原因：

> 因为当前的隔离级别是RR,所以在事务开启后，执行了第一次查询后，视图就已经被创建，期间的视图不会改变
>
> 当执行 update t set c =0 where id = c; 之前弄一个事务将所有的数据都改了，导致id = c 条件不满足，从 query ok,0 rows affected 可以看出，该操作没有满足条件的行
>
> 因为 update 是一个当前读的操作，所以会拿最新的数据，执行时就发现没有行满足改条件了
>
> 最后的select * from t 是因为 update 没有更新行，视图不变

#### 普通索引和唯一索引

##### 查询过程

```mysql
select id from T where k=5
```

对于普通索引来说，查找到满足条件的第一个记录 (5,500) 后，需要查找下一个记录，直到碰到第一个不满足 k=5 条件的记录。

对于唯一索引来说，由于索引定义了唯一性，查找到第一个满足条件的记录后，就会停止继续检索。

那么，这个不同带来的性能差距会有多少呢？答案是，微乎其微。

> 你知道的，InnoDB 的数据是按数据页为单位来读写的。也就是说，当需要读一条记录的时候，并不是将这个记录本身从磁盘读出来，而是以页为单位，将其整体读入内存。在 InnoDB 中，每个数据页的大小默认是 16KB

> 因为引擎是按页读写的，所以说，当找到 k=5 的记录的时候，它所在的数据页就都在内存里了。那么，对于普通索引来说，要多做的那一次“查找和判断下一条记录”的操作，就只需要一次指针寻找和一次计算。

> 当然，如果 k=5 这个记录刚好是这个数据页的最后一个记录，那么要取下一个记录，必须读取下一个数据页，这个操作会稍微复杂一些。但是，我们之前计算过，对于整型字段，一个数据页可以放近千个 key，因此出现这种情况的概率会很低。所以，我们计算平均性能差异时，仍可以认为这个操作成本对于现在的 CPU 来说可以忽略不计。

##### 更新过程

change buffer （可以持久化的数据）

> 当需要更新一个数据页时，如果数据页在内存中就直接更新，而如果这个数据页还没有在内存中的话，在不影响数据一致性的前提下，InnoDB 会将这些更新操作缓存在 change buffer 中，这样就不需要从磁盘中读入这个数据页了。

> 在下次查询需要访问这个数据页的时候，将数据页读入内存，然后执行 change buffer 中与这个页有关的操作。通过这种方式就能保证这个数据逻辑的正确性。

需要说明的是，虽然名字叫作 change buffer，实际上它是可以持久化的数据。也就是说，change buffer 在内存中有拷贝，也会被写入到磁盘上。

将 change buffer 中的操作应用到原数据页，得到最新结果的过程称为 merge；

merge的触发方式：

- 访问这个数据页会触发 merge 
- 系统有后台线程会定期 merge
- 在数据库正常关闭（shutdown）的过程中，也会执行 merge 操作。



显然，如果能够将更新操作先记录在 change buffer，减少随机读磁盘，语句的执行速度会得到明显的提升。

而且，数据读入内存是需要占用 buffer pool 的，所以这种方式还能够避免占用内存，提高内存利用率。

##### 什么条件下可以使用 change buffer  (只有普通索引可以使用)

对于唯一索引来说，所有的更新操作都要先判断这个操作是否违反唯一性约束

> 必须要将数据页读入内存才能判断。如果都已经读入到内存了，那直接更新内存会更快，就没必要使用 change buffer 了
>
> 唯一索引的更新就不能使用 change buffer，实际上也只有普通索引可以使用。

changge buffer的设置：

> change buffer 用的是 buffer pool 里的内存，因此不能无限增大。change buffer 的大小，可以通过参数 innodb_change_buffer_max_size 来动态设置。这个参数设置为 50 的时候，表示 change buffer 的大小最多只能占用 buffer pool 的 50%。

##### 如果要在这张表中插入一个新记录 (4,400) 的话，InnoDB 的处理流程是怎样的。

这个记录要更新的目标页在内存中

- 对于唯一索引来说，找到 3 和 5 之间的位置，判断到没有冲突，插入这个值，语句执行结束；
- 对于普通索引来说，找到 3 和 5 之间的位置，插入这个值，语句执行结束

这个记录要更新的目标页不在内存中

- 对于唯一索引来说，需要将数据页读入内存，判断到没有冲突，插入这个值，语句执行结束；
- 对于普通索引来说，则是将更新记录在 change buffer，语句执行就结束了

将数据从磁盘读入内存涉及随机 IO 的访问，是数据库里面成本最高的操作之一。change buffer 因为减少了随机磁盘访问，所以对更新性能的提升是会很明显的。

##### change buffer 的使用场景

change buffer 只限于用在普通索引的场景下，而不适用于唯一索引

> 因为 merge 的时候是真正进行数据更新的时刻，而 change buffer 的主要目的就是将记录的变更动作缓存下来，所以在一个数据页做 merge 之前，change buffer 记录的变更越多（也就是这个页面上要更新的次数越多），收益就越大

页面在写完以后马上被访问到的概率比较小，此时 change buffer 的使用效果最好。这种业务模型常见的就是**账单类、日志类**的系统。

反过来，假设一个业务的更新模式是写入之后马上会做查询，那么即使满足了条件，将更新先记录在 change buffer，但之后由于马上要访问这个数据页，会立即触发 merge 过程。

> 这样随机访问 IO 的次数不会减少，反而增加了 change buffer 的维护代价。所以，对于这种业务模式来说，change buffer 反而起到了副作用。

##### 索引选择和实践

其实，这两类索引在查询能力上是没差别的，主要考虑的是对更新性能的影响。所以，我建议你尽量选择普通索引。

在实际使用中，你会发现，普通索引和 change buffer 的配合使用，对于数据量大的表的更新优化还是很明显的

##### change buffer 和 redo log

WAL 提升性能的核心机制，是尽量减少随机读写

> redo log 主要节省的是随机写磁盘的 IO 消耗（转成顺序写）

>  change buffer 主要节省的则是随机读磁盘的 IO 消耗

change buffer 和redo log 的更新流程:

```mysql
insert into t(id,k) values(id1,k1),(id2,k2); -- k 是普通索引
```

> k1 所在的数据页在内存 (InnoDB buffer pool) 中，k2 所在的数据页不在内存中

<img src="https://static001.geekbang.org/resource/image/98/a3/980a2b786f0ea7adabef2e64fb4c4ca3.png?wh=1142*856" alt="img" style="zoom:50%;" />

> 内存、redo log（ib_log_fileX）、 数据表空间（t.ibd）、系统表空间（ibdata1）。
>
> buffer pool 是 mysql innodb引擎使用的内存总空间，里面有change buffer 、page信息、索引一二层信息等
>
> redo log 日志文件
>
> 你会看到，执行这条更新语句的成本很低，就是写了两处内存，然后写了一处磁盘（两次操作合在一起写了一次磁盘），而且还是顺序写的

change buffer 和redo log 的查询流程:

```mysql
select * from t where k in (k1, k2)
```

> 如果读语句发生在更新语句后不久，内存中的数据都还在，那么此时的这两个读操作就与系统表空间（ibdata1）和 redo log（ib_log_fileX）无关了

<img src="https://static001.geekbang.org/resource/image/6d/8e/6dc743577af1dbcbb8550bddbfc5f98e.png?wh=1142*856" alt="img" style="zoom:50%;" />

> 读 Page 1 的时候，直接从内存返回
>
> 要读 Page 2 的时候，需要把 Page 2 从磁盘读入内存中，然后应用 change buffer 里面的操作日志，生成一个正确的版本并返回结果。（需要先marge）



#### MySQL为什么有时候会选错索引？

```mysql
set long_query_time=0;
select * from t where a between 10000 and 20000; /*Q1*/
select * from t force index(a) where a between 10000 and 20000;/*Q2*/
```

> 第一句，是将慢查询日志的阈值设置为 0，表示这个线程接下来的语句都会被记录入慢查询日志中；

> 第二句，Q1 是 session B 原来的查询；

> 第三句，Q2 是加了 force index(a) 来和 session B 原来的查询语句执行情况对比

![img](https://static001.geekbang.org/resource/image/7c/f6/7c58b9c71853b8bba1a8ad5e926de1f6.png?wh=1221*325)

Q1 扫描了 10 万行，显然是走了全表扫描，执行时间是 40 毫秒。

Q2 扫描了 10001 行，执行了 21 毫秒。也就是说，我们在没有使用 force index 的时候，MySQL 用错了索引，导致了更长的执行时间。

**这个例子对应的是我们平常不断地删除历史数据和新增数据的场景**



##### 优化器的逻辑

选择索引是优化器的工作。

而优化器选择索引的目的，是找到一个最优的执行方案，并用最小的代价去执行语句

> 扫描行数是影响执行代价的因素之一。扫描的行数越少，意味着访问磁盘数据的次数越少，消耗的 CPU 资源越少
>
> 扫描行数并不是唯一的判断标准，优化器还会结合**是否使用临时表、是否排序**等因素进行综合判断。包含：回表获取信息也会考虑进去

扫描行数是怎么判断的?

> MySQL 在真正开始执行语句之前，并不能精确地知道满足这个条件的记录有多少条，而只能根据统计信息来估算记录数。
>
> 索引的“区分度" :  一个索引上不同的值越多，这个索引的区分度就越好 ,我们称之为“基数”（cardinality）
>
> 也就是说，这个基数越大，索引的区分度越好。

MySQL 是怎样得到索引的基数的呢？

> 采样统计

analyze table t 命令，可以用来重新统计索引信息

> 所以在实践中，如果你发现 explain 的结果预估的 rows 值跟实际情况差距比较大，可以采用这个方法来处理

##### 索引选择异常和处理

- 一种方法是，像我们第一个例子一样，采用 force index 强行选择一个索引

> MySQL 会根据词法解析的结果分析出可能可以使用的索引作为候选项，然后在候选列表中依次判断每个索引需要扫描多少行。如果 force index 指定的索引在候选索引列表中，就直接选择这个索引，不再评估其他索引的执行代价。

- 第二种方法就是，我们可以考虑修改语句，引导 MySQL 使用我们期望的索引

- 第三种方法是，在有些场景下，我们可以新建一个更合适的索引，来提供给优化器做选择，或删掉误用的索引

思考题：

如果某次写入使用了 change buffer 机制，之后主机异常重启，是否会丢失 change buffer 和数据。

> 这个问题的答案是不会丢失。虽然是只更新内存，但是在事务提交的时候，我们把 change buffer 的操作也记录到 redo log 里了，所以崩溃恢复的时候，change buffer 也能找回来

思考题:

merge 的过程是否会把数据直接写回磁盘

> merge 的执行流程是这样的：

> 从磁盘读入数据页到内存（老版本的数据页）；
>
> 从 change buffer 里找出这个数据页的 change buffer 记录 (可能有多个），依次应用，得到新版数据页；
>
> 写 redo log。这个 redo log 包含了数据的变更和 change buffer 的变更。
>
> 到这里 merge 过程就结束了。这时候，数据页和内存中 change buffer 对应的磁盘位置都还没有修改，属于脏页，之后各自刷回自己的物理数据，就是另外一个过程了

#### 怎么给字符串字段加索引？

前缀索引 和 整个字符串的索引

```mysql
alter table SUser add index index1(email);
alter table SUser add index index2(email(6));

select id,name,email from SUser where email='zhangssxyz@xxx.com';
```

如果使用的是 index1（即 email 整个字符串的索引结构），执行顺序是这样的：

> 从 index1 索引树找到满足索引值是’zhangssxyz@xxx.com’的这条记录，取得 ID2 的值；
>
> 到主键上查到主键值是 ID2 的行，判断 email 的值是正确的，将这行记录加入结果集；
>
> 取 index1 索引树上刚刚查到的位置的下一条记录，发现已经不满足 email='zhangssxyz@xxx.com’的条件了，循环结束。
>
> 这个过程中，只需要回主键索引取一次数据，所以系统认为只扫描了一行。



如果使用的是 index2（即 email(6) 索引结构），执行顺序是这样的：

> 从 index2 索引树找到满足索引值是’zhangs’的记录，找到的第一个是 ID1；
>
> 到主键上查到主键值是 ID1 的行，判断出 email 的值不是’zhangssxyz@xxx.com’，这行记录丢弃；
>
> 取 index2 上刚刚查到的位置的下一条记录，发现仍然是’zhangs’，取出 ID2，再到 ID 索引上取整行然后判断，这次值对了，将这行记录加入结果集；
>
> 重复上一步，直到在 idxe2 上取到的值不是’zhangs’时，循环结束。
>
> 在这个过程中，要回主键索引取 4 次数据，也就是扫描了 4 行。

通过这个对比，你很容易就可以发现，使用前缀索引后，可能会导致查询语句读数据的次数变多。

**使用前缀索引，定义好长度，就可以做到既节省空间，又不用额外增加太多的查询成本**



我们在建立索引时关注的是区分度，区分度越高越好。因为区分度越高，意味着重复的键值越少

```mysql
select 
  count(distinct left(email,4)）as L4,
  count(distinct left(email,5)）as L5,
  count(distinct left(email,6)）as L6,
  count(distinct left(email,7)）as L7,
from SUser;
```

> 当然，使用前缀索引很可能会损失区分度，所以你需要预先设定一个可以接受的损失比例，比如 5%。然后，在返回的 L4~L7 中，找出不小于 L * 95% 的值，假设这里 L6、L7 都满足，你就可以选择前缀长度为 6。

##### 前缀索引对覆盖索引的影响

> 如果使用 index1（即 email 整个字符串的索引结构）的话，可以利用覆盖索引，从 index1 查到结果后直接就返回了，不需要回到 ID 索引再去查一次。
>
> 而如果使用 index2（即 email(6) 索引结构）的话，就不得不回到 ID 索引再去判断 email 字段的值。
>
> 即使你将 index2 的定义修改为 email(18) 的前缀索引，这时候虽然 index2 已经包含了所有的信息，但 InnoDB 还是要回到 id 索引再查一下，因为系统并不确定前缀索引的定义是否截断了完整信息。

使用前缀索引就用不上覆盖索引对查询性能的优化了

索引选取的越长，占用的磁盘空间就越大，相同的数据页能放下的索引值就越少，搜索的效率也就会越低

如果我们能够确定业务需求里面只有按照身份证进行等值查询的需求，还有没有别的处理方法呢？这种方法，既可以占用更小的空间，也能达到相同的查询效率

- 第一种方式是使用倒序存储
- 第二种方式是使用 hash 字段

##### 小结

- 直接创建完整索引，这样可能比较占用空间；
- 创建前缀索引，节省空间，但会增加查询扫描次数，并且不能使用覆盖索引；
- 倒序存储，再创建前缀索引，用于绕过字符串本身前缀的区分度不够的问题；
- 创建 hash 字段索引，查询性能稳定，有额外的存储和计算消耗，跟第三种方式一样，都不支持范围扫描。

#### 为什么我的MySQL会“抖”一下？

可能原因是：mysql正在刷“脏页”到磁盘中

- 当内存数据页跟磁盘数据页内容不一致的时候，我们称这个内存页为“脏页”。

- 内存数据写入到磁盘后，内存和磁盘上的数据页的内容就一致了，称为“干净页”

##### 什么情况会引发数据库的 flush 过程呢？

- 第一种场景是： InnoDB 的 redo log 写满了。这时候系统会停止所有更新操作，把 checkpoint 往前推进，redo log 留出空间可以继续写
- 第二种场景是：系统内存不足。当需要新的内存页，而内存不够用的时候，就要淘汰一些数据页，空出内存给别的数据页使用。如果淘汰的是“脏页”，就要先将脏页写到磁盘
- 第三种场景是： MySQL 认为系统“空闲”的时候
- 第四种场景是：MySQL 正常关闭的情况。这时候，MySQL 会把内存的脏页都 flush 到磁盘上



第三种情况是属于 MySQL 空闲时的操作，这时系统没什么压力，而第四种场景是数据库本来就要关闭了。这两种情况下，你不会太关注“性能”问题

> 第一种是“redo log 写满了，要 flush 脏页”，这种情况是 InnoDB 要尽量避免的。因为出现这种情况的时候，整个系统就不能再接受更新了，所有的更新都必须堵住。如果你从监控上看，这时候更新数会跌为 0。

> 第二种是“内存不够用了，要先将脏页写到磁盘”，这种情况其实是常态。InnoDB 用缓冲池（buffer pool）管理内存，缓冲池中的内存页有三种状态：
>
> 1、还没有使用的；
>
> 2、使用了并且是干净页；
>
> 3、使用了并且是脏页



当要读入的数据页没有在内存的时候，就必须到缓冲池中申请一个数据页。

这时候只能把最久不使用的数据页从内存中淘汰掉：如果要淘汰的是一个干净页，就直接释放出来复用；但如果是脏页呢，就必须将脏页先刷到磁盘，变成干净页后才能复用。



所以，刷脏页虽然是常态，但是出现以下这两种情况，都是会明显影响性能的：

> 一个查询要淘汰的脏页个数太多，会导致查询的响应时间明显变长；
>
> 日志写满，更新全部堵住，写性能跌为 0，这种情况对敏感业务来说，是不能接受的。



flush脏页时：如何标志redo log中的记录，以至于redo log重放时，不会执行已经flush的change buffer中的数据

http://mysql.taobao.org/monthly/2015/05/01/

>为了管理脏页，在 Buffer Pool 的每个instance上都维持了一个flush list，flush list 上的 page 按照修改这些 page 的LSN号进行排序。
>
>也就是说在重放redo log时，redo log 上的page 的LSN小于 page的LSN时，就不需要再执行了，因为已经执行过了 (类似于版本号的概念)

##### InnoDB 刷脏页的控制策略

正确地告诉 InnoDB 所在主机的 IO 能力

> 用到 innodb_io_capacity 这个参数了，它会告诉 InnoDB 你的磁盘能力。这个值我建议你设置成磁盘的 IOPS

InnoDB 的刷盘速度就是要参考这两个因素：一个是脏页比例，一个是 redo log 写盘速度。

> 参数 innodb_max_dirty_pages_pct 是脏页比例上限，默认值是 75%

> 要尽量避免这种情况，你就要合理地设置 innodb_io_capacity 的值，并且平时要多关注脏页比例，不要让它经常接近 75%。

在准备刷一个脏页的时候，如果这个数据页旁边的数据页刚好是脏页，就会把这个“邻居”也带着一起刷掉

> innodb_flush_neighbors 参数就是用来控制这个行为的，值为 1 的时候会有上述的“连坐”机制
>
> 值为 0 时表示不找邻居，自己刷自己的。