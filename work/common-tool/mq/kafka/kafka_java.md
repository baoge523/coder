

# kafka

## 一、client

### producer

### consumer

使用样例：

```java
Properties props = new Properties();
props.setProperty("bootstrap.servers", "localhost:9092");
props.setProperty("group.id", "test");
props.setProperty("enable.auto.commit", "true");
props.setProperty("auto.commit.interval.ms", "1000");
props.setProperty("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
props.setProperty("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
KafkaConsumer<String, String> consumer = new KafkaConsumer<>(props);
consumer.subscribe(Arrays.asList("foo", "bar"));   // 设置主题
while (true) {
    ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(100));
    for (ConsumerRecord<String, String> record : records)
        System.out.printf("offset = %d, key = %s, value = %s%n", record.offset(), record.key(), record.value());
}
```

#### KafkaConsumer构造器

```java
KafkaConsumer(ConsumerConfig config, Deserializer<K> keyDeserializer, Deserializer<V> valueDeserializer) {
        try {
            // 组重平衡相关的配置
            GroupRebalanceConfig groupRebalanceConfig = new GroupRebalanceConfig(config,
                    GroupRebalanceConfig.ProtocolType.CONSUMER);
            // 消费者组id
            this.groupId = Optional.ofNullable(groupRebalanceConfig.groupId);
            // 客户端id
            this.clientId = config.getString(CommonClientConfigs.CLIENT_ID_CONFIG);
            // 日志上下文
            LogContext logContext;

            // If group.instance.id is set, we will append it to the log context.
            if (groupRebalanceConfig.groupInstanceId.isPresent()) {
                logContext = new LogContext("[Consumer instanceId=" + groupRebalanceConfig.groupInstanceId.get() +
                        ", clientId=" + clientId + ", groupId=" + groupId.orElse("null") + "] ");
            } else {
                logContext = new LogContext("[Consumer clientId=" + clientId + ", groupId=" + groupId.orElse("null") + "] ");
            }

            this.log = logContext.logger(getClass());
            // 是否开启自动提交
            boolean enableAutoCommit = config.maybeOverrideEnableAutoCommit();
            groupId.ifPresent(groupIdStr -> {
                if (groupIdStr.isEmpty()) {
                    log.warn("Support for using the empty group id by consumers is deprecated and will be removed in the next major release.");
                }
            });

            log.debug("Initializing the Kafka consumer");
            // 请求超时时间：request.timeout.ms
            this.requestTimeoutMs = config.getInt(ConsumerConfig.REQUEST_TIMEOUT_MS_CONFIG);
            this.defaultApiTimeoutMs = config.getInt(ConsumerConfig.DEFAULT_API_TIMEOUT_MS_CONFIG);
            this.time = Time.SYSTEM;
            // 指标
            this.metrics = buildMetrics(config, time, clientId);
            // 重试的阻塞时间
            this.retryBackoffMs = config.getLong(ConsumerConfig.RETRY_BACKOFF_MS_CONFIG);
            // 拦截器集合：主要在消息返回时做增强拦截
            List<ConsumerInterceptor<K, V>> interceptorList = (List) config.getConfiguredInstances(
                    ConsumerConfig.INTERCEPTOR_CLASSES_CONFIG,
                    ConsumerInterceptor.class,
                    Collections.singletonMap(ConsumerConfig.CLIENT_ID_CONFIG, clientId));
            this.interceptors = new ConsumerInterceptors<>(interceptorList);

            // key 反序列化器
            if (keyDeserializer == null) {
                this.keyDeserializer = config.getConfiguredInstance(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, Deserializer.class);
                this.keyDeserializer.configure(config.originals(Collections.singletonMap(ConsumerConfig.CLIENT_ID_CONFIG, clientId)), true);
            } else {
                config.ignore(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG);
                this.keyDeserializer = keyDeserializer;
            }
            // value 反序列化器
            if (valueDeserializer == null) {
                this.valueDeserializer = config.getConfiguredInstance(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, Deserializer.class);
                this.valueDeserializer.configure(config.originals(Collections.singletonMap(ConsumerConfig.CLIENT_ID_CONFIG, clientId)), false);
            } else {
                config.ignore(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG);
                this.valueDeserializer = valueDeserializer;
            }
            // offset 重置策略  LATEST, EARLIEST, NONE
            OffsetResetStrategy offsetResetStrategy = OffsetResetStrategy.valueOf(config.getString(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG).toUpperCase(Locale.ROOT));
            /**
             *  订阅者信息
             */
            this.subscriptions = new SubscriptionState(logContext, offsetResetStrategy);
            // 资源监听器
            ClusterResourceListeners clusterResourceListeners = configureClusterResourceListeners(keyDeserializer,
                    valueDeserializer, metrics.reporters(), interceptorList);
            // 消费者元数据信息
            this.metadata = new ConsumerMetadata(retryBackoffMs,                        // 重试阻塞时间
                    config.getLong(ConsumerConfig.METADATA_MAX_AGE_CONFIG),
                    !config.getBoolean(ConsumerConfig.EXCLUDE_INTERNAL_TOPICS_CONFIG),  // 排除内部主题 取反
                    config.getBoolean(ConsumerConfig.ALLOW_AUTO_CREATE_TOPICS_CONFIG),  // 是否允许自动创建主题
                    subscriptions, logContext, clusterResourceListeners);               // 订阅者、日志、资源监听器

            // 消费者连接的kafka cluster 地址
            List<InetSocketAddress> addresses = ClientUtils.parseAndValidateAddresses(
                    config.getList(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG), config.getString(ConsumerConfig.CLIENT_DNS_LOOKUP_CONFIG));
            this.metadata.bootstrap(addresses);
            String metricGrpPrefix = "consumer";

            FetcherMetricsRegistry metricsRegistry = new FetcherMetricsRegistry(Collections.singleton(CLIENT_ID_METRIC_TAG), metricGrpPrefix);
            // 消费者的channel构建器
            ChannelBuilder channelBuilder = ClientUtils.createChannelBuilder(config, time, logContext);
            // 隔离级别 读未提交 读已提交
            IsolationLevel isolationLevel = IsolationLevel.valueOf(
                    config.getString(ConsumerConfig.ISOLATION_LEVEL_CONFIG).toUpperCase(Locale.ROOT));
            Sensor throttleTimeSensor = Fetcher.throttleTimeSensor(metrics, metricsRegistry);
            // 内部心跳：heartbeat.interval.ms
            int heartbeatIntervalMs = config.getInt(ConsumerConfig.HEARTBEAT_INTERVAL_MS_CONFIG);

            ApiVersions apiVersions = new ApiVersions();
            /**
             * 网络客户端
             * 主要是用来与kafka cluster 交互信息的
             * 比如：
             *   获取元数据
             *   拉取消息
             *   提交位移
             */
            NetworkClient netClient = new NetworkClient(
                    new Selector(config.getLong(ConsumerConfig.CONNECTIONS_MAX_IDLE_MS_CONFIG), metrics, time, metricGrpPrefix, channelBuilder, logContext),
                    this.metadata,
                    clientId,
                    100, // a fixed large enough value will suffice for max in-flight requests
                    config.getLong(ConsumerConfig.RECONNECT_BACKOFF_MS_CONFIG),
                    config.getLong(ConsumerConfig.RECONNECT_BACKOFF_MAX_MS_CONFIG),
                    config.getInt(ConsumerConfig.SEND_BUFFER_CONFIG),
                    config.getInt(ConsumerConfig.RECEIVE_BUFFER_CONFIG),
                    config.getInt(ConsumerConfig.REQUEST_TIMEOUT_MS_CONFIG),
                    config.getLong(ConsumerConfig.SOCKET_CONNECTION_SETUP_TIMEOUT_MS_CONFIG),
                    config.getLong(ConsumerConfig.SOCKET_CONNECTION_SETUP_TIMEOUT_MAX_MS_CONFIG),
                    ClientDnsLookup.forConfig(config.getString(ConsumerConfig.CLIENT_DNS_LOOKUP_CONFIG)),
                    time,
                    true,
                    apiVersions,
                    throttleTimeSensor,
                    logContext);
            /**
             * 包装了一下NetworkClient，底层干活的还是NetworkClient
             */
            this.client = new ConsumerNetworkClient(
                    logContext,         // 日志上下文
                    netClient,          // 网络客户端
                    metadata,           // 元数据
                    time,               // 时间
                    retryBackoffMs,     // 重试阻塞时间
                    config.getInt(ConsumerConfig.REQUEST_TIMEOUT_MS_CONFIG),  // 请求超时时间
                    heartbeatIntervalMs); //Will avoid blocking an extended period of time to prevent heartbeat thread starvation

            // 分区分配器 默认是 org.apache.kafka.clients.consumer.RangeAssignor
            this.assignors = getAssignorInstances(config.getList(ConsumerConfig.PARTITION_ASSIGNMENT_STRATEGY_CONFIG),
                    config.originals(Collections.singletonMap(ConsumerConfig.CLIENT_ID_CONFIG, clientId)));

            // no coordinator will be constructed for the default (null) group id
            // group id 为null时，coordinator 也为null
            this.coordinator = !groupId.isPresent() ? null :
                new ConsumerCoordinator(groupRebalanceConfig,   //组配置
                        logContext,                             // 日志上下文
                        this.client,                            // 消费者网络客户端工具
                        assignors,                              // 分区分配器
                        this.metadata,                          // 元数据
                        this.subscriptions,                     // 订阅者
                        metrics,                                // 指标
                        metricGrpPrefix,
                        this.time,
                        enableAutoCommit,                       // 是否自动提交
                        config.getInt(ConsumerConfig.AUTO_COMMIT_INTERVAL_MS_CONFIG), //自动提交时间
                        this.interceptors,                      // 拦截器
                        config.getBoolean(ConsumerConfig.THROW_ON_FETCH_STABLE_OFFSET_UNSUPPORTED));
            /**
             * 获取数据
             */
            this.fetcher = new Fetcher<>(
                    logContext,                // 日志上下文
                    this.client,               // 消费者网络客户端工具
                    config.getInt(ConsumerConfig.FETCH_MIN_BYTES_CONFIG),     // 拿取最小的byte大小
                    config.getInt(ConsumerConfig.FETCH_MAX_BYTES_CONFIG),     // 拿取最大的byte大小
                    config.getInt(ConsumerConfig.FETCH_MAX_WAIT_MS_CONFIG),   // 最大的等待时间
                    config.getInt(ConsumerConfig.MAX_PARTITION_FETCH_BYTES_CONFIG),  // 分区获取的最大byte大小
                    config.getInt(ConsumerConfig.MAX_POLL_RECORDS_CONFIG),    // 默认一次性最多拿500条消息记录
                    config.getBoolean(ConsumerConfig.CHECK_CRCS_CONFIG),
                    config.getString(ConsumerConfig.CLIENT_RACK_CONFIG),
                    this.keyDeserializer,
                    this.valueDeserializer,
                    this.metadata,
                    this.subscriptions,
                    metrics,
                    metricsRegistry,
                    this.time,
                    this.retryBackoffMs,
                    this.requestTimeoutMs,
                    isolationLevel,
                    apiVersions);

            this.kafkaConsumerMetrics = new KafkaConsumerMetrics(metrics, metricGrpPrefix);

            config.logUnused();
            AppInfoParser.registerAppInfo(JMX_PREFIX, clientId, metrics, time.milliseconds());
            log.debug("Kafka consumer initialized");
        } catch (Throwable t) {
            throw new KafkaException("Failed to construct kafka consumer", t);
        }
    }
```



