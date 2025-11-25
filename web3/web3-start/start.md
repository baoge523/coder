
## 密码学原理


## 数据结构
区块链
以链式的方式存放交易，当前block的hash是由上一个hash+当前数据计算出来的

merkel tree
merkel proof

可以验证一个交易是否在merkel root中，但不能验证一个交易不在merkel root中（理论上，是可以通过hash有序的手段处理的，但是没这个必要）

## 协议

理解如何表示交易的

铸币（挖矿）产生的币

消费-转币
输入：币来源，来源的publicKey
输出：接受者的publicKey          这些用于验证币的合法性

币特性中的共识协议   consensus in bitcoin

接受的区块是扩展最长区块 ---> longest valid chain



## BTC-实现
UTXO：Unspent Transaction Output  作用就是用于防止 double spending attack
用于记录交易输出

## bitcoin network
simple robust
p2p 平等



## 以太坊  --- 得重新看看，现在看的可能有点过时了
比特币和以太坊 都是加密货币 - 去中心化   多国交易更方便

double spending attack
replay attack

比特币  BTC   比拼算力，挖矿工具专业性
以太坊  ETH   

proof of work  工作量证明       proof of state

以太坊支持智能合约    智能合约 类似于 现实中的合同


以太坊的账户： 分两种账户
exterally owned account  外部账户
    balance  how many money you have
    nonce    use record transaction times

smart contract account  智能合约账户    不能主动发起交易，但在一次交易中可用调用其他的智能合约账户
    code     code of account; no change
    storage  storage state; changed by each transaction



状态树、交易树、收据树   数据结构 -> MPT

### 权益证明
比特币、以太坊 都是基于工作证明的方式进行了，有一个缺点就是浪费电

proof of work 工作量证明 POW

在2022年9月15日 以太坊的共识方式从pow 转变成了pos，从工作量证明转换成了权益证明

POS 权益证明 ： 通过用户压置币的方式来验证交易和创建新区块
    说白了：就是压得越多挖矿越容易，但是压置的币需要有一定的冷却时间

### 学习
NFT Non-Fungible token 非同质化通证 。 是一种代表独立资产的数字通证，与比特币和以太网等同质化通证不同，每个NFT都是独一无二的，不能与其他的NFT互换

DAO Decentralized autonomous organization：去中心化自治组织
    通过智能合约在区块链上运行的组织，DAO的决策和管理过程是去中心化的，由所有持有通证的成员共同参与和投票决定
