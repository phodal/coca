package todo

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/phodal/coca/languages/comment"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/adapter/shell"
	"github.com/phodal/coca/pkg/application/git"
	"github.com/phodal/coca/pkg/application/todo/astitodo"
	"path/filepath"
	"strconv"
	"strings"
)

type TodoApp struct {
}

func NewTodoApp() TodoApp {
	return TodoApp{

	}
}

type TodoDetail struct {
	Date     string
	FileName string
	Author   string
	Line     string
	Assignee string
	Message  string
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
	var CodeFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".py") || strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".ts") || strings.HasSuffix(path, ".java")
	}

	files := cocafile.GetFilesWithFilter(path, CodeFileFilter)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("parse java call: " + displayName)

		is, _ := antlr.NewFileStream(file)
		lexer := NewCommentLexer(is)

		for _, token := range lexer.GetAllTokens() {
			COMMENT := 1
			LINE_COMMENT := 2
			PYTHON_COMMENT := 3

			// based on `JavaLexer.tokens` file
			if token.GetTokenType() == COMMENT ||
				token.GetTokenType() == LINE_COMMENT ||
				token.GetTokenType() == PYTHON_COMMENT {

				todo := astitodo.ParseComment(token, file)
				if todo != nil {
					todos = append(todos, todo)
				}
			}
		}
	}

	return todos
}
