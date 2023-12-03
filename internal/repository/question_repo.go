package repository

import (
	"community-qa-platform/internal/models"
	"database/sql"
)

type QuestionRepo struct {
	db *sql.DB
}

// NewQuestionRepo 创建一个新的QuestionRepo实例
func NewQuestionRepo(db *sql.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

// CreateQuestion 创建新的问题
func (q *QuestionRepo) CreateQuestion(question *models.Question) error {
	sqlStr := `INSERT INTO questions (user_id, title, content) VALUES (?, ?, ?)`

	_, err := q.db.Exec(sqlStr, question.UserID, question.Title, question.Content)
	if err != nil {
		return err
	}

	return nil
}

// GetQuestion 根据ID获取问题
func (q *QuestionRepo) GetQuestion(id int64) (*models.Question, error) {
	sqlStr := `SELECT id, user_id, title, content FROM questions WHERE id = ?`

	row := q.db.QueryRow(sqlStr, id)

	question := &models.Question{}
	err := row.Scan(&question.ID, &question.UserID, &question.Title, &question.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 问题不存在
		}
		return nil, err
	}

	return question, nil
}

// UpdateQuestion 更新问题
func (q *QuestionRepo) UpdateQuestion(question *models.Question) error {
	sqlStr := `UPDATE questions SET user_id = ?, title = ?, content = ? WHERE id = ?`

	_, err := q.db.Exec(sqlStr, question.UserID, question.Title, question.Content, question.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteQuestion 删除问题
func (q *QuestionRepo) DeleteQuestion(id int64) error {
	sqlStr := `DELETE FROM questions WHERE id = ?`

	_, err := q.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
