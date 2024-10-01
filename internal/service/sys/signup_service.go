package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/handleerror"
)

type SignupService struct{}

// Signup 注册
func (ss *SignupService) Signup(userModel sys.User) (data sys.User, errMsg string) {
	err := database.DB.Create(&userModel).Error
	errMsg = handleerror.GormError(err)
	if errMsg != "" {
		return sys.User{}, errMsg
	}
	return userModel, errMsg
}
