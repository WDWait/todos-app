package database

import (
	"log"
	"os"
	"todos-app/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局变量
var DB *gorm.DB

// InitDB 初始化
func InitDB() {
	var err error
	dsn := os.Getenv("DB_DSN")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = DB.AutoMigrate(&model.Todo{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}
	log.Println("Connected to MySQL database")
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting sql.DB: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	} else {
		log.Println("Database connection closed successfully")
	}
}
