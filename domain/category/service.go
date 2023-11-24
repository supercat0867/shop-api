package category

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mime/multipart"
	"shop-api/utils/csv_helper"
	"shop-api/utils/pagination"
)

// Service 商品分类服务
type Service struct {
	r Repository
}

// NewCategoryService 实例化商品分类服务
func NewCategoryService(r Repository) *Service {
	// 迁移商品分类表
	r.Migrate()
	// 插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// Create 创建单个商品分类
func (c *Service) Create(category *Category) error {
	// 检查商品分类是否已存在
	exist, _ := c.r.GetByName(category.Name)
	if exist.ID != 0 {
		return ErrCategoryNameExists
	}
	err := c.r.Create(category)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改商品分类
func (c *Service) Update(category *Category) error {
	// 检查是否存在除了当前分类以外同名的分类
	var exist Category
	result := c.r.db.Where("Name = ? AND ID <> ?", category.Name, category.ID).First(&exist)

	// 检查是否找到记录并且没有错误
	if result.Error == nil && exist.ID != 0 {
		return ErrCategoryNameExists
	}
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	// 更新商品分类
	return c.r.Update(category)
}

// BulkCreate 通过csv批量创建商品分类
func (c *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := c.r.BulkCreate(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetAll 获取全部商品分类（分页）
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}

// GetCategoryByID 根据商品分类Id查找分类
func (c *Service) GetCategoryByID(id uint) (Category, error) {
	return c.r.GetById(id)
}
