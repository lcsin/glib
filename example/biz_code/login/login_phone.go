package login

import "fmt"

// PhoneLogin 手机号登录
type PhoneLogin struct {
	Phone     string
	InputCode string
}

// Login 手机号登录
func (p *PhoneLogin) Login() error {
	messageCode := "123456"
	if messageCode == p.InputCode {
		return nil
	}
	return fmt.Errorf("用户名或密码错误")
}
