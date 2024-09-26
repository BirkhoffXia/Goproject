package user_test

import (
	"testing"

	"gitlab.com/go-course-project/go13/vblog/apps/user"
)

// d:\goprojects\vblog\apps\user\interface_test.go:12: $2a$10$Lr50RE.J.bR8C4thAq0ufeoCp9n9cZ822dDw0fHaOyzrS22mlFmjy
// $2a$10$Lr50RE.J.bR8C4thAq0ufeoCp9n9cZ822dDw0fHaOyzrS22mlFmjy
// 2a:Bcrypt、10:Rounds(4-31)、Salt:Lr50RE.J.bR8C4thAq0ufe(128bits 22 chars)、oCp9n9cZ822dDw0fHaOyzrS22mlFmjy:Hash(138 bits 31 chars)

// Bcrycpt有个4个变量
// 1. saltRound:正数、代表Hash杂凑次数，数值越高越安全，默认10次
// 2. myPassword：明文密码字符串
// 3. salt：盐，一个128bits随机字符串，22字符
// 4. myHash：经过明文密码password和盐salt进行hash，个人的理解是默认10次下，循环加盐10次，得到myHash
func TestHashedPassword(t *testing.T) {
	// req := user.NewCreateUserRequest()
	// req.Password = "123456"
	// req.HashedPassword()
	// t.Log(req.Password)

	// t.Log(req.CheckPassword("1234561"))
	req := user.NewCreateUserRequest()
	req.Password = "123456"
	u := user.NewUser(req)
	t.Log(u.Password)
	t.Log(u.CheckPassword("123456"))

}
