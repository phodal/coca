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
[d00f0124d] Phodal Huang 2019-12-19 update files
0       0       core/domain/bs/BadSmellApp.go

[1d00f0124b] Phodal Huang 2019-12-19 update files
1       1       cmd/bs.go
0       0       core/domain/bs/BadSmellApp.go

[d00f04111b] Phodal Huang 2019-12-18 refactor: move bs to adapter
1       1       cmd/bs.go
5       5       core/{domain => adapter}/bs/BadSmellApp.go

[d00f01214b] Phodal Huang 2019-12-19 update files
1       1       cmd/bs.go
0       0       core/adapter/bs/BadSmellApp.go
`)
			summary := gitt.GetTeamSummary(messages)
			Expect(summary[0].EntityName).To(Equal("core/adapter/bs/BadSmellApp.go"))
			Expect(summary[1].EntityName).To(Equal("cmd/bs.go"))
			Expect(len(summary)).To(Equal(2))
		})

		It("support for first path change", func() {
			messages := gitt.BuildMessageByInput(`
[333] Phodal Huang 2019-12-19 update files
0       0       src/domain/gitt/README.md

[d00f0124d] Phodal Huang 2019-12-19 update files
0       0       {src => core}/domain/gitt/README.md

`)
			summary := gitt.GetTeamSummary(messages)
			Expect(summary[0].EntityName).To(Equal("core/domain/gitt/README.md"))
			Expect(len(summary)).To(Equal(1))
		})

		It("should update child", func() {
			result := gitt.BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       adapter/JavaCallListener.go

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       adapter/{ => call}/JavaCallListener.go

`)
			summary := gitt.GetTeamSummary(result)
			Expect(summary[0].EntityName).To(Equal("adapter/call/JavaCallListener.go"))
			Expect(len(summary)).To(Equal(1))
		})

		It("should enable handle start move", func() {
			result := gitt.BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens => src/language/java/JavaLexer.tokens

`)
			summary := gitt.GetTeamSummary(result)
			Expect(len(summary)).To(Equal(1))
			Expect(summary[0].EntityName).To(Equal("src/language/java/JavaLexer.tokens"))
		})
	})
})