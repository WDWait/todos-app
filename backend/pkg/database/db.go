package database

import (
	"database/sql"
	"log"
	"os"
)

// DB 全局变量
var DB *sql.DB

// InitDB 初始化
func InitDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	open, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	if err = open.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	log.Println("Connected to MySQL database")
}
