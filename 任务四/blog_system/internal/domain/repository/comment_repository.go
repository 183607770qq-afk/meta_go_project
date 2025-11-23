package repository

import (
	"blog_system/internal/domain/entity"
	"context"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) error
	GetCommentByPostyID(ctx context.Context, id uint) (*[]entity.Comment, error)
}
