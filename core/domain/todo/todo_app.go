package todo

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain/gitt"
	"github.com/phodal/coca/core/domain/todo/astitodo"
	"github.com/phodal/coca/core/support"
	. "github.com/phodal/coca/languages/java"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
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
	todos := buildComment(path)
	return todos
}

func (a TodoApp) BuildWithGitHistory(todos []*astitodo.TODO) []TodoDetail {
	var todoList []TodoDetail = nil

	for _, todo := range todos {
		lineOutput := runGitGetLog(todo.Line, todo.Filename)

		todoDetail := &TodoDetail{
			Date:     "",
			FileName: todo.Filename,
			Author:   "",
			Line:     strconv.Itoa(todo.Line),
			Assignee: todo.Assignee,
			Message:  todo.Message,
		}
		commitMessages := gitt.BuildMessageByInput(lineOutput)

		if len(commitMessages) > 0 {
			commit := commitMessages[0]
			todoDetail.Date = commit.Date
			todoDetail.Author = commit.Author
		}
		todoList = append(todoList, *todoDetail)
	}

	return todoList
}

func buildComment(path string) []*astitodo.TODO {
	var todos []*astitodo.TODO
	files := support.GetJavaFiles(path)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		//abs, _ := filepath.Abs(file)
		fmt.Println("Start parse java call: " + displayName)

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

func runGitGetLog(line int, fileName string) string {
	// git log -1 -L2:README.md --pretty="format:[%h] %aN %ad %s" --date=short  --numstat
	historyArgs := []string{"log", "-1", "-L" + strconv.Itoa(line) + ":" + fileName, "--pretty=\"format:[%h] %aN %ad %s\"", "--date=short", "--numstat", "--summary"}
	cmd := exec.Command("git", historyArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	split := strings.Split(string(out), "\n")
	output := split[0] + "\n "
	return output
}
