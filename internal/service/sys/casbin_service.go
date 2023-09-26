package sys

import (
	"fmt"
	"github.com/yahahaff/rapide/pkg/casbin"
)

type CasbinService struct{}

// AddCabinPolicy 添加policy
func (cs *CasbinService) AddCabinPolicy(ptype, roleID, uri, method string) error {
	_, err := casbin.Enforcer.AddNamedPolicy(ptype, roleID, uri, method)

	if err != nil {
		goto doError
	}

	err = casbin.Enforcer.LoadPolicy()
	if err != nil {
		goto doError
	}
	return err //可以不写，当有错误时会自动往下执行
doError:
	fmt.Println(err)
	return err
}

// GetPolicy P list
func (cs *CasbinService) GetPolicy() [][]string {
	policy := casbin.Enforcer.GetPolicy()
	return policy
}
