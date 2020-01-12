package bs_domain

import "sort"

type BadSmellModel struct {
	File        string `json:"EntityName,omitempty"`
	Line        string `json:"Line,omitempty"`
	Bs          string `json:"BS,omitempty"`
	Description string `json:"Description,omitempty"`
	Size        int    `size:"Description,omitempty"`
}

func SortSmellByType(models []BadSmellModel, filterFunc func(key string) bool) map[string][]BadSmellModel {
	sortSmells := make(map[string][]BadSmellModel)
	for _, model := range models {
		sortSmells[model.Bs] = append(sortSmells[model.Bs], model)
	}

	for key, smells := range sortSmells {
		if filterFunc(key) {
			sort.Slice(smells, func(i, j int) bool {
				return smells[i].Size > (smells[j].Size)
			})

			sortSmells[key] = smells
		}
	}

	return sortSmells
}

func FilterBadSmellList(models []BadSmellModel, ignoreRules map[string]bool) []BadSmellModel {
	var results []BadSmellModel
	for _, model := range models {
		if !ignoreRules[model.Bs] {
			results = append(results, model)
		}
	}
	return results
}
