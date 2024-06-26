package api

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedOn   string
	Completed   bool
}

/*
Create a new todo struct

Parameters:
  - id				string	: UUID of the todo item
  - title			string	: Title of the todo item
  - description		string	: Description of the todo item
  - created_on		string	: Date the todo item was created
  - completed		string	: Whether the todo item is completed or not

Returns:
  - *Todo					: Pointer to a new todo item
*/
func NewTodoFromDB(id string, title string, description string, created_on string, completed string) *Todo {
	// Parse strings into proper types
	uuidType, err := uuid.Parse(id)
	if err != nil {
		log.Fatalf("Failed to parse UUID: %v", err)
	}

	// Parse "completed" string into a boolean
	completedBool, err := strconv.ParseBool(completed)
	if err != nil {
		// Instead of failing, log a warning and proceed with default value of false
		log.Printf("WARNING: Failed to parse boolean: %v", err)
		log.Println("       : proceeding with default value - false")
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

/*
Get a list of todo items from the database

Parameters:
  - rdb		*redis.Client	: Redis database client
  - r		*http.Request	: HTTP request object

Returns:
  - []*Todo					: List of pointers for todo items
*/
func GetTodos(rdb *redis.Client, r *http.Request) []*Todo {
	// Get a list of keys (UUIDs) and store them in 'keys'
	keys, err := rdb.Keys(r.Context(), "*").Result()
	if err != nil {
		log.Fatalf("Failed to get keys: %v", err)
	}

	var todoList []*Todo
	for _, key := range keys {
		// Get todo items based on UUID (key)
		result, err := rdb.HGetAll(r.Context(), key).Result()
		if err != nil {
			log.Fatalf("Failed to get key: %v", err)
		}

		// Create todo struct and add to todoList for templating
		todo := NewTodoFromDB(
			key,
			result["title"],
			result["description"],
			result["created_on"],
			result["completed"])

		todoList = append(todoList, todo)
	}

	// Sort page data by creation date
	sort.Slice(todoList, func(i, j int) bool { return todoList[i].CreatedOn < todoList[j].CreatedOn })
	return todoList
}

/*
Create a new todo entry in the database using form data

Paramaters:
  - rdb		*redis.Client	: Redis database client
  - r		*http.Request	: HTTP request object
*/
func AddTodo(rdb *redis.Client, r *http.Request) {
	err := rdb.HMSet(r.Context(), uuid.New().String(), map[string]interface{}{
		"title":       r.FormValue("title"),
		"description": r.FormValue("description"),
		"created_on":  time.Now().Format(time.RFC1123),
		"completed":   "false",
	}).Err()

	if err != nil {
		log.Fatalf("Failed to create new todo: %v", err)
	}
	log.Println("Added new todo item to database")
}

/*
Update title of an item in the database

Parameters:
  - rdb		*redis.Client	: Redis database client
  - r		*http.Request	: HTTP request object
*/
func UpdateTitle(rdb *redis.Client, r *http.Request) {
	// Get ID from query parameters
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]

	err := rdb.HSet(r.Context(), id, "title", r.FormValue("title")).Err()
	if err != nil {
		log.Fatalf("Failed to update title: %v", err)
	}
	log.Printf("Updated todo list: title -> %s", r.FormValue("title"))
}

/*
Update description of an item in the database

Parameters:
  - rdb		*redis.Client	: Redis database client
  - r		*http.Request	: HTTP request object
*/
func UpdateDescription(rdb *redis.Client, r *http.Request) {
	// Get ID from query parameters
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]

	err := rdb.HSet(r.Context(), id, "description", r.FormValue("description")).Err()
	if err != nil {
		log.Fatalf("Failed to update description: %v", err)
	}
	log.Printf("Updated todo list: description -> %s", r.FormValue("description"))
}

/*
Update completion status of an item in the database

Parameters:
  - rdb		*redis.Client	: Redis database client
  - r		*http.Request	: HTTP request object
*/
func UpdateCompletion(rdb *redis.Client, r *http.Request) {
	// Get ID from query parameters
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]

	err := rdb.HSet(r.Context(), id, "completed", r.FormValue("completed")).Err()
	if err != nil {
		log.Fatalf("Failed to update completion: %v", err)
	}
	log.Printf("Updated todo list: completed -> %s", r.FormValue("completed"))
}

/*
Delete a todo item from the database

Parameters:
- rdb	*redis.Client	: Redis database client
- r		*http.Request	: HTTP request object
*/
func DeleteTodo(rdb *redis.Client, r *http.Request) {
	// Get ID from query parameters
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]

	_, err := rdb.Del(r.Context(), id).Result()
	if err != nil {
		log.Fatalf("Failed to delete todo: %v", err)
	}
	log.Println("Deleted todo item from database")
}
