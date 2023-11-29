package handlers

import (
	"community-qa-platform/internal/models"
	"community-qa-platform/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AnswerHandler 用于处理与回答相关的请求
type AnswerHandler struct {
	service *service.AnswerService
}

// NewAnswerHandler 创建一个新的AnswerHandler实例
func NewAnswerHandler(service *service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		service: service,
	}
}

// CreateAnswer 处理创建回答的请求
func (h *AnswerHandler) CreateAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.BindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateAnswer(&answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, answer)
}

// GetAnswer 处理获取单个回答的请求
func (h *AnswerHandler) GetAnswer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的回答ID"})
		return
	}

	answer, err := h.service.GetAnswerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// UpdateAnswer 处理更新回答的请求
// DeleteAnswer 处理删除回答的请求
// ... 其他函数
