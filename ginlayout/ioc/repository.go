package ioc

import (
	"log"
	"os"
	"time"

	"github.com/lcsin/glib/ginlayout/config"
	"github.com/lcsin/glib/ginlayout/internal/repository/gen/query"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	query.SetDefault(db)
	return db
}

func InitRedis() redis.Cmdable {
	addr := config.Cfg.Redis.Addr
	passwd := config.Cfg.Redis.Passwd
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
	})
}
