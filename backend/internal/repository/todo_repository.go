package repository

import (
	"errors"
	"fmt"
	"todos-app/internal/model"
	"todos-app/pkg/database"

	"gorm.io/gorm"
)

// TodoRepository 先定义一个Todos的结构体，后续可以给这个结构体添加方法。可以看作一个空接口
type TodoRepository struct{}

// NewTodoRepository 这个方法会返回当前的TodoRepository结构体，因为它用的了“&”
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

// Delete 删除待办事项
func (r *TodoRepository) Delete(id int) error {
	tx := database.DB.Delete(&model.Todo{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Update 更新待办事项
func (r *TodoRepository) Update(id int, todo *model.Todo) error {
	var existingTodo model.Todo
	// 先检查记录是否存在
	first := database.DB.First(&existingTodo, id)
	if first.Error != nil {
		return fmt.Errorf("todo not found")
	} else if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		return first.Error
	}

	// 使用 Model + Updates 进行更新
	// 注意：这里传入的是指针，GORM 会根据非零值字段进行更新
	database.DB.Model(&existingTodo).Updates(model.Todo{
		Title:     todo.Title,
		Completed: todo.Completed,
	})
	// 确保 updated_at 字段被正确更新
	// GORM 通常会自动处理 updated_at，但如果需要强制更新，可以这样：
	// result = database.DB.Model(&existingTodo).Updates(map[string]interface{}{"title": todo.Title, "completed": todo.Completed})
	// 或者，先查询再赋值再 Save
	return nil
}

// Create 创建新的待办事项
func (r *TodoRepository) Create(todo *model.Todo) error {
	create := database.DB.Create(todo)
	if create.Error != nil {
		return create.Error
	}
	// GORM 通常会在 Create 后自动将生成的 ID 赋值给 todo.ID
	return nil
}

// GetByID 根据 ID 获取单个待办事项
func (r *TodoRepository) GetByID(id int) (*model.Todo, error) {
	var todo model.Todo
	// 使用 GORM 的 First 方法，它会返回第一条记录
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("todo not found")
	}

	return &todo, nil
}

// GetAll 获取所有待办事项，按创建时间倒序排列
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	// 使用 GORM 的 Order 和 Find 方法
	result := database.DB.Order("created_at DESC").Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

/*
func (r *TodoRepository) Delete(id int) error {
	result, err := database.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}

func (r *TodoRepository) Update(id int, todo *model.Todo) error {
	result, err := database.DB.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", todo.Title, todo.Completed, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	result, err := database.DB.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = int(id)
	return nil
}

// GetByID 更具id获取
func (r *TodoRepository) GetByID(id int) (*model.Todo, error) {
	row := database.DB.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = ?", id)
	var todo model.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("todo not found")
	} else if err != nil {
		return nil, err
	}
	return &todo, nil
}

// GetAll 获取全部todos
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	// 开发中select语句一般不要写 *，虽然你要查全部字段
	rows, err := database.DB.Query("SELECT id, title, completed, created_at, updated_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	// 很重要! Query()后必须关闭资源。你可能疑惑为什么QueryRow() 不用close()，很简单，QueryRow() 里默认做了Close()操作。
	// 你可以这样写：defer func(rows *sql.Rows) {
	//	err := rows.Close()
	//	if err != nil {
	//		log.Fatalf("Error closing rows: %v", err)
	//	}
	//}(rows)
	// 但是一般都简化为
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

*/

/*
	是否在疑惑什么时候返回值用*，什么时候不用呢？

	在 Go 语言中，这两个方法返回值的指针使用差异主要与「零值」和「是否可能返回空结果」有关：

	1,*GetByID 为什么返回 model.Todo
		当根据 ID 找不到对应记录时，需要明确返回「空」状态（nil）
		如果返回 model.Todo（非指针），即使没有找到记录，也会返回一个「零值对象」（所有字段都是默认值）
		调用者无法区分「找到了一个全是默认值的记录」和「根本没找到记录」，所以用指针的 nil 状态来表示「未找到」

	2,GetAll 为什么返回 [] model.Todo（非指针切片）
		切片本身是引用类型，其底层结构包含指针
		当没有数据时，返回一个空切片（[]model.Todo{}）比返回 nil 更符合使用习惯
		空切片可以安全地进行遍历操作（不会 panic），而 nil 切片虽然在 Go 中也能遍历，但返回空切片更能表达「存在一个结果集但里面没有数据」的语义

	简单说就是：单个对象需要用指针区分「存在 / 不存在」，而切片本身的特性使其无需指针就能清晰表达空状态。
*/
