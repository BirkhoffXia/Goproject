package impl

import (
	"context"

	"gitlab.com/go-course-project/go13/vblog/apps/token"
)

// 一般存放一些复用的函数和操作
func (i *TokenServiceImpl) getToken(ctx context.Context, accessToken string) (*token.Token, error) {
	tk := token.NewToken(false)

	err := i.db.WithContext(ctx).
		Model(&token.Token{}).
		Where("access_token = ?", accessToken).
		First(tk).
		Error

	if err != nil {
		return nil, err
	}
	return tk, nil
}
