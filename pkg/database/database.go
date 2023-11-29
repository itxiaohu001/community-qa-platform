package database

import (
	"community-qa-platform/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DB 是一个全局数据库连接实例
var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB(db *config.Db) {
	if db == nil {
		log.Fatalf("nil db config")
	}
	var err error

	// 打开数据库连接
	switch db.DBType {
	case "mysql":
		connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, db.DBName)
		DB, err = sql.Open("mysql", connectStr)
	case "postgres":
		connectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			db.Host, db.Port, db.User, db.Password, db.DBName)
		DB, err = sql.Open("postgres", connectStr)
	default:
		connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, db.DBName)
		DB, err = sql.Open("mysql", connectStr)
	}

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// 测试数据库连接
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Database connected")
}
