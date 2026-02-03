package validators

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"rapide/pkg/logger"
	"rapide/pkg/response"
)

type FieldError struct {
	Key   string `json:"Key"`
	Error string `json:"Error"`
}

func Validate(c *gin.Context, obj interface{}) bool {

	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		logger.WarnString("requests", "error", fmt.Sprintf("参数校验错误: %s", c.Request.URL))
		response.Abort400(c, "请求解析错误,请确认请求格式是否正确.上传文件请使用 multipart 标头,参数请使用JSON格式")
		return false
	}
	// 创建验证器实例
	v, ok := c.Get("validator")
	if !ok {
		// 处理验证器对象不存在的情况
		return false
	}
	// 使用验证器对象进行验证
	if err := v.(*validator.Validate).Struct(obj); err != nil {
		// 处理验证失败的情况
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// 处理验证器错误
			return false
		}

		fieldErrors := make(map[string][]string)
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			errMsg := extractValidationErrorMessage(e)
			fieldErrors[e.Field()] = append(fieldErrors[e.Field()], errMsg)
		}

		response.ValidationError(c, fieldErrors)
		return false
	}
	return true
}

func extractValidationErrorMessage(err validator.FieldError) string {
	errMsg := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.ActualTag())
	return errMsg
}
