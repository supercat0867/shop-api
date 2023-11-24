package admin

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// NewAdminRepository 实例化管理员仓库
func NewAdminRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migrate 迁移管理员表
func (r *Repository) Migrate() {
	err := r.db.AutoMigrate(&Admin{})
	if err != nil {
		log.Printf("管理员表迁移错误：%v", err)
	}
}

// InsertSampleData 添加测试数据
func (r *Repository) InsertSampleData() {
	admin := NewAdmin("admin", "初始管理员", "123456", "123456")
	r.db.FirstOrCreate(&admin, Admin{Username: admin.Username})
}

// Create 创建管理员
func (r *Repository) Create(admin *Admin) error {
	result := r.db.Create(admin)
	return result.Error
}

// Update 更新管理员信息
func (r *Repository) Update(admin *Admin) error {
	return r.db.Save(admin).Error
}

// GetByUserName 通过用户名查找管理员
func (r *Repository) GetByUserName(username string) (Admin, error) {
	var admin Admin
	err := r.db.Where("Username = ?", username).First(&admin).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, nil
}

// GetByUserId 通过id查找管理员
func (r *Repository) GetByUserId(id uint) (Admin, error) {
	var admin Admin
	err := r.db.First(&admin, id).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, nil
}
