package model

type User struct {
	ID       int64  `gorm:"comment:用户id;primaryKey;autoIncrement"`
	Email    string `gorm:"comment:邮箱;unique"`
	Username string `gorm:"comment:用户名;size:256"`
	Passwd   string `gorm:"comment:密码;size:256"`
	Age      int8   `gorm:"comment:年龄"`

	CreatedTime int64 `gorm:"comment:创建时间"`
	UpdatedTime int64 `gorm:"comment:更新时间"`
}

func (u User) TableName() string {
	return "user_tbl"
}
