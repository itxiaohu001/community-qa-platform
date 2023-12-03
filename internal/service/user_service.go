package service

import (
	"community-qa-platform/internal/models"
	"community-qa-platform/internal/repository"
	"community-qa-platform/internal/utils"
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
	user.PassWord = utils.GetMd5ByStr(user.PassWord)

	return u.repo.CreateUser(user)
}

func (u *UserService) UpdateUser(user *models.User) error {
	return u.repo.UpdateUser(user)
}

func (u *UserService) GetUserByID(id int64) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserService) GetUserByName(name string) (*models.User, error) {
	return u.repo.GetUserByName(name)
}

func (u *UserService) CheckUserNamePassword(n, p string) (bool, error) {
	user, err := u.repo.GetUserByName(n)
	if err != nil {
		return false, err
	}

	if utils.GetMd5ByStr(p) == user.PassWord {
		return true, nil
	}

	return false, nil
}

func (u *UserService) DeleteUser(id int64) error {
	return u.repo.DeleteUser(id)
}
