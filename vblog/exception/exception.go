package exception

import (
	"fmt"

	"github.com/infraboard/mcube/tools/pretty"
)

func NewAPIException(code int, reason string) *APIException {
	return &APIException{
		Code:   code,
		Reason: reason,
	}
}

// error的自定义实现
// 通过 API 直接序列号化{}
type APIException struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Reason   string `json:"reason"`
	Message  string `json:"message"`
}

func (e *APIException) Error() string {
	return fmt.Sprintf("%s, %s", e.Reason, e.Message)
}

func (e *APIException) String() string {
	return pretty.ToJSON(e)
}

// 设计为链式调用 New().WithMessage()
func (e *APIException) WithMessage(msg string) *APIException {
	e.Message = msg
	return e
}

// 设计为链式调用 New().WithHttpCode()
func (e *APIException) WithHttpCode(code int) *APIException {
	e.HttpCode = code
	return e
}

// 设计为链式调用 New().WithMessage()
func (e *APIException) WithMessagef(format string, a ...any) *APIException {
	e.Message = fmt.Sprintf(format, a...)
	return e
}

// 给一个异常判断的方法
func IsException(err error, e *APIException) bool {
	if targe, ok := err.(*APIException); ok {
		return targe.Code == e.Code
	}

	return false
}
