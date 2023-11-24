package product

import "errors"

var (
	ErrProductNotExists  = errors.New("商品不存在或已下架！")
	ErrProductNotExists2 = errors.New("商品不存在！")
	ErrCreateProduct     = errors.New("数据库异常，商品创建失败！")
)
