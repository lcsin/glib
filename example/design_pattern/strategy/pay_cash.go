package strategy

import "fmt"

// Cash 现金支付
type Cash struct{}

// Pay 现金支付
func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}
