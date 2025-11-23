package database

import (
	"blog_system/internal/domain/entity"
	"blog_system/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db: db}
}
func (c *commentRepository) Create(ctx context.Context, comment *entity.Comment) error {
	return c.db.Create(comment).Error
}
func (c *commentRepository) GetCommentByPostyID(ctx context.Context, id uint) (*[]entity.Comment, error) {
	var comments []entity.Comment
	err := c.db.Model(&entity.Comment{}).Where("post_id", id).Find(&comments).Error
	if err != nil {
		return nil, err  
	}
	return &comments, nil
}
