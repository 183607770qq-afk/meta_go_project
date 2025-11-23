package handler

import (
    // "net/http"
    
    "blog_system/internal/app/service"
    "blog_system/pkg/response"
    
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

// Register 用户注册
// @Summary 用户注册
// @Description 创建新用户账户
// @Tags auth
// @Accept json
// @Produce json
// @Param request body service.RegisterRequest true "注册信息"
// @Success 201 {object} response.SuccessResponse{data=service.RegisterResponse}
// @Failure 400 {object} response.ErrorResponse
// @Router /api/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
    var req service.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "请求参数错误: "+err.Error())
        return
    }
    
    resp, err := h.userService.Register(c.Request.Context(), &req)
    if err != nil {
        response.BadRequest(c, err.Error())
        return
    }
    
    response.Created(c, resp, "注册成功")
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录获取token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body service.LoginRequest true "登录信息"
// @Success 200 {object} response.SuccessResponse{data=service.LoginResponse}
// @Failure 400 {object} response.ErrorResponse
// @Router /api/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
    var req service.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "请求参数错误: "+err.Error())
        return
    }
    
    resp, err := h.userService.Login(c.Request.Context(), &req)
    if err != nil {
        response.BadRequest(c, err.Error())
        return
    }
    
    response.OK(c, resp, "登录成功")
}

// GetProfile 获取用户资料
// @Summary 获取用户资料
// @Description 获取当前登录用户资料
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.SuccessResponse{data=service.ProfileResponse}
// @Failure 401 {object} response.ErrorResponse
// @Router /api/users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        response.Unauthorized(c, "用户未认证")
        return
    }
    
    resp, err := h.userService.GetProfile(c.Request.Context(), userID.(uint))
    if err != nil {
        response.BadRequest(c, err.Error())
        return
    }
    
    response.OK(c, resp, "获取成功")
}