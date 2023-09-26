package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(sys.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(sys.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel sys.User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel sys.User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}

// GetById 通过 用户ID 获取用户Model
func GetById(idstr string) (userModel sys.User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

// GetByLoginId 通过 loginId 获取用户Model
func GetByLoginId(loginId string) (userModel sys.User) {
	database.DB.Where("name", loginId).First(&userModel)
	return
}

// GetOtpSecret 通过 loginId 获取用户opt密钥用于验证OPT
func GetOtpSecret(loginId string) (userModel sys.User) {
	database.DB.First(&userModel, "name", loginId)
	return
}

// UpdateOpt 更新OTP info
func UpdateOpt(name, key, url string) (userModel sys.User) {
	result := database.DB.First(&userModel, "name", name)
	if result.Error != nil {
		return
	}
	dataToUpdate := sys.User{
		OtpSecret:  key,
		OtpAuthUrl: url,
	}
	result.Updates(dataToUpdate)
	return
}

// DisableOpt 关闭OTP info
func DisableOpt(name string) (userModel sys.User) {
	result := database.DB.First(&userModel, "name", name)
	if result.Error != nil {
		return
	}
	dataToUpdate := sys.User{
		OtpEnabled: false,
	}
	result.Updates(dataToUpdate)
	return
}

// SetOptStatus 设置用户OPT 状态
func SetOptStatus(name string) (userModel sys.User) {
	result := database.DB.First(&userModel, "name", name)
	if result.Error != nil {
		return
	}
	dataToUpdate := sys.User{
		OtpEnabled:  true,
		OtpVerified: true,
	}
	result.Updates(dataToUpdate)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel sys.User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

// UserAll 获取所有用户数据
func UserAll() (users []sys.User) {
	database.DB.Find(&users)
	return
}

// UserDeletelById 删除用户
func UserDeletelById(id int) {
	database.DB.Where("id=?", id).Delete(&sys.User{})

}
