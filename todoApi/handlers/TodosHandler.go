package handlers

import (
	"encoding/json"
	"fmt"
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

func GetTodoHandler(w http.ResponseWriter, r *http.Request, mu *sync.RWMutex, todos helper.TodosMap) {
	id, err := helper.GetIdParam(r)

	if err != nil {
		mesgErr := "ID invalido"
		RespondWithError(w, r, http.StatusBadRequest, &mesgErr)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	t, ok := todos[id]

	if !ok {
		mess := fmt.Sprintf("Todo with id %s not found", id)
		RespondWithError(w, r, http.StatusNotFound, &mess)
	}

	ResponseWithJson(w, r, http.StatusOK, helper.FormatTodo(t))
}
