package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleError 错误处理
func HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrResponse{Message: err.Error()})
	c.Abort()
	return
}
