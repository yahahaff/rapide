package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/yahahaff/rapide/internal/models/sys"
	casbinPkg "github.com/yahahaff/rapide/pkg/casbin"
	"github.com/yahahaff/rapide/pkg/database"
	"log"
	"strconv"
)

// SetupCasbinEnforcer 创建casbin的enforcer
func SetupCasbinEnforcer() {

	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", `some(where (p.eft == allow))`)
	m.AddDef("m", "m", `g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || p.obj == "*"`)

	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(database.DB, &sys.CasbinRule{}, "sys_casbin_rule")
	if err != nil {
		log.Fatalf("gormadapter.NewAdapterByDBWithCustomTable err: %v", err)
	}
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("casbin.NewEnforcer err: %v", err)
		return
	}

	enforcer.EnableLog(false)
	// 添加自定义函数
	//err = enforcer.AddFunction("ApiMatch", apiMatchFunc)
	if err != nil {
		log.Fatalf("enforcer.AddFunction err: %v", err)
		return
	}

	casbinPkg.Enforcer = enforcer

	//初始化api 到casBinPolicy
	var roles []sys.Role
	if err := database.DB.Preload("Menus").Find(&roles).Error; err != nil {
		log.Fatalf("Failed to retrieve roles: %v", err)
	}
	for _, role := range roles {

		// 第一个角色默认为admin角色,为admin角色添加访问所有权限
		if role.ID == 1 {
			policy := []string{strconv.Itoa(int(role.ID)), "*", "*"}
			_, err := enforcer.AddPolicy(policy)
			if err != nil {
				return
			}
			continue
		}

		// 其他角色,添加权限
		err := addPoliciesForMenu(enforcer, &role)
		if err != nil {
			log.Fatalf("Failed to add policies for role: %v", err)
		}
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return
	}

}

func addPoliciesForMenu(enforcer *casbin.Enforcer, role *sys.Role) error {
	for _, menu := range role.Menus {
		if menu.IsMenu {
			// Variable is true
			continue
		}
		// 根据菜单的URL生成对应的Casbin策略规则
		policy := []string{strconv.Itoa(int(role.ID)), menu.URL, menu.Method}
		_, err := enforcer.AddPolicy(policy)
		if err != nil {
			return err
		}
	}
	return nil
}
