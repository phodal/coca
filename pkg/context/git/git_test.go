package git

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_should_success_parse_log(t *testing.T) {
	//t.Parallel()
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[828fe39523] Rossen Stoyanchev 2019-12-04 Consistently use releaseBody in DefaultWebClient
5       3       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/ClientResponse.java
1       1       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/DefaultWebClient.java
9       3       spring-webflux/core/main/java/org/springframework/web/reactive/function/client/WebClient.java
6       11      core/docs/asciidoc/web/webflux-webclient.adoc

`)
	g.Expect(result[0].Rev).To(Equal("828fe39523"))
	g.Expect(result[0].Date).To(Equal("2019-12-04"))
	g.Expect(result[0].Author).To(Equal("Rossen Stoyanchev"))
	g.Expect(result[0].Message).To(Equal("Consistently use releaseBody in DefaultWebClient"))
	g.Expect(len(result[0].Changes)).To(Equal(4))
}

func Test_identify_file_move(t *testing.T) {
	//t.Parallel()
	g := NewGomegaWithT(t)

	messages := BuildMessageByInput(`
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
	summary := GetTeamSummary(messages)
	g.Expect(summary[0].EntityName).To(Equal("core/adapter/bs/BadSmellApp.go"))
	g.Expect(summary[1].EntityName).To(Equal("cmd/bs.go"))
	g.Expect(len(summary)).To(Equal(2))
}

func Test_identify_move_to_directory(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       adapter/JavaCallListener.go

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       adapter/{ => call}/JavaCallListener.go

`)
	summary := GetTeamSummary(result)
	g.Expect(summary[0].EntityName).To(Equal("adapter/call/JavaCallListener.go"))
	g.Expect(len(summary)).To(Equal(1))
}

func Test_handle_for_delete(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 refactor: extract vars
0       0       adapter/JavaCallListener.go
0       0       adapter/JavaCallListener2.go

[ef9165c] Phodal Huang 2019-12-18 refactor: extract vars
0       0       adapter/JavaCallListener2.go
 delete mode 100644 adapter/JavaCallListener2.go

`)
	summary := GetTeamSummary(result)
	g.Expect(len(summary)).To(Equal(1))
	g.Expect(summary[0].EntityName).To(Equal("adapter/JavaCallListener.go"))
}

func Test_identify_direct_move(t *testing.T) {
	//t.Parallel()
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens => src/language/java/JavaLexer.tokens

`)
	summary := GetTeamSummary(result)
	g.Expect(len(summary)).To(Equal(1))
	g.Expect(summary[0].EntityName).To(Equal("src/language/java/JavaLexer.tokens"))
}

func TestCalculateCodeAge(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens => src/language/java/JavaLexer.tokens

`)
	summary := CalculateCodeAge(result)
	g.Expect(len(summary)).To(Equal(1))
}

func TestGetTopAuthors(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens => src/language/java/JavaLexer.tokens

`)
	summary := GetTopAuthors(result)
	g.Expect(len(summary)).To(Equal(1))
	g.Expect(summary[0].Name).To(Equal("Phodal Huang"), )

}

func TestBasicSummary(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens

[ef9165c] Phodal Huang 2019-12-18 fefactor: extract vars
0       0       language/java/JavaParser.tokens => src/language/java/JavaLexer.tokens

`)
	summary := BasicSummary(result)
	g.Expect(summary.Authors).To(Equal(1))
	g.Expect(summary.Entities).To(Equal(2))
}

func Test_identify_first_file_move(t *testing.T) {
	g := NewGomegaWithT(t)

	messages := BuildMessageByInput(`
[333333333] Phodal Huang 2019-12-19 update files
0       0       src/domain/gitt/README.md

[d00f0124d] Phodal Huang 2019-12-19 update files
0       0       {src => core}/domain/gitt/README.md

`)
	summary := GetTeamSummary(messages)
	g.Expect(summary[0].EntityName).To(Equal("core/domain/gitt/README.md"))
	g.Expect(len(summary)).To(Equal(1))
}

func TestChangeModel(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[c24069b] Phodal HUANG 2019-10-25 fix: fix test
7       0       README.md
13      0       learn_go_suite_test.go
3       3       imp/imp_test.go => learn_go_test.go
 create mode 100644 learn_go_suite_test.go
 rename imp/imp_test.go => learn_go_test.go (70%)
 delete mode 100644 adapter/call/visitor/JavaCallVisitor.go

`)

	g.Expect(result[0].Changes[0].File).To(Equal("adapter/call/visitor/JavaCallVisitor.go"))
	g.Expect(result[0].Changes[0].Mode).To(Equal("delete"))
}


func Test_ShouldReturnRelatedFiles(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[ef9165d] Phodal Huang 2019-12-18 refactor: extract vars
0       0       adapter/JavaCallListener.java
0       0       adapter/JavaCallListener2.java
0       0       adapter/JavaCallListener3.java
0       0       adapter/JavaCallListener5.java

[ef9165c] Phodal Huang 2019-12-18 refactor: extract vars
0       0       adapter/JavaCallListener.java
0       0       adapter/JavaCallListener2.java
0       0       adapter/JavaCallListener3.java
0       0       adapter/JavaCallListener4.java
`)
	relatedConfig := []byte(`{
	"minSupport": 0.1,
	"minConfidence": 0.9,
	"minLift": 0,
	"maxLength": 3
}
`)

	relatedFiles := GetRelatedFiles(result, relatedConfig)

	g.Expect(len(relatedFiles)).To(Equal(7))
}

