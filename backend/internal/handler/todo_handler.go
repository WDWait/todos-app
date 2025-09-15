package handler

import (
	"net/http"
	"strconv"
	"todos-app/internal/model"
	"todos-app/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		service: service.NewTodoService(),
	}
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo item by ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Produce json
// @Success 204 "No Content"
// @Router /api/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := h.service.DeleteTodo(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo item by ID
// @Tags todos
// @Accept json
// @Param id path int true "Todo ID"
// @Param todo body model.Todo true "Todo object"
// @Produce json
// @Success 200 {object} model.Todo
// @Router /api/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTodo(id, &todo); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Param todo body model.Todo true "Todo object"
// @Produce json
// @Success 201 {object} model.Todo
// @Router /api/todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, todo)
}

// GetTodoByID godoc
// @Summary Get todo by ID
// @Description Get a todo by its ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Produce json
// @Success 200 {object} model.Todo
// @Router /api/todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	todo, err := h.service.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todo)
}

// GetAllTodos godoc
// @Summary Get all todos
// @Description Get a list of all todos
// @Tags todos
// @Produce json
// @Success 200 {array} model.Todo
// @Router /api/todos [get]
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
