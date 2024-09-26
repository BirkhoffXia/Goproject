package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/go-course-project/go13/vblog/exception"
)

// API成功, 返回数据
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// API失败 返回错误, API Excepiton
func Failed(c *gin.Context, err error) {
	// 构造异常数据
	var resp *exception.APIException
	if e, ok := err.(*exception.APIException); ok {
		resp = e
	} else {
		resp = exception.NewAPIException(
			500,
			http.StatusText(http.StatusInternalServerError),
		).WithMessage(err.Error()).WithHttpCode(500)
	}

	// 返回异常
	c.JSON(resp.HttpCode, resp)

	// 再中断
	c.Abort()
}
