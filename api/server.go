package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

type Config struct {
	URL     string
	PORT    int
	DB_PORT int
}

/* Create a HTTP server that listens for requests based on config */
func Serve(config *Config) {
	// Create a file server to serve static files
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fileServer))

	// Define routes
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/todo/", TodoIDHandler)

	// Initialize database connection
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.URL, config.DB_PORT),
	})

	// Start server
	log.Printf("Listening on http://%s:%d...", config.URL, config.PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), nil)
}

/* Serves the main webpage with a list of todo items */
func RootHandler(w http.ResponseWriter, r *http.Request) {

	// Fetches all todo items and serves main webpage
	if r.Method == http.MethodGet {
		// Fetch all todo items from database
		todoList := GetTodos(rdb, r)

		// serve main webpage with todo items as template data
		tmpl := template.Must(template.ParseFiles("./web/index.html"))
		tmpl.Execute(w, todoList)
	}

	// Adds new todo item to database
	if r.Method == http.MethodPost {
		// TODO validate form data in helper function
		AddTodo(rdb, r)

		// Redirect to main webpage
		r.Method = "GET"
		r.URL.Path = "/"
		RootHandler(w, r)
	}
}

/* Serves an individual + editable todo item webpage */
func TodoIDHandler(w http.ResponseWriter, r *http.Request) {
	// Update the todo item
	if r.Method == http.MethodPost {
		// TODO server-side form data validation

		// Update title
		if r.FormValue("title") != "" {
			UpdateTitle(rdb, r)
		}

		// Update description
		if r.FormValue("description") != "" {
			UpdateDescription(rdb, r)
		}

		// Update completion status
		if r.FormValue("completed") != "" {
			UpdateCompletion(rdb, r)
		}
	}

	// Delete the todo item
	if r.Method == http.MethodDelete {
		DeleteTodo(rdb, r)
	}
}
