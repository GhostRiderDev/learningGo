package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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
		return
	}

	ResponseWithJson(w, r, http.StatusOK, helper.FormatTodo(t))
}

func UpdateTodoHanlder(w http.ResponseWriter, r *http.Request, mu *sync.RWMutex, todos helper.TodosMap, validate *validator.Validate) {
	id, err := helper.GetIdParam(r)
	if err != nil {
		messError := "Id invalido"
		RespondWithError(w, r, http.StatusBadRequest, &messError)
	}

	mu.Lock()
	defer mu.Unlock()

	todoFound, ok := todos[id]

	if !ok {
		messError := fmt.Sprintf("Todo with id: %s not found", id)
		RespondWithError(w, r, http.StatusBadRequest, &messError)
	}

	respT := &models.Todo{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(respT); err != nil {
		err := err.Error()
		RespondWithError(w, r, http.StatusBadRequest, &err)
		return
	}

	err = validate.Struct(respT)

	if err != nil {
		err := err.Error()
		RespondWithError(w, r, http.StatusBadRequest, &err)
		return
	}

	todoFound.Title = respT.Title
	todoFound.IsCompleted = respT.IsCompleted

	ResponseWithJson(w, r, http.StatusOK, helper.FormatTodo(todoFound))
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request, mu *sync.RWMutex, todos helper.TodosMap) {
	id, err := helper.GetIdParam(r)
	if err != nil {
		messError := err.Error()
		RespondWithError(w, r, http.StatusBadRequest, &messError)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	t, ok := todos[id]
	if !ok {
		messErr := fmt.Sprintf("todo with id '%v' not found", id)
		RespondWithError(w, r, http.StatusNotFound, &messErr)
		return
	}

	delete(todos, t.Id)
	log.Printf("::: %v ::: %v %v", http.StatusNoContent, r.Method, r.URL)
	w.WriteHeader(http.StatusNoContent)
}
