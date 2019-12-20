package test_test

import (
	"coca/core/domain/gitt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Git Parser", func() {
	Context("Test for Range", func() {
		It("support for long commits", func() {
			result := gitt.BuildMessageByInput(`
[828fe39523] Rossen Stoyanchev 2019-12-04 Consistently use releaseBody in DefaultWebClient
5       3       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/ClientResponse.java
1       1       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/DefaultWebClient.java
9       3       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/WebClient.java
6       11      core/docs/asciidoc/web/webflux-webclient.adoc

`)
			Expect(result[0].Rev).To(Equal("828fe39523"))
			Expect(result[0].Date).To(Equal("2019-12-04"))
			Expect(result[0].Author).To(Equal("Rossen Stoyanchev"))
			Expect(result[0].Message).To(Equal("Consistently use releaseBody in DefaultWebClient"))
			Expect(len(result[0].Changes)).To(Equal(4))
			Expect(result[0].Changes[0].File).To(Equal("spring-webflux/core/main/java/org/springframework/web/reactive/function/client/ClientResponse.java"))
		})
	})

	Context("Test for Move file", func() {
		It("should have a current file move update", func() {
			messages := gitt.BuildMessageByInput(`
[d00f0124b] Phodal Huang 2019-12-19 update files
1       1       cmd/bs.go
0       0       core/bs/models/BadSmellApp.go

[d00f0234b] Phodal Huang 2019-12-19 update files
12       0       core/bs/models/BadSmellApp.go

[d00f04b] Phodal Huang 2019-12-18 refactor: move bs to adapter
1       1       cmd/bs.go
5       5       core/{domain => adapter}/bs/BadSmellApp.go
`)
			summary := gitt.GetTeamSummary(messages)
			Expect(summary[2].EntityName).To(Equal("core/{domain => adapter}/bs/BadSmellApp.go"))
		})


//		It("should update child", func() {
//			result := gitt.BuildMessageByInput(`
//[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
//13      11      cmd/analysis.go
//10      8       cmd/api.go
//0       0       cmd/{call_graph.go => call.go}
//0       0       cmd/{git_cmd.go => git.go}
//`)
//			Expect(result).To(Equal("2019-12-04"))
//		})
	})
})