#### KafkaConsumer的重要组件

##### SubscriptionState

###### 实例化

```java
this.subscriptions = new SubscriptionState(logContext, offsetResetStrategy); // 日志上下文、offset重置策略(LATEST, EARLIEST, NONE)
```

```java
public SubscriptionState(LogContext logContext, OffsetResetStrategy defaultResetStrategy) {
        this.log = logContext.logger(this.getClass());       // 日志
        this.defaultResetStrategy = defaultResetStrategy;    // 重置策略
        this.subscription = new HashSet<>();								 // 当前消费者订阅的主题列表
        this.assignment = new PartitionStates<>();					 // 当前消费者消费的分区列表状态 （包含分区位移、分区leader epoch）
        this.groupSubscription = new HashSet<>();						 // 当前消费者的所属组订阅的所有主题列表
        this.subscribedPattern = null;         							 // 正则的方式订阅主题
        this.subscriptionType = SubscriptionType.NONE;       // 消费类型：NONE, AUTO_TOPICS, AUTO_PATTERN, USER_ASSIGNED
    }
```

###### 作用

```java
// A class for tracking the topics, partitions, and offsets for the consumer
```

用来记录消费者的主题信息、分区信息、位移信息

###### 重要属性

```java
private Set<String> subscription;      // 主题列表
private Set<String> groupSubscription; // 组订阅的所有主题列表
private final PartitionStates<TopicPartitionState> assignment;  // 分区状态 （当前消费者消费的分区）
private final OffsetResetStrategy defaultResetStrategy;    // 位移重置策略
private ConsumerRebalanceListener rebalanceListener;      // 重平衡监听器
```

