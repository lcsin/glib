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
	db, err := gorm.Open(mysql.Open(config.Cfg.MySQL.DNS), &gorm.Config{
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
	//addr := viper.Get("redis.addr").(string)
	//passwd := viper.Get("redis.passwd").(string)
	return redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Passwd,
	})
}
