package impl

import (
	"gitlab.com/go-course-project/go13/vblog/conf"
	"gorm.io/gorm"
)

// 用户业务定义层(对业务的抽象)，由lmpl模块来完成具体的功能实现

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		// 获取全局的DB对象
		// 前提：配置对象准备完成
		db: conf.C().DB(),
	}
}

// 怎么实现user.Service接口?
// 定义UserServiceImpl来实现接口
type UserServiceImpl struct {
	// 依赖了一个数据库操作的链接池对象
	db *gorm.DB
}