PartitionStates

> 用来保存主题分区 和 主题分区对应的状态

```java
private final LinkedHashMap<TopicPartition, S> map = new LinkedHashMap<>();  // s 就是 TopicPartitionState
```

TopicPartitionState

> 用来保存主题分区的状态

```java
private FetchState fetchState;  // 分区的拉取状态  INITIALIZING  FETCHING  AWAIT_RESET  AWAIT_VALIDATION

//包含的信息：（offset、offsetEpoch、currentLeader（leader、epoch））
private FetchPosition position; // last consumed position  上一次消费的位移信息  

private Long highWatermark; // the high watermark from last fetch  上次拉取的高水位
private Long logStartOffset; // the log start offset 日志开始位移
private Long lastStableOffset; // 开启事务的情况下的 可见位移
private boolean paused;  // 分区是否被用户暂停
private OffsetResetStrategy resetStrategy;  // 重置策略
```



##### ConsumerMetadata

###### 实例化

```java
 this.metadata = new ConsumerMetadata(retryBackoffMs,                        // 重试阻塞时间
                    config.getLong(ConsumerConfig.METADATA_MAX_AGE_CONFIG),  // 元数据过期时间，最多可用多长时间
                    !config.getBoolean(ConsumerConfig.EXCLUDE_INTERNAL_TOPICS_CONFIG),  // 排除内部主题 取反
                    config.getBoolean(ConsumerConfig.ALLOW_AUTO_CREATE_TOPICS_CONFIG),  // 是否允许自动创建主题
                    subscriptions, 					// 订阅者
                    logContext, 						// 日志
                    clusterResourceListeners);              // 资源监听器

// 消费者连接的kafka cluster 地址
List<InetSocketAddress> addresses = ClientUtils.parseAndValidateAddresses(
                    config.getList(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG),
  									config.getString(ConsumerConfig.CLIENT_DNS_LOOKUP_CONFIG));

// 将地址保存在metadata中
this.metadata.bootstrap(addresses); 
```

###### 作用

> 管理着消费者的元信息

##### ConsumerNetworkClient

###### 实例化

