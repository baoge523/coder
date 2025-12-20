
## 节点的结构
geth 主要分为两个client
    - execution client  矿工client
    - consensus client 共识client

### execution client
transaction handle
transaction gossip
EVM
state management

1、接受用户的交易信息，执行交易，维护状态树（MPT）
2、接受公链上的block，验证block中的交易的合法性，更新状态树 - re-execution transaction

### consensus client
consensus logic
fork choice   分叉选择，总是选择最长链
block gossip

1、接受来自peer的block
2、运行分叉选择算法，保证总是选择最丰富积累证明的链路
3、和execution client 是不同的p2p网络
4、自己的p2p网络共享block
5、不会提交block， 则是validator干的事情 （consensus client可选的成为validator）


### validator
需要32ETH作为押金（当然可以更多，越多会容易获得记账权），成为validator，可以提交block赢得冲块奖励