package main

import (
	"github.com/ttamre/go.do/api"
)

// Simple main function for easy customization
func main() {
	config := api.Config{
		URL:     "localhost", // host url for server and database
		PORT:    5000,        // server port
		DB_PORT: 5001,        // redis port
	}

	api.Serve(&config)
}
