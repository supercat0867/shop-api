package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop-api/domain/category"
	"time"
)

// Product 商品模型
type Product struct {
	ID             uuid.UUID `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt    `gorm:"index"`
	Name           string            `gorm:"comment:商品名"`
	Desc           string            `gorm:"comment:商品描述"`
	StockCount     int               `gorm:"comment:商品库存量"`
	Price          float32           `gorm:"comment:商品价格"`
	CategoryID     uint              `gorm:"comment:商品分类id"`
	Category       category.Category `json:"-"`
	CoverImageUrl  string            `gorm:"comment:商品封面图URL"`
	ImageURLs      string            `gorm:"comment:商品图片URLs;type:text"`
	IsDiscontinued bool              `gorm:"comment:是否下架"`
}

// NewProduct 实例化商品
func NewProduct(name, desc string, cid uint, price float32, count int) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		Price:      price,
		CategoryID: cid,
		StockCount: count,
	}
}
