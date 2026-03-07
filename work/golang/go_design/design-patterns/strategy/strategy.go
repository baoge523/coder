package main

import "fmt"

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// AlipayStrategy 支付宝支付
type AlipayStrategy struct {
	account string
}

func NewAlipayStrategy(account string) *AlipayStrategy {
	return &AlipayStrategy{account: account}
}

func (a *AlipayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用支付宝账户 %s 支付 %.2f 元", a.account, amount)
}

// WechatStrategy 微信支付
type WechatStrategy struct {
	account string
}

func NewWechatStrategy(account string) *WechatStrategy {
	return &WechatStrategy{account: account}
}

func (w *WechatStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用微信账户 %s 支付 %.2f 元", w.account, amount)
}

// CreditCardStrategy 信用卡支付
type CreditCardStrategy struct {
	cardNumber string
}

func NewCreditCardStrategy(cardNumber string) *CreditCardStrategy {
	return &CreditCardStrategy{cardNumber: cardNumber}
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用信用卡 %s 支付 %.2f 元", c.cardNumber, amount)
}

// PaymentContext 支付上下文
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) string {
	return p.strategy.Pay(amount)
}

func main() {
	fmt.Println("=== 策略模式 - 支付系统 ===\n")
	
	amount := 199.99
	
	// 使用支付宝支付
	alipay := NewAlipayStrategy("user@alipay.com")
	context := NewPaymentContext(alipay)
	fmt.Println(context.ExecutePayment(amount))
	
	// 切换到微信支付
	wechat := NewWechatStrategy("user_wechat_id")
	context.SetStrategy(wechat)
	fmt.Println(context.ExecutePayment(amount))
	
	// 切换到信用卡支付
	creditCard := NewCreditCardStrategy("**** **** **** 1234")
	context.SetStrategy(creditCard)
	fmt.Println(context.ExecutePayment(amount))
}
