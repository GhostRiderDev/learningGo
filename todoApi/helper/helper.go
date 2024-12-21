package helper

import (
	"github.com/ghostriderdev/goTodoApi/models"
	"github.com/google/uuid"
)

type FT map[string]interface{}
type TodosMap map[uuid.UUID]*models.Todo


func FormatTodo(t *models.Todo) *FT {
	return &FT{
		"id":           t.Id,
		"title":        t.Title,
		"is_completed": t.IsCompleted,
	}
}