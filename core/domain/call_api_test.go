package domain

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSortApi(t *testing.T) {
	g := NewGomegaWithT(t)
	var apis []CallApi
	api3 := &CallApi{ "get","/blog","home", 3}
	api2 := &CallApi{ "get","/blog","home", 5}
	api5 := &CallApi{ "get","/blog","home", 2}
	apis = append(apis, *api3)
	apis = append(apis, *api2)
	apis = append(apis, *api5)

	g.Expect(apis[0].Size).To(Equal(3))

	SortApi(apis)

	g.Expect(apis[0].Size).To(Equal(2))
}