package service

import (
	"log"

	"gin-swagger/internal/models"
)

func GetUserByID(uid int64) (*models.User, error) {
	return &models.User{
		UID:  uid,
		Name: "张三",
		Age:  20,
	}, nil
}

func AddUser(user models.User) error {
	log.Println("add user ok: ", user)
	return nil
}
