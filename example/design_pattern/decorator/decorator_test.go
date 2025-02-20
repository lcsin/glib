package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = NewAddDecorator(c, 10)
	c = NewMulDecorator(c, 8)
	res := c.Calc()

	fmt.Println(res) // 80
}
