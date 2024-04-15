package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

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
		// Fetch all todo items from database
		todoList := api.GetTodos(rdb, r)

		// serve main webpage with todo items as template data
		tmpl := template.Must(template.ParseFiles("./web/index.html"))
		tmpl.Execute(w, todoList)
	}

	// Adds new todo item to database
	if r.Method == "POST" {
		// TODO validate form data in helper function
		api.AddTodo(rdb, r)

		// Redirect to main webpage
		r.Method = "GET"
		r.URL.Path = "/"
		RootHandler(w, r)
	}
}

/* Serves an individual + editable todo item webpage */
func TodoIDHandler(w http.ResponseWriter, r *http.Request) {
	// Update the todo item
	if r.Method == "POST" {
		// TODO server-side form data validation

		// Update title
		if r.FormValue("title") != "" {
			api.UpdateTitle(rdb, r)
		}

		// Update description
		if r.FormValue("description") != "" {
			api.UpdateDescription(rdb, r)
		}

		// Update completion status
		if r.FormValue("completed") != "" {
			api.UpdateCompletion(rdb, r)
		}
	}

	// Delete the todo item
	if r.Method == "DELETE" {
		api.DeleteTodo(rdb, r)
	}
}
