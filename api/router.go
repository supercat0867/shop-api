package api

import (
	"github.com/gin-gonic/gin"
	"log"
	adminApi "shop-api/api/admin"
	categoryApi "shop-api/api/category"
	productApi "shop-api/api/product"
	"shop-api/config"
	"shop-api/domain/admin"
	"shop-api/domain/category"
	"shop-api/domain/product"
	"shop-api/utils/database"
	"shop-api/utils/middleware"
)

// Databases 数据库
type Databases struct {
	adminRepository    *admin.Repository
	categoryRepository *category.Repository
	productRepository  *product.Repository
}

// AppConfig 配置信息
var AppConfig = &config.Configuration{}

// CreateDBs 获取配置信息、创建数据库
func CreateDBs() *Databases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalf("配置文件读取失败：%v", err)
	}
	AppConfig = conf
	db := database.NewMySqlDB(AppConfig.DatabaseURL)
	return &Databases{
		adminRepository:    admin.NewAdminRepository(db),
		categoryRepository: category.NewCategoryRepository(db),
		productRepository:  product.NewProductRepository(db),
	}
}

// RegisterHandlers 注册所有控制器
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreateDBs()
	RegisterAdminController(r, dbs)
	RegisterCategoryController(r, dbs)
	RegisterProductController(r, dbs)
}

// RegisterAdminController 注册管理员控制器
func RegisterAdminController(r *gin.Engine, dbs Databases) {
	adminService := admin.NewAdminService(*dbs.adminRepository)
	adminController := adminApi.NewAdminController(adminService, AppConfig)
	adminGroup := r.Group("/admin")
	adminGroup.POST("", adminController.CreateAdmin)
	adminGroup.POST("/login", adminController.Login)
	adminGroup.PATCH("/passwd", middleware.AuthUserMiddleware(AppConfig.SecretKey), adminController.ChangePassword)
}

// RegisterCategoryController 注册商品分类控制器
func RegisterCategoryController(r *gin.Engine, dbs Databases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := r.Group("/category")
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.POST("", middleware.AuthUserMiddleware(AppConfig.SecretKey), categoryController.CreateCategory)
	categoryGroup.POST("/upload", middleware.AuthUserMiddleware(AppConfig.SecretKey), categoryController.BulkCreateCategory)
	categoryGroup.PATCH("", middleware.AuthUserMiddleware(AppConfig.SecretKey), categoryController.Update)
}

// RegisterProductController 注册商品控制器
func RegisterProductController(r *gin.Engine, dbs Databases) {
	productService := product.NewProductService(*dbs.productRepository)
	productController := productApi.NewProductController(productService)
	productGroup := r.Group("/product")
	productGroup.GET("", productController.GetProducts)
	productGroup.POST("", middleware.AuthUserMiddleware(AppConfig.SecretKey), productController.CreateProduct)
	productGroup.PATCH("", middleware.AuthUserMiddleware(AppConfig.SecretKey), productController.UpdateProduct)
	productGroup.PUT("/discontinue", middleware.AuthUserMiddleware(AppConfig.SecretKey), productController.DiscontinueProduct)
	productGroup.DELETE("", middleware.AuthUserMiddleware(AppConfig.SecretKey), productController.DeleteProduct)
}
