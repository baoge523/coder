## 同步模型

### 全节点 full
snap full nodes  从一个相对最近的block开始同步，然后最多保持128个block在内存中的数据 -- 即总是保留最近的128个block的数据，维护内部的状态树 --速度快
full nodes  从创世block开始同步，然后最多也保持128个block在内存中 -- 即总是保留最近的128个block的数据，维护内部的状态树 -- 速度慢
archive nodes  归档节点，记录了从创世block到现在的所有block信息，到2022年就已经有12TB，目前未知
archive snap nodes 可以新建一个snap archive节点，表示从初始化的block（最近的block）开始一直记录，不会滚动掉数据


### 轻节点 light
好像不可用 -- 待确认

### consensus layer syncing

#### optimistic syncing
在execution client 验证block之前，optimistic syncing 下载block
在optimistic sync 下载在block期间，认为下载的block数据总是正确的，然后补充性的验证所有下载的block
当节点是optimistic sync时，是不允许证明和提交（产出）block，原因是他们不能保证他们的链首是合法的，正确的
#### checkpoint syncing
sync fast

consensus client 会从一个可信任的来源抓去一个检查点，该检查点会提供同步的状态，之后才会切换成全同步模式，然后验证每一个区块是真实的

