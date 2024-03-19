package api

import (
	"context"

	"github.com/google/uuid"
)

type Server interface {
	GetTodo(context.Context, uuid.UUID) (*Todo, error)
	GetTodos(context.Context, uuid.UUID) (*TodoList, error)
	UpdateTitle(context.Context, *Todo, string) error
	UpdateDescription(context.Context, *Todo, string) error
	UpdateCompleted(context.Context, *Todo, bool) error
	Delete(context.Context, *Todo) error
	DeleteList(context.Context, *TodoList) error
}

type TodoServer struct{ url string }

// Accepts a URL
// Returns a pointer to a new TodoServer{}
func NewTodoServer(url string) *TodoServer { return &TodoServer{url: url} }

// Accepts a context and the UUID of a Todo{} item
// Returns a pointer to the todo item
func GetTodo(ctx context.Context, id uuid.UUID) (*Todo, error) {
	return nil, nil
}

// Accepts a context and the UUID of a TodoList{}
// Returns a pointer to the TodoList{}
func GetTodos(ctx context.Context, id uuid.UUID) (*TodoList, error) {
	return nil, nil
}
