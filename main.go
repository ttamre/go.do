package main

import (
	"fmt"

	"github.com/ttamre/go.do/api"
)

func main() {
	todoList := api.NewTodoList()

	for i := 0; i < 5; i++ {
		todo := api.NewTodo(fmt.Sprintf("title%d", i+1), fmt.Sprintf("desc%d", i+1), todoList.ID)
		api.Add(todoList, todo)
	}

	api.Delete(todoList, todoList.List[1])
	api.UpdateTitle(todoList.List[0], "UpdatedTitle")
	api.UpdateDescription(todoList.List[1], "UpdatedDescription")
	api.Complete(todoList.List[2])

	api.PrintList(todoList)
}
