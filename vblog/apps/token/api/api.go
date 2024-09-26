// 业务API开发

// 使用Gin做开发API的接口: 接口的状态管理(Cookie)

// LogIn: 登录，令牌的颁发
// Token服务颁发Token
// 颁发完成后, 使用SetCookie 通知前端(浏览器), 把cookie设置到本地(前端)

// LogOut: 登出, 令牌的销毁
// Token服务销毁Token
// 使用SetCookie 通知前端 从新设置Cookie为""

package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/go-course-project/go13/vblog/apps/token"
	"gitlab.com/go-course-project/go13/vblog/conf"
	"gitlab.com/go-course-project/go13/vblog/response"
)

// 来实现对外提供 RESTFUL 接口
type TokenApiHandler struct {
	svc token.Service
}

func NewTokenApiHandler(svc token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		svc: svc,
	}
}

// 如何为Handler添加路径, 如果把路由注册给 Http Server
// 需要一个Root Router: path prefix: /vblog/api/v1
func (h *TokenApiHandler) Registry(rr gin.IRouter) {
	// 每个业务模块 都需要往Gin Engine对象注册路由
	// r := gin.Default()
	// rr := r.Group("/vblog/api/v1")

	// 模块路径
	// /vblog/api/v1/tokens
	mr := rr.Group(token.AppName)
	mr.POST("/", h.Login)
	mr.DELETE("/", h.Logout)
}

// 登录
func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1.解析用户请求
	// http的请求可以放到哪里，放Body
	// io.ReadAll(c.Request.Body)
	// defer c.Request.Body.Close()

	// Body 必须是Json
	req := token.NewIssueTokenRequest("", "")
	if err := c.Bind(req); err != nil {
		response.Failed(c, err)
		return
	}

	// 2.业务逻辑处理
	tk, err := h.svc.IssueToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2.1 设置 Cookie
	c.SetCookie(
		"token",
		tk.AccessToken,
		tk.AccessTokenExpiredAt,
		"/",
		conf.C().Application.Domain,
		false,
		true,
	)

	// 3.返回处理的结果
	response.Success(c, tk)
}

// 退出
func (h *TokenApiHandler) Logout(c *gin.Context) {
	// 1. 解析用户请求
	// token为了安全 存放在Cookie获取自定义Header中
	accessToken := token.GetAccessTokenFromHttp(c.Request)
	req := token.NewRevokeTokenRequest(accessToken, c.Query("refresh_token"))
	// 2. 业务逻辑处理
	_, err := h.svc.RevokeToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2.1 删除前端的cookie
	c.SetCookie(
		token.TOKEN_COOKIE_KEY,
		"",
		-1,
		"/",
		conf.C().Application.Domain,
		false,
		true,
	)

	// 3. 返回处理的结果
	response.Success(c, "退出成功")
}
