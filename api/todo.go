package api

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	ListID      uuid.UUID
	Title       string
	Description string
	Completed   bool
}

func NewTodo(Title string, Description string, ListID uuid.UUID) *Todo {
	// INSERT INTO todos (id, list_id, title, description, completed) VALUES (uuid.New(), ListID, Title, Description, false)
	return &Todo{
		ID:          uuid.New(),
		ListID:      uuid.New(),
		Title:       Title,
		Description: Description,
		Completed:   false,
	}
}

func UpdateTitle(Item *Todo, Title string) {
	// UPDATE todos SET title = Title WHERE id = Item.ID
	Item.Title = Title
}

func UpdateDescription(Item *Todo, Description string) {
	// UPDATE todos SET description = Description WHERE id = Item.ID
	Item.Description = Description
}

func Complete(Item *Todo) {
	// UPDATE todos SET completed = true WHERE id = Item.ID
	Item.Completed = true
}

func Print(Item *Todo) {
	fmt.Printf("Title: %s\nDescription: %s\nCompleted: %t\n\n", Item.Title, Item.Description, Item.Completed)
}
