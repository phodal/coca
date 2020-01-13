package testhelper

import (
	"fmt"
	"os/exec"
)

func ResetGitDir(codePath string) {
	cmd := exec.Command("git", "checkout", "--ignore-skip-worktree-bits", "--", codePath)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println("cmd.Run() failed with: ", err)
	}
}
