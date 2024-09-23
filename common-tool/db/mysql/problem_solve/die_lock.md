# mysql 死锁问题分析
查看mysql死锁信息的命令 （保留最近一次的死锁信息）
> show engine innodb status

查看到死锁信息如下:
例子1：
```text
2024-07-30 16:16:09 7f453b997700 *** 
(1) TRANSACTION: TRANSACTION 8590711621, ACTIVE 2 sec inserting mysql tables in use 1, 
locked 1 LOCK WAIT 7 lock struct(s), heap size 1184, 3 row lock(s), undo log entries 5 MySQL thread id 806605821, 
OS thread handle 0x7f45433ef700, 

query id 8674469325 30.170.225.64 root 
update INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6164994,339773,'') 
*** (1) WAITING FOR THIS LOCK TO BE GRANTED: 
RECORD LOCKS space id 2152631 page no 25 n bits 920 index `i_policy_id` of table `domain_policy`.`policy_notice_relation` trx id 8590711621 
lock_mode X insert intention waiting Record lock, heap no 1 PHYSICAL RECORD: n_fields 1; 
compact format; info bits 0 0: len 8; hex 73757072656d756d; asc supremum;; 
*** (2) TRANSACTION: TRANSACTION 8590711620, ACTIVE 2 sec inserting mysql tables in use 1, locked 1 6 lock struct(s), 
heap size 1184, 3 row lock(s), undo log entries 4 MySQL thread id 806605756, OS thread handle 0x7f453b997700, 

query id 8674469500 30.170.225.64 root update 
INSERT INTO `policy_notice_relation` (`policy_id`,`notice_id`,`bind_levels`) VALUES (6164993,339773,'') 
*** (2) HOLDS THE LOCK(S): RECORD LOCKS space id 2152631 page no 25 n bits 920 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` trx id 8590711620 lock_mode X Record lock, 
heap no 1 PHYSICAL RECORD: n_fields 1; compact format; info bits 0 0: len 8; hex 73757072656d756d; asc supremum;; 
*** (2) WAITING FOR THIS LOCK TO BE GRANTED: RECORD LOCKS space id 2152631 page no 25 n bits 920 
index `i_policy_id` of table `domain_policy`.`policy_notice_relation` 
trx id 8590711620 lock_mode X insert intention waiting Record lock, 
heap no 1 PHYSICAL RECORD: n_fields 1; compact format; info bits 0 0: len 8; 
hex 73757072656d756d; asc supremum;; *** WE ROLL BACK TRANSACTION (2) 

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
*** (1) TRANSACTION: TRANSACTION 14151871647, ACTIVE 0 sec inserting mysql tables in use 1, locked 1 LOCK WAIT 3 lock struct(s), 
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
'100037935268','1726216700', '1726216700'), ('','162333' *** 
(1) WAITING FOR THIS LOCK TO BE GRANTED: RECORD LOCKS space id 295 page no 7517 n bits 344 
index idx_igid_iid_rg of table `StormCloudConf`.`cInstanceItem` 
trx id 14151871647 lock_mode X locks gap before rec insert intention waiting Record lock, heap no 150 PHYSICAL 
RECORD: n_fields 4; compact format; info bits 0 0: len 8; hex 8000000000027a1d; asc z ;; 1: len 30; 
hex 316534326463356636333732343462393365623730663430623261653738; asc 1e42dc5f637244b93eb70f40b2ae78; (total 32 bytes); 
2: len 2; hex 7368; asc sh;; 3: len 8; hex 80000000001602f0; asc ;; 

*** (2) TRANSACTION: TRANSACTION 14151871645, ACTIVE 0 sec fetching rows mysql tables in use 2, locked 2 9 lock struct(s), 
heap size 1136, 321 row lock(s) MySQL thread id 13466165, OS thread handle 139402062837504, 
query id 17137453229 30.167.46.192 cloud_monitor Sending data 
update cInstanceGroup set lastEditUin=100037935268,lastModifyTime=1726216700,
instanceSum=(select count(1) from cInstanceItem where instanceGroupId=162333) where id=162333 
*** (2) HOLDS THE LOCK(S): RECORD LOCKS space id 295 page no 7517 n bits 344 
index idx_igid_iid_rg of table `StormCloudConf`.`cInstanceItem` trx id 14151871645 lock mode S Record lock, 
heap no 1 PHYSICAL RECORD: n_fields 1; compact format; info bits 0 0: len 8; hex 73757072656d756d; asc supremum;; 
Record lock, heap no 108 PHYSICAL RECORD: n_fields 4; compact format; info bits 0 0: len 8; hex 8000000000027a1d; 
asc z ;; 1: len 30; hex 303262313363343331333335623730323361386265336365303461313937; asc 02b13c431335b7023a8be3ce04a197; (total 32 bytes); 
2: len 2; hex 7368; asc sh;; 3: len 8; hex 80000000001602d1; asc ;; Record lock, heap no 109 PHYSICAL RECORD: n_fields 4; 
compact format; info bits 0 0: len 8; hex 8000000000027a1d; asc z ;; 1: len 30; hex 303637616439623935383036613333306630363364356539346132356332; 
   *** WE ROLL BACK
```