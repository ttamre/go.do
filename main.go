package main

import (
	"fmt"

	"github.com/ttamre/go.do/api"
)

func main() {
	user := api.NewUser("test", "password")
	todoList := api.CreateTodoList(user, "list title")

	for i := 0; i < 5; i++ {
		api.CreateTodo(user, todoList, fmt.Sprintf("title%d", i+1), fmt.Sprintf("desc%d", i+1))
	}

	api.Delete(todoList, todoList.List[1])
	api.UpdateTitle(todoList.List[0], "UpdatedTitle")
	api.UpdateDescription(todoList.List[1], "UpdatedDescription")
	api.Complete(todoList.List[2])

	// api.PrintList(todoList)
}
