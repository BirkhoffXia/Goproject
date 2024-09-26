package unittest_test

import (
	"os"
	"strconv"
	"testing"

	"gitlab.com/go-course-project/go13/skills/unittest"
)

// 这个单元测试需要读取外部变量
// 这里的ARG1/ARG2 如果传递给单元测试
// 当我们点击debug test时 怎么注入自定义环境变量
func TestSum(t *testing.T) {
	// read file,文件有相对路径和绝对路径
	// 使用环境变量 读取和使用
	a1 := os.Getenv("ARG1")
	a2 := os.Getenv("ARG2")
	a1I, _ := strconv.Atoi(a1)
	a2I, _ := strconv.Atoi(a2)

	t.Log(unittest.Sum(a1I, a2I))
}

// Running tool: D:\Software Install\Go1220\bin\go.exe test -timeout 30s -run ^TestSum$ gitlab.com/go-course-project/go13/skills/unittest

// === RUN   TestSum
//     d:\goprojects\skills\unittest\add_test.go:15: 3
// --- PASS: TestSum (0.00s)
// PASS
// ok      gitlab.com/go-course-project/go13/skills/unittest       0.512s
