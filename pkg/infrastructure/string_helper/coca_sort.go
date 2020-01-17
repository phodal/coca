package string_helper

import (
	"github.com/yourbasic/radix"
)

// from: https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
func SortWord(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}

	radix.SortSlice(pl, func(i int) string { return pl[i].Key })
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
