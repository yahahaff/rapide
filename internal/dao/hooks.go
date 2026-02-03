package dao

import (
	"rapide/internal/models/sys"
	"rapide/pkg/hash"
	"gorm.io/gorm"
)

type User struct {
	sys.User
}

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用,用于密码加密
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
