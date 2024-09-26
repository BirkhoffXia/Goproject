package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/go-course-project/go13/vblog/apps/token/api"
	token_impl "gitlab.com/go-course-project/go13/vblog/apps/token/impl"
	user_impl "gitlab.com/go-course-project/go13/vblog/apps/user/impl"
)

// BIRKHOFF VLOB PROCESS[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//  - using env:	export GIN_MODE=release
//  - using code:	gin.SetMode(gin.ReleaseMode)

// [GIN-debug] POST   /vblog/api/v1/token.AppName/ --> gitlab.com/go-course-project/go13/vblog/apps/token/api.(*TokenApiHandler).Login-fm (3 handlers)
// [GIN-debug] DELETE /vblog/api/v1/token.AppName/ --> gitlab.com/go-course-project/go13/vblog/apps/token/api.(*TokenApiHandler).Logout-fm (3 handlers)
// [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
// Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
// [GIN-debug] Listening and serving HTTP on :8080

// 使用 POSTMAN进行测试 1.登录 2.退出
func main() {
	fmt.Printf("BIRKHOFF VLOB PROCESS")

	// user service impl
	usvc := user_impl.NewUserServiceImpl()

	// token service impl
	tsvc := token_impl.NewTokenServiceImpl(usvc)

	// api
	TokenApiHandler := api.NewTokenApiHandler(tsvc)

	engine := gin.Default()

	rr := engine.Group("/vblog/api/v1")
	TokenApiHandler.Registry(rr)

	//
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
