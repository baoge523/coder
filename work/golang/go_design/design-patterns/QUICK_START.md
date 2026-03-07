# 快速开始指南

## 环境要求
- Go 1.21 或更高版本

## 运行示例

### 1. 初始化模块
```bash
cd design-patterns
go mod tidy
```

### 2. 运行单个设计模式示例

#### 创建型模式
```bash
# 单例模式
cd singleton && go run singleton.go

# 工厂模式
cd factory && go run factory.go

# 建造者模式
cd builder && go run builder.go

# 原型模式
cd prototype && go run prototype.go
```

#### 结构型模式
```bash
# 适配器模式
cd adapter && go run adapter.go

# 装饰器模式
cd decorator && go run decorator.go

# 代理模式
cd proxy && go run proxy.go

# 外观模式
cd facade && go run facade.go
```

#### 行为型模式
```bash
# 策略模式
cd strategy && go run strategy.go

# 观察者模式
cd observer && go run observer.go

# 责任链模式
cd chain && go run chain.go

# 模板方法模式
cd template && go run template.go

# 命令模式
cd command && go run command.go
```

### 3. 运行测试
```bash
# 运行单例模式测试
cd singleton && go test -v

# 运行所有测试（如果有）
go test ./... -v
```

## 学习路径建议

### 初学者
1. 单例模式 (Singleton) - 最简单，理解单例概念
2. 工厂模式 (Factory) - 理解对象创建
3. 策略模式 (Strategy) - 理解算法封装
4. 观察者模式 (Observer) - 理解事件通知

### 进阶
1. 建造者模式 (Builder) - 复杂对象构建
2. 装饰器模式 (Decorator) - 动态功能扩展
3. 适配器模式 (Adapter) - 接口转换
4. 代理模式 (Proxy) - 访问控制

### 高级
1. 责任链模式 (Chain of Responsibility) - 请求处理链
2. 命令模式 (Command) - 请求封装
3. 模板方法模式 (Template Method) - 算法骨架
4. 外观模式 (Facade) - 子系统封装

## 实际应用场景

### Web 开发
- 中间件：责任链模式、装饰器模式
- 路由：策略模式
- 数据库连接：单例模式
- API 客户端：适配器模式

### 微服务
- 服务发现：观察者模式
- 配置管理：单例模式
- 请求处理：责任链模式
- 服务代理：代理模式

### 系统设计
- 日志系统：单例模式、装饰器模式
- 缓存系统：代理模式
- 消息队列：观察者模式
- 工作流引擎：责任链模式、命令模式

## 常见问题

### Q: Go 语言没有继承，如何实现设计模式？
A: Go 使用接口和组合来实现设计模式，这种方式更加灵活。

### Q: 哪些设计模式在 Go 中最常用？
A: 单例、工厂、建造者、策略、装饰器、适配器最常用。

### Q: 如何选择合适的设计模式？
A: 根据具体问题选择：
- 对象创建问题 → 创建型模式
- 对象组合问题 → 结构型模式
- 对象交互问题 → 行为型模式

## 扩展阅读
- 《设计模式：可复用面向对象软件的基础》
- 《Head First 设计模式》
- Go 官方文档：https://golang.org/doc/
