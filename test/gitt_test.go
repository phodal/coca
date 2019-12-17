package test_test


import (
	"coca/core/domain/gitt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Git Parser", func() {
	Context("Test for Range", func() {
		It("should be a novel", func() {
			result := gitt.ParseLog(`
[828fe39523] Rossen Stoyanchev 2019-12-04 Consistently use releaseBody in DefaultWebClient
5       3       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/ClientResponse.java
1       1       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/DefaultWebClient.java
9       3       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/WebClient.java
6       11      core/docs/asciidoc/web/webflux-webclient.adoc
`)
			Expect(result.Date).To(Equal("2019-12-04"))
		})
	})
})