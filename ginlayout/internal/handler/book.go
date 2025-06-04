package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/glib/ginlayout/internal/domain"
	"github.com/lcsin/glib/ginlayout/internal/service"
	"github.com/sirupsen/logrus"
)

var _ IHandler = (*BookHandler)(nil)

type BookHandler struct {
	svc service.IBookService
}

func NewBookHandler(svc service.IBookService) *BookHandler {
	return &BookHandler{svc: svc}
}

func (b *BookHandler) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/v1/books")

	g.POST("/add", b.Add)
	g.PUT("/edit", b.Edit)
	g.GET("/:id", b.GetByID)
	g.POST("/page", b.GetByPage)
}

func (b *BookHandler) GetByPage(c *gin.Context) {
	type Req struct {
		Page     int64 `json:"page"`
		PageSize int64 `json:"pageSize"`
	}

	var req Req
	if err := c.Bind(&req); err != nil {
		logrus.WithError(err).Error("request bind error")
		ResponseError(c, -400, err.Error())
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	books, total, err := b.svc.FindByPage(c, req.Page, req.PageSize)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"page":     req.Page,
			"pageSize": req.PageSize,
		}).WithError(err).Error("find by page error")
		ResponseError(c, -500, err.Error())
		return
	}

	ResponseOK(c, gin.H{"list": books, "total": total})
}

func (b *BookHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.WithFields(logrus.Fields{
			"id": id,
		}).Error("id is empty")
		ResponseError(c, -400, "参数无效")
		return
	}
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"bookId": bookId,
		}).WithError(err).Error("parse int error")
		ResponseError(c, -400, err.Error())
		return
	}

	book, err := b.svc.FindByID(c, bookId)
	if err != nil {
		logrus.WithError(err).Error("find by id error")
		ResponseError(c, -500, err.Error())
		return
	}

	ResponseOK(c, book)
}

func (b *BookHandler) Edit(c *gin.Context) {
	type Req struct {
		ID          int64  `json:"id"`
		Title       string `json:"title"`
		Author      string `json:"author"`
		Price       int64  `json:"price"`
		PublishTime int64  `json:"publishTime"`
	}

	var req Req
	if err := c.Bind(&req); err != nil {
		ResponseError(c, -400, err.Error())
		return
	}

	if err := b.svc.EditByID(c, &domain.Book{
		ID:          req.ID,
		Title:       req.Title,
		Author:      req.Author,
		Price:       req.Price,
		PublishDate: req.PublishTime,
	}); err != nil {
		ResponseError(c, -500, err.Error())
		return
	}

	ResponseOK(c, nil)
}

func (b *BookHandler) Add(c *gin.Context) {
	type Req struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		Price       int64  `json:"price"`
		PublishTime int64  `json:"publishTime"`
	}

	var req Req
	if err := c.Bind(&req); err != nil {
		ResponseError(c, -400, err.Error())
		return
	}

	if err := b.svc.Add(c, &domain.Book{
		Title:       req.Title,
		Author:      req.Author,
		Price:       req.Price,
		PublishDate: req.PublishTime,
	}); err != nil {
		ResponseError(c, -500, err.Error())
		return
	}

	ResponseOK(c, nil)
}
