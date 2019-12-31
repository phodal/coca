package cmd_util

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/core/domain"
	"log"
	"os/exec"
)

func GetDepsFromJson(depPath string) []domain.JClassNode {
	var parsedDeps []domain.JClassNode
	file := ReadFile(depPath)
	if file == nil {
		log.Fatal("lost file:" + depPath)
	}
	_ = json.Unmarshal(file, &parsedDeps)

	return parsedDeps
}

func ConvertToSvg(name string) {
	cmd := exec.Command("dot", []string{"-Tsvg", config.CocaConfig.ReporterPath + "/" + name + ".dot", "-o", config.CocaConfig.ReporterPath + "/" + name + ".svg"}...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd.Run() failed with:", err)
	}
}

