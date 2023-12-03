package models

type Question struct {
	ID      int64  `db:"id"`
	UserID  int64  `db:"user_id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}
