package api

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	ListID      uuid.UUID
	Title       string
	Description string
	Completed   bool
}

func NewTodo(userID uuid.UUID, listID uuid.UUID, title string, description string) *Todo {
	// INSERT INTO todos (id, list_id, title, description, completed) VALUES (uuid.New(), ListID, Title, Description, false)
	return &Todo{
		ID:          uuid.New(),
		UserID:      userID,
		ListID:      listID,
		Title:       title,
		Description: description,
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

func Complete(item *Todo) {
	// UPDATE todos SET completed = true WHERE id = Item.ID
	item.Completed = true
}

// func Print(item *Todo) {
// 	fmt.Printf("Title: %s\nDescription: %s\nCompleted: %t\n\n", item.Title, item.Description, item.Completed)
// }
