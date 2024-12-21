package helper

import (
	"errors"
	"net/http"
	"strings"

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

func GetIdParam(r *http.Request) (uuid.UUID, error) {
	// Remove the "/api/v1/todos/" prefix to get the UUID part
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/todos/")
	if id == "" {
		return uuid.Nil, errors.New("missing todo id in url")
	}
	return uuid.Parse(id)
}
