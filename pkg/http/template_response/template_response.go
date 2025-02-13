package template_response

import (
	"github.com/gin-gonic/gin"

	"github.com/andys920605/meme-coin/pkg/http/gcontext"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

const (
	okCode    = 0
	okMessage = "ok"
)

// Empty returns an empty response with code 0 and message "ok".
// Example:
//
//	{
//		"code": 0,
//		"msg": "ok",
//		"data": {}
//	}
func Empty() Response {
	return New(okCode, okMessage, struct{}{})
}

// OK returns a response with code 0 and message "ok".
// The data field is set to the given data.
// Example:
//
//	{
//		"code": 0,
//		"msg": "ok",
//		"data": {
//			"key": "value"
//		}
func OK(data any) Response {
	return New(okCode, okMessage, data)
}

// Error returns a response with the given code and message.
// The data field is set to nil.
// Example:
//
//		{
//			"code": 100,
//			"msg": "error message"
//	 }
func Error(code int, message string) Response {
	return New(code, message, nil)
}

func New(code int, message string, data any) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (r Response) To(c *gin.Context, status int) {
	c.Set(gcontext.ContextKeyRespStatus, status)
	c.Set(gcontext.ContextKeyRespBody, r)
}
