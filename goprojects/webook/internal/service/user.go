package service

import (
	"context"
	"errors"

	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(ctx context.Context, u domain.User) error
	Login(ctx context.Context, email, passwd string) (*domain.User, error)
	Profile(ctx context.Context, uid int64) (*domain.User, error)
	Edit(ctx context.Context, u domain.User) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (us *UserService) Register(ctx context.Context, u domain.User) error {
	// 校验邮箱是否已经被注册
	// 先查再注册的情况下，理论上会有并发问题，但是我觉得邮箱冲突的并发问题发生的可能性不大。
	_, err := us.repo.GetByEmail(ctx, u.Email)
	switch err {
	case gorm.ErrRecordNotFound: // 说明邮箱没重复，走正常的注册流程
		// 密码加密
		password, err := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcrypt.DefaultCost)
		if err != nil {
			zap.L().Error("bcrypt.GenerateFromPassword", zap.Error(err))
			return err
		}
		u.Passwd = string(password)
		return us.repo.Create(ctx, u)
	default:
		zap.L().Error("us.repo.GetByEmail", zap.Error(err))
		return errors.New("系统错误")
	}
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

	user.Passwd = ""
	return user, nil
}

func (us *UserService) Profile(ctx context.Context, uid int64) (*domain.User, error) {
	user, err := us.repo.GetByID(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		zap.L().Error("us.repo.GetByID Error", zap.Error(err))
		return nil, errors.New("系统错误")
	}

	user.Passwd = ""
	return user, nil
}

func (us *UserService) Edit(ctx context.Context, u domain.User) error {
	if err := us.repo.ModifyByID(ctx, u); err != nil {
		zap.L().Error("us.repo.ModifyByID", zap.Error(err))
		return errors.New("系统错误")
	}
	return nil
}
