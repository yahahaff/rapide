// Package response 响应处理工具
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义响应结构体
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message"`
}

// ResponseOption 响应选项函数类型
type ResponseOption func(*Response)

// WithCode 设置响应码
func WithCode(code int) ResponseOption {
	return func(r *Response) { r.Code = code }
}

// WithMessage 设置消息
func WithMessage(message string) ResponseOption {
	return func(r *Response) { r.Message = message }
}

// WithData 设置数据
func WithData(data interface{}) ResponseOption {
	return func(r *Response) { r.Data = data }
}

// WithError 设置错误
func WithError(err interface{}) ResponseOption {
	return func(r *Response) { r.Error = err }
}

// respond 核心响应方法
func respond(c *gin.Context, status int, opts ...ResponseOption) {
	resp := &Response{Code: 0} // 默认成功
	for _, opt := range opts {
		opt(resp)
	}

	if status >= 400 {
		c.AbortWithStatusJSON(status, resp)
	} else {
		c.JSON(status, resp)
	}
}

// OK 成功响应
func OK(c *gin.Context, data interface{}, message ...string) {
	opts := []ResponseOption{
		WithData(data),
		WithMessage(defaultMessage("OK", message...)),
	}
	respond(c, http.StatusOK, opts...)
}

// Success 简单成功响应
func Success(c *gin.Context, message ...string) {
	respond(c, http.StatusOK, WithMessage(defaultMessage("Success", message...)))
}

// Error 错误响应
func Error(c *gin.Context, status int, opts ...ResponseOption) {
	// 确保错误响应有错误码
	hasCode := false
	for _, opt := range opts {
		if _, ok := interface{}(opt).(func(*Response)); ok {
			hasCode = true
			break
		}
	}
	if !hasCode {
		opts = append(opts, WithCode(1))
	}

	respond(c, status, opts...)
}

// 以下是常用错误响应的快捷方法

// BadRequest 400错误
func Abort400(c *gin.Context, message ...string) {
	Error(c, http.StatusBadRequest,
		WithMessage(defaultMessage("请求解析错误,请确认请求格式是否正确", message...)))
}

// Unauthorized 401错误
func Abort401(c *gin.Context, message ...string) {
	Error(c, http.StatusUnauthorized,
		WithMessage(defaultMessage("未认证", message...)))
}

// Forbidden 403错误
func Abort403(c *gin.Context, message ...string) {
	Error(c, http.StatusForbidden,
		WithMessage(defaultMessage("权限不足", message...)))
}

// NotFound 404错误
func Abort404(c *gin.Context, message ...string) {
	Error(c, http.StatusNotFound,
		WithMessage(defaultMessage("数据不存在", message...)))
}

// InternalServerError 500错误
func Abort500(c *gin.Context, message ...string) {
	Error(c, http.StatusInternalServerError,
		WithMessage(defaultMessage("服务器内部错误,请稍后再试", message...)))
}

// ValidationError 验证错误
func ValidationError(c *gin.Context, errors map[string][]string, message ...string) {
	Error(c, http.StatusBadRequest,
		WithMessage(defaultMessage("参数验证错误", message...)),
		WithData(errors))
}

// defaultMessage 默认消息处理
func defaultMessage(defaultMsg string, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return defaultMsg
}
