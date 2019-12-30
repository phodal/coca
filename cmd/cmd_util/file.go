package cmd_util

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/infrastructure"
	"log"
	"os/exec"
)

func GetDepsFromJson(depPath string) []models.JClassNode {
	var parsedDeps []models.JClassNode
	file := infrastructure.ReadFile(depPath)
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

