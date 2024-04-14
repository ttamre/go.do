package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/ttamre/go.do/api"
)

const URL = "localhost"
const PORT = 5000
const DB_PORT = 5001

var rdb *redis.Client

func main() {
	// Create a file server to serve static files
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fileServer))

	// Define routes
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/todo/", TodoIDHandler)

	// Initialize database connection
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", URL, DB_PORT),
	})

	// Start server
	log.Printf("Listening on http://%s:%d...", URL, PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

/* Serves the main webpage with a list of todo items */
func RootHandler(w http.ResponseWriter, r *http.Request) {

	// Fetches all todo items and serves main webpage
	if r.Method == "GET" {
		// Get a list of keys (UUIDs) and store them in 'keys'
		keys, err := rdb.Keys(r.Context(), "*").Result()
		if err != nil {
			log.Fatalf("Failed to get keys: %v", err)
		}

		var todoList []api.Todo
		for _, key := range keys {
			// Get todo items based on UUID (key)
			result, err := rdb.HGetAll(r.Context(), key).Result()
			if err != nil {
				log.Fatalf("Failed to get key: %v", err)
			}

			// Create todo struct and add to todoList for templating
			todo := api.NewTodoFromDB(
				key,
				result["title"],
				result["description"],
				result["created_on"],
				result["completed"])

			todoList = append(todoList, *todo)
		}

		// Sort page data by creation date
		sort.Slice(todoList, func(i, j int) bool { return todoList[i].CreatedOn < todoList[j].CreatedOn })

		// serve main webpage with template data
		tmpl := template.Must(template.ParseFiles("./web/index.html"))
		tmpl.Execute(w, todoList)
	}

	// Adds new todo item to database
	if r.Method == "POST" {
		// TODO validate form data
		if r.FormValue("title") == "" {
			log.Println("Title is required")
			return
		}

		// Create new todo item
		// NOTE: We could save memory by directly inserting form values into the database,
		// but using a constructor allows for easier validation and error handling
		todo := api.NewTodo(r.FormValue("title"), r.FormValue("description"))

		// Add new todo to database
		err := rdb.HMSet(r.Context(), uuid.New().String(), map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"created_on":  todo.CreatedOn,
			"completed":   fmt.Sprintf("%t", todo.Completed),
		}).Err()

		if err != nil {
			log.Fatalf("Failed to create new todo: %v", err)
		}

		log.Println("Added new todo item to database")

		// Redirect to main webpage
		r.Method = "GET"
		r.URL.Path = "/"
		RootHandler(w, r)
	}
}

/* Serves an individual + editable todo item webpage */
func TodoIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get Todo ID from URL
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]

	// =============================================================
	// NOTE: Unused function, editing is now accessible on home page
	//       therefore, no need to GET individual todo item webpages
	// =============================================================
	// Get todo item from database and serve it with the webpage
	// if r.Method == "GET" && !(strings.Contains(r.URL.Path, "static")) {
	// 	// Fetch todos from database
	// 	result, err := rdb.HGetAll(r.Context(), id).Result()
	// 	if err != nil {
	// 		log.Fatalf("Failed to get todo: %v", err)
	// 	}

	// 	// Create todo struct from database data
	// 	todo := api.NewTodoFromDB(
	// 		id,
	// 		result["title"],
	// 		result["description"],
	// 		result["created_on"],
	// 		result["completed"])

	// 	// serve todo webpage with template data
	// 	tmpl := template.Must(template.ParseFiles("./web/todo.html"))
	// 	tmpl.Execute(w, todo)
	// }
	// =============================================================

	// Update the todo item
	if r.Method == "POST" {
		// TODO server-side form data validation

		// Update title
		if r.FormValue("title") != "" {
			err := rdb.HSet(r.Context(), id, "title", r.FormValue("title")).Err()
			if err != nil {
				log.Fatalf("Failed to update title: %v", err)
			}
			log.Printf("Updated todo list: title -> %s", r.FormValue("title"))
		}

		// Update description
		if r.FormValue("description") != "" {
			err := rdb.HSet(r.Context(), id, "description", r.FormValue("description")).Err()
			if err != nil {
				log.Fatalf("Failed to update description: %v", err)
			}
			log.Printf("Updated todo list: description -> %s", r.FormValue("description"))
		}

		// Update completion status
		if r.FormValue("completed") != "" {
			err := rdb.HSet(r.Context(), id, "completed", r.FormValue("completed")).Err()
			if err != nil {
				log.Fatalf("Failed to update completion: %v", err)
			}
			log.Printf("Updated todo list: completed -> %s", r.FormValue("completed"))
		}
	}

	// Delete the todo item
	if r.Method == "DELETE" {
		_, err := rdb.Del(r.Context(), id).Result()
		if err != nil {
			log.Fatalf("Failed to delete todo: %v", err)
		}
		log.Println("Deleted todo item")
	}
}
