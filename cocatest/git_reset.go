package cocatest

import (
	"fmt"
	"log"
	"os/exec"
)

func ResetGitDir(codePath string) {
	cmd := exec.Command("git", "checkout", "--ignore-skip-worktree-bits", "--", codePath)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}
