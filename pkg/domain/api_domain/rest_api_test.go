package api_domain

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_FilterRestApi(t *testing.T) {
	g := NewGomegaWithT(t)
	var apis []RestAPI
	blogApi := &RestAPI{"/blog", "", "", "", "", nil, "", "",}
	homeApi := &RestAPI{"/home", "", "", "", "", nil, "", "",}
	apis = append(apis, *blogApi)
	apis = append(apis, *homeApi)

	filteredApi := FilterApiByPrefix("/blog", apis)

	g.Expect(len(filteredApi)).To(Equal(1))
	g.Expect(filteredApi[0].Uri).To(Equal("/blog"))
}
