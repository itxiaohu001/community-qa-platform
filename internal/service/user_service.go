package service

import (
	"community-qa-platform/internal/models"
	"community-qa-platform/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

// CreateUser 创建用户
func (u *UserService) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserService) UpdateUser(user *models.User) error {
	return u.repo.UpdateUser(user)
}

func (u *UserService) GetUser(id int64) (*models.User, error) {
	return u.repo.GetUser(id)
}

func (u *UserService) DeleteUser(id int64) error {
	return u.repo.DeleteUser(id)
}
