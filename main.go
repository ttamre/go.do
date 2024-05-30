package main

import (
	"github.com/ttamre/go.do/api"
)

// Simple main function for easy customization
func main() {
	config := api.Config{
		Host:       "localhost", // server + database host
		ListenAddr: 5000,        // server port
		RedisAddr:  5001,        // database port
	}

	api.Serve(&config)
}
