package response

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

type SuccessResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func Success(c *gin.Context, code int, data interface{}, message string) {
    c.JSON(code, SuccessResponse{
        Code:    code,
        Message: message,
        Data:    data,
    })
}

func Error(c *gin.Context, code int, message string) {
    c.JSON(code, ErrorResponse{
        Code:    code,
        Message: message,
    })
}

// 快捷方法
func OK(c *gin.Context, data interface{}, message string) {
    Success(c, http.StatusOK, data, message)
}

func Created(c *gin.Context, data interface{}, message string) {
    Success(c, http.StatusCreated, data, message)
}

func BadRequest(c *gin.Context, message string) {
    Error(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
    Error(c, http.StatusUnauthorized, message)
}

func Forbidden(c *gin.Context, message string) {
    Error(c, http.StatusForbidden, message)
}

func NotFound(c *gin.Context, message string) {
    Error(c, http.StatusNotFound, message)
}

func InternalServerError(c *gin.Context, message string) {
    Error(c, http.StatusInternalServerError, message)
}