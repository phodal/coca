package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/core/domain"
	"log"
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

