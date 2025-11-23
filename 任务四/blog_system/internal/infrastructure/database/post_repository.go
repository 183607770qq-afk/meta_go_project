package database

import (
	"blog_system/internal/domain/entity"
	"blog_system/internal/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}
func (r *postRepository) Create(ctx context.Context, post *entity.Post) error {
	return r.db.Create(post).Error
}
func (r *postRepository) FindByID(ctx context.Context, id uint) (*entity.Post, error) {
	var post entity.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}
func (r *postRepository) ListPosts(ctx context.Context, page, pageSize int) ([]*entity.Post, int64, error) {
	var posts []*entity.Post
	var total int64
	err := r.db.Model(&entity.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}
	fmt.Println("Total posts:", posts)
	return posts, total, nil
}
func (r *postRepository) Update(ctx context.Context, post *entity.Post) error {
	return r.db.Model(&entity.Post{}).Where("id = ?", post.ID).Updates(post).Error
}
func (r *postRepository) Delete(ctx context.Context, id uint) error {
	var post entity.Post
	err := r.db.Debug().First(&post, id).Error
	fmt.Println(post, "data")
	
	if err != nil {
		return err
	}

	return r.db.Delete(&entity.Post{}, id).Error
}