```java
// 网络客户端
NetworkClient netClient = new NetworkClient(
                    new Selector(config.getLong(ConsumerConfig.CONNECTIONS_MAX_IDLE_MS_CONFIG), metrics, time, metricGrpPrefix, channelBuilder, logContext),              // nio 选择器
                    this.metadata,        // 元数据
                    clientId,             // 客户端id
                    100, // a fixed large enough value will suffice for max in-flight requests  发送中的最大请求数
                    config.getLong(ConsumerConfig.RECONNECT_BACKOFF_MS_CONFIG),     // 重连时的基础等待时间 默认50ms
                    config.getLong(ConsumerConfig.RECONNECT_BACKOFF_MAX_MS_CONFIG), // 重连时的最大等待时间 默认1s
                    config.getInt(ConsumerConfig.SEND_BUFFER_CONFIG),               // 发送的buffer大小
                    config.getInt(ConsumerConfig.RECEIVE_BUFFER_CONFIG),            // 接收的buffer大小
                    config.getInt(ConsumerConfig.REQUEST_TIMEOUT_MS_CONFIG),        // 请求超时时间 默认40s
                    config.getLong(ConsumerConfig.SOCKET_CONNECTION_SETUP_TIMEOUT_MS_CONFIG), // 客户端等待套接字连接建立所需的时间。如果在超时之前没有建立连接，客户端将关闭套接字通道 默认10s
                    config.getLong(ConsumerConfig.SOCKET_CONNECTION_SETUP_TIMEOUT_MAX_MS_CONFIG), // 客户端等待套接字连接建立的最长时间。连接设置超时将随着每次连续连接失败而呈指数增长，直至此最大值。为了避免连接风暴，将对超时应用0.2的随机化因子，导致计算值低于20%和高于20%之间的随机范围。 默认 30s
                    ClientDnsLookup.forConfig(config.getString(ConsumerConfig.CLIENT_DNS_LOOKUP_CONFIG)), // 控制客户端如何使用DNS查找 默认use_all_dns_ips
                    time,
                    true,
                    apiVersions,
                    throttleTimeSensor,
                    logContext);

// -----------------------------------------------------------------------------------------------------
// 相当于就是一个请求工具
 this.client = new ConsumerNetworkClient(
                    logContext,         // 日志上下文
                    netClient,          // 网络客户端
                    metadata,           // 元数据
                    time,               // 时间
                    retryBackoffMs,     // 重试阻塞时间
                    config.getInt(ConsumerConfig.REQUEST_TIMEOUT_MS_CONFIG),  // 请求超时时间
                    heartbeatIntervalMs);  // 内部心跳时间
```

###### 作用

> consumer所有与kafka cluster的交互：
>
> 比如：建立连接、找Coordinator请求、心跳请求、加入组请求、离组请求、拉取数据请求、元数据请求

###### 核心方法

send(Node node, AbstractRequest.Builder<?> requestBuilder,int requestTimeoutMs)

```java
  public RequestFuture<ClientResponse> send(Node node,       // kafka cluster 节点
                                            AbstractRequest.Builder<?> requestBuilder,   // 请求信息
                                            int requestTimeoutMs) {     // 请求超时时间
      long now = time.milliseconds();
      RequestFutureCompletionHandler completionHandler = new RequestFutureCompletionHandler();
      // 构造请求
      ClientRequest clientRequest = client.newClientRequest(node.idString(), requestBuilder, now, true,
          requestTimeoutMs, completionHandler);
      // 将请求放到未发送的集合中
      unsent.put(node, clientRequest);

      // wakeup the client in case it is blocking in poll so that we can send the queued request
      client.wakeup(); // 唤醒客户端
      return completionHandler.future;
  }
```

> 调用send方法时，会将请求放到unsent集合中，等执行poll方法时，才会真正发送到kafka cluster 服务去

poll(Timer timer, PollCondition pollCondition, boolean disableWakeup)

```java
public void poll(Timer timer, PollCondition pollCondition, boolean disableWakeup) {
        // there may be handlers which need to be invoked if we woke up the previous call to poll
        firePendingCompletedRequests();
  
        lock.lock();
        try {
            // Handle async disconnects prior to attempting any sends
            handlePendingDisconnects();

            // 发送目前所有的请求
            long pollDelayMs = trySend(timer.currentTimeMs());

            // .... 省略了一些步骤

            // try again to send requests since buffer space may have been
            // cleared or a connect finished in the poll
            trySend(timer.currentTimeMs());

            // 处理失败过期的请求
            failExpiredRequests(timer.currentTimeMs());

            // 清空unsent集合
            unsent.clean();
        } finally {
            lock.unlock();
        }
    }
```

> 发送unsent里面的所有网络请求

##### ConsumerPartitionAssignor

###### 实例化

```java
// 分区分配器 默认是 org.apache.kafka.clients.consumer.RangeAssignor
this.assignors = getAssignorInstances(config.getList(ConsumerConfig.PARTITION_ASSIGNMENT_STRATEGY_CONFIG),
                    config.originals(Collections.singletonMap(ConsumerConfig.CLIENT_ID_CONFIG, clientId)));
```

###### 作用

