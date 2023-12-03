package service

import (
	"community-qa-platform/internal/models"
	"community-qa-platform/internal/repository"
)

type QuestionService struct {
	repo *repository.QuestionRepo
}

// NewQuestionService 创建一个新的 QuestionService 实例
func NewQuestionService(repo *repository.QuestionRepo) *QuestionService {
	return &QuestionService{repo: repo}
}

// CreateQuestion 创建一个问题
func (q *QuestionService) CreateQuestion(question *models.Question) error {
	return q.repo.CreateQuestion(question)
}

// UpdateQuestion 更新回答
func (q *QuestionService) UpdateQuestion(question *models.Question) error {
	return q.repo.UpdateQuestion(question)
}

// GetQuestion 获取回答
func (q *QuestionService) GetQuestion(id int64) (*models.Question, error) {
	return q.repo.GetQuestion(id)
}

// DeleteQuestion 删除回答
func (q *QuestionService) DeleteQuestion(id int64) error {
	return q.repo.DeleteQuestion(id)
}
