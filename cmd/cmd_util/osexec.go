package cmd_util

import (
	"fmt"
	"github.com/phodal/coca/cmd/config"
	"os/exec"
)

func ConvertToSvg(name string) {
	cmd := exec.Command("dot", []string{"-Tsvg", config.CocaConfig.ReporterPath + "/" + name + ".dot", "-o", config.CocaConfig.ReporterPath + "/" + name + ".svg"}...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd.Run() failed with:", err)
	}
}

