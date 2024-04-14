package api

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   string    `json:"created_on"`
	Completed   bool      `json:"completed"`
}

func NewTodo(title string, description string) *Todo {
	return &Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		CreatedOn:   time.Now().Format(time.RFC1123),
		Completed:   false,
	}
}

func NewTodoFromDB(id string, title string, description string, created_on string, completed string) *Todo {
	// Parse strings into proper types
	uuidType, err := uuid.Parse(id)
	if err != nil {
		log.Fatalf("Failed to parse UUID: %v", err)
	}

	completedBool, err := strconv.ParseBool(completed)
	if err != nil {
		log.Printf("Failed to parse completed: %v", err)
		completedBool = false
	}

	return &Todo{
		ID:          uuidType,
		Title:       title,
		Description: description,
		CreatedOn:   created_on,
		Completed:   completedBool,
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
