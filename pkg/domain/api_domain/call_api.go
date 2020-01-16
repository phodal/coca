package api_domain

import "sort"

type CallAPI struct {
	HTTPMethod string
	URI        string
	Caller     string
	Size       int
}

func SortAPIs(callAPIs []CallAPI) {
	sort.Slice(callAPIs, func(i, j int) bool {
		return callAPIs[i].Size < callAPIs[j].Size
	})
}
