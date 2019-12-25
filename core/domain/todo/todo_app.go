package todo

import (
	"fmt"
	"github.com/phodal/coca/core/domain/todo/astitodo"
	"log"
)

type TodoApp struct {
}

func NewTodoApp() TodoApp {
	return *&TodoApp{

	}
}

func (a TodoApp) AnalysisPath(path string) {
	todos, err := astitodo.Extract(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, todo := range todos {
		fmt.Println(todo.Message, todo.Line, todo.Filename)
	}
}
