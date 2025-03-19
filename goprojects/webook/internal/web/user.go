package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/service"
	"github.com/lcsin/webook/pkg/api"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (u *UserHandler) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/users/v1")

	g.POST("/register", u.Register)
	g.POST("/login", u.Login)
	g.POST("/logout", u.Logout)
	g.POST("/profile", u.Profile)
	g.POST("/edit", u.Edit)
}

func (u *UserHandler) Register(c *gin.Context) {
	type RegisterReq struct {
		Email         string `json:"email"`
		Passwd        string `json:"passwd"`
		ConfirmPasswd string `json:"confirm_passwd"`
		Username      string `json:"username"`
		Age           int8   `json:"age"`
	}

	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	// todo 校验邮箱格式和密码格式
	if req.Passwd != req.ConfirmPasswd {
		api.ResponseError(c, -1, "两次输入的密码不一致")
		return
	}

	if err := u.svc.Register(c, domain.User{
		Email:    req.Email,
		Passwd:   req.Passwd,
		Username: req.Username,
		Age:      req.Age,
	}); err != nil {
		api.ResponseError(c, -1, "系统错误")
		return
	}

	api.ResponseOK(c, nil)
}

func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email  string `json:"email"`
		Passwd string `json:"passwd"`
	}

	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	user, err := u.svc.Login(c, req.Email, req.Passwd)
	if err != nil {
		api.ResponseError(c, -1, err.Error())
		return
	}
	user.Passwd = ""

	api.ResponseOK(c, user)
}

func (u *UserHandler) Logout(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {

}

func (u *UserHandler) Edit(c *gin.Context) {

}
