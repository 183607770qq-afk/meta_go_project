package repository

import (
    "context"
    
    "blog_system/internal/domain/entity"
)

type PostRepository interface {
    Create(ctx context.Context, post *entity.Post) error
    FindByID(ctx context.Context, id uint) (*entity.Post, error)
	ListPosts(ctx context.Context, page, pageSize int) ([]*entity.Post, int64, error)
    // FindByAuthorID(ctx context.Context, authorID uint) ([]*entity.Post, error)
    Update(ctx context.Context, post *entity.Post) error
    Delete(ctx context.Context, id uint) error
    // List(ctx context.Context, page, pageSize int) ([]*entity.Post, int64, error)
}