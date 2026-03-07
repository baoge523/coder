package main

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(message string)
	GetID() string
}

// Subject 主题接口
type Subject interface {
	Attach(observer Observer)
	Detach(observerID string)
	Notify(message string)
}

// ConcreteObserver 具体观察者
type ConcreteObserver struct {
	id   string
	name string
}

func NewObserver(id, name string) *ConcreteObserver {
	return &ConcreteObserver{id: id, name: name}
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] 收到通知: %s\n", o.name, message)
}

func (o *ConcreteObserver) GetID() string {
	return o.id
}

// ConcreteSubject 具体主题
type ConcreteSubject struct {
	observers map[string]Observer
}

func NewSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make(map[string]Observer),
	}
}

func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers[observer.GetID()] = observer
	fmt.Printf("观察者 %s 已订阅\n", observer.GetID())
}

func (s *ConcreteSubject) Detach(observerID string) {
	delete(s.observers, observerID)
	fmt.Printf("观察者 %s 已取消订阅\n", observerID)
}

func (s *ConcreteSubject) Notify(message string) {
	fmt.Printf("\n发送通知: %s\n", message)
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// 实际应用示例：新闻订阅系统
type NewsPublisher struct {
	*ConcreteSubject
	newsTitle string
}

func NewNewsPublisher() *NewsPublisher {
	return &NewsPublisher{
		ConcreteSubject: NewSubject(),
	}
}

func (n *NewsPublisher) PublishNews(title string) {
	n.newsTitle = title
	n.Notify(fmt.Sprintf("新闻发布: %s", title))
}

func main() {
	fmt.Println("=== 观察者模式 - 新闻订阅系统 ===\n")
	
	// 创建新闻发布者
	publisher := NewNewsPublisher()
	
	// 创建订阅者
	user1 := NewObserver("001", "张三")
	user2 := NewObserver("002", "李四")
	user3 := NewObserver("003", "王五")
	
	// 订阅新闻
	publisher.Attach(user1)
	publisher.Attach(user2)
	publisher.Attach(user3)
	
	// 发布新闻
	publisher.PublishNews("Go 1.22 正式发布！")
	
	// 取消订阅
	fmt.Println()
	publisher.Detach("002")
	
	// 再次发布新闻
	publisher.PublishNews("Kubernetes 新版本更新")
}