> 当一个消费者做为组内的leader时，会为组内的所有消费者分配分区，通过分区分配器来分配
>
> RangeAssignor 范围分配 （ 默认）
>
> RoundRobinAssignor 轮询分配
>
> StickyAssignor  具有粘性分配



##### ConsumerCoordinator

###### 实例化

```java
this.coordinator = !groupId.isPresent() ? null :
                new ConsumerCoordinator(groupRebalanceConfig,   //组配置
                        logContext,                             // 日志上下文
                        this.client,                            // 消费者网络客户端工具
                        assignors,                              // 分区分配器
                        this.metadata,                          // 元数据
                        this.subscriptions,                     // 订阅者
                        metrics,                                // 指标
                        metricGrpPrefix,
                        this.time,
                        enableAutoCommit,                       // 是否自动提交
                        config.getInt(ConsumerConfig.AUTO_COMMIT_INTERVAL_MS_CONFIG), //自动提交时间
                        this.interceptors,                      // 拦截器
                        config.getBoolean(ConsumerConfig.THROW_ON_FETCH_STABLE_OFFSET_UNSUPPORTED));
```

###### 作用

> 与kafka cluster 的coordinator之间的交互：心跳、加入组、离组、位移提交、元数据等

###### 核心方法

poll(Timer timer, boolean waitForJoinGroup)

```java
/**
     * 拉取事件：确保coordinator可以连接、消费者加入组，同时开启了自动提交会提交offset
     *
     * 1、可能需要更新订阅者的元数据信息
     * 2、执行已经完成offset提交的回调方法
     * 3、判断订阅者是否是指定主题或者正则订阅
     *      1、判断coordinator是否可以连接
     *      2、是否需要加入组
     *           启动心跳线程
     *           消费者加入组 --- 处理响应时，如果自己是组的leader，那么就会执行分区分配，会将分配好的信息再发送给coordinator
     *  4、如果开启自动提交 且 当前时间已经超过了自动提交时间， 那么提交offset
     */
    public boolean poll(Timer timer, boolean waitForJoinGroup) {
        // 是否需要更新元数据
        maybeUpdateSubscriptionMetadata();
				// 执行已经完成提交的回调
        invokeCompletedOffsetCommitCallbacks();

        if (subscriptions.hasAutoAssignedPartitions()) {
            if (protocol == null) {
                throw new IllegalStateException("User configured " + ConsumerConfig.PARTITION_ASSIGNMENT_STRATEGY_CONFIG +
                    " to empty while trying to subscribe for group protocol to auto assign partitions");
            }
            // 更新poll的时间：用于心跳时的判断，如果更新时间超过了配置的时间，就会离开组
            pollHeartbeat(timer.currentTimeMs());
            // 未知的coordinator 并且 coordinator没有准备好
            if (coordinatorUnknown() && !ensureCoordinatorReady(timer)) {
                return false;
            }
						// 判断该消费者是否加入组了
            if (rejoinNeededOrPending()) {
                // 订阅主题的方式：是正则吗？
                if (subscriptions.hasPatternSubscription()) {
                    if (this.metadata.timeToAllowUpdate(timer.currentTimeMs()) == 0) {
                        this.metadata.requestUpdate();
                    }

                    if (!client.ensureFreshMetadata(timer)) {
                        return false;
                    }

                    maybeUpdateSubscriptionMetadata();
                }
                // 确保消费者组是激活状态
                // 启动心跳线程
                // 消费者加入组 --- 处理响应时，如果自己是组的leader，那么就会执行分区分配，会将分配好的信息再发送给coordinator
                if (!ensureActiveGroup(waitForJoinGroup ? timer : time.timer(0L))) {
                    timer.update(time.milliseconds());
                    return false;
                }
            }
        } else {
            if (metadata.updateRequested() && !client.hasReadyNodes(timer.currentTimeMs())) {
                client.awaitMetadataUpdate(timer);
            }
        }
        // 自动提交：开启自动提交且超过了下一次提交的时间，才会提交offset
        maybeAutoCommitOffsetsAsync(timer.currentTimeMs());
        return true;
    }
```

commitOffsetsAsync 异步提交

```java
public void commitOffsetsAsync(final Map<TopicPartition, OffsetAndMetadata> offsets, final OffsetCommitCallback callback) {
        // 执行已经完成的commit的回调
        invokeCompletedOffsetCommitCallbacks();
        // 向coordinator提交
        if (!coordinatorUnknown()) {
            doCommitOffsetsAsync(offsets, callback);
        } else {
            pendingAsyncCommits.incrementAndGet();
            // 找Coordinator
            lookupCoordinator().addListener(new RequestFutureListener<Void>() {
                @Override
                public void onSuccess(Void value) {
                    pendingAsyncCommits.decrementAndGet();
                    doCommitOffsetsAsync(offsets, callback); // 获取coordinator后，再次提交
                    client.pollNoWakeup();
                }

                @Override
                public void onFailure(RuntimeException e) {
                    pendingAsyncCommits.decrementAndGet();
                    completedOffsetCommits.add(new OffsetCommitCompletion(callback, offsets,
                            new RetriableCommitFailedException(e)));
                }
            });
        }
        // 真正的触发提交请求
        client.pollNoWakeup();
    }
```

