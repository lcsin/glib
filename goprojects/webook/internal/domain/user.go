package domain

import (
	"encoding/gob"

	"github.com/golang-jwt/jwt/v5"
)

// 在使用 gin 的 sessions 插件时，如果 save() 的是一个结构体
// 可能会抛出：gob: type not registered for interface 的错误
// 解决方法是：在序列化之前进行一次注册
func init() {
	gob.Register(&User{})
}

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Passwd       string `json:"passwd"`
	Username     string `json:"username"`
	Age          int8   `json:"age"`
	RegisterTime int64  `json:"register_time"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	UID          int64
	Email        string
	Username     string
	Age          int8
	RegisterTime int64
}
