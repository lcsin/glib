package login

import (
	"testing"
)

func TestPasswdLogin(t *testing.T) {
	build := NewILoginBuilder().
		WithEmail("1847@qq.com").
		WithInputPasswd("1234").Build(Passwd)
	if err := build.Login(); err != nil {
		t.Fatalf("登录失败：%v", err)
	}
	t.Log("登录成功")
}

func TestPhoneLogin(t *testing.T) {
	build := NewILoginBuilder().
		WithPhone("187xxxxx").
		WithInputCode("123456").Build(PhoneCode)
	if err := build.Login(); err != nil {
		t.Fatalf("登录失败：%v", err)
	}
	t.Log("登录成功")
}