Map<TopicPartition, OffsetAndMetadata> fetchCommittedOffsets(final Set<TopicPartition> partitions,final Timer timer)  获取分区位移信息

maybeLeaveGroup(String leaveReason) 离开组（其父AbstractCoordinator中定义）

```java
public synchronized RequestFuture<Void> maybeLeaveGroup(String leaveReason) {
        RequestFuture<Void> future = null;

        if (isDynamicMember() && !coordinatorUnknown() &&
            state != MemberState.UNJOINED && generation.hasMemberId()) {
           
            log.info("Member {} sending LeaveGroup request to coordinator {} due to {}",
                generation.memberId, coordinator, leaveReason);
            // 离组请求构造器
            LeaveGroupRequest.Builder request = new LeaveGroupRequest.Builder(
                rebalanceConfig.groupId,   // 组id
                Collections.singletonList(new MemberIdentity().setMemberId(generation.memberId)) // 成员id
            );
            // 发送离组请求到unsent中，并设置响应处理器
            future = client.send(coordinator, request).compose(new LeaveGroupResponseHandler(generation));
            // 将unsent中的请求发送到node上
            client.pollNoWakeup();
        }
        /**
         * 重置状态
         */
        resetGenerationOnLeaveGroup();

        return future;
    }
```

HeartbeatThread 心跳线程 （其父AbstractCoordinator中定义）

```java
public void run() {
            try {
                log.debug("Heartbeat thread started");
                while (true) {
                    synchronized (AbstractCoordinator.this) {
                        if (closed)
                            return;

                        if (!enabled) {
                            AbstractCoordinator.this.wait();
                            continue;
                        }

                        // 没有加入组：重试
                        if (state.hasNotJoinedGroup() || hasFailed()) {
                            disable();
                            continue;
                        }

                        client.pollNoWakeup();
                        long now = time.milliseconds();

                        // 未知的coordinator
                        if (coordinatorUnknown()) {
                            if (findCoordinatorFuture != null) {
                                clearFindCoordinatorFuture();
                                // backoff properly
                                AbstractCoordinator.this.wait(rebalanceConfig.retryBackoffMs);
                            } else {
                                lookupCoordinator(); // 找coordinator
                            }
                        } else if (heartbeat.sessionTimeoutExpired(now)) {  // 没有收到心跳响应
                            markCoordinatorUnknown("session timed out without receiving a "
                                    + "heartbeat response");
                        } else if (heartbeat.pollTimeoutExpired(now)) {   // poll间隔时间太长，导致失效
                            String leaveReason = "consumer poll timeout has expired. This means the time between subsequent calls to poll() " +
                                                    "was longer than the configured max.poll.interval.ms, which typically implies that " +
                                                    "the poll loop is spending too much time processing messages. " +
                                                    "You can address this either by increasing max.poll.interval.ms or by reducing " +
                                                    "the maximum size of batches returned in poll() with max.poll.records.";
                            // 离开组
                            maybeLeaveGroup(leaveReason);
                        } else if (!heartbeat.shouldHeartbeat(now)) { // 更新心跳时间
                            // poll again after waiting for the retry backoff in case the heartbeat failed or the
                            // coordinator disconnected
                            AbstractCoordinator.this.wait(rebalanceConfig.retryBackoffMs);
                        } else {
                            // 更新时间
                            heartbeat.sentHeartbeat(now);
                            // 发送心跳
                            final RequestFuture<Void> heartbeatFuture = sendHeartbeatRequest();
                            // 响应回调
                            heartbeatFuture.addListener(new RequestFutureListener<Void>() {
                                @Override
                                public void onSuccess(Void value) {
                                    synchronized (AbstractCoordinator.this) {
                                        heartbeat.receiveHeartbeat(); // 接收到心跳的响应
                                    }
                                }

                                @Override
                                public void onFailure(RuntimeException e) {
                                    synchronized (AbstractCoordinator.this) {
                                        if (e instanceof RebalanceInProgressException) {
                                            // it is valid to continue heartbeating while the group is rebalancing. This
                                            // ensures that the coordinator keeps the member in the group for as long
                                            // as the duration of the rebalance timeout. If we stop sending heartbeats,
                                            // however, then the session timeout may expire before we can rejoin.
                                            heartbeat.receiveHeartbeat();
                                        } else if (e instanceof FencedInstanceIdException) {
                                            log.error("Caught fenced group.instance.id {} error in heartbeat thread", rebalanceConfig.groupInstanceId);
                                            heartbeatThread.failed.set(e);
                                        } else {
                                            heartbeat.failHeartbeat();
                                            // wake up the thread if it's sleeping to reschedule the heartbeat
                                            AbstractCoordinator.this.notify();
                                        }
                                    }
                                }
                            });
                        }
                    }
                }
            } catch (AuthenticationException e) {
                log.error("An authentication error occurred in the heartbeat thread", e);
                this.failed.set(e);
            } catch (GroupAuthorizationException e) {
                log.error("A group authorization error occurred in the heartbeat thread", e);
                this.failed.set(e);
            } catch (InterruptedException | InterruptException e) {
                Thread.interrupted();
                log.error("Unexpected interrupt received in heartbeat thread", e);
                this.failed.set(new RuntimeException(e));
            } catch (Throwable e) {
                log.error("Heartbeat thread failed due to unexpected error", e);
                if (e instanceof RuntimeException)
                    this.failed.set((RuntimeException) e);
                else
                    this.failed.set(new RuntimeException(e));
            } finally {
                log.debug("Heartbeat thread has closed");
            }
        }
```



