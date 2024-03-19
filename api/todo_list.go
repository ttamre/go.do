package api

import (
	"github.com/google/uuid"
)

type TodoList struct {
	ID   uuid.UUID
	List []*Todo
}

func NewTodoList() *TodoList {
	// INSERT INTO todo_lists (id) VALUES (uuid.New())
	return &TodoList{ID: uuid.New(), List: []*Todo{}}
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

func PrintList(todoList *TodoList) {
	for _, item := range todoList.List {
		Print(item)
	}
}
