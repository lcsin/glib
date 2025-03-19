package domain

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Passwd       string `json:"passwd"`
	Username     string `json:"username"`
	Age          int8   `json:"age"`
	RegisterTime int64  `json:"register_time"`
}
