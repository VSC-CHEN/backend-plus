package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//封装的结构
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Success = 200 // 成功

	NoLogin         = 401 // 未登录
	ParamErr        = 402 // 参数错误
	NoAuth          = 403 // 无权限
	NotFind         = 404 // 找不到
	TokenExpired    = 405 // Token过期
	TooManyRequests = 429 // 请求过于频繁

	ServerErr = 500 // 服务器错误
	TokenErr  = 501 // Token分发失败
	SQLErr    = 502 // sql错误
)

//封装反馈的结果
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
//开始封装
func Succeed(c *gin.Context) {
	Result(Success, map[string]interface{}{}, "操作成功", c)
}

func SucceedWithInfor(message string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, message, c)
}

func SucceedWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func SucceedWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

func FailWithInfor(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
