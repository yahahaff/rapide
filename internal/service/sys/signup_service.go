package sys

import (
	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	sysMod "github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/handleerror"
)

type SignupService struct{}

// IsPhoneExist 判断手机号是否存在
func (ss *SignupService) IsPhoneExist(phone string) (sysMod.User, error) {
	sysDao.IsPhoneExist(phone)
	return sysMod.User{}, nil
}

// IsEmailExist 判断手机号是否存在
func (ss *SignupService) IsEmailExist(email string) (sysMod.User, error) {
	sysDao.IsPhoneExist(email)
	return sysMod.User{}, nil
}

// Signup 注册
func (ss *SignupService) Signup(userModel sysMod.User) (data sysMod.User, errMsg string) {
	err := database.DB.Create(&userModel).Error
	errMsg = handleerror.GormError(err)
	if errMsg != "" {
		return sysMod.User{}, errMsg
	}
	return userModel, errMsg
}
