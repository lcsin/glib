package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/internal/repository"
	"github.com/lcsin/webook/internal/repository/dao"
	"github.com/lcsin/webook/internal/service"
	"github.com/lcsin/webook/internal/web"
	"github.com/lcsin/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	server := initWebServer()

	userHandler := initUserHandler(db)
	userHandler.RegisterRoutes(server)

	server.Run(":8080")
}

func initUserHandler(db *gorm.DB) *web.UserHandler {
	userDAO := dao.NewUserDAO(db)
	userRepository := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	return userHandler
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(middleware.DefaultCors())
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return server
}

func initDB() *gorm.DB {
	dns := "root:root@tcp(localhost:13306)/webook?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}
	if err = dao.InitTable(db); err != nil {
		panic(err)
	}
	return db
}
