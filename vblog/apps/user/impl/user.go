package impl

import (
	"context"

	"gitlab.com/go-course-project/go13/vblog/apps/user"
)

// 实现 user.Service
// 怎么判断这个服务有没有实现这个接口？
// &UserServiceImpl{} 是会分配内存, 怎么才能不分配内存
// nil 如何生命 *UserServiceImpl 的 nil
// (*UserServiceImpl)(nil) ---> int8 1  int32(1)  (int32)(1)
// nil 就是一个*UserServiceImpl的空指针
var _ user.Service = (*UserServiceImpl)(nil)

// 用户创建
//
//	u, err := i.CreateUser(ctx, req)
func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	in *user.CreateUserRequest) (*user.User, error) {
	// 1. 校验用户的参数请求
	// 以下写的代码可以独立出来
	// if in.Username == "" {
	// 	return nil,fmt.Errorf("Username Required")
	// }
	// if in.Password == "" {
	// 	return nil,fmt.Errorf("Password Required")
	// }
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// 2. 创建用户实例对象
	u := user.NewUser(in)

	// Hash 完成后入库

	// 3. 把对象持久化-存放数据库中
	// ORM: ORM 需要定义这个对象 存在哪个表中，以及struct和数据库里表字段的映射关系
	// Object ---> Row
	// 比如 create user时到达了4秒的时候，请求还没返回，用户就取消了请求，后端会因为请求退出而结束么？
	// 程序里 并没有终端数据库操作的能力，通过WithContext携带上ctx
	// WithContext(ctx) 方法将上下文 ctx 与后续的数据库操作绑定，这样可以在请求的整个生命周期中跟踪和管理。
	if err := i.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}

	// 4.返回创建好的对象
	return u, nil
}

// 查询用户列表，对象列表[{}]
// 查询数据库里，多行记录
func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (*user.UserSet, error) {

	// 1.构建一个mysql 条件查询语句 select * from users where ...
	query := i.db.WithContext(ctx).Model(&user.User{})

	// 2. 构造条件 where username = ""
	if in.Username != "" {
		// query 会生成一个新的语句，不会修改query对象本身
		query = query.Where("username = ?", in.Username)
	}

	set := user.NewUserSet()

	// 统计当前有多少个？
	// select COUNT(*) from users where ...
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}
	// 3. 做真正的分页查询：sql LIMIT 10 offset,limit
	//  LIMIT 20,20 这个是查询的第2页
	// 使用Find 把多行数据查询出来，使用[]User 接受返回
	err = query.
		Limit(in.Limit()).
		Offset(in.Offset()).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

// 查询用户详情, 通过Id查询,
func (i *UserServiceImpl) DescribeUser(
	ctx context.Context,
	in *user.DescribeUserRequest) (*user.User, error) {

	query := i.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", in.UserId)

	// 准备一个对象 接受数据库的返回
	u := user.NewUser(user.NewCreateUserRequest())
	if err := query.First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
