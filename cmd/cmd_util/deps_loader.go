package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

func GetDepsFromJson(depPath string) []core_domain.CodeDataStruct {
	var parsedDeps []core_domain.CodeDataStruct
	file := ReadFile(depPath)
	_ = json.Unmarshal(file, &parsedDeps)

	return parsedDeps
}

