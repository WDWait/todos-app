package main

import (
	"log"
	"os"
	"todos-app/internal/handler"
	"todos-app/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 初始化数据库
	database.InitDB()
	defer database.DB.Close()

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"db":     "connected",
		})
	})

	// API 路由器
	api := r.Group("/api")
	{
		todoHandler := handler.NewTodoHandler()
		todos := api.Group("/todos")
		{
			todos.GET("", todoHandler.GetAllTodos)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.POST("", todoHandler.CreateTodo)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}

	// 获取端口
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// 启动
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
