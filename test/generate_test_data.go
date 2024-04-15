/*
Helper script to populate redis database with test data

Usage: go run generate_test_data.go -u [URL] -p [PORT] -f [JSON_FILE]

Version:	1.0
Author: 	Tem Tamre
Contact: 	temtamre@gmail.com
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/ttamre/go.do/api"
)

func main() {
	// Process command line arguments
	var (
		URL, FILE string
		DB_PORT   int
	)
	flag.StringVar(&URL, "u", "127.0.0.1", "database url (default: 127.0.0.1)")
	flag.IntVar(&DB_PORT, "p", 5001, "database port (default: 5001)")
	flag.StringVar(&FILE, "f", "test_data.json", "test data (default: test_data.json)")

	// Initialize context and database connection
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", URL, DB_PORT),
	})

	// Generate test data and add it to database
	jsonData, err := os.ReadFile(FILE)
	if err != nil {
		log.Printf("Failed to read data file: %v\nGenerating data instead...", err)

		// If we can't read the file, generate test data instead instead of giving up
		for i := 1; i <= 10; i++ {

			// Make a mock HTTP request to hold our form data
			formData := url.Values{"title": {"Test Title " + fmt.Sprint(i)}, "description": {"Test Description " + fmt.Sprint(i)}}
			req, err := http.NewRequest("POST", "http://localhost:5000/", strings.NewReader(formData.Encode()))
			if err != nil {
				log.Fatalf("Failed to create mock HTTP request: %v", err)
			}

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			api.AddTodo(rdb, req)
		}
		log.Println("Test data generated successfully")
		return
	}

	// Unmarshal JSON data
	var todo api.Todo
	err = json.Unmarshal(jsonData, &todo)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	log.Println(todo)

}
