package impl

import (
	"context"
	"fmt"

	"gitlab.com/go-course-project/go13/vblog/apps/token"
	"gitlab.com/go-course-project/go13/vblog/apps/user"
)

// 登录： 颁发令牌
func (i *TokenServiceImpl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	// 1.1 确认用户密码是否正确
	req := user.NewQueryUserRequest()
	req.Username = in.Username
	// 面向接口， 面向具体的业务逻辑， 进行抽象编程
	us, err := i.user.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		return nil, fmt.Errorf("用户名或者密码错误")
	}

	// 1.2 校验密码是否正确
	u := us.Items[0]
	if err := us.Items[0].CheckPassword(in.Password); err != nil {
		return nil, err
	}

	// 2.正确的请求下 就颁发用户令牌
	tk := token.NewToken(in.RemindMe)
	tk.UserId = fmt.Sprintf("%d", u.Id)
	tk.UserName = u.Username
	tk.Role = u.Role

	// 3.保存用户Token
	err = i.db.WithContext(ctx).Model(&token.Token{}).Create(tk).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 退出：撤消令牌，把这个令牌删除-直接删除
// 明确的结构返回
func (i *TokenServiceImpl) RevokeToken(ctx context.Context, in *token.RevokeTokenRequest) (*token.Token, error) {

	// 查询Token
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	// Refresh 确认
	err = tk.CheckRefreshToken(in.RefreshToken)
	if err != nil {
		return nil, err
	}
	// Find Token & Delete
	err = i.db.WithContext(ctx).
		Where("access_token = ?", in.AccessToken).
		Where("refresh_token = ?", in.RefreshToken).
		Delete(&token.Token{}).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 校验令牌
// 依赖User模块来检验，用户的密码是否正确
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	// 1.查询Token, 判断令牌是否存在
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	// 2.判断令牌是否过期
	if err := tk.ValidateExpired(); err != nil {
		return nil, err
	}

	// 3.令牌合法返回令牌
	return tk, nil
}
