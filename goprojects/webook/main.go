package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/config"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository"
	"github.com/lcsin/webook/internal/repository/cache"
	"github.com/lcsin/webook/internal/repository/dao"
	"github.com/lcsin/webook/internal/service"
	"github.com/lcsin/webook/internal/web"
	"github.com/lcsin/webook/internal/web/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := initDB()
	rdb := InitRedis()
	server := initWebServer()

	userHandler := initUserHandler(db, rdb)
	userHandler.RegisterRoutes(server)

	server.Run(":8080")
}

func initUserHandler(db *gorm.DB, cmd redis.Cmdable) *web.UserHandler {
	userDAO := dao.NewUserDAO(db)
	userCache := cache.NewUserCache(cmd)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	return userHandler
}

func initWebServer() *gin.Engine {
	//store, err := redis.NewStore(16, "tcp", "127.0.0.1:16379", "",
	//	[]byte("08092c221370c1ddca1db0ab89cf61b7"), []byte("c0c5091f15628a94ead9f5b7184d918a"))
	//if err != nil {
	//	panic(err)
	//}

	server := gin.Default()
	server.Use(
		middleware.DefaultCors(),
		//middleware.Session("ssid", store),
		middleware.NewLoginValidatorBuilder("uid").
			IgnorePath("/users/v1/register").
			IgnorePath("/users/v1/login").
			JWT([]byte(config.Cfg.Jwt.Secret), &domain.UserClaims{}),
	)
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return server
}

func initDB() *gorm.DB {
	dns := "root:root@tcp(localhost:13306)/webook?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			Colorful:      true,
			LogLevel:      logger.Info,
		}),
	})
	if err != nil {
		panic(err)
	}
	if err = dao.InitTable(db); err != nil {
		panic(err)
	}
	return db
}

func InitRedis() redis.Cmdable {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:16379",
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return client
}
