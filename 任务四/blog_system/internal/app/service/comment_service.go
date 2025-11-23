package service

import (
	"blog_system/internal/domain/entity"
	"blog_system/internal/domain/repository"
	"context"
)
type CommentService struct {
	commentRepo repository.CommentRepository
}
func NewCommentService(commentRepo repository.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}

}
type CreateCommentRequest struct {
	Content string
	PostID  uint
	UserID  uint
}
func (s *CommentService) CreateComment(ctx context.Context, req *CreateCommentRequest) error {
	comment := &entity.Comment{
		Content: req.Content,
		PostID:  req.PostID,
		UserID:  req.UserID,
	}
	return s.commentRepo.Create(ctx, comment)
}
func (s *CommentService) GetCommentByPostyID(ctx context.Context, id uint) (*[]entity.Comment, error) {
	return s.commentRepo.GetCommentByPostyID(ctx, id)
}