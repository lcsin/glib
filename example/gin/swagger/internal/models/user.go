package models

type User struct {
	UID  int64  `json:"uid" example:"1"`
	Name string `json:"name" example:"zhangsan"`
	Age  int64  `json:"age" example:"20"`
}
