package impl_test

import (
	"context"
	"testing"

	"gitlab.com/go-course-project/go13/vblog/apps/token"

	"gitlab.com/go-course-project/go13/vblog/apps/token/impl"
	ui "gitlab.com/go-course-project/go13/vblog/apps/user/impl"
)

var (
	i   token.Service
	ctx = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123456")
	req.RemindMe = true
	tk, err := i.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tk)
}

// DELETE FROM `tokens` WHERE access_token = 'crp19id19qon7f6p9pk0' AND refresh_token = 'crp19id19qon7f6p9pkg'
func TestRevokeToken(t *testing.T) {
	req := token.NewRevokeTokenRequest("crp19id19qon7f6p9pk0", "crp19id19qon7f6p9pkg")
	tk, err := i.RevokeToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("crp1l8t19qoqk414rg7g")
	tk, err := i.ValidateToken(ctx, req)

	// exception.IsException(err, token.ErrAccessTokenExpired)

	// if e, ok := err.(*exception.APIException); ok {
	// 	t.Log(e.String())
	// 	// 判断该异常是不是TokenExpired异常
	// 	if e.Code == token.ErrAccessTokenExpired.Code {
	// 		t.Log(e.String())
	// 	}
	// }

	if err != nil {
		t.Fatal(err)
	}

	t.Log(tk)
}
func init() {
	// 加载被测试对象, i 就是User Service接口的具体实现对象
	i = impl.NewTokenServiceImpl(ui.NewUserServiceImpl())
}