##### Fetcher

###### 实例化

```java
this.fetcher = new Fetcher<>(
                    logContext,                // 日志上下文
                    this.client,               // 消费者网络客户端工具
                    config.getInt(ConsumerConfig.FETCH_MIN_BYTES_CONFIG),     // 拿取最小的byte大小
                    config.getInt(ConsumerConfig.FETCH_MAX_BYTES_CONFIG),     // 拿取最大的byte大小
                    config.getInt(ConsumerConfig.FETCH_MAX_WAIT_MS_CONFIG),   // 最大的等待时间
                    config.getInt(ConsumerConfig.MAX_PARTITION_FETCH_BYTES_CONFIG),  // 分区获取的最大byte大小
                    config.getInt(ConsumerConfig.MAX_POLL_RECORDS_CONFIG),    // 默认一次性最多拿500条消息记录
                    config.getBoolean(ConsumerConfig.CHECK_CRCS_CONFIG),
                    config.getString(ConsumerConfig.CLIENT_RACK_CONFIG),
                    this.keyDeserializer,            // key 反序列化器
                    this.valueDeserializer,          // value 反序列化器
                    this.metadata,                  // 元数据
                    this.subscriptions,             // 订阅者信息
                    metrics,                        // 指标
                    metricsRegistry,
                    this.time,
                    this.retryBackoffMs,            // 重试阻塞时间
                    this.requestTimeoutMs,          // 请求超时时间
                    isolationLevel,                 // 隔离级别
                    apiVersions);                   // api版本
```

###### 作用

> 从 kafka cluster 获取消息

###### 核心方法

sendFetches()  发送获取message请求到unsent中

```java
public synchronized int sendFetches() {
        // Update metrics in case there was an assignment change
        sensors.maybeUpdateAssignment(subscriptions);
        // 构造可发送分区的fetch请求：指定了分区的偏移、分区的fetch.size
        Map<Node, FetchSessionHandler.FetchRequestData> fetchRequestMap = prepareFetchRequests();
        // 循环处理准备好的fetch请求数据
        for (Map.Entry<Node, FetchSessionHandler.FetchRequestData> entry : fetchRequestMap.entrySet()) {
            // 节点信息
            final Node fetchTarget = entry.getKey();
            final FetchSessionHandler.FetchRequestData data = entry.getValue();
            /**
             * 构造请求对象
             * 注意： 这里是基于node来的，可能这个请求的是node上面的多个分区中的数据
             */
            final FetchRequest.Builder request = FetchRequest.Builder
                    .forConsumer(this.maxWaitMs, this.minBytes, data.toSend())    // consumer信息   data.toSend() 返回的是Map<TopicPartition, PartitionData>  PartitionData记录了分区的起始偏移
                    .isolationLevel(isolationLevel)                               // 隔离级别
                    .setMaxBytes(this.maxBytes)                                   // 最大的限制
                    .metadata(data.metadata())                                    // 元数据
                    .toForget(data.toForget())                                    // 目标主题分区
                    .rackId(clientRackId);                                        // 客户端机架id

            log.debug("Sending {} {} to broker {}", isolationLevel, data, fetchTarget);

            /**
             * send请求
             * 1、封装clientRequest
             * 2、unsent.put(node, clientRequest); 将请求放到未发送的请求集合中
             */
            RequestFuture<ClientResponse> future = client.send(fetchTarget, request);
          
            this.nodesWithPendingFetchRequests.add(entry.getKey().id());
            /**
             * 添加监听器，用来处理响应成功或者失败的后置操作
             * 注意：
             *      都是单个节点里面的分区数据
             *
             * 成功：
             *   将返回的数据封装成CompletedFetch，放入到completedFetches集合中
             *   completedFetches.add(new CompletedFetch(partition, partitionData,
             *                                             metricAggregator, batches, fetchOffset, responseVersion));
             */
            future.addListener(new RequestFutureListener<ClientResponse>() {
                @Override
                public void onSuccess(ClientResponse resp) {
                    synchronized (Fetcher.this) {
                        try {
                            @SuppressWarnings("unchecked")
                            FetchResponse<Records> response = (FetchResponse<Records>) resp.responseBody();
                            FetchSessionHandler handler = sessionHandler(fetchTarget.id());
                            if (handler == null) {
                                log.error("Unable to find FetchSessionHandler for node {}. Ignoring fetch response.",
                                        fetchTarget.id());
                                return;
                            }
                            // handler不能处理，也直接返回
                            if (!handler.handleResponse(response)) {
                                return;
                            }
                            // 获取主题分区集合
                            Set<TopicPartition> partitions = new HashSet<>(response.responseData().keySet());
                            FetchResponseMetricAggregator metricAggregator = new FetchResponseMetricAggregator(sensors, partitions);

                            // 遍历响应结果
                            for (Map.Entry<TopicPartition, FetchResponse.PartitionData<Records>> entry : response.responseData().entrySet()) {
                                // 主题分区
                                TopicPartition partition = entry.getKey();
                                // 获取主题分区的请求数据
                                FetchRequest.PartitionData requestData = data.sessionPartitions().get(partition);
                                if (requestData == null) {
                                    String message;
                                    if (data.metadata().isFull()) {
                                        message = MessageFormatter.arrayFormat(
                                                "Response for missing full request partition: partition={}; metadata={}",
                                                new Object[]{partition, data.metadata()}).getMessage();
                                    } else {
                                        message = MessageFormatter.arrayFormat(
                                                "Response for missing session request partition: partition={}; metadata={}; toSend={}; toForget={}",
                                                new Object[]{partition, data.metadata(), data.toSend(), data.toForget()}).getMessage();
                                    }

                                    // Received fetch response for missing session partition
                                    throw new IllegalStateException(message);
                                } else {
                                    // 请求数据里面的offset，也就是开始offset
                                    long fetchOffset = requestData.fetchOffset;
                                    // 响应数据
                                    FetchResponse.PartitionData<Records> partitionData = entry.getValue();

                                    log.debug("Fetch {} at offset {} for partition {} returned fetch data {}",
                                            isolationLevel, fetchOffset, partition, partitionData);

                                    Iterator<? extends RecordBatch> batches = partitionData.records().batches().iterator();
                                    short responseVersion = resp.requestHeader().apiVersion();
                                    // 成功了，就往completedFetches添加对象
                                    completedFetches.add(new CompletedFetch(partition, partitionData,
                                            metricAggregator, batches, fetchOffset, responseVersion));
                                }
                            }

                            sensors.fetchLatency.record(resp.requestLatencyMs());
                        } finally {
                            nodesWithPendingFetchRequests.remove(fetchTarget.id());
                        }
                    }
                }

                @Override
                public void onFailure(RuntimeException e) {
                    synchronized (Fetcher.this) {
                        try {
                            FetchSessionHandler handler = sessionHandler(fetchTarget.id());
                            if (handler != null) {
                                handler.handleError(e);
                            }
                        } finally {
                            nodesWithPendingFetchRequests.remove(fetchTarget.id());
                        }
                    }
                }
            });

        }
        return fetchRequestMap.size();
    }
```

