package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/ghostriderdev/goTodoApi/config"
	"github.com/ghostriderdev/goTodoApi/handlers"
	"github.com/ghostriderdev/goTodoApi/helper"
	"github.com/go-playground/validator/v10"
)

var (
	todos    = make(helper.TodosMap)
	validate = validator.New(validator.WithRequiredStructEnabled())
	mu       = &sync.RWMutex{}
)

func handlerOne(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetTodosHandler(w, r, mu, todos)
	case http.MethodPost:
		handlers.CreateTodoHandler(w, r, mu, todos, validate)
	default:
		message := "Method not allowed"
		handlers.RespondWithError(w, r, http.StatusMethodNotAllowed, &message)

	}
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/api/v1/health", handlers.HealthHandler)
	server.HandleFunc("/api/v1/todos", handlerOne)

	log.Println("Server is running at http://localhost:6060/")
	log.Fatal(http.ListenAndServe(":6060", config.Cors(server)))
}
