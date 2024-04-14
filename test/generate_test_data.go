/*
Helper script to populate redis database with test data

Usage: go run generate_test_data.go -u [URL] -p [PORT] -f [JSON_FILE]

Version:	1.0
Author: 	Tem Tamre
Contact: 	temtamre@gmail.com
*/

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/ttamre/go.do/api"
)

func main() {
	// Process command line arguments
	var (
		url, file string
		db_port   int
	)
	flag.StringVar(&url, "u", "127.0.0.1", "database url (default: 127.0.0.1)")
	flag.IntVar(&db_port, "p", 5001, "database port (default: 5001)")
	flag.StringVar(&file, "f", "test_data.json", "test data (default: test_data.json)")

	// Initialize context and database connection
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", url, db_port),
	})

	// Generate test data and add it to database
	jsonData, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Failed to read data file: %v\nGenerating data instead...", err)

		// If we can't read the file, generate test data instead instead of giving up
		for i := 1; i <= 10; i++ {
			todo := api.NewTodo(fmt.Sprintf("Todo %d", i), fmt.Sprintf("Description for Todo %d", i))
			err := rdb.HMSet(ctx, todo.ID.String(), map[string]interface{}{
				"title":       todo.Title,
				"description": todo.Description,
				"created_on":  todo.CreatedOn,
				"completed":   strconv.FormatBool(todo.Completed),
			}).Err()

			if err != nil {
				log.Fatalf("Failed to create test data: %v", err)
			}
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
