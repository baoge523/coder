#!/bin/bash

# Go 设计模式示例运行脚本

echo "=========================================="
echo "Go 设计模式示例演示"
echo "=========================================="

# 创建型模式
echo -e "\n\n【创建型模式】\n"

echo ">>> 1. 单例模式 (Singleton)"
cd singleton && go run singleton.go
cd ..

echo -e "\n>>> 2. 工厂模式 (Factory)"
cd factory && go run factory.go
cd ..

echo -e "\n>>> 3. 建造者模式 (Builder)"
cd builder && go run builder.go
cd ..

echo -e "\n>>> 4. 原型模式 (Prototype)"
cd prototype && go run prototype.go
cd ..

# 结构型模式
echo -e "\n\n【结构型模式】\n"

echo ">>> 5. 适配器模式 (Adapter)"
cd adapter && go run adapter.go
cd ..

echo -e "\n>>> 6. 装饰器模式 (Decorator)"
cd decorator && go run decorator.go
cd ..

echo -e "\n>>> 7. 代理模式 (Proxy)"
cd proxy && go run proxy.go
cd ..

echo -e "\n>>> 8. 外观模式 (Facade)"
cd facade && go run facade.go
cd ..

# 行为型模式
echo -e "\n\n【行为型模式】\n"

echo ">>> 9. 策略模式 (Strategy)"
cd strategy && go run strategy.go
cd ..

echo -e "\n>>> 10. 观察者模式 (Observer)"
cd observer && go run observer.go
cd ..

echo -e "\n>>> 11. 责任链模式 (Chain of Responsibility)"
cd chain && go run chain.go
cd ..

echo -e "\n>>> 12. 模板方法模式 (Template Method)"
cd template && go run template.go
cd ..

echo -e "\n>>> 13. 命令模式 (Command)"
cd command && go run command.go
cd ..

echo -e "\n\n=========================================="
echo "所有示例运行完成！"
echo "=========================================="
