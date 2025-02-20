package strategy

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(*PaymentContext)
}

// PaymentContext 支付上下文
type PaymentContext struct {
	Name, CardID string
	Money        int
}

// Payment 支付实体
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

// NewPayment 构造函数
func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

// Pay 支付实体的支付方法
func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}
