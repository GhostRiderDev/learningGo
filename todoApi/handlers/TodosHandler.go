package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ghostriderdev/goTodoApi/helper"
	"github.com/ghostriderdev/goTodoApi/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request, mu *sync.RWMutex, todos helper.TodosMap) {
	mu.RLock()

	defer mu.RUnlock()

	t := make([]*models.Todo, 0, len(todos))

	for key, _ := range todos {
		t = append(t, todos[key])
	}

	ResponseWithJson(w, r, http.StatusOK, helper.FT{"todos": t})
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request, mu *sync.RWMutex, todos helper.TodosMap, validate *validator.Validate) {
	t := &models.Todo{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(t); err != nil {
		errMsg := err.Error()
		RespondWithError(w, r, http.StatusBadRequest, &errMsg)
	}


	err := validate.Struct(t)

	if err != nil {
		errMsg := err.Error()
		RespondWithError(w, r, http.StatusBadRequest, &errMsg)
		return
	}

	t.Id = uuid.New()
	mu.Lock()
	defer mu.Unlock()

	todos[t.Id] = t
	ResponseWithJson(w, r, http.StatusCreated, helper.FormatTodo(t))
}
