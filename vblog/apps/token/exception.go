package token

import (
	"gitlab.com/go-course-project/go13/vblog/exception"
)

// 这个模块定义的业务异常
// token expired %f minutes
// 约定俗成： ErrXXXX 来进行自定义异常定义，方便快速在包里搜索
var (
	ErrAccessTokenExpired  = exception.NewAPIException(5000, "Access Token Expired")
	ErrRefreshTokenExpired = exception.NewAPIException(6000, "Refresh Token Expired")
)
