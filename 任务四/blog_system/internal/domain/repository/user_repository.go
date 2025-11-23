package repository

import (
    "context"
    
    "blog_system/internal/domain/entity"
)

type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
    FindByID(ctx context.Context, id uint) (*entity.User, error)
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
    FindByName(ctx context.Context, name string) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context, page, pageSize int) ([]*entity.User, int64, error)
}