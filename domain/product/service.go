package product

import (
	"github.com/google/uuid"
	"shop-api/utils/pagination"
)

// Service 商品服务
type Service struct {
	r Repository
}

// NewProductService 实例化商品服务
func NewProductService(r Repository) *Service {
	// 迁移商品表
	r.Migrate()
	// 插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// GetAll 获取全部商品（分页）
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

// SearchProduct 查找商品
func (c *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := c.r.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

// Create 创建商品
func (c *Service) Create(p *Product) error {
	return c.r.Create(p)
}

// Update 更新商品信息
func (c *Service) Update(id uuid.UUID, p *Product) error {
	return c.r.Update(id, p)
}

// Discontinued 上架/下架商品
func (c *Service) Discontinued(id string) error {
	pid, err := uuid.Parse(id)
	if err != nil {
		return ErrProductNotExists2
	}
	return c.r.DiscontinuedByID(pid)
}

// Delete 软删除商品
func (c *Service) Delete(id string) error {
	pid, err := uuid.Parse(id)
	if err != nil {
		return ErrProductNotExists2
	}
	return c.r.DeleteByID(pid)
}
