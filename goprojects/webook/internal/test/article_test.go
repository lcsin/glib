package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/webook/cmd"
	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/ioc"
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
//func (a *ArticleTestSuite) TearDownTest() {
//	// 清空所有数据，并且自增主键恢复到 1
//	a.db.Exec("TRUNCATE TABLE article_tbl")
//}

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
	type Response[T any] struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    T      `json:"data"`
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
		wantResult Response[int64]
	}{
		// 第一个测试用例
		{
			name:   "新建帖子-创建成功",
			before: func(t *testing.T) {},
			after:  func(t *testing.T) {},
			req: Req{
				Title:   "我的标题",
				Content: "我的内容",
			},
			wantCode: http.StatusOK,
			wantResult: Response[int64]{
				Code:    0,
				Message: "ok",
				Data:    1,
			},
		}, {
			name:   "更新帖子-更新成功",
			before: func(t *testing.T) {},
			after:  func(t *testing.T) {},
			req: Req{
				ID:      1,
				Title:   "我的新标题",
				Content: "我的新内容",
			},
			wantCode: http.StatusOK,
			wantResult: Response[int64]{
				Code:    0,
				Message: "ok",
				Data:    1,
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

			var result Response[int64]
			err = json.NewDecoder(resp.Body).Decode(&result)
			require.NoError(t, err)
			assert.Equal(t, tc.wantResult, result)
			tc.after(t)
		})
	}
}
