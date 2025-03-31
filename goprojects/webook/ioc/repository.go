package ioc

import (
	"log"
	"os"
	"time"

	"github.com/lcsin/webook/config"
	"github.com/lcsin/webook/internal/repository/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化MySQL
func InitDB() *gorm.DB {
	dsn := config.Cfg.MySQL.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			Colorful:      true,
			LogLevel:      logger.Warn,
		}),
	})
	if err != nil {
		panic(err)
	}
	// 初始化数据库表
	if err = db.AutoMigrate(&model.User{}, &model.ArticleWriter{}, &model.ArticleReader{}); err != nil {
		panic(err)
	}

	return db
}

func InitTestDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:13306)/webook?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			Colorful:      true,
			LogLevel:      logger.Warn,
		}),
	})
	if err != nil {
		panic(err)
	}
	// 初始化数据库表
	if err = db.AutoMigrate(&model.User{}, &model.ArticleWriter{}, &model.ArticleReader{}); err != nil {
		panic(err)
	}

	return db
}

// InitRedis 初始化Redis
func InitRedis() redis.Cmdable {
	addr := config.Cfg.Redis.Addr
	passwd := config.Cfg.Redis.Passwd
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
	})
}
