package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ttamre/go.do/api"
)

func TestNewTodo(t *testing.T) {
	todo := api.NewTodo("Test Title", "Test Description", uuid.New())
	assert.NotNil(t, todo)
	assert.NotEqual(t, uuid.Nil, todo.ID)
	assert.NotEqual(t, uuid.Nil, todo.ListID)
	assert.Equal(t, "Test Title", todo.Title)
	assert.Equal(t, "Test Description", todo.Description)
	assert.False(t, todo.Completed)
}

func TestUpdateTitle(t *testing.T) {
	todo := &api.Todo{ID: uuid.New(), Title: "Old Title"}
	api.UpdateTitle(todo, "New Title")
	assert.Equal(t, "New Title", todo.Title)
}

func TestUpdateDescription(t *testing.T) {
	todo := &api.Todo{ID: uuid.New(), Description: "Old Description"}
	api.UpdateDescription(todo, "New Description")
	assert.Equal(t, "New Description", todo.Description)
}

func TestComplete(t *testing.T) {
	todo := &api.Todo{ID: uuid.New(), Completed: false}
	api.Complete(todo)
	assert.True(t, todo.Completed)
}

func TestPrint(t *testing.T) {
	// no assertions, just verify that it;s formatted correctly and doesn't crash
	todo := &api.Todo{ID: uuid.New(), Title: "Test Print", Description: "todo_test.go", Completed: true}
	api.Print(todo)
}
