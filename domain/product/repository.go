package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

// Repository 商品仓库
type Repository struct {
	db *gorm.DB
}

// NewProductRepository 实例化商品仓库
func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migrate 迁移商品表
func (r *Repository) Migrate() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		log.Printf("商品表迁移错误：%v", err)
	}
}

// InsertSampleData 插入测试数据
func (r *Repository) InsertSampleData() {
	products := []Product{
		{Name: "Product1", Desc: "测试商品1", CategoryID: 1, Price: 998, StockCount: 3},
		{Name: "Product2", Desc: "测试商品2", CategoryID: 2, Price: 998, StockCount: 4},
	}
	for _, c := range products {
		r.db.FirstOrCreate(&c, Product{Name: c.Name})
	}
}

// Create 创建商品
func (r *Repository) Create(p *Product) error {
	result := r.db.Create(p)
	if result.Error != nil {
		return ErrCreateProduct
	}
	return nil
}

// FindByID 根据商品id查找商品（过滤掉已下架商品）
func (r *Repository) FindByID(id uuid.UUID) (*Product, error) {
	var product *Product
	err := r.db.Where("IsDiscontinued = ?", 0).First(&product, id).First(&product).Error
	if err != nil {
		return nil, ErrProductNotExists
	}
	return product, nil
}

// FindByIDAll 根据商品id查找商品
func (r *Repository) FindByIDAll(id uuid.UUID) (*Product, error) {
	var product *Product
	err := r.db.First(&product, id).First(&product).Error
	if err != nil {
		return nil, ErrProductNotExists2
	}
	return product, nil
}

// GetAll 查询所有商品（分页）（过滤掉已下架商品）
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64
	r.db.Where("IsDiscontinued = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

// SearchByString 搜索商品返回分页结果（过滤掉已下架商品）
func (r *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where("IsDiscontinued = ?", 0).Where(
		"Name LIKE ? OR `Desc` Like ?", convertedStr, convertedStr).
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

// Update 更新商品信息
func (r *Repository) Update(id uuid.UUID, updateProduce *Product) error {
	savedProduct, err := r.FindByID(id)
	if err != nil {
		return err
	}
	err = r.db.Model(&savedProduct).Updates(updateProduce).Error
	return err
}

// DeleteByID 通过商品id删除商品(软删除)
func (r *Repository) DeleteByID(id uuid.UUID) error {
	result := r.db.Where("ID = ?", id).Delete(&Product{})
	return result.Error
}

// DiscontinuedByID 通过商品id下架/上架商品
func (r *Repository) DiscontinuedByID(id uuid.UUID) error {
	currentProduct, err := r.FindByIDAll(id)
	if err != nil {
		return err
	}
	currentProduct.IsDiscontinued = !currentProduct.IsDiscontinued
	err = r.db.Save(currentProduct).Error
	return err
}
