package token

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	// 模块名称
	AppName = "tokens"
)

// Token Service接口定义
type Service interface {
	// 登录、颁发令牌
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)

	// 退出：撤消令牌
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)

	// 校验令牌
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

// 颁发令牌的请求
type IssueTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// 延长Token的有效期为1周
	RemindMe bool `json:"remind_me"`
}

func NewRevokeTokenRequest(accessToken, refreshToken string) *RevokeTokenRequest {
	return &RevokeTokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

// 撤销令牌的请求
type RevokeTokenRequest struct {
	AccessToken string
	// 他们配对的 撤销时 需要验证他们是不是一对的
	RefreshToken string
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}

type ValidateTokenRequest struct {
	AccessToken string
}

func GetAccessTokenFromHttp(req *http.Request) string {
	// 自定义头, 头叫什么名字: Authorization: xxx Bearer xxxx
	// auth header = ah
	ah := req.Header.Get(TOKEN_HEADER_KEY)
	if ah != "" {
		hv := strings.Split(ah, " ")
		if len(hv) > 1 {
			return hv[1]
		}
	}

	// 再读Cookie
	ck, err := req.Cookie(TOKEN_COOKIE_KEY)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	val, _ := url.QueryUnescape(ck.Value)
	return val
}
