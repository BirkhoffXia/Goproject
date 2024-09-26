package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/tools/pretty"
	"golang.org/x/crypto/bcrypt"
)

const (
	// 模块名称
	AppName = "users"
)

// validator库 validator.New() 校验对象，全局单列模式
// "github.com/go-playground/validator/v10"
var (
	v = validator.New()
)

// user.Service
// 面向对象
// 接口定义，一定要考虑兼容性、接口的参数不能变
type Service interface {
	//用户创建
	// CreateUser(username,password,role string,label map[string]string) 不用这种 万一新增一个参数 会破坏接口传参
	// 设计CreateUserRequest，可以扩展对象，而不会影响接口的定义
	// 1.这个接口支持取消吗？ 要支持取消该如何？
	// 2.这个接口支持Trace，TraceId怎么传递
	// 中间件参数，取消Trace/ ... 怎么产生怎么传递
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表, 对象列表 [{}]
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// 查询用户详情, 通过Id查询,
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)

	//作业：
	//用户修改
	//用户删除
}

// 为了避免对象内部出现很空指针, 指针对象为初始化, 为该对象提供一个构造函数
// 还能做一些相关兼容，补充默认值的功能, New+对象名称()
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		// Role:  ROLE_VISITOR,
		// Role:  "member",
		Role:  ROLE_MEMBER,
		Label: map[string]string{},
	}
}

// 用户创建的参数
type CreateUserRequest struct {
	Username string `json:"username" validate:"required" gorm:"column:username"`
	Password string `json:"password" validate:"required" gorm:"column:password"`
	Role     Role   `json:"role" gorm:"column:role"`
	// https://gorm.io/docs/serializer.html 内部转化为Json 存到字段里去
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}

// 创建一个 Hash 密码存储于数据库中
func (c *CreateUserRequest) hashedPassword() {
	hp, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	c.Password = string(hp)
}

func (c *CreateUserRequest) CheckPassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(pass))
}

func (req *CreateUserRequest) Validate() error {
	// if req.Username == "" {
	// 	return fmt.Errorf("Username Required")
	// }
	// if req.Password == "" {
	// 	return fmt.Errorf("Password Required")
	// }

	return v.Struct(req)
}

// 构造函数
func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

// 查询用户列表
type QueryUserRequest struct {
	// 分页大小，一个多少个
	PageSize int
	// 当前页，查询哪一页的数据
	PageNumber int
	// 根据name查询用户
	Username string
}

func (req *QueryUserRequest) Limit() int {
	return req.PageSize
}

// 1,0
// 2,20,
// 3,20 * 2
// 4,20 * 3
func (req *QueryUserRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)
}

type UserSet struct {
	// 总共有多少个
	Total int64 `json:"total"`
	// 当前查询的数据清单
	Items []*User `json:"items"`
}

func (u *UserSet) String() string {
	return pretty.ToJSON(u)
}

// 初始化UserSet
func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

// 初始化UserSet
func NewDescribeUserRequest(uid int) *DescribeUserRequest {
	return &DescribeUserRequest{
		UserId: uid,
	}
}

type DescribeUserRequest struct {
	UserId int
}
