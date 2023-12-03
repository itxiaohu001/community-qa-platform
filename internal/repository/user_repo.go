package repository

import (
	"community-qa-platform/internal/models"
	"database/sql"
	"errors"
)

type UserRepo struct {
	db *sql.DB
}

// NewUserRepo 创建一个UserRepo实例
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// CreateUser 创建一个用户
func (u *UserRepo) CreateUser(user *models.User) error {
	sqlStr := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`

	_, err := u.db.Exec(sqlStr, user.UserName, user.PassWord, user.Email)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser 更新用户
func (u *UserRepo) UpdateUser(user *models.User) error {
	sqlStr := `UPDATE users SET username=?, password=?, email=? WHERE id=?`

	_, err := u.db.Exec(sqlStr, user.UserName, user.PassWord, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID 根据ID获取用户
func (u *UserRepo) GetUserByID(id int64) (*models.User, error) {
	sqlStr := `SELECT id, username, password, email FROM users WHERE id=?`

	row := u.db.QueryRow(sqlStr, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetUserByName(name string) (*models.User, error) {
	sqlStr := `SELECT id, username, password, email FROM users WHERE username=?`

	row := u.db.QueryRow(sqlStr, name)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return user, nil
}

// DeleteUser 删除用户
func (u *UserRepo) DeleteUser(id int64) error {
	sqlStr := `DELETE FROM users WHERE id=?`

	_, err := u.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
