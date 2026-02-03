package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	sysDao "rapide/internal/dao/sys"
	"rapide/internal/models/sys"
	sysReq "rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/pkg/database"
	"rapide/pkg/response"
)

type RoleController struct {
	controllers.BaseAPIController
}

func (rc *RoleController) GetRole(c *gin.Context) {
	// 1. 验证表单
	request := sysReq.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 10 // 设置默认值
	}

	// 处理页码参数
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 获取角色列表和总数
	roles, total := sysDao.GetRoles(page, pageSize)

	// 构造返回数据
	responseData := map[string]interface{}{
		"result":   roles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	response.OK(c, responseData)
}

func (rc *RoleController) AddRole(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.RoleAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，创建数据
	RoleModel := sys.Role{
		RoleName: request.Name,
		Sort:     request.Value,  // 注意：request.Value 对应 Sort 字段
		Remark:   request.Remark, // 注意：request.Remark 对应 Remark 字段
		Status:   request.Status,
		// RoleValue 和 RoleCode 需要根据业务逻辑生成，这里暂时使用默认值
		RoleValue: request.Name,
		RoleCode:  request.Name,
	}

	// 3. 创建角色
	RoleModel.Create()
	if RoleModel.ID == 0 {
		response.Abort500(c, "创建角色失败，请稍后尝试~")
		return
	}

	// 4. 处理权限关联
	if len(request.Permissions) > 0 {
		// 将字符串数组转换为uint64数组
		menuIDs := make([]uint64, 0, len(request.Permissions))
		for _, perm := range request.Permissions {
			var menuID uint64
			if _, err := fmt.Sscan(perm, &menuID); err == nil {
				menuIDs = append(menuIDs, menuID)
			}
		}

		// 分配菜单权限
		if len(menuIDs) > 0 {
			if err := RoleModel.AssignMenus(menuIDs); err != nil {
				response.Abort500(c, "分配菜单权限失败，请稍后尝试~")
				return
			}
		}
	}

	response.OK(c, RoleModel)
}

func (rc *RoleController) DeleteRoleById(c *gin.Context) {

	// 1. 从URL路径中获取id参数
	idStr := c.Param("id")
	var id int
	if _, err := fmt.Sscan(idStr, &id); err != nil || id <= 0 {
		response.Abort500(c, "无效的角色ID")
		return
	}

	// 2. 删除数据
	sysDao.RoleDeletelById(id)
	response.Success(c)

}

// UpdateRole 更新角色信息
func (rc *RoleController) UpdateRole(c *gin.Context) {
	// 1. 从URL路径中获取id参数
	idStr := c.Param("id")
	var id uint64
	if _, err := fmt.Sscan(idStr, &id); err != nil || id <= 0 {
		response.Abort500(c, "无效的角色ID")
		return
	}

	// 2. 解析和验证请求体
	request := sysReq.RoleUpdateRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 3. 根据ID查询角色
	var role sys.Role
	if err := database.DB.First(&role, id).Error; err != nil {
		response.Abort500(c, "角色不存在")
		return
	}

	// 4. 更新角色字段
	// 状态字段：无论值是什么，只要请求中包含就更新
	role.Status = request.Status
	if request.Name != "" {
		role.RoleName = request.Name
		role.RoleValue = request.Name
		role.RoleCode = request.Name
	}
	if request.Remark != "" {
		role.Remark = request.Remark
	}
	// 排序字段：无论值是什么，只要请求中包含就更新
	role.Sort = request.Sort

	// 5. 更新角色权限
	if len(request.Permissions) > 0 {
		// 将字符串数组转换为uint64数组
		menuIDs := make([]uint64, 0, len(request.Permissions))
		for _, perm := range request.Permissions {
			var menuID uint64
			if _, err := fmt.Sscan(perm, &menuID); err == nil {
				menuIDs = append(menuIDs, menuID)
			}
		}

		// 分配菜单权限
		if len(menuIDs) > 0 {
			if err := role.AssignMenus(menuIDs); err != nil {
				response.Abort500(c, "更新角色权限失败")
				return
			}
		}
	}

	// 6. 保存更新
	if err := role.Update(); err != nil {
		response.Abort500(c, "更新角色失败")
		return
	}

	response.OK(c, role)
}

// GetRolePermissions 获取角色权限
func (rc *RoleController) GetRolePermissions(c *gin.Context) {
	// 1. 从URL路径中获取角色ID
	idStr := c.Param("id")
	var roleID uint64
	if _, err := fmt.Sscan(idStr, &roleID); err != nil {
		response.Abort500(c, "无效的角色ID")
		return
	}

	// 2. 查询角色
	var role sys.Role
	if err := database.DB.Preload("Menus").First(&role, roleID).Error; err != nil {
		response.Abort500(c, "角色不存在")
		return
	}

	// 3. 提取菜单ID列表
	permissionIDs := make([]uint64, 0, len(role.Menus))
	for _, menu := range role.Menus {
		permissionIDs = append(permissionIDs, menu.ID)
	}

	// 4. 返回响应
	response.OK(c, gin.H{
		"permissions": permissionIDs,
	})
}

// UpdateRolePermissions 更新角色权限
func (rc *RoleController) UpdateRolePermissions(c *gin.Context) {
	// 1. 从URL路径中获取角色ID
	idStr := c.Param("id")
	var roleID uint64
	if _, err := fmt.Sscan(idStr, &roleID); err != nil {
		response.Abort500(c, "无效的角色ID")
		return
	}

	// 2. 解析请求体
	type PermissionRequest struct {
		Permissions []uint64 `json:"permissions"`
	}
	var request PermissionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Abort500(c, "请求参数解析失败")
		return
	}

	// 3. 查询角色
	var role sys.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		response.Abort500(c, "角色不存在")
		return
	}

	// 4. 更新角色权限
	if err := role.AssignMenus(request.Permissions); err != nil {
		response.Abort500(c, "更新角色权限失败")
		return
	}

	// 5. 返回响应
	response.OK(c, gin.H{
		"permissions": request.Permissions,
	})
}
