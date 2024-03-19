package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ttamre/go.do/api"
)

func TestNewTodoList(t *testing.T) {
	// Test NewTodoList function
	todoList := api.NewTodoList()
	assert.NotNil(t, todoList)
	assert.NotEqual(t, uuid.Nil, todoList.ID)
	assert.Empty(t, todoList.List)
}

func TestAdd(t *testing.T) {
	// Test Add function
	todoList := api.NewTodoList()
	item := &api.Todo{ID: uuid.New(), Title: "Test Title", Description: "Test Description", Completed: false}
	api.Add(todoList, item)
	assert.Len(t, todoList.List, 1)
	assert.Equal(t, item, todoList.List[0])
}

func TestDelete(t *testing.T) {
	// Test Delete function
	todoList := api.NewTodoList()
	item := &api.Todo{ID: uuid.New(), Title: "Test Title", Description: "Test Description", Completed: false}
	api.Add(todoList, item)
	api.Delete(todoList, item)
	assert.Len(t, todoList.List, 0)
}

func TestPrintList(t *testing.T) {
	// Test PrintList function (no assertions, just verify that it doesn't crash)
	todoList := api.NewTodoList()
	item := &api.Todo{ID: uuid.New(), Title: "Test Print", Description: "todo_list_test.go", Completed: false}
	api.Add(todoList, item)
	api.PrintList(todoList)
}
