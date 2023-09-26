package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/logger"
	"go.uber.org/zap"
)

type CabinController struct {
	api.BaseAPIController
}

// AddCasbin 新增权限
// @Summary 新增权限
// @Security Bearer
// @Param data body sys.CasbinAddRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/permissions [post]
func (ctrl *CabinController) AddCasbin(c *gin.Context) {
	request := sys.CasbinAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	err := service.Entrance.SysService.CasbinService.AddCabinPolicy(request.Type, request.RoleID, request.Uri, request.Method)
	if err != nil {
		logger.Error("casbin", zap.String("error", fmt.Sprint(err)))
		return
	}
	response.Success(c)
}

// GetPolicy 获取策略
// @Summary 获取策略
// @Security Bearer
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/permissions [get]
func (ctrl *CabinController) GetPolicy(c *gin.Context) {
	data := service.Entrance.SysService.CasbinService.GetPolicy()
	response.OK(c, data)
}
