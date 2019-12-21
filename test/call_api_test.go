package test_test

import (
	"coca/core/domain/call_graph"
	"coca/core/models"
	"coca/core/support"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Git Parser", func() {
	Context("Test for Range", func() {
		It("should be a novel", func() {
			var parsedDeps []models.JClassNode
			analyser := call_graph.NewCallGraph()

			file := support.ReadFile("_fixtures/call_api_test.json")
			if file == nil {
				log.Fatal("lost file:")
			}

			_ = json.Unmarshal(file, &parsedDeps)

			dotContent := analyser.Analysis("com.phodal.pholedge.book.BookController.createBook", *&parsedDeps)

			Expect(dotContent).To(Equal(`digraph G {
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookFactory.create";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getIsbn";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getName";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookRepository.save";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.Book.getId";
"com.phodal.pholedge.book.BookController.createBook" -> "com.phodal.pholedge.book.BookService.createBook";
}
`))
		})
	})
})