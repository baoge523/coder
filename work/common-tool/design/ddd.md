### DDD 领域驱动模型

#### 概念


```text
实体（Entity）：有唯一标识（ID）和生命周期的对象。例如 User（UserId 是标识）、Order（OrderId 是标识）。

∙运用：User类中包含 ChangeEmail()、VerifyPassword()等方法，而不仅仅是 getter/setter。

∙值对象（Value Object）：没有唯一标识，通过其属性值来识别的对象。通常是不可变的。例如 Money（包含金额和货币）、Address。

∙运用：将 Address建模为一个包含 Country, Province, City, Detail的值对象，而不是在 User实体中平铺几个字符串字段。这提高了内聚性和表现力。

∙聚合（Aggregate）：这是DDD中最重要、最难掌握的概念。它是一组相关实体和值对象的集合，有一个根实体（Aggregate Root） 作为对外的唯一访问点。

∙目的：维护数据的一致性边界（强一致性边界）。

∙运用：Order（聚合根）和 OrderItem是典型的聚合。外界不能直接操作 OrderItem，必须通过 Order聚合根的方法（如 Order.AddItem(...)）来操作。这保证了在添加商品时，订单总价能同步更新等业务规则。

∙领域服务（Domain Service）：当某个业务逻辑不适合放在实体或值对象中时（它不属于任何一个实体），使用领域服务。例如，“资金转账”操作涉及两个账户实体，它就是一个典型的领域服务 TransferService.Transfer(...)。

∙仓储（Repository）：提供类似集合的接口，用于持久化和检索聚合。只定义在聚合根上。

∙运用：定义接口 IOrderRepository在领域层，包含 GetById, Add, Remove等方法。其具体实现（如 OrderRepository）在基础设施层，用 EF Core 或 MyBatis 等实现。

∙领域事件（Domain Event）：表示领域中发生的某个重要事情。用于解耦不同聚合或甚至不同限界上下文之间的交互。

∙运用：当 Order被支付后，会发布一个 OrderPaidEvent。然后一个独立的 handler 可以监听这个事件，去触发发送邮件、通知物流等后续操作，而不需要支付核心逻辑关心这些。


```


#### 在日常开发中的应用

1、在代码分层架构上
```text
云监控的代码分层：
1、entity层，包含了实力的字段定义和操作实体的方法定义； ---  entity（实体） + Value Object（值对象）

2、usecase层，业务功能的实现层，只依赖于entity； --- 领域服务（Domain Service），聚合（Aggregate）

3、pkg层，非核心业务的三方依赖

4、其他的辅助业务功能的代码

5、cmd层，代码启动入口


```

2、在数据内敛上
```text

更多的体现的entity和Value Object的定义上，将有业务意义的不直接使用单独的字段，而是通过定义的值对象，该值对象包含了自己的一些操作方法，比如address定义成值对象，被多个entity使用（county、city。。。）

```