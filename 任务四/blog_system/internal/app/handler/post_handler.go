package handler

import (
	"blog_system/internal/app/service"
	"blog_system/pkg/response"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req service.CreatePostRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}
	if err := h.postService.CreatePost(c.Request.Context(), &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Created(c, nil, "创建文章成功")
}
func (h *PostHandler) GetPostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}
	data, err := h.postService.GetPostByID(c.Request.Context(), uint(id))
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if data == nil {
		response.NotFound(c, "文章不存在")
		return

	}
	response.OK(c, data, "获取文章成功")
}
func (h *PostHandler) ListPosts(c *gin.Context) {
	var req service.RequestListPosts
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}
	data, total, err := h.postService.ListPosts(c.Request.Context(), req.Page, req.PageSize)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	fmt.Println(data)

	m := make(map[string]interface{})
	m["total"] = total
	m["posts"] = data
	response.OK(c, m, "获取文章列表成功")
}
func (h *PostHandler) UpdatePost(c *gin.Context) {
	req := service.UpdatePostRequest{}
	userID, exist := c.Get("userID")
	if !exist {
		response.BadRequest(c, "用户未认证")
	}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}
	if userID != req.UserId {
		response.Forbidden(c, "没有权限修改该文章")
		return

	}
	err := h.postService.UpdatePost(c.Request.Context(), &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.OK(c, nil, "更新文章成功")

}
func (h *PostHandler) DeletePost(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		response.BadRequest(c, "用户未认证")
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "请求参数错误: "+err.Error())
	}
	data, err := h.postService.GetPostByID(c.Request.Context(), uint(id))
	if err != nil {
		response.NotFound(c, "文章不存在")
		return
	}

	if userID != data.UserID {
		response.Forbidden(c, "没有权限删除该文章")
		return

	}
	err = h.postService.DeletePost(c.Request.Context(), uint(id))
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.OK(c, nil, "删除文章成功")

}
