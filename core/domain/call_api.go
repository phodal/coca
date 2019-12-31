package domain

import "sort"

type CallApi struct {
	HttpMethod string
	Uri        string
	Caller     string
	Size       int
}

func SortApi(apis []CallApi) {
	sort.Slice(apis, func(i, j int) bool {
		return apis[i].Size < apis[j].Size
	})
}
