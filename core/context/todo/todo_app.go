package todo

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/cocafile"
	"github.com/phodal/coca/core/adapter/shell"
	"github.com/phodal/coca/core/context/git"
	"github.com/phodal/coca/core/context/todo/astitodo"
	. "github.com/phodal/coca/languages/java"
	"path/filepath"
	"strconv"
)

type TodoApp struct {
}

func NewTodoApp() TodoApp {
	return *&TodoApp{

	}
}

type TodoDetail struct {
	Date     string
	FileName string
	Author   string
	Line     string
	Assignee string
	Message  []string
}

func (a TodoApp) AnalysisPath(path string) []*astitodo.TODO {
	todos := buildComments(path)
	return todos
}

func (a TodoApp) BuildWithGitHistory(todos []*astitodo.TODO) []TodoDetail {
	var todoList []TodoDetail = nil

	for _, todo := range todos {
		lineOutput := shell.RunGitGetLog(todo.Line, todo.Filename)

		todoDetail := &TodoDetail{
			Date:     "",
			FileName: todo.Filename,
			Author:   "",
			Line:     strconv.Itoa(todo.Line),
			Assignee: todo.Assignee,
			Message:  todo.Message,
		}
		commitMessages := git.BuildMessageByInput(lineOutput)

		if len(commitMessages) > 0 {
			commit := commitMessages[0]
			todoDetail.Date = commit.Date
			todoDetail.Author = commit.Author
		}
		todoList = append(todoList, *todoDetail)
	}

	return todoList
}

func buildComments(path string) []*astitodo.TODO {
	var todos []*astitodo.TODO
	files := cocafile.GetJavaFiles(path)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Refactoring parse java call: " + displayName)

		is, _ := antlr.NewFileStream(file)
		lexer := NewJavaLexer(is)

		for _, token := range lexer.GetAllTokens() {
			COMMENT_TOKEN_INDEX := 109
			COMMENT_LINE_TOKNE_INDEX := 110
			// based on `JavaLexer.tokens` file
			if token.GetTokenType() == COMMENT_TOKEN_INDEX || token.GetTokenType() == COMMENT_LINE_TOKNE_INDEX {
				todo := astitodo.ParseComment(token, file)
				if todo != nil {
					todos = append(todos, todo)
				}
			}
		}
	}

	return todos
}
