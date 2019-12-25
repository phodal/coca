package todo

import (
	"fmt"
	"github.com/phodal/coca/core/domain/gitt"
	"github.com/phodal/coca/core/domain/todo/astitodo"
	"log"
	"os/exec"
	"strconv"
	"strings"
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
		lineOutput := runGitGetLog(todo.Line, todo.Filename)

		commitMessages := gitt.BuildMessageByInput(lineOutput)
		if len(commitMessages) > 0 {
			commit := commitMessages[0]
			fmt.Println(commit.Date, todo.Filename, commit.Author, todo.Line)
		}
	}
}

func runGitGetLog(line int, fileName string) string {
	// git log -1 -L2:README.md --pretty="format:[%h] %aN %ad %s" --date=short  --numstat
	historyArgs := []string{"log", "-1", "-L" + strconv.Itoa(line) + ":" + fileName, "--pretty=format:[%h] %aN %ad %s", "--date=short", "--numstat", "--summary"}
	cmd := exec.Command("git", historyArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	split := strings.Split(string(out), "\n")
	output := split[0] + "\n "
	return output
}
