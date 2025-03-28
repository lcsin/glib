package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/cmd"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository/model"
	"github.com/lcsin/webook/ioc"
	"github.com/lcsin/webook/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// 测试套件
type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

// 执行测试之前的初始化配置
func (a *ArticleTestSuite) SetupSuite() {
	engine := gin.Default()
	// 模拟用户登录
	engine.Use(func(c *gin.Context) {
		c.Set("uid", &domain.UserClaims{
			UID: 1,
		})
	})
	// 初始化handler
	handler := cmd.InitTestArticleHandler()
	handler.RegisterRoutes(engine)
	a.server = engine
	a.db = ioc.InitTestDB()
}

// 每一个测试都会执行
func (a *ArticleTestSuite) TearDownTest() {
	// 清空所有数据，并且自增主键恢复到 1
	a.db.Exec("TRUNCATE TABLE article_tbl")
}

func TestArticle(t *testing.T) {
	suite.Run(t, &ArticleTestSuite{})
}

// 测试新建帖子
func (a *ArticleTestSuite) TestEdit() {
	type Req struct {
		ID      int64
		Title   string
		Content string
	}
	testCases := []struct {
		name string
		// 要提前准备的数据
		before func(t *testing.T)
		// 验证并且删除数据
		after func(t *testing.T)
		// 请求参数
		req Req
		// 预期响应
		wantCode int
		// 预期返回结果
		wantResult pkg.Response[int64]
	}{
		{
			name:   "新建帖子成功",
			before: func(t *testing.T) {},
			after:  func(t *testing.T) {},
			req: Req{
				Title:   "我的标题",
				Content: "我的内容",
			},
			wantCode: http.StatusOK,
			wantResult: pkg.Response[int64]{
				Code:    0,
				Message: "ok",
				Data:    1,
			},
		}, {
			name: "更新自己的帖子成功",
			before: func(t *testing.T) {
				a.db.Create(&model.Article{
					AuthorID:    1,
					Title:       "我的标题",
					Content:     "我的内容",
					CreatedTime: time.Now().UnixMilli(),
				})
			},
			after: func(t *testing.T) {},
			req: Req{
				ID:      1,
				Title:   "我的新标题",
				Content: "我的新内容",
			},
			wantCode: http.StatusOK,
			wantResult: pkg.Response[int64]{
				Code:    0,
				Message: "ok",
				Data:    1,
			},
		}, {
			name: "更新别人的帖子失败",
			before: func(t *testing.T) {
				a.db.Create(&model.Article{
					ID:          3,
					AuthorID:    2,
					Title:       "2的帖子",
					Content:     "2的帖子",
					CreatedTime: 123,
					UpdatedTime: 123,
				})
			},
			after: func(t *testing.T) {},
			req: Req{
				ID:      3,
				Title:   "1的帖子",
				Content: "1的帖子",
			},
			wantCode: http.StatusOK,
			wantResult: pkg.Response[int64]{
				Code:    -1,
				Message: "系统错误",
				Data:    0,
			},
		},
	}

	for _, tc := range testCases {
		a.T().Run(tc.name, func(t *testing.T) {
			tc.before(t)
			data, err := json.Marshal(tc.req)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/articles/v1/edit", bytes.NewReader(data))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			a.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != http.StatusOK {
				return
			}

			var result pkg.Response[int64]
			err = json.NewDecoder(resp.Body).Decode(&result)
			require.NoError(t, err)
			assert.Equal(t, tc.wantResult, result)
			tc.after(t)
		})
	}
}

// 测试发布帖子
func (a *ArticleTestSuite) TestPublish() {
	testCases := []pkg.Case{
		{
			Name: "发布帖子成功",
			Before: func(t *testing.T) {
				a.db.Create(&model.Article{
					AuthorID:    1,
					Title:       "帖子1",
					Content:     "帖子1",
					Status:      0,
					CreatedTime: time.Now().UnixMilli(),
				})
			},
			After: func(t *testing.T) {
				var article model.Article
				err := a.db.Where("id = ?", 1).First(&article).Error
				assert.NoError(t, err)
				assert.Equal(t, article.Status, int8(domain.ArticlePublished))
			},
			ExpCode: http.StatusOK,
			ExpResult: pkg.Response[int64]{
				Code:    0,
				Message: "ok",
				Data:    0,
			},
		},
	}

	for _, tc := range testCases {
		a.T().Run(tc.Name, func(t *testing.T) {
			tc.Before(t)

			req, err := tc.HttpTest(http.MethodPost, "/articles/v1/release/1", nil, nil)
			assert.NoError(t, err)
			a.server.ServeHTTP(tc.Response, req)
			assert.Equal(t, tc.ExpCode, tc.Response.Code)

			var result pkg.Response[int64]
			err = tc.ResponseBodyDecoder(&result)
			require.NoError(t, err)
			assert.Equal(t, tc.ExpResult, result)

			tc.After(t)
		})
	}
}
