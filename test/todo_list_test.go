package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ttamre/go.do/api"
)

func TestNewTodoList(t *testing.T) {
	// only calling user because it's required for NewTodoList
	testUser := api.NewUser("test", "password")
	todoList := api.NewTodoList(testUser, "Test List")

	assert.NotNil(t, todoList)
	assert.NotEqual(t, uuid.Nil, todoList.ID)
	assert.NotEqual(t, uuid.Nil, todoList.UserID)
	assert.NotNil(t, todoList.Title)
	assert.Empty(t, todoList.List)
}

func TestAdd(t *testing.T) {
	testUser := api.NewUser("test", "password")
	todoList := api.NewTodoList(testUser, "Test List")
	item := &api.Todo{ID: uuid.New(), Title: "Test Title", Description: "Test Description", Completed: false}

	api.Add(todoList, item)

	assert.Len(t, todoList.List, 1)
	assert.Equal(t, item, todoList.List[0])
}

func TestDelete(t *testing.T) {
	testUser := api.NewUser("test", "password")
	todoList := api.NewTodoList(testUser, "Test List")
	item := &api.Todo{ID: uuid.New(), Title: "Test Title", Description: "Test Description", Completed: false}

	api.Add(todoList, item)
	api.Delete(todoList, item)

	assert.Len(t, todoList.List, 0)
}

func TestUpdate(t *testing.T) {
	testUser := api.NewUser("test", "password")
	todoList := api.NewTodoList(testUser, "Test List")
	api.Update(todoList, "Updated Title")

	assert.Equal(t, "Updated Title", todoList.Title)
}

// func TestPrintList(t *testing.T) {
// 	// no assertions, just verify that it's formatted correctly and doesn't crash
// 	testUser := api.NewUser("test", "password")
// 	todoList := api.NewTodoList(testUser, "Test List")
// 	item := &api.Todo{ID: uuid.New(), Title: "Test Print", Description: "todo_list_test.go", Completed: false}

// 	api.Add(todoList, item)
// 	api.PrintList(todoList)
// }
