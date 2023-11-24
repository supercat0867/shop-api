package category

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

// Repository 商品分类仓库
type Repository struct {
	db *gorm.DB
}

// NewCategoryRepository 实例化商品分类仓库
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migrate 迁移商品分类表
func (r *Repository) Migrate() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Printf("商品分类表迁移错误：%v", err)
	}
}

// InsertSampleData 插入测试数据
func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "Category1", Desc: "Category1"},
		{Name: "Category2", Desc: "Category2"},
	}
	for _, c := range categories {
		r.db.FirstOrCreate(&c, Category{Name: c.Name})
	}
}

// Create 创建商品分类
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByName 通过名称查询商品分类
func (r *Repository) GetByName(name string) (Category, error) {
	var category Category
	result := r.db.Where("Name = ?", name).First(&category)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return category, ErrCategoryNameNotExists
	}
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

// GetById 通过id查询商品分类
func (r *Repository) GetById(id uint) (Category, error) {
	var category Category
	result := r.db.First(&category, id)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return category, ErrCategoryNameNotExists
	}
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

// BulkCreate 批量创建商品分类
func (r *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categories).Count(&count).Error
	return int(count), err
}

// GetAll 获取所有商品分类(分页)
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
}

// Update 更新商品分类
func (r *Repository) Update(c *Category) error {
	err := r.db.Save(c).Error
	return err
}

// Delete todo 删除商品分类，检查是否存在改分类的商品，如果存在就报错，请先删除或修改分类中的商品
// Delete 删除商品分类
func (r *Repository) Delete(c *Category) {

}
