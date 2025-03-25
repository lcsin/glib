package ioc

import (
	"log"
	"os"
	"time"

	"github.com/lcsin/webook/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化MySQL
func InitDB() *gorm.DB {
	//dns := viper.GetString("mysql.dsn")
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

	return db
}

// InitRedis 初始化Redis
func InitRedis() redis.Cmdable {
	//addr := viper.GetString("redis.addr")
	//passwd := viper.GetString("redis.passwd")
	addr := config.Cfg.Redis.Addr
	passwd := config.Cfg.Redis.Passwd
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
	})
}
