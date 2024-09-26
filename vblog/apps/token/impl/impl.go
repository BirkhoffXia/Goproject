package impl

import (
	"gitlab.com/go-course-project/go13/vblog/apps/token"
	"gitlab.com/go-course-project/go13/vblog/apps/user"
	"gitlab.com/go-course-project/go13/vblog/conf"
	"gorm.io/gorm"
)

// 用户业务定义层(对业务的抽象)，由lmpl模块来完成具体的功能实现

var (
	_ token.Service = (*TokenServiceImpl)(nil)
)

func NewTokenServiceImpl(UserServiceImpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		// 获取全局的DB对象
		// 前提：配置对象准备完成
		db:   conf.C().DB(),
		user: UserServiceImpl,
	}
}

// 怎么实现token.Service接口?
// 定义TokenServiceImpl来实现接口
type TokenServiceImpl struct {
	// 依赖了一个数据库操作的链接池对象
	db *gorm.DB

	// 依赖user.Service ， 没有使用 UserServiceImpl 具体实现 为了解耦
	// 依赖接口， 不要接口的具体实现
	user user.Service
}
