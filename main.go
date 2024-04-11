package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/ttamre/go.do/api"
)

const URL = "localhost"
const PORT = 5000
const DB_PORT = 5001

var rdb *redis.Client

/*
	/
		GET		serve webpage (list of todos)

	/todo
		POST	create todo							reached by "new todo" button from GET /

	/todo/{id}
		GET		serve webpage (edit todo)			reached by selecting a todo from list in GET /
		POST	update todo							reached by selecting a todo from list in GET /
		DELETE	delete todo							reached by selecting a todo from list in GET /

*/

func main() {
	// Create a file server to serve static files
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fileServer))

	// Define routes
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/todo", TodoHandler)
	http.HandleFunc("/todo/", TodoIDHandler)

	// Initialize database connection
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", URL, DB_PORT),
	})

	// Start server
	log.Printf("Listening on http://%s:%d...", URL, PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

type PageData struct {
	Data []api.Todo
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get Todos from database
		data := PageData{Data: []api.Todo{
			*api.NewTodo("Title 1", "Description 1"),
			*api.NewTodo("Title 2", "Description 2"),
		}}

		log.Println(data)

		// serve main webpage with template data
		log.Println("GET /")
		tmpl := template.Must(template.ParseFiles("./web/index.html"))
		tmpl.Execute(w, data)
	}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// create a new todo
		// 	reachable from: main webpage (new button)
		log.Printf("POST /todo %s %s", r.FormValue("title"), r.FormValue("description"))

		// Redirect to main webpage
		r.Method = "GET"
		RootHandler(w, r)
	}
}

func TodoIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get Todo ID from URL
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]

	if r.Method == "GET" {
		// Serve webpage for individual todo
		// 	reachable from: main webpage (by clicking on a todo)

		// Fetch todos from database

		// Serve webpage with todo template data
		log.Printf("GET /todo/%s", id)
		http.ServeFile(w, r, "./web/todo.html")
	}
	if r.Method == "POST" {
		// Update todo
		// 	reachable from: individual todo page (update button)
		if r.FormValue("title") != "" {
			log.Printf("POST /todo/%s %s", id, r.FormValue("title"))
		}
		if r.FormValue("description") != "" {
			log.Printf("POST /todo/%s %s", id, r.FormValue("description"))
		}
	}
	if r.Method == "DELETE" {
		// Delete todo
		// 	reachable from: individual todo page (delete button)
		log.Printf("DELETE /todo/%s", id)
	}
}
