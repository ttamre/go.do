package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ttamre/go.do/api"
)

func TestNewUser(t *testing.T) {
	username := "testuser"
	password := "testpassword"
	user := api.NewUser(username, password)
	assert.NotNil(t, user, "NewUser returned nil")
	assert.NotEqual(t, uuid.Nil, user.ID)
	assert.Equal(t, username, user.Username, "Unexpected username")
	assert.Equal(t, password, user.Password, "Unexpected password")
	assert.Empty(t, user.TodoLists)
}

func TestUpdateUsername(t *testing.T) {
	username := "testuser"
	newUser := api.NewUser(username, "password")
	newUsername := "updateduser"
	api.UpdateUsername(newUser, newUsername)
	assert.Equal(t, newUsername, newUser.Username, "Unexpected username")
}

func TestUpdatePassword(t *testing.T) {
	username := "testuser"
	password := "password"
	newUser := api.NewUser(username, password)
	newPassword := "newpassword"
	api.UpdatePassword(newUser, newPassword)
	assert.Equal(t, newPassword, newUser.Password, "Unexpected password")
}

func TestDeleteUser(t *testing.T) {}

func TestCreateTodoList(t *testing.T) {
	username := "testuser"
	user := api.NewUser(username, "password")
	title := "Test TodoList"
	todoList := api.CreateTodoList(user, title)
	assert.NotNil(t, todoList, "CreateTodoList returned nil")
	assert.Equal(t, title, todoList.Title, "Unexpected title")
	assert.Equal(t, 1, len(user.TodoLists), "Unexpected number of TodoLists")
}

func TestCreateTodo(t *testing.T) {
	username := "testuser"
	user := api.NewUser(username, "password")
	todoList := api.CreateTodoList(user, "Test TodoList")
	title := "Test Todo"
	description := "Test description"
	todo := api.CreateTodo(user, todoList, title, description)
	assert.NotNil(t, todo, "CreateTodo returned nil")
	assert.Equal(t, title, todo.Title, "Unexpected title")
	assert.Equal(t, description, todo.Description, "Unexpected description")
}

func TestGetTodoLists(t *testing.T) {
	username := "testuser"
	user := api.NewUser(username, "password")
	todoLists := api.GetTodoLists(user)
	assert.Empty(t, todoLists, "Unexpected number of TodoLists")
}
