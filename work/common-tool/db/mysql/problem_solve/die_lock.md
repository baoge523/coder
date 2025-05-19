# mysql 死锁问题分析
查看mysql死锁信息的命令 （保留最近一次的死锁信息）
> show engine innodb status

查看到死锁信息如下:
例子1：
这里造成死锁的原因是：并发insert时，两个事务都锁定了同一个gap lock，导致的相互等待
transaction1:
INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6164994,339773,'')
transaction2:
INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6164993,339773,'')

transaction1 and transaction2 都锁着住了policy_id index (6164992,正无穷)的 gap lock,导致的相互等待

解决方式：
1、串行执行
2、同一个事务中执行

```sql
CREATE TABLE `policy_notice_relation` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`policy_id` int(11) NOT NULL,
	`notice_id` int(11) NOT NULL,
	`bind_levels` varchar(100) DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `i_notice_id` (`notice_id`),
	KEY `i_policy_id` (`policy_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 11873506 CHARSET = utf8
```

```text
2024-07-30 16:16:09 7f453b997700 *** 
(1) TRANSACTION: 
TRANSACTION 8590711621, ACTIVE 2 sec inserting mysql tables in use 1, 
locked 1 LOCK WAIT 7 lock struct(s), heap size 1184, 3 row lock(s), undo log entries 5 MySQL thread id 806605821, 
OS thread handle 0x7f45433ef700, query id 8674469325 30.170.225.64 root 
update INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6164994,339773,'') 

(1) WAITING FOR THIS LOCK TO BE GRANTED: 
RECORD LOCKS space id 2152631 page no 25 n bits 920 index `i_policy_id` of table `domain_policy`.`policy_notice_relation` trx id 8590711621 
lock_mode X insert intention waiting Record lock


(2) TRANSACTION: 
TRANSACTION 8590711620, ACTIVE 2 sec inserting mysql tables in use 1, locked 1 6 lock struct(s), 
heap size 1184, 3 row lock(s), undo log entries 4 MySQL thread id 806605756, OS thread handle 0x7f453b997700, 
query id 8674469500 30.170.225.64 root update 
INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6164993,339773,'') 

(2) HOLDS THE LOCK(S): 
RECORD LOCKS space id 2152631 page no 25 n bits 920 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` trx id 8590711620 lock_mode X Record lock

(2) WAITING FOR THIS LOCK TO BE GRANTED: 
RECORD LOCKS space id 2152631 page no 25 n bits 920 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` 
trx id 8590711620 lock_mode X insert intention waiting Record lock

```
例子2：
```text
2024-08-06 13:28:45 7f4549935700 *** (1) TRANSACTION: TRANSACTION 8594815465, 
ACTIVE 1 sec inserting mysql tables in use 1, locked 1 LOCK WAIT 6 lock struct(s), 
heap size 1184, 3 row lock(s), undo log entries 4 MySQL thread id 828917232, OS thread handle 0x7f45831f7700, 

query id 8970435318 30.170.225.64 root update 
INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6174556,339773,'') 
*** (1) WAITING FOR THIS LOCK TO BE GRANTED: RECORD LOCKS space id 2152631 page no 25 n bits 976 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` 
trx id 8594815465 lock_mode X insert intention waiting Record lock, 
heap no 1 PHYSICAL RECORD: n_fields 1; compact format; info bits 0 0: len 8; hex 73757072656d756d; asc supremum;; 
*** (2) TRANSACTION: TRANSACTION 8594815468, ACTIVE 1 sec inserting mysql tables in use 1, locked 1 6 lock struct(s), 
heap size 1184, 3 row lock(s), undo log entries 4 MySQL thread id 828917621, OS thread handle 0x7f4549935700, 
query id 8970435389 30.170.225.64 root update 
INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6174557,339773,'') 
*** (2) HOLDS THE LOCK(S): RECORD LOCKS space id 2152631 page no 25 n bits 976 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` 
trx id 8594815468 lock_mode X Record lock, heap no 1 PHYSICAL RECORD: n_fields 1; compact format; 
info bits 0 0: len 8; hex 73757072656d756d; asc supremum;; 
*** (2) WAITING FOR THIS LOCK TO BE GRANTED: RECORD LOCKS space id 2152631 page no 25 n bits 976 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` 
trx id 8594815468 lock_mode X insert intention waiting Record lock, heap no 1 PHYSICAL RECORD: n_fields 1; 
compact format; info bits 0 0: len 8; hex 73757072656d756d; asc supremum;; 
*** WE ROLL BACK TRANSACTION (2) 
```

