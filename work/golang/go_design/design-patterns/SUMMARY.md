# Go 设计模式总结

## 项目概述
本项目包含 13 个常用的 Go 语言设计模式实现，每个模式都包含：
- 详细的说明文档 (README.md)
- 可运行的示例代码
- 实际应用场景

## 设计模式分类

### 一、创建型模式 (Creational Patterns)
创建型模式关注对象的创建机制，试图以适当的方式创建对象。

#### 1. 单例模式 (Singleton)
- **目的**：确保一个类只有一个实例
- **实现**：使用 `sync.Once` 保证线程安全
- **应用**：数据库连接池、配置管理器、日志记录器
- **文件**：`singleton/singleton.go`

#### 2. 工厂模式 (Factory)
- **目的**：定义创建对象的接口
- **实现**：简单工厂和工厂方法
- **应用**：对象创建逻辑复杂时
- **文件**：`factory/factory.go`

#### 3. 建造者模式 (Builder)
- **目的**：分步骤构建复杂对象
- **实现**：链式调用
- **应用**：配置对象、复杂实体构建
- **文件**：`builder/builder.go`

#### 4. 原型模式 (Prototype)
- **目的**：通过克隆创建对象
- **实现**：深拷贝和浅拷贝
- **应用**：对象创建成本高时
- **文件**：`prototype/prototype.go`

### 二、结构型模式 (Structural Patterns)
结构型模式关注类和对象的组合，形成更大的结构。

#### 5. 适配器模式 (Adapter)
- **目的**：转换接口使其兼容
- **实现**：包装不兼容的接口
- **应用**：第三方库集成、接口统一
- **文件**：`adapter/adapter.go`

#### 6. 装饰器模式 (Decorator)
- **目的**：动态添加功能
- **实现**：包装对象并扩展功能
- **应用**：中间件、功能增强
- **文件**：`decorator/decorator.go`

#### 7. 代理模式 (Proxy)
- **目的**：控制对象访问
- **实现**：代理对象控制真实对象
- **应用**：延迟加载、访问控制、缓存
- **文件**：`proxy/proxy.go`

#### 8. 外观模式 (Facade)
- **目的**：简化复杂系统接口
- **实现**：提供统一的高层接口
- **应用**：子系统封装、API 简化
- **文件**：`facade/facade.go`

### 三、行为型模式 (Behavioral Patterns)
行为型模式关注对象之间的通信和职责分配。

#### 9. 策略模式 (Strategy)
- **目的**：算法可以互相替换
- **实现**：定义算法族并封装
- **应用**：支付方式、排序算法
- **文件**：`strategy/strategy.go`

#### 10. 观察者模式 (Observer)
- **目的**：一对多依赖关系
- **实现**：发布-订阅机制
- **应用**：事件系统、消息通知
- **文件**：`observer/observer.go`

#### 11. 责任链模式 (Chain of Responsibility)
- **目的**：多个对象处理请求
- **实现**：处理器链
- **应用**：审批流程、中间件链
- **文件**：`chain/chain.go`

#### 12. 模板方法模式 (Template Method)
- **目的**：定义算法骨架
- **实现**：抽象类定义步骤
- **应用**：数据处理流程、测试框架
- **文件**：`template/template.go`

#### 13. 命令模式 (Command)
- **目的**：将请求封装为对象
- **实现**：命令对象封装操作
- **应用**：撤销/重做、事务操作
- **文件**：`command/command.go`

## 使用统计

### 最常用的模式（推荐优先学习）
1. **单例模式** - 几乎所有项目都会用到
2. **工厂模式** - 对象创建的标准方式
3. **策略模式** - 算法封装的最佳实践
4. **装饰器模式** - Go 中间件的基础
5. **适配器模式** - 接口集成必备

### Go 语言特色实现
- **接口优先**：Go 的隐式接口实现使模式更灵活
- **组合优于继承**：使用嵌入而非继承
- **函数式特性**：函数作为一等公民简化某些模式
- **并发支持**：Channel 和 Goroutine 增强观察者模式

## 学习建议

### 初学者路径
1. 从单例模式开始，理解基本概念
2. 学习工厂模式，掌握对象创建
3. 实践策略模式，理解算法封装
4. 尝试装饰器模式，体验功能扩展

### 进阶路径
1. 深入建造者模式，构建复杂对象
2. 掌握适配器模式，处理接口兼容
3. 学习代理模式，实现访问控制
4. 研究责任链模式，构建处理链

### 实战建议
1. **不要过度设计**：只在需要时使用设计模式
2. **理解原理**：知道为什么用比知道怎么用更重要
3. **结合实际**：在真实项目中应用才能深刻理解
4. **持续重构**：随着需求变化调整设计

## 项目结构
```
design-patterns/
├── README.md                    # 项目总览
├── QUICK_START.md              # 快速开始指南
├── SUMMARY.md                  # 本文件
├── PATTERNS_COMPARISON.md      # 模式对比
├── go.mod                      # Go 模块文件
├── run_all.sh                  # 运行所有示例脚本
│
├── singleton/                  # 单例模式
│   ├── README.md
│   ├── singleton.go
│   └── singleton_test.go
│
├── factory/                    # 工厂模式
│   ├── README.md
│   └── factory.go
│
├── builder/                    # 建造者模式
│   ├── README.md
│   └── builder.go
│
├── prototype/                  # 原型模式
│   ├── README.md
│   └── prototype.go
│
├── adapter/                    # 适配器模式
│   ├── README.md
│   └── adapter.go
│
├── decorator/                  # 装饰器模式
│   ├── README.md
│   └── decorator.go
│
├── proxy/                      # 代理模式
│   ├── README.md
│   └── proxy.go
│
├── facade/                     # 外观模式
│   ├── README.md
│   └── facade.go
│
├── strategy/                   # 策略模式
│   ├── README.md
│   └── strategy.go
│
├── observer/                   # 观察者模式
│   ├── README.md
│   └── observer.go
│
├── chain/                      # 责任链模式
│   ├── README.md
│   └── chain.go
│
├── template/                   # 模板方法模式
│   ├── README.md
│   └── template.go
│
└── command/                    # 命令模式
    ├── README.md
    └── command.go
```

## 运行示例

### 运行单个模式
```bash
cd design-patterns/singleton
go run singleton.go
```

### 运行所有模式
```bash
cd design-patterns
chmod +x run_all.sh
./run_all.sh
```

### 运行测试
```bash
cd design-patterns/singleton
go test -v
```

## 扩展资源

### 推荐书籍
- 《设计模式：可复用面向对象软件的基础》（GoF）
- 《Head First 设计模式》
- 《Go 语言设计模式》

### 在线资源
- Go 官方文档：https://golang.org/doc/
- Go by Example：https://gobyexample.com/
- Refactoring Guru：https://refactoring.guru/design-patterns

## 贡献指南
欢迎提交 Issue 和 Pull Request 来改进本项目！

## 许可证
MIT License

---

**最后更新**：2024年
**Go 版本**：1.21+
