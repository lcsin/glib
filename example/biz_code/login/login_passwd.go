package login

import "fmt"

// PasswdLogin 密码登录
type PasswdLogin struct {
	Email       string
	InputPasswd string
}

// Login 密码登录
func (p *PasswdLogin) Login() error {
	passwd := "123456"
	if passwd == p.InputPasswd {
		return nil
	}
	return fmt.Errorf("用户名或密码错误")
}
