package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-api/domain/category"
	"shop-api/utils/api_helper"
	"shop-api/utils/pagination"
)

// Controller 商品分类控制器
type Controller struct {
	categoryService *category.Service
}

// NewCategoryController 实例化商品分类控制器
func NewCategoryController(service *category.Service) *Controller {
	return &Controller{
		categoryService: service,
	}
}

// CreateCategory godoc
// @Summary 创建单个商品分类
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param CreateCategoryRequest body CreateCategoryRequest true "category information"
// @Success 201 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /category [post]
func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err := c.categoryService.Create(category.NewCategory(req.Name, req.Desc))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, api_helper.Response{Code: 201, Message: "success"})
}

// BulkCreateCategory godoc
// @Summary 上传csv批量创建商品分类
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param file formData file true "csv文件"
// @Success 200 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /category/upload [post]
func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Code: 200, Message: fmt.Sprintf("%s上传成功，新增%d个商品分类", fileHeader.Filename, count)})
}

// GetCategories godoc
// @Summary 获取全部商品分类（分页）
// @Tags Category
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "页面大小"
// @Success 200 {object} pagination.Pages
// @Router /category [get]
func (c *Controller) GetCategories(g *gin.Context) {
	page := pagination.NewFromRequest(g, -1)
	page = c.categoryService.GetAll(page)
	g.JSON(http.StatusOK, page)
}

// Update godoc
// @Summary 修改商品分类
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header"
// @Param UpdateCategoryRequest body UpdateCategoryRequest true "category information"
// @Success 200 {object} api_helper.Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /category [patch]
func (c *Controller) Update(g *gin.Context) {
	var req UpdateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	// 查找当前的分类名
	thisCategory, err := c.categoryService.GetCategoryByID(req.ID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	thisCategory.Name = req.Name
	thisCategory.Desc = req.Desc
	err = c.categoryService.Update(&thisCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Code: 200, Message: "success"})
}
