package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"log"
	"os/exec"
)

func GetDepsFromJson(depPath string) []models.JClassNode {
	var parsedDeps []models.JClassNode
	file := support.ReadFile(depPath)
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
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

