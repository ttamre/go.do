package api

import (
	"github.com/google/uuid"
)

type TodoList struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Title  string
	List   []*Todo
}

// Accepts a pointer to a user and a title
// Returns a pointer to a new TodoList{}
func NewTodoList(user *User, title string) *TodoList {
	// INSERT INTO todo_lists (id, title) VALUES (uuid.New(), title)
	return &TodoList{ID: uuid.New(), UserID: user.ID, Title: title, List: []*Todo{}}
}

func Add(todoList *TodoList, item *Todo) {
	// INSERT INTO todos (id, list_id, title, description, completed) VALUES (Item.ID, List.ID, Item.Title, Item.Description, Item.Completed)
	todoList.List = append(todoList.List, item)
}

func Delete(todoList *TodoList, item *Todo) {
	// DELETE FROM todos WHERE id = Item.ID
	for i, todo := range todoList.List {
		if todo.ID == item.ID {
			todoList.List = append(todoList.List[:i], todoList.List[i+1:]...)
			break
		}
	}
	item = &Todo{}
}

func Update(todoList *TodoList, title string) {
	// UPDATE todos SET title = Title WHERE id = Item.ID
	todoList.Title = title
}

// func PrintList(todoList *TodoList) {
// 	for _, item := range todoList.List {
// 		Print(item)
// 	}
// }
