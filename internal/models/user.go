package models

type User struct {
	ID       int64  `db:"id"`
	UserName string `db:"username"`
	PassWord string `db:"password"`
	Email    string `db:"email"`
}
