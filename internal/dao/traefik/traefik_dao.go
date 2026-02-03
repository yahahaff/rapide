package traefik

import (
	"rapide/internal/models/traefik"
	"rapide/pkg/database"
)

// TraefikDAO Traefik数据访问对象
type TraefikDAO struct{}

// NewTraefikDAO 创建TraefikDAO实例
func NewTraefikDAO() *TraefikDAO {
	return &TraefikDAO{}
}

// GetAllRouters 获取所有启用的路由
func (dao *TraefikDAO) GetAllRouters() ([]traefik.TraefikRouter, error) {
	var routers []traefik.TraefikRouter
	result := database.DB.Where("status = ?", "enabled").Find(&routers)
	return routers, result.Error
}

// GetAllServices 获取所有启用的服务
func (dao *TraefikDAO) GetAllServices() ([]traefik.TraefikService, error) {
	var services []traefik.TraefikService
	result := database.DB.Where("status = ?", "enabled").Find(&services)
	return services, result.Error
}

// GetAllMiddlewares 获取所有启用的中间件
func (dao *TraefikDAO) GetAllMiddlewares() ([]traefik.TraefikMiddleware, error) {
	var middlewares []traefik.TraefikMiddleware
	result := database.DB.Where("status = ?", "enabled").Find(&middlewares)
	return middlewares, result.Error
}

// CreateRouter 创建路由
func (dao *TraefikDAO) CreateRouter(router *traefik.TraefikRouter) error {
	return database.DB.Create(router).Error
}

// CreateService 创建服务
func (dao *TraefikDAO) CreateService(service *traefik.TraefikService) error {
	return database.DB.Create(service).Error
}

// CreateMiddleware 创建中间件
func (dao *TraefikDAO) CreateMiddleware(middleware *traefik.TraefikMiddleware) error {
	return database.DB.Create(middleware).Error
}

// UpdateRouter 更新路由
func (dao *TraefikDAO) UpdateRouter(router *traefik.TraefikRouter) error {
	return database.DB.Save(router).Error
}

// UpdateService 更新服务
func (dao *TraefikDAO) UpdateService(service *traefik.TraefikService) error {
	return database.DB.Save(service).Error
}

// UpdateMiddleware 更新中间件
func (dao *TraefikDAO) UpdateMiddleware(middleware *traefik.TraefikMiddleware) error {
	return database.DB.Save(middleware).Error
}

// DeleteRouter 删除路由
func (dao *TraefikDAO) DeleteRouter(name string) error {
	return database.DB.Where("name = ?", name).Delete(&traefik.TraefikRouter{}).Error
}

// DeleteService 删除服务
func (dao *TraefikDAO) DeleteService(name string) error {
	return database.DB.Where("name = ?", name).Delete(&traefik.TraefikService{}).Error
}

// DeleteMiddleware 删除中间件
func (dao *TraefikDAO) DeleteMiddleware(name string) error {
	return database.DB.Where("name = ?", name).Delete(&traefik.TraefikMiddleware{}).Error
}
