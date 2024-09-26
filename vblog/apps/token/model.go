package token

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/tools/pretty"
	"github.com/rs/xid"
	"gitlab.com/go-course-project/go13/vblog/apps/user"
)

const (
	DEFAULT_EXPIRED_AT = 2 * 60 * 60
	WEEK_EXPIRED_AT    = 7 * 24 * 60 * 60
)

// 生成一个令牌
func NewToken(remindMe bool) *Token {
	// AccessTokenExpiredAt 默认
	atet := DEFAULT_EXPIRED_AT

	if remindMe {
		// 7天过期时间
		atet = WEEK_EXPIRED_AT
	}

	t := &Token{
		// 直接使用uuid库来，生成一个随机字符串
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  atet,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: atet * 4,
		CreatedAt:             time.Now().Unix(),
	}

	return t
}

/*
// 默认时间

	{
		"user_id": "17",
		"username": "admin",
		"access_token": "crp1jmt19qomid6hlq90",
		"access_token_token_expired_at": 7200,
		"refresh_token": "crp1jmt19qomid6hlq9g",
		"refresh_token_expired_at": 28800,
		"created_at": 1727142363,
		"updated_at": 1727142363,
		"role": 0
	}

// RemindMe 为true 时间

	{
		"user_id": "17",
		"username": "admin",
		"access_token": "crp1l8t19qoqk414rg7g",
		"access_token_token_expired_at": 604800,
		"refresh_token": "crp1l8t19qoqk414rg80",
		"refresh_token_expired_at": 2419200,
		"created_at": 1727142563,
		"updated_at": 1727142563,
		"role": 0
	  }
*/
type Token struct {
	// 该Token是颁发
	UserId string `json:"user_id" gorm:"colume:user_id"`
	// 人的名称 user_name
	UserName string `json:"username" gorm:"column:username"`
	// 颁发给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token" gorm:"colume:access_token"`
	// 过期时间(2h),单位是秒
	AccessTokenExpiredAt int `json:"access_token_token_expired_at" gorm:"colume:access_token_token_expired_at"`
	// 刷新Token
	RefreshToken string `json:"refresh_token" gorm:"colume:refresh_token"`
	// 刷新Token过期时间(7h)
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"colume:refresh_token_expired_at"`

	// 创建时间
	CreatedAt int64 `json:"created_at" gorm:"colume:created_at"`
	// 更新实现
	UpdatedAt int64 `json:"updated_at" gorm:"colume:updated_at"`

	// 额外补充信息，gorm忽略处理
	Role user.Role `json:"role" gorm:"-"`
}

//

func (t *Token) CheckRefreshToken(refreshToken string) error {
	if t.RefreshToken != refreshToken {
		return fmt.Errorf("Refresh Token Not Correct !!!")
	}
	return nil
}

// 校验Token是否过期
func (t *Token) ValidateExpired() error {
	// 颁发时间 + refresh_token过期
	refreshExpiredTime := time.Unix(t.CreatedAt, 0).Add(time.Duration(t.RefreshTokenExpiredAt) * time.Second)
	// 和当前时间比较
	// now - refreshExpiredTime
	rDelta := time.Since(refreshExpiredTime).Minutes()
	if rDelta > 0 {
		// return fmt.Errorf("Refresh Token Has Been Expired %f minutes", rDelta)
		return ErrRefreshTokenExpired.WithMessagef("Refresh Token Has Been Expired %f minutes", rDelta)
	}

	// 颁发时间 + access_token过期
	accessExpiredTime := time.Unix(t.CreatedAt, 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)

	aDelta := time.Since(accessExpiredTime).Minutes()
	if aDelta > 0 {
		// return fmt.Errorf("Access Token Has Been Expired %f minutes", aDelta)
		return ErrAccessTokenExpired.WithMessagef("Access Token Has Been Expired %f minutes", aDelta)
	}

	return nil
}

func (t *Token) TableName() string {
	return "tokens"
}

func (u *Token) String() string {
	return pretty.ToJSON(u)
}
