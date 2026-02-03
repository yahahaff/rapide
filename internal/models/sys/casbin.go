package sys

import (
	"rapide/internal/models"
)

// CasbinRule is the Gorm model for Casbin rule.
type CasbinRule struct {
	models.BaseModel
	PType string `json:"p_type" gorm:"size:100;index:idx_casbin_rule_p_type;column:ptype"`
	V0    string `json:"v0" gorm:"size:100;index:idx_casbin_rule_unique"`
	V1    string `json:"v1" gorm:"size:100;index:idx_casbin_rule_unique"`
	V2    string `json:"v2" gorm:"size:100;index:idx_casbin_rule_unique"`
	V3    string `json:"v3" gorm:"size:100;index:idx_casbin_rule_unique"`
	V4    string `json:"v4" gorm:"size:100;index:idx_casbin_rule_unique"`
	V5    string `json:"v5" gorm:"size:100;index:idx_casbin_rule_unique"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (CasbinRule) TableName() string {
	return "sys_casbin_rule"
}
