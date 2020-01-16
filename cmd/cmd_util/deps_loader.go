package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/pkg/domain/jdomain"
)

func GetDepsFromJson(depPath string) []jdomain.JClassNode {
	var parsedDeps []jdomain.JClassNode
	file := ReadFile(depPath)
	_ = json.Unmarshal(file, &parsedDeps)

	return parsedDeps
}

