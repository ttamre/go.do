package main

import (
	"flag"

	"github.com/ttamre/go.do/api"
)

// Simple main function for easy customization
func main() {
	// Command line flags
	listenAddr := flag.Int("listenAddr", 5000, "Port to listen on")
	redisAddr := flag.Int("redisAddr", 5001, "Port for Redis database")
	flag.Parse()

	config := api.Config{
		Host:       "localhost", // server + database host
		ListenAddr: *listenAddr, // server port
		RedisAddr:  *redisAddr,  // database port
	}

	api.Serve(&config)
}
