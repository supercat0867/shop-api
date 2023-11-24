package product

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"shop-api/domain/product"
	"shop-api/utils/api_helper"
	"shop-api/utils/pagination"
)

type Controller struct {
	productService *product.Service
}

// NewProductController 实例化商品控制器
func NewProductController(service *product.Service) *Controller {
	return &Controller{
		productService: service,
	}
}

// GetProducts godoc
// @Summary 获取全部商品（分页）
// @Tags Product
// @Accept json
// @Produce json
// @Param qt query string false "搜索匹配的sku和商品名"
// @Param page query int false "页码"
// @Param pageSize query int false "页面大小"
// @Success 200 {object} pagination.Pages
// @Router /product [get]
func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromRequest(g, -1)
	queryText := g.Query("qt")
	if queryText != "" {
		page = c.productService.SearchProduct(queryText, page)
	} else {
		page = c.productService.GetAll(page)
	}
	g.JSON(http.StatusOK, page)
}

// CreateProduct godoc
// @Summary 创建单个商品
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param CreateProductRequest body CreateProductRequest true "商品信息"
// @Success 201 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /product [post]
func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	// 实例化商品
	newProduct := product.NewProduct(req.Name, req.Desc, req.CategoryID, req.Price, req.Count)
	err := c.productService.Create(newProduct)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, api_helper.Response{Code: 201, Message: "success"})
}

// UpdateProduct godoc
// @Summary 修改商品信息
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param UpdateProductRequest body UpdateProductRequest true "商品信息"
// @Success 200 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /product [patch]
func (c *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	updateProduct := product.NewProduct(req.Name, req.Desc, req.CategoryID, req.Price, req.Count)
	u, err := uuid.Parse(req.ID)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err = c.productService.Update(u, updateProduct)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Code: 200, Message: "success"})
}

// DiscontinueProduct godoc
// @Summary 上架/下架商品
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param Request body Request true "商品ID"
// @Success 200 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /product/discontinue [put]
func (c *Controller) DiscontinueProduct(g *gin.Context) {
	var req Request
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err := c.productService.Discontinued(req.ID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Code: 200, Message: "success"})
}

// DeleteProduct godoc
// @Summary 删除商品（软删除）
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param Request body Request true "商品ID"
// @Success 200 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /product [delete]
func (c *Controller) DeleteProduct(g *gin.Context) {
	var req Request
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err := c.productService.Delete(req.ID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Code: 200, Message: "success"})
}
