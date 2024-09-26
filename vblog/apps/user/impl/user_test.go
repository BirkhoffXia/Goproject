package impl_test

import (
	"context"
	"testing"

	"gitlab.com/go-course-project/go13/vblog/apps/user"
	"gitlab.com/go-course-project/go13/vblog/apps/user/impl"
)

var (
	i   user.Service
	ctx = context.Background()
)

func TestCreateUser(t *testing.T) {
	// 使用构造函数创建请求对象
	// user.CreateUserRequest()容易空指针
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	req.Role = user.ROLE_ADMIN

	// 单元测试报错 异常处理
	// [49.448ms] [rows:0] INSERT INTO `users` (`created_at`,`updated_at`,`username`,`password`,`role`,`label`) VALUES (1726626910,1726626910,'admin','123456','admin','{}')
	u, err := i.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 自己进行期望对比，进行单元测试报错
	if u == nil {
		t.Fatal("User not created")
	}

	//正常打印对象
	t.Log(u)
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	userlist, err := i.QueryUser(ctx, req)
	//直接报错中断单元流程并且失败
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userlist)
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequest(16)
	userlist, err := i.DescribeUser(ctx, req)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(userlist)
	t.Log(userlist.CheckPassword("123456")) //检查密码是否对应相同

}

func init() {
	// 加载被测试对象, i 就是User Service接口的具体实现对象
	i = impl.NewUserServiceImpl()
}
