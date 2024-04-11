package api

import (
	"github.com/redis/go-redis/v9"
)

type Server interface {
	NewTodo() *Todo
	DeleteTodo(*Todo) error
	UpdateTitle(*Todo, string) error
	UpdateDescription(*Todo, string) error
	UpdateCompleted(*Todo) error
}

type TodoServer struct {
	Url  string
	Port string
	Db   *redis.Client
}
