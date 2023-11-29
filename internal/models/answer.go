package models

// Answer 表示回答的数据模型
type Answer struct {
	ID         int64  `db:"id"`
	QuestionID int64  `db:"question_id"`
	UserID     int64  `db:"user_id"`
	Content    string `db:"content"`
	// 根据实际情况添加其他字段
}
