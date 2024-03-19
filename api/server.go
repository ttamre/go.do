package api

import (
	"context"

	"github.com/google/uuid"
)

type Server interface {
	GetTodo(context.Context, uuid.UUID) (*Todo, error)
	GetTodos(context.Context, uuid.UUID) ([]*Todo, error)
	UpdateTitle(context.Context, *Todo, string) error
	UpdateDescription(context.Context, *Todo, string) error
	UpdateCompleted(context.Context, *Todo, bool) error
	Delete(context.Context, *Todo) error
	DeleteList(context.Context, []*Todo) error
}

type TodoServer struct{ url string }

func NewTodoServer(url string) *TodoServer { return &TodoServer{url: url} }

/* Get a TODO item by UUID */
func GetTodo(ctx context.Context, id uuid.UUID) (*Todo, error) {
	// SELECT * FROM todos WHERE id = id
	return nil, nil
}

/* Get all TODO items by UUID */
func GetTodos(ctx context.Context, id uuid.UUID) ([]*Todo, error) {
	// SELECT * FROM todos WHERE id = id
	return nil, nil
}
