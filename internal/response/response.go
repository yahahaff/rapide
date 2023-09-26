// Package response 响应处理工具
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Option 优雅可选参数
type Option struct {
	f func(*Response)
}

func WithCode(code int) Option {
	return Option{func(mp *Response) {
		mp.Code = code
	}}
}

func WithMessage(message string) Option {
	return Option{func(mp *Response) {
		mp.Message = message
	}}
}

func WithData(data interface{}) Option {
	return Option{func(mp *Response) {
		mp.Data = data
	}}
}

// OK 响应成功的JSON数据
func OK(c *gin.Context, data interface{}, message ...string) {

	response := Response{
		Code:    0,
		Message: defaultMessage("OK", message...),
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

// Error 错误响应
func Error(c *gin.Context, ops ...Option) {
	mp := &Response{}
	for _, do := range ops {
		do.f(mp)
	}

	response := Response{
		Code:    defaultCode(1, mp.Code),
		Message: defaultMessage("Error", mp.Message),
		Data:    mp.Data,
	}
	c.JSON(http.StatusOK, response)
}

// Success 响应成功的JSON
func Success(c *gin.Context) {
	response := Response{
		Code:    0,
		Message: "Success",
	}

	c.JSON(http.StatusOK, response)
}

// Abort401 Unauthorized 响应 401，未传参 msg 时使用默认消息
func Abort401(c *gin.Context, msg ...string) {
	response := Response{
		Code:    1,
		Message: defaultMessage("未认证", msg...),
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Abort404 响应 404，未传参 msg 时使用默认消息
func Abort404(c *gin.Context, msg ...string) {
	response := Response{
		Code:    1,
		Message: defaultMessage("数据不存在", msg...),
	}
	c.AbortWithStatusJSON(http.StatusNotFound, response)
}

// Abort403 响应 403，未传参 msg 时使用默认消息
func Abort403(c *gin.Context, msg ...string) {
	response := Response{
		Code:    1,
		Message: defaultMessage("权限不足", msg...),
	}
	c.AbortWithStatusJSON(http.StatusForbidden, response)
}

// Abort500 响应 500，未传参 msg 时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	response := Response{
		Code:    1,
		Message: defaultMessage("服务器内部错误,请稍后再试", msg...),
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, response)
}

// BadRequest 响应 400，传参 err 对象，未传参 msg 时使用默认消息
// 在解析用户请求，请求的格式或者方法不符合预期时调用
func BadRequest(c *gin.Context, msg ...string) {
	response := Response{
		Code:    1,
		Message: defaultMessage("请求解析错误,请确认请求格式是否正确. 参数请使用JSON格式。", msg...),
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func ValidationError(c *gin.Context, errors map[string][]string, msg ...string) {
	response := Response{
		Code:    1,
		Message: defaultMessage("参数验证错误", msg...),
		Data:    errors,
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}

func defaultCode(defaultCode int, cod ...int) (code int) {
	if len(cod) > 1 {
		code = cod[0]
	} else {
		code = defaultCode
	}
	return
}
