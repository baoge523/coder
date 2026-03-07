package main

import "fmt"

// Target 目标接口 - 5V电压
type Target interface {
	Request() string
}

// Adaptee 被适配者 - 220V电压
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "220V 电压"
}

// Adapter 适配器
type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	voltage := a.adaptee.SpecificRequest()
	return fmt.Sprintf("适配器将 %s 转换为 5V 电压", voltage)
}

// 实际应用示例：支付接口适配

// PaymentProcessor 统一支付接口
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// AlipayAPI 支付宝原始API
type AlipayAPI struct{}

func (a *AlipayAPI) SendPayment(money float64) string {
	return fmt.Sprintf("支付宝支付: ¥%.2f", money)
}

// WechatAPI 微信原始API
type WechatAPI struct{}

func (w *WechatAPI) Pay(amount float64) string {
	return fmt.Sprintf("微信支付: ¥%.2f", amount)
}

// AlipayAdapter 支付宝适配器
type AlipayAdapter struct {
	alipay *AlipayAPI
}

func NewAlipayAdapter() *AlipayAdapter {
	return &AlipayAdapter{alipay: &AlipayAPI{}}
}

func (a *AlipayAdapter) ProcessPayment(amount float64) string {
	return a.alipay.SendPayment(amount)
}

// WechatAdapter 微信适配器
type WechatAdapter struct {
	wechat *WechatAPI
}

func NewWechatAdapter() *WechatAdapter {
	return &WechatAdapter{wechat: &WechatAPI{}}
}

func (w *WechatAdapter) ProcessPayment(amount float64) string {
	return w.wechat.Pay(amount)
}

// PaymentService 支付服务
type PaymentService struct {
	processor PaymentProcessor
}

func NewPaymentService(processor PaymentProcessor) *PaymentService {
	return &PaymentService{processor: processor}
}

func (p *PaymentService) MakePayment(amount float64) {
	result := p.processor.ProcessPayment(amount)
	fmt.Println(result)
}

func main() {
	fmt.Println("=== 适配器模式 ===\n")
	
	// 电源适配器示例
	fmt.Println("电源适配器:")
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)
	fmt.Println(adapter.Request())
	
	// 支付接口适配示例
	fmt.Println("\n支付接口适配:")
	
	// 使用支付宝
	alipayAdapter := NewAlipayAdapter()
	service := NewPaymentService(alipayAdapter)
	service.MakePayment(100.00)
	
	// 使用微信
	wechatAdapter := NewWechatAdapter()
	service = NewPaymentService(wechatAdapter)
	service.MakePayment(200.00)
}
