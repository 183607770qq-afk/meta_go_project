package service

import (
	"blog_system/internal/domain/entity"
	"blog_system/internal/domain/repository"
	"context"
)

//	type PostServiceInterface interface {
//		CreatePost(ctx context.Context) error
//	}
type PostService struct {
	postRope repository.PostRepository
}


func NewPostService(postRope repository.PostRepository) *PostService {
	return &PostService{postRope: postRope}
}

type CreatePostRequest struct {
	Title   string
	Content string
	UserId  uint
}
type UpdatePostRequest struct {
	ID      uint
	Title   string
	Content string
	UserId  uint
}

func (s *PostService) CreatePost(ctx context.Context, req *CreatePostRequest) error {
	post := &entity.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserId,
	}
	return s.postRope.Create(ctx, post)
}
func (s *PostService) GetPostByID(ctx context.Context, id uint) (*entity.Post, error) {
	return s.postRope.FindByID(ctx, id)
}

type RequestListPosts struct {
	Page     int
	PageSize int
}

func (s *PostService) ListPosts(ctx context.Context, page, pageSize int) ([]*entity.Post, int64, error) {
	return s.postRope.ListPosts(ctx, page, pageSize)
}
func (s *PostService) UpdatePost(ctx context.Context, req *UpdatePostRequest) error {
	post := &entity.Post{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserId}
	return s.postRope.Update(ctx, post)
}

func (s *PostService) DeletePost(context context.Context, id uint) error {
	return s.postRope.Delete(context, id)
}