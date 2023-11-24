package admin

import "shop-api/utils/hash"

// Service 服务
type Service struct {
	r Repository
}

// NewAdminService 实例化管理员服务
func NewAdminService(r Repository) *Service {
	// 迁移管理员表
	r.Migrate()
	// 插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// Create 创建管理员
func (c *Service) Create(admin *Admin) error {
	// 检查密码是否匹配
	if admin.Password != admin.Password2 {
		return ErrMismatchedPassword
	}
	// 检查管理员是否存在
	_, err := c.r.GetByUserName(admin.Username)
	if err == nil {
		return ErrUserIsExist
	}
	// 创建管理员
	err = c.r.Create(admin)
	return err
}

// GetAdminByID 通过id查找管理员信息
func (c *Service) GetAdminByID(userId uint) (Admin, error) {
	user, err := c.r.GetByUserId(userId)
	if err != nil {
		return Admin{}, ErrUserNotExist
	}
	return user, nil
}

// CheckUserAndPassword 检查管理员是否存在并检验密码是否正确
func (c *Service) CheckUserAndPassword(username, password string) (Admin, error) {
	user, err := c.r.GetByUserName(username)
	if err != nil {
		return Admin{}, ErrUserNotExist
	}
	match := hash.CheckPassword(password, user.Password)
	if !match {
		return Admin{}, ErrPasswordNotCorrect
	}
	return user, nil
}

// ChangePassword 更改管理员密码
func (c *Service) ChangePassword(user *Admin, oldPassword, newPassword string) error {
	// 检查密码是否匹配
	if newPassword != user.Password2 {
		return ErrMismatchedPassword
	}
	// 检查密码是否匹配
	match := hash.CheckPassword(oldPassword, user.Password)
	if !match {
		return ErrPasswordNotCorrect
	}
	// 加密新密码
	hashPassword, _ := hash.EncryptPassword(newPassword)
	user.Password = hashPassword
	// 更新用户信息
	return c.r.Update(user)
}

// Update 更新管理员信息
func (c *Service) Update(user *Admin) error {
	return c.r.Update(user)
}
