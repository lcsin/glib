package service

import (
	"context"
	"errors"

	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) Register(ctx context.Context, u domain.User) error {
	// 校验邮箱是否已经被注册
	user, _ := us.repo.GetByEmail(ctx, u.Email)
	if user.ID > 0 {
		return errors.New("该邮箱已经被注册")
	}

	// 密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Passwd = string(password)

	return us.repo.Create(ctx, u)
}

func (us *UserService) Login(ctx context.Context, email, passwd string) (*domain.User, error) {
	user, err := us.repo.GetByEmail(ctx, email)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("用户名不存在或密码错误")
	}

	// 比对密码是否一致
	if err = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); err != nil {
		return nil, errors.New("用户名不存在或密码错误")
	}

	return user, nil
}