例子3：
```text
2024-09-13 16:38:20 0x7ec9126fb700 
*** (1) TRANSACTION: 
TRANSACTION 14151871647, ACTIVE 0 sec inserting mysql tables in use 1, locked 1 LOCK WAIT 3 lock struct(s), 
heap size 1136, 2 row lock(s), undo log entries 2 MySQL thread id 13466162, OS thread handle 139399709333248, 
query id 17137453231 30.167.45.123 cloud_monitor update 
replace into cInstanceItem(id, instanceGroupId, region, uniqueId, appAddress, deviceName, deviceIP, uuid, eventUniqueId, 
information, eventInfo, lastEditUin, lastModifyTime, createTime) 
values ('','162333', 'sh', 'da5d2041b1abd4a4f55682d978f31b66', '2', '', '', 
'namespace=kgcagent-prod&tenant=pulsar-dm25z9qvv4xq&topic=kgcagent-cancel-bf0731a6-57e2-4ffc-af7d-217dd3337330', '',
'{"namespace":"kgcagent-prod","tenant":"pulsar-dm25z9qvv4xq","topic":"kgcagent-cancel-bf0731a6-57e2-4ffc-af7d-217dd3337330"}',
'','100037935268','1726216700', '1726216700'), 
('','162333', 'sh', '1c64cf919d8ef417a98fdfcfc0dea68e', '2', '', '', 
'namespace=kgcagent-prod&tenant=pulsar-dm25z9qvv4xq&topic=kgcagent-c36ae2f2-6e45-40c1-a628-e5d4266b4134', '',
'{"namespace":"kgcagent-prod","tenant":"pulsar-dm25z9qvv4xq","topic":"kgcagent-c36ae2f2-6e45-40c1-a628-e5d4266b4134"}','',
'100037935268','1726216700', '1726216700'), ('','162333' 

*** (1) WAITING FOR THIS LOCK TO BE GRANTED: 
RECORD LOCKS space id 295 page no 7517 n bits 344 
index idx_igid_iid_rg of table `StormCloudConf`.`cInstanceItem` 
trx id 14151871647 lock_mode X locks gap before rec insert intention waiting Record lock, heap no 150 PHYSICAL 

*** (2) TRANSACTION: 
TRANSACTION 14151871645, ACTIVE 0 sec fetching rows mysql tables in use 2, locked 2 9 lock struct(s), 
heap size 1136, 321 row lock(s) MySQL thread id 13466165, OS thread handle 139402062837504, 
query id 17137453229 30.167.46.192 cloud_monitor Sending data 
update cInstanceGroup set lastEditUin=100037935268,lastModifyTime=1726216700,
instanceSum=(select count(1) from cInstanceItem where instanceGroupId=162333) where id=162333 
*** (2) HOLDS THE LOCK(S): 
RECORD LOCKS space id 295 page no 7517 n bits 344 
index idx_igid_iid_rg of table `StormCloudConf`.`cInstanceItem` trx id 14151871645 lock mode S Record lock, 
heap no 1 PHYSICAL RECORD: 
```
lock mode S 表示共享锁

### 例子4
该例子的分析结论是：
> UPDATE `notice_users` SET `deleted_at`='2025-05-14 15:15:21.523' WHERE app_id = 1321948919 AND notice_id = 2795031 AND `notice_users`.`deleted_at` IS NULL

该sql语句在执行时，mysql会锁住`notice_id`、`app_id`、`id`索引的相关行

死锁的本质，就是两个或者多个transaction获取锁的顺序不一致，导致的相关依赖
事务1、mysql优化器选择了notice_id索引，然后持有了notice_id 和 primary key(id)的相关行的锁，准备获取app_id索引上的相关行的锁， -- 这里的锁，可能是行锁 + 间隙锁
事务2、mysql优化器选择了app_id索引，然后只有app_id相关行的锁，准备获取primary key(id)相关行锁，与事务1冲突

前提是：
1、在这两个update语句时，mysql优化器在判断notice_id 和app_id索引时，发现两个的查询成本差不多
2、且获取锁的顺序不一致导致

解决死锁的方式：
1、添加一个 app_id、notice_id的复合索引
2、人工干预，让mysql优化器选择同一个索引检索，比如选择app_id



