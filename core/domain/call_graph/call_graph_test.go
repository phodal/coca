package call_graph_test

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/domain/call_graph"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"path/filepath"
	"testing"
)


func Test_should_generate_correct_files(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := call_graph.NewCallGraph()
	codePath := "../../../_fixtures/call_api_test.json"
	codePath = filepath.FromSlash(codePath)

	file := support.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	dotContent := analyser.Analysis("com.phodal.pholedge.book.BookController.createBook", *&parsedDeps)

	g.Expect(dotContent).To(Equal(`digraph G {
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookFactory.create";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getIsbn";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getName";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookRepository.save";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.Book.getId";
"com.phodal.pholedge.book.BookController.createBook" -> "com.phodal.pholedge.book.BookService.createBook";
}
`))

}
