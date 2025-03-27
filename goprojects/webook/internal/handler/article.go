package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/service"
	"github.com/lcsin/webook/pkg"
)

var _ IHandler = (*ArticleHandler)(nil)

type ArticleHandler struct {
	svc service.IArticleService
}

func NewArticleHandler(svc service.IArticleService) *ArticleHandler {
	return &ArticleHandler{svc: svc}
}

func (a *ArticleHandler) RegisterRoutes(v1 *gin.Engine) {
	g := v1.Group("/articles/v1")

	g.POST("/edit", a.Edit)
	g.POST("/detail/:id", a.Detail)
	g.POST("/delete", a.Delete)
}

func (a *ArticleHandler) Edit(c *gin.Context) {
	type Req struct {
		ID      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := c.Bind(&req); err != nil {
		return
	}

	uid, err := GetContextJwtUID(c)
	if err != nil {
		pkg.ResponseError(c, -1, err.Error())
		return
	}

	aid, err := a.svc.Save(c, domain.Article{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			ID: uid,
		},
	})
	if err != nil {
		pkg.ResponseError(c, -1, "系统错误")
		return
	}
	pkg.ResponseOK(c, aid)
}

func (a *ArticleHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		pkg.ResponseError(c, -1, "参数无效")
		return
	}

	article, err := a.svc.Detail(c, articleId)
	if err != nil {
		pkg.ResponseError(c, -1, err.Error())
		return
	}

	pkg.ResponseOK(c, article)
}

func (a *ArticleHandler) Delete(c *gin.Context) {

}