fetchedRecords()  获取message记录

```java
public Map<TopicPartition, List<ConsumerRecord<K, V>>> fetchedRecords() {
        Map<TopicPartition, List<ConsumerRecord<K, V>>> fetched = new HashMap<>();
        Queue<CompletedFetch> pausedCompletedFetches = new ArrayDeque<>();
        int recordsRemaining = maxPollRecords;

        try {
            while (recordsRemaining > 0) {
                // 如果nextInLineFetch为空，或者说已经被消费了，那么就去找下一个不为空，且没有被消费的已经完成请求消息信息
                if (nextInLineFetch == null || nextInLineFetch.isConsumed) {
                    CompletedFetch records = completedFetches.peek();
                    if (records == null) break;  // 一个完成的都没有，直接退出

                    if (records.notInitialized()) { // 如果没有初始化，就初始化一下
                        try {
                            nextInLineFetch = initializeCompletedFetch(records);
                        } catch (Exception e) {
                            FetchResponse.PartitionData<Records> partition = records.partitionData;
                            if (fetched.isEmpty() && (partition.records() == null || partition.records().sizeInBytes() == 0)) {
                                completedFetches.poll();
                            }
                            throw e;
                        }
                    } else {
                        nextInLineFetch = records;
                    }
                    // 出栈
                    completedFetches.poll();
                } else if (subscriptions.isPaused(nextInLineFetch.partition)) {
                    log.debug("Skipping fetching records for assigned partition {} because it is paused", nextInLineFetch.partition);
                    pausedCompletedFetches.add(nextInLineFetch);
                    nextInLineFetch = null;
                } else {
                     /**
                     * 获取message记录
                     * 将已获取消息的分区的offset更新到subscriptions中
                     */
                    List<ConsumerRecord<K, V>> records = fetchRecords(nextInLineFetch, recordsRemaining);

                    if (!records.isEmpty()) {
                        // 主题分区
                        TopicPartition partition = nextInLineFetch.partition;

                        // 将获取的records放入到已获取的map中
                        List<ConsumerRecord<K, V>> currentRecords = fetched.get(partition);
                        if (currentRecords == null) {
                            fetched.put(partition, records);
                        } else {
                            List<ConsumerRecord<K, V>> newRecords = new ArrayList<>(records.size() + currentRecords.size());
                            newRecords.addAll(currentRecords);
                            newRecords.addAll(records);
                            fetched.put(partition, newRecords);
                        }
                        // 迭代
                        recordsRemaining -= records.size();
                    }
                }
            }
        } catch (KafkaException e) {
            if (fetched.isEmpty())
                throw e;
        } finally {
            // add any polled completed fetches for paused partitions back to the completed fetches queue to be
            // re-evaluated in the next poll
            completedFetches.addAll(pausedCompletedFetches);
        }

        return fetched;
    }
```



## 二、server

