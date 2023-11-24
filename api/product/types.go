package product

// CreateProductRequest 创建商品请求
type CreateProductRequest struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	CategoryID uint    `json:"cid"`
	Count      int     `json:"count"`
}

// UpdateProductRequest 更新商品请求
type UpdateProductRequest struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	CategoryID uint    `json:"cid"`
	Count      int     `json:"count"`
}

// Request 商品ID请求
type Request struct {
	ID string `json:"id"`
}
