package shell

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func RunGitGetLog(line int, fileName string) string {
	// git log -1 -L2:README.md --pretty="format:[%h] %aN %ad %s" --date=short
	historyArgs := []string{"log", "-1", "-L" + strconv.Itoa(line) + ":" + fileName, "--pretty=\"format:[%h] %aN %ad %s\"", "--date=short"}
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
