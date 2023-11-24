package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	// DefaultPageSize 默认页面大小
	DefaultPageSize = 20
	// MaxPageSize 最大页面限制
	MaxPageSize = 1000
	PageVar     = "page"
	PageSizeVar = "pageSize"
)

// Pages 分页结构体
type Pages struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	PageCount  int         `json:"pageCount"`
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

// New 实例化分页结构体
func New(page, pageSize, total int) *Pages {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		// 计算总页数
		pageCount = (total + pageSize - 1) / pageSize
		// 页码超出总页数
		if page > pageCount {
			page = pageCount
		}
	}
	if page <= 0 {
		page = 1
	}
	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// ParseInt 类型转换
func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// NewFromRequest 根据gin请求实例化分页结构体
func NewFromRequest(c *gin.Context, count int) *Pages {
	page := ParseInt(c.Query(PageVar), 1)
	pageSize := ParseInt(c.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// Offset 计算偏移
func (p *Pages) Offset() int {
	return p.Page - 1
}

// Limit 计算limit
func (p *Pages) Limit() int {
	return p.PageSize
}
