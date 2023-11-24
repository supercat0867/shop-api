package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"shop-api/config"
	"shop-api/domain/admin"
	"shop-api/utils/api_helper"
	jwthelper "shop-api/utils/jwt"
	"strconv"
	"time"
)

// Controller 管理员控制器
type Controller struct {
	adminService *admin.Service
	appConfig    *config.Configuration
}

// NewAdminController 实例化管理员控制器
func NewAdminController(service *admin.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		adminService: service,
		appConfig:    appConfig,
	}
}

// CreateAdmin godoc
// @Summary 创建管理员
// @Tags Auth
// @Accept json
// @Product json
// @Param CreateUserRequest body CreateUserRequest true "管理员信息"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} api_helper.ErrResponse
// @Router /admin [post]
func (c *Controller) CreateAdmin(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	newUser := admin.NewAdmin(req.Username, req.Nickname, req.Password, req.Password2)
	err := c.adminService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, CreateUserResponse{Username: req.Username})
}

// Login godoc
// @Summary 管理员登录
// @Tags Auth
// @Accept json
// @Product json
// @Param LoginRequest body LoginRequest true "管理员账号密码"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} api_helper.ErrResponse
// @Router /admin/login [post]
func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	currentUser, err := c.adminService.CheckUserAndPassword(req.Username, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	decodedClaims := jwthelper.VerifyToken(currentUser.Token, c.appConfig.JwtSettings.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"nickname": currentUser.NickName,
				"roleId":   currentUser.RoleID,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp":      time.Now().Add(48 * time.Hour).Unix(),
			})
		token, err := jwthelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
		if err != nil {
			api_helper.HandleError(g, api_helper.ErrGenerateJwt)
			return
		}
		currentUser.Token = token
	}
	currentUser.LastIP = g.ClientIP()
	timeNow := time.Now()
	currentUser.LastLogin = &timeNow
	err = c.adminService.Update(&currentUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, LoginResponse{
		Username: currentUser.Username,
		UserId:   currentUser.ID,
		Token:    currentUser.Token,
	})
}

// ChangePassword godoc
// @Summary 修改密码
// @Tags Auth
// @Accept json
// @Product json
// @Param Authorization header string true "Authorization header"
// @Param ChangePasswordRequest body ChangePasswordRequest true "旧密码、新密码、重复密码"
// @Success 200 {object} Response
// @Failure 400 {object} api_helper.ErrResponse
// @Router /admin/passwd [patch]
func (c *Controller) ChangePassword(g *gin.Context) {
	var req ChangePasswordRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	// 获取用户id
	userId := api_helper.GetUserId(g)
	// 获取用户信息
	userInfo, err := c.adminService.GetAdminByID(userId)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userInfo.Password2 = req.Password2
	err = c.adminService.ChangePassword(&userInfo, req.OldPassword, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, Response{Message: "success"})
}