```sql
CREATE TABLE `notice_users` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`app_id` int(11) NOT NULL,
	`notice_id` int(11) NOT NULL,
	PRIMARY KEY (`id`),
	KEY `i_app_id` (`app_id`),
	KEY `i_notice_id` (`notice_id`)
) ENGINE = InnoDB CHARSET = utf8
```
```text
Transaction1:

UPDATE `notice_users` SET `deleted_at`='2025-05-14 15:15:21.523' WHERE app_id = 1321948919 AND notice_id = 2795030 AND `notice_users`.`deleted_at` IS NULL

LockRequest:
RECORD LOCKS space id 234 page no 54454 n bits 896 index i_app_id of table `domain_policy`.`notice_users` trx id 733248075 lock_mode X locks rec but not gap waiting

Transaction2:

UPDATE `notice_users` SET `deleted_at`='2025-05-14 15:15:21.523' WHERE app_id = 1321948919 AND notice_id = 2795031 AND `notice_users`.`deleted_at` IS NULL

LockRequest:
RECORD LOCKS space id 234 page no 36468 n bits 168 index PRIMARY of table `domain_policy`.`notice_users` trx id 733248074 lock_mode X locks rec but not gap waiting

LockHold:
RECORD LOCKS space id 234 page no 54454 n bits 896 index i_app_id of table `domain_policy`.`notice_users` trx id 733248074 lock_mode X locks rec but not gap
```
```text
2025-05-14 15:15:21 0x7f1df7381700
*** (1) TRANSACTION:
TRANSACTION 733248075, ACTIVE 0 sec fetching rows
mysql tables in use 3, locked 3
LOCK WAIT 8 lock struct(s), heap size 1136, 7 row lock(s), undo log entries 1
MySQL thread id 11304439, OS thread handle 139770761574144, query id 24411137617 21.22.153.124 root updating
UPDATE `notice_users` SET `deleted_at`='2025-05-14 15:15:21.523' WHERE app_id = 1321948919 AND notice_id = 2795030 AND `notice_users`.`deleted_at` IS NULL
*** (1) WAITING FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 234 page no 54454 n bits 896 index i_app_id of table `domain_policy`.`notice_users` trx id 733248075 lock_mode X locks rec but not gap waiting
Record lock, heap no 98 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 4; hex cecb56f7; asc   V ;;
 1: len 4; hex 8036045b; asc  6 [;;

*** (2) TRANSACTION:
TRANSACTION 733248074, ACTIVE 0 sec fetching rows
mysql tables in use 3, locked 3
10 lock struct(s), heap size 1136, 8 row lock(s), undo log entries 1
MySQL thread id 11304537, OS thread handle 139766678427392, query id 24411137613 30.173.60.214 root updating
UPDATE `notice_users` SET `deleted_at`='2025-05-14 15:15:21.523' WHERE app_id = 1321948919 AND notice_id = 2795031 AND `notice_users`.`deleted_at` IS NULL
*** (2) HOLDS THE LOCK(S):
RECORD LOCKS space id 234 page no 54454 n bits 896 index i_app_id of table `domain_policy`.`notice_users` trx id 733248074 lock_mode X locks rec but not gap
Record lock, heap no 98 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 4; hex cecb56f7; asc   V ;;
 1: len 4; hex 8036045b; asc  6 [;;

Record lock, heap no 100 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 4; hex cecb56f7; asc   V ;;
 1: len 4; hex 8036045d; asc  6 ];;

*** (2) WAITING FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 234 page no 36468 n bits 168 index PRIMARY of table `domain_policy`.`notice_users` trx id 733248074 lock_mode X locks rec but not gap waiting
Record lock, heap no 29 PHYSICAL RECORD: n_fields 26; compact format; info bits 0
 0: len 4; hex 8036045d; asc  6 ];;
 1: len 6; hex 00000e187d2d; asc     }-;;
 2: len 7; hex 2500000e1d1427; asc %     ';;
 3: len 4; hex cecb56f7; asc   V ;;
 4: len 12; hex 313030303333393731373235; asc 100033971725;;
 5: len 4; hex 802aa616; asc  *  ;;
 6: len 4; hex 55534552; asc USER;;
 7: len 17; hex 31383233343033362c3137353639373035; asc 18234036,17569705;;
 8: len 0; hex ; asc ;;
 9: len 0; hex ; asc ;;
 10: len 4; hex 80000000; asc     ;;
```