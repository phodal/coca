package visual

import "github.com/modernizing/coca/pkg/domain/core_domain"

type DData struct {
	Nodes []DNode `json:"nodes,omitempty"`
	Links []DLink `json:"links,omitempty"`
}

type DNode struct {
	ID    string `json:"id,omitempty"`
	Group int    `json:"group,omitempty"`
}

type DLink struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
	Value  int    `json:"value,omitempty"`
}

func FromDeps(deps []core_domain.CodeDataStruct) DData {
	var data DData
	nodeMap := make(map[string]DNode)
	sourceTargetMap := make(map[string]int)
	var links []DLink
	var groupIndex = 0

	for _, dep := range deps {
		groupIndex++
		nodeMap[dep.GetClassFullName()] = DNode{
			ID:    dep.GetClassFullName(),
			Group: groupIndex,
		}
		for _, pkg := range dep.FunctionCalls {
			if pkg.BuildClassFullName() != "" {
				nodeMap[pkg.BuildClassFullName()] = DNode{
					ID:    pkg.BuildClassFullName(),
					Group: groupIndex,
				}
				links = append(links, DLink{
					Source: dep.GetClassFullName(),
					Target: pkg.BuildClassFullName(),
					Value:  1,
				})
				sourceTargetMap[dep.GetClassFullName()+".coca."+pkg.BuildClassFullName()]++
			}
		}
	}

	for _, value := range nodeMap {
		data.Nodes = append(data.Nodes, value)
	}
	for _, link := range links {
		link.Value = sourceTargetMap[link.Source+".coca."+link.Target]
	}

	data.Links = links
	return data
}
