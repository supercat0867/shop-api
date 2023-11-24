package category

import "gorm.io/gorm"

// Category 商品分类模型
type Category struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Desc     string
	IsActive bool
}

// NewCategory 实例化商品分类
func NewCategory(name string, desc string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}
}
