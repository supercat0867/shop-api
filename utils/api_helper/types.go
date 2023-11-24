package api_helper

import "github.com/pkg/errors"

// Response 通用响应
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrResponse 响应错误结构体
type ErrResponse struct {
	Message string `json:"msg"`
}

var (
	ErrInvalidBody = errors.New("请求参数不合法")
	ErrGenerateJwt = errors.New("token生成失败")
)
