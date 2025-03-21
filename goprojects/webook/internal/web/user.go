package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/service"
	"github.com/lcsin/webook/pkg"
)

const (
	SessionUser = "session_user"
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
		pkg.ResponseError(c, -1, "两次输入的密码不一致")
		return
	}

	if err := u.svc.Register(c, domain.User{
		Email:    req.Email,
		Passwd:   req.Passwd,
		Username: req.Username,
		Age:      req.Age,
	}); err != nil {
		pkg.ResponseError(c, -1, err.Error())
		return
	}

	pkg.ResponseOK(c, nil)
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
		pkg.ResponseError(c, -1, err.Error())
		return
	}

	// 登录成功，设置session
	session := sessions.Default(c)
	session.Set(SessionUser, user)
	if err = session.Save(); err != nil {
		pkg.ResponseError(c, -1, "系统错误")
		return
	}

	pkg.ResponseOK(c, user)
}

func (u *UserHandler) Logout(c *gin.Context) {
	// 删除session和cookie
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: -1, Path: "/"})
	session.Clear()
	session.Save()

	// 清除cookie
	//c.SetCookie(SessionKey, "", -1, "/", c.Request.Host, false, false)
	pkg.ResponseOK(c, nil)
}

func (u *UserHandler) Profile(c *gin.Context) {
	// 直接从session中获取
	session := sessions.Default(c)
	user, ok := session.Get(SessionUser).(*domain.User)
	if ok && user != nil {
		pkg.ResponseOK(c, user)
		return
	}

	profile, err := u.svc.Profile(c, c.GetInt64("uid"))
	if err != nil {
		pkg.ResponseError(c, -1, err.Error())
		return
	}

	pkg.ResponseOK(c, profile)
}

func (u *UserHandler) Edit(c *gin.Context) {
	type EditReq struct {
		Username string `json:"username"`
		Age      int8   `json:"age"`
	}

	var req EditReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	if err := u.svc.Edit(c, domain.User{
		ID:       c.GetInt64("uid"),
		Username: req.Username,
		Age:      req.Age,
	}); err != nil {
		pkg.ResponseError(c, -1, err.Error())
		return
	}

	pkg.ResponseOK(c, nil)
}
