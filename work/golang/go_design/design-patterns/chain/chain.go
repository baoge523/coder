package main

import "fmt"

// Request 请求
type Request struct {
	Level   int
	Content string
}

// Handler 处理器接口
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request *Request)
}

// BaseHandler 基础处理器
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) HandleNext(request *Request) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// TeamLeaderHandler 组长处理器
type TeamLeaderHandler struct {
	BaseHandler
}

func (t *TeamLeaderHandler) Handle(request *Request) {
	if request.Level <= 1 {
		fmt.Printf("组长处理请求: %s (级别: %d)\n", request.Content, request.Level)
	} else {
		fmt.Println("组长: 权限不足，转交上级")
		t.HandleNext(request)
	}
}

// ManagerHandler 经理处理器
type ManagerHandler struct {
	BaseHandler
}

func (m *ManagerHandler) Handle(request *Request) {
	if request.Level <= 2 {
		fmt.Printf("经理处理请求: %s (级别: %d)\n", request.Content, request.Level)
	} else {
		fmt.Println("经理: 权限不足，转交上级")
		m.HandleNext(request)
	}
}

// DirectorHandler 总监处理器
type DirectorHandler struct {
	BaseHandler
}

func (d *DirectorHandler) Handle(request *Request) {
	if request.Level <= 3 {
		fmt.Printf("总监处理请求: %s (级别: %d)\n", request.Content, request.Level)
	} else {
		fmt.Println("总监: 权限不足，转交上级")
		d.HandleNext(request)
	}
}

// CEOHandler CEO处理器
type CEOHandler struct {
	BaseHandler
}

func (c *CEOHandler) Handle(request *Request) {
	fmt.Printf("CEO处理请求: %s (级别: %d)\n", request.Content, request.Level)
}

func main() {
	fmt.Println("=== 责任链模式 - 请假审批系统 ===\n")
	
	// 创建处理器
	teamLeader := &TeamLeaderHandler{}
	manager := &ManagerHandler{}
	director := &DirectorHandler{}
	ceo := &CEOHandler{}
	
	// 构建责任链
	teamLeader.SetNext(manager).SetNext(director).SetNext(ceo)
	
	// 测试不同级别的请求
	requests := []*Request{
		{Level: 1, Content: "请假1天"},
		{Level: 2, Content: "请假3天"},
		{Level: 3, Content: "请假7天"},
		{Level: 4, Content: "请假30天"},
	}
	
	for i, req := range requests {
		fmt.Printf("\n请求 %d:\n", i+1)
		teamLeader.Handle(req)
	}
}
