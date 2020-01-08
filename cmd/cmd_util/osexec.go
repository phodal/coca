package cmd_util

import (
	"fmt"
	"github.com/phodal/coca/cmd/config"
	"os/exec"
)

func ConvertToSvg(name string) {
	reporter_path := config.CocaConfig.ReporterPath
	cmd := exec.Command("dot", "-Tsvg", reporter_path+"/"+name+".dot", "-o", reporter_path+"/"+name+".svg")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd.Run() failed with:", err)
	}
}
