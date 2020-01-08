package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/core/domain"
)

func GetDepsFromJson(depPath string) []domain.JClassNode {
	var parsedDeps []domain.JClassNode
	file := ReadFile(depPath)
	_ = json.Unmarshal(file, &parsedDeps)

	return parsedDeps
}

