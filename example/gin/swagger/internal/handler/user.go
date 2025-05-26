package handler

import (
	"strconv"

	"gin-swagger/internal/models"
	"gin-swagger/internal/service"
	"gin-swagger/pkg"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(r gin.IRouter) {
	u := r.Group("/users")
	u.GET("/:uid", GetUserByID)
	u.POST("", AddUser)
}

// AddUser 添加用户
// @Summary 添加用户
// @Tag 用户管理
// @Accept json
// @Produce json
// @Param req body models.User true "用户信息"
// @Success 200 {object} pkg.Response "请求成功"
// @Router /users [POST]
func AddUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBind(&req); err != nil {
		pkg.ResponseError(c, -500, "参数错误")
		return
	}

	if err := service.AddUser(req); err != nil {
		pkg.ResponseError(c, -500, "系统错误")
		return
	}

	pkg.ResponseOK(c, nil)
}

// GetUserByID 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Tag 用户管理
// @Accept json
// @Produce json
// @Param uid path int64 true "用户uid"
// @Success 200 {object} pkg.Response{data=models.User} "请求成功"
// @Router /users/{uid} [GET]
func GetUserByID(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		pkg.ResponseError(c, -400, "缺少uid")
		return
	}
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		pkg.ResponseError(c, -500, "系统错误")
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		pkg.ResponseError(c, -500, "系统错误")
		return
	}

	pkg.ResponseOK(c, user)
}
