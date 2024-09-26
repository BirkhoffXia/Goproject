package user

// Go 什么是枚举？ 为什么要使用枚举？
// 枚举 把所有可选项 一一列举出来
// 到底有哪些值 可传 是需要程序例举出来
// Role、 Admin/Member  Role string ==> "", 看代码的人不知道Role到了应该传怎么样的值
// 枚举核心能力： 约束 只能传递 列举除
// ROLE_ADMIN / ROLE_MEMBER

// 通过声明一种自定义类型来声明一种类型
type Role int

// 通过定义满足类型的常量，来定义满足这个类型的列表
const (
	// 当 值为0的时候，就是默认值
	// 枚举命名风格
	ROLE_ADMIN Role = iota //默认是0
	ROLE_MEMBER
)
