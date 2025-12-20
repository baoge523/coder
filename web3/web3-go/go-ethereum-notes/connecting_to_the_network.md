## go-eth node to connecting to the network

network:
    mainnet
    testnet
        - goerli : Goerli proof-of-authority test network
        - sepolia : epolia proof-of-work test network

但是在2022年时，ethereum 使用了proof-of-stake（权益证明），所以goerli、sepolia也过度到了POS - 所以都需要consensus client


### finding peers 如何找到其他的nodes节点，并建议连接

1、Geth会持续的尝试和其他的node建立连接，如果有足够多的连接，也是可以不在继续向其他node建立连接，此时会开启一个 Internet-facing server
但是它仍然会接受其他的node的连接

--- 这个足够多的连接，到底是多少？ 所有nodes的一半以上么？ -- 参考peer limit

Geth finds peers using the discovery protocol.： Geth通过discovery协议来发现peers的

2、当一个新的节点加入一个网络时，会从一个名叫bootnode的节点中获取node列表信息；bootnode是硬编码到Geth中的，但是我们也可以通过在启动node时指定
> geth --bootnodes enode://pubkey1@ip1:port1,enode://pubkey2@ip2:port2,enode://pubkey3@ip3:port3

3、如果是启动的一个测试node用于本地使用，那么可以通过--nodiscover 参数指定，不发现其他节点


### connectivity problems
1、local time might be incorrect;  an accurate clock is required to participate in the Geth,  一个准确的时钟是必须具有的在Geth中
2、some firewall configurations can prohibit UDP traffic   prohibit 阻止
3、light mode often leads to connectivity issues, because there are few nodes running light servers
    Note：the light mode does not currently work on proof-of-stake networks

4、the public test network Geth is connection might be deprecated or have a low number of  active nodes that are hard to find.
   in this case,the best action is switch to alternative test network


### checking connectivity
net.listening 查看当前节点是否在监听
net.peerCount 查看已经建立连接的peer的个数

admin.peers 查看已经建立连接的peer的详细信息
admin.nodeInfo 查看本地节点的信息

### custom network
自定义网络主要用于开发使用，对于开发而言是很有用的，它不需要连接公开的测试环境或者生产环境，不需要考虑和其他矿工竞争才能创建block

指定一个私有网络，我们可以在启动node时，使用 --networkid参数指定一个chainID，该chainID只要是一个不存在的即可（就是你直接指定的，chainID中不存在的）
同时也需要创建一个genesis.json文件  --- 更多详情，请参考私网模块

### static nodes
静态nodes表示，node每次总是连接指定的peers

### peer limit
指定连接的peer的数量，默认是50个，但是我们也可以通过 --maxpeers 来指定数量


### trust nodes
可信任的节点
可通过在config.toml配置文件中，在TrustedNodes中指定可信任的节点，然后通过配置文件的方式启动node
当然也可以在控制台中通过admin.addTrustedPeer()、admin.removeTrustedPeer() 操作可信任的节点

