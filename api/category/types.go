package category

// CreateCategoryRequest 创建商品分类请求
type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// UpdateCategoryRequest 更新商品分类请求
type UpdateCategoryRequest struct {
	ID   uint   `json:"ID"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
