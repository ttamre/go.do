package api

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedOn   string
	Completed   bool
}

func NewTodo(title string, description string) *Todo {
	// INSERT INTO todos (id, list_id, title, description, completed) VALUES (uuid.New(), ListID, Title, Description, false)
	return &Todo{
		Title:       title,
		Description: description,
		CreatedOn:   time.Now().String(),
		Completed:   false,
	}
}

func UpdateTitle(item *Todo, title string) {
	// UPDATE todos SET title = Title WHERE id = Item.ID
	item.Title = title
}

func UpdateDescription(item *Todo, description string) {
	// UPDATE todos SET description = Description WHERE id = Item.ID
	item.Description = description
}

func UpdateCompletion(item *Todo) {
	// UPDATE todos SET completed = true WHERE id = Item.ID
	item.Completed = !item.Completed
}
