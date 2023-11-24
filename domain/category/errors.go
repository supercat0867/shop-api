package category

import "errors"

var (
	ErrCategoryNameExists    = errors.New("已存在此商品分类")
	ErrCategoryNameNotExists = errors.New("不存在此商品分类")
)
