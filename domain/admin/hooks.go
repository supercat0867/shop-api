package admin

import (
	"gorm.io/gorm"
	"shop-api/utils/hash"
)

// BeforeCreate 创建管理员，给明文密码加盐哈希加密
func (u *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	hashPassword, err := hash.EncryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword
	return nil
}
