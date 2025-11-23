package handler

import (
	"blog_system/internal/app/service"
	"blog_system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req service.CreateCommentRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}
	if err := h.commentService.CreateComment(c.Request.Context(), &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Created(c, nil, "创建评论成功")

}
func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	// 实现获取文章评论的处理逻辑
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}
	data, err := h.commentService.GetCommentByPostyID(c.Request.Context(), uint(id))
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.OK(c, data, "获取评论成功")
}
