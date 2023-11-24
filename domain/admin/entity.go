package admin

import (
	"gorm.io/gorm"
	"time"
)

// Admin 管理员模型
type Admin struct {
	gorm.Model
	Username  string     `gorm:"comment:登录名"`
	Password  string     `gorm:"comment:密码"`
	Password2 string     `gorm:"-"`
	NickName  string     `gorm:"comment:昵称"`
	Token     string     `gorm:"comment:token"`
	RoleID    uint       `gorm:"comment:角色id"`
	LastLogin *time.Time `gorm:"comment:上次登录时间"`
	LastIP    string     `gorm:"comment:上次登录ip"`
}

// NewAdmin 实例化管理员
func NewAdmin(username, nickname, password, password2 string) *Admin {
	return &Admin{
		Username:  username,
		NickName:  nickname,
		Password:  password,
		Password2: password2,
	}
}
