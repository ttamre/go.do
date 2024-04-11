package api

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Password  string
	TodoLists []*TodoList
}

// Accepts a a username and password
// Inserts new user into database
// Returns a pointer to the new user
func NewUser(username string, password string) *User {
	// INSERT INTO users (id, username, password) VALUES (uuid.New(), username, password)
	return &User{ID: uuid.New(), Username: username, Password: password, TodoLists: []*TodoList{}}
}

// Accepts a user and a username
// Updates the user's username in database and in memory
func UpdateUsername(user *User, username string) {
	// UPDATE users SET username = username WHERE id = user.ID
	user.Username = username
}

// Accepts a user and a password
// Updates the user's password in database and in memory
func UpdatePassword(user *User, password string) {
	// UPDATE users SET password = password WHERE id = user.ID
	user.Password = password
}

// Accepts a user
// Deletes user from database
func DeleteUser(user *User) {
	// DELETE FROM users WHERE id = user.ID
}

/* TODO list functions */
func CreateTodoList(user *User, title string) *TodoList {
	// INSERT INTO todo_lists (id, title) VALUES (uuid.New(), title)
	todoList := NewTodoList(user, title)
	user.TodoLists = append(user.TodoLists, todoList)
	return todoList
}

func CreateTodo(user *User, list *TodoList, title string, description string) *Todo {
	// INSERT INTO todos (id, user_id, list_id title, description, completed) VALUES (uuid.New(), UserID, ListID, Title, Description, false)
	todo := NewTodo(user.ID, list.ID, title, description)
	Add(list, todo)
	return todo

}

func GetTodoLists(user *User) []*TodoList {
	// SELECT * FROM todo_lists WHERE user_id = user.ID
	return user.TodoLists
}
