package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeCreate 创建商品，生成唯一id
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	pid := uuid.New()
	p.ID = pid
	return nil
}
