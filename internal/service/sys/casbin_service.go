package sys

import (
	"fmt"
	"rapide/pkg/casbin"
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

// GetPolicyList GetPolicy P list
func (cs *CasbinService) GetPolicyList() [][]string {
	policies, err := casbin.Enforcer.GetPolicy()
	if err != nil {
		panic(err)
	}
	return policies
}

// GetPolicyListByRole 获取指定角色的策略规则
func (cs *CasbinService) GetPolicyListByRole(roleName string) [][]string {
	policies, err := casbin.Enforcer.GetFilteredPolicy(0, roleName) // 索引 0 表示按策略的第一个字段进行过滤，也就是角色名称
	if err != nil {
		panic(err)
	}
	return policies
}
