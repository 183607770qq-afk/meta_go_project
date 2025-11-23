package service

import (
	"context"
	"errors"

	"blog_system/internal/domain/entity"
	"blog_system/internal/domain/repository"
	"blog_system/pkg/auth"
)

type UserService struct {
    userRepo   repository.UserRepository
    jwtService *auth.JWTService
}

func NewUserService(userRepo repository.UserRepository, jwtService *auth.JWTService) *UserService {
    return &UserService{
        userRepo:   userRepo,
        jwtService: jwtService,
    }
}

type RegisterRequest struct {
    Name string `json:"Name" binding:"required,min=3,max=50"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type RegisterResponse struct {
    ID       uint   `json:"id"`
    Name string `json:"Name"`
    Email    string `json:"email"`
    Token    string `json:"token"`
}

func (s *UserService) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
    // 检查用户名是否已存在
    existingUser, err := s.userRepo.FindByName(ctx, req.Name)
    if err == nil && existingUser != nil {
        return nil, errors.New("用户名已存在")
    }
    
    // 检查邮箱是否已存在
    existingUser, err = s.userRepo.FindByEmail(ctx, req.Email)
    if err == nil && existingUser != nil {
        return nil, errors.New("邮箱已存在")
    }
    
    // 创建用户
    user := &entity.User{
        Name: req.Name,
        Email:    req.Email,
    }
    
    if err := user.SetPassword(req.Password); err != nil {
        return nil, errors.New("密码加密失败")
    }
    
    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, errors.New("创建用户失败")
    }
    
    // 生成JWT token
    token, err := s.jwtService.GenerateToken(user.ID, user.Name)
    if err != nil {
        return nil, errors.New("生成token失败")
    }
    
    return &RegisterResponse{
        ID:       user.ID,
        Name: user.Name,
        Email:    user.Email,
        Token:    token,
    }, nil
}

type LoginRequest struct {
    Name string `json:"Name" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
    ID       uint   `json:"id"`
    Name string `json:"Name"`
    Email    string `json:"email"`
    Token    string `json:"token"`
}

func (s *UserService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
    // 查找用户
    user, err := s.userRepo.FindByName(ctx, req.Name)
    if err != nil {
        return nil, errors.New("用户不存在")
    }
    
    // 验证密码
    if !user.CheckPassword(req.Password) {
        return nil, errors.New("密码错误")
    }
    
    // 生成JWT token
    token, err := s.jwtService.GenerateToken(user.ID, user.Name)
    if err != nil {
        return nil, errors.New("生成token失败")
    }
    
    return &LoginResponse{
        ID:       user.ID,
        Name: user.Name,
        Email:    user.Email,
        Token:    token,
    }, nil
}

type ProfileResponse struct {
    ID        uint   `json:"id"`
    Name  string `json:"Name"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
}

func (s *UserService) GetProfile(ctx context.Context, userID uint) (*ProfileResponse, error) {
    user, err := s.userRepo.FindByID(ctx, userID)
    if err != nil {
        return nil, errors.New("用户不存在")
    }
    
    return &ProfileResponse{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
    }, nil
}