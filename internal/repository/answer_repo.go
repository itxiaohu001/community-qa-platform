package repository

import (
	"community-qa-platform/internal/models"
	"database/sql"
)

type AnswerRepo struct {
	db *sql.DB
}

// NewAnswerRepo 创建一个新的AnswerRepo实例
func NewAnswerRepo(db *sql.DB) *AnswerRepo {
	return &AnswerRepo{
		db: db,
	}
}

// CreateAnswer 创建新的回答
func (repo *AnswerRepo) CreateAnswer(answer *models.Answer) error {
	sqlStr := `INSERT INTO answers (question_id, user_id, content) VALUES (?, ?, ?)`

	_, err := repo.db.Exec(sqlStr, answer.QuestionID, answer.UserID, answer.Content)
	if err != nil {
		return err
	}

	return nil
}

// GetAnswerByID 根据ID获取回答
func (repo *AnswerRepo) GetAnswerByID(id int64) (*models.Answer, error) {
	sqlStr := `SELECT id, question_id, user_id, content FROM answers WHERE id = ?`

	row := repo.db.QueryRow(sqlStr, id)

	answer := &models.Answer{}
	err := row.Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 回答不存在
		}
		return nil, err
	}

	return answer, nil
}

// UpdateAnswer 更新回答
func (repo *AnswerRepo) UpdateAnswer(answer *models.Answer) error {
	sqlStr := `UPDATE answers SET question_id = ?, user_id = ?, content = ? WHERE id = ?`

	_, err := repo.db.Exec(sqlStr, answer.QuestionID, answer.UserID, answer.Content, answer.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAnswer 删除回答
func (repo *AnswerRepo) DeleteAnswer(id int64) error {
	sqlStr := `DELETE FROM answers WHERE id = ?`

	_, err := repo.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}

// 其他可能的函数，例如获取特定问题的所有回答等
