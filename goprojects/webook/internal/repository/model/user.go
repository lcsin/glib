package model

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"unique"`
	Username string `gorm:"size:256"`
	Passwd   string `gorm:"size:256"`
	Age      int8

	CreatedTime int64
	UpdatedTime int64
}

func (u User) TableName() string {
	return "user_tbl"
}
