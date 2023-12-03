package repo

import (
	"community-qa-platform/internal/config"
	"community-qa-platform/internal/models"
	"community-qa-platform/internal/repository"
	"community-qa-platform/pkg/database"
	"testing"
)

func NewDefaultRepo() *repository.UserRepo {
	dbConf := config.InitConfig("./config.json")
	database.InitDB(dbConf.Db)
	return repository.NewUserRepo(database.DB)
}

func TestCreateUserRepo(t *testing.T) {
	repo := NewDefaultRepo()

	user := &models.User{
		UserName: "xiaohu4",
		PassWord: "fsafdasdf",
		Email:    "test4@qq.com",
	}
	if err := repo.CreateUser(user); err != nil {
		t.Error(err)
		return
	}
}
