package admin

import "github.com/pkg/errors"

var (
	ErrMismatchedPassword = errors.New("两次密码不相同")
	ErrUserIsExist        = errors.New("已存在此管理员")
	ErrUserNotExist       = errors.New("不存在此管理员")
	ErrPasswordNotCorrect = errors.New("密码错误")
)
