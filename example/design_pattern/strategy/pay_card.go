package strategy

import "fmt"

// Bank 银行卡支付
type Bank struct{}

// Pay 银行卡支付
func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}
