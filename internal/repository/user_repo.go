package repository

import (
	"community-qa-platform/internal/models"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

// NewUserRepo 创建一个UserRepo实例
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// CreateUserRepo 创建一个用户
func (u *UserRepo) CreateUserRepo(user *models.User) error {
	sqlStr := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`

	_, err := u.db.Exec(sqlStr, user.UserName, user.PassWord, user.Email)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserRepo 更新用户
func (u *UserRepo) UpdateUserRepo(user *models.User) error {
	sqlStr := `UPDATE users SET username=?, password=?, email=? WHERE id=?`

	_, err := u.db.Exec(sqlStr, user.UserName, user.PassWord, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetUserRepo 根据ID获取用户
func (u *UserRepo) GetUserRepo(id int64) (*models.User, error) {
	sqlStr := `SELECT id, username, password, email FROM users WHERE id=?`

	row := u.db.QueryRow(sqlStr, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return user, nil
}

// DeleteUserRepo 删除用户
func (u *UserRepo) DeleteUserRepo(id int64) error {
	sqlStr := `DELETE FROM users WHERE id=?`

	_, err := u.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
