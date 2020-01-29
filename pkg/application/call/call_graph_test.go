package call_test

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cocatest/testhelper"
	"github.com/phodal/coca/pkg/application/api"
	"github.com/phodal/coca/pkg/application/call"
	api_domain2 "github.com/phodal/coca/pkg/domain/api_domain"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func Test_ShouldBuildSuccessDataFromJson(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := call.NewCallGraph()
	codePath := "../../../_fixtures/call/call_api_test.json"
	codePath = filepath.FromSlash(codePath)

	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	dotContent := analyser.Analysis("com.phodal.pholedge.book.BookController.createBook", parsedDeps)

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

func Test_ShouldBuildSuccessDataFromSourceData(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/java/examples/api"
	callNodes, identifiersMap, identifiers := testhelper.BuildAnalysisDeps(codePath)

	diMap := core_domain.BuildDIMap(identifiers, identifiersMap)
	app := new(api.JavaApiApp)
	restApis := app.AnalysisPath(codePath, callNodes, identifiersMap, diMap)

	analyser := call.NewCallGraph()
	dotContent, apis := analyser.AnalysisByFiles(restApis, callNodes, diMap)

	api_domain2.SortAPIs(apis)
	g.Expect(len(apis)).To(Equal(4))
	g.Expect(apis[0].Size).To(Equal(4))
	g.Expect(apis[1].Size).To(Equal(7))
	g.Expect(apis[2].Size).To(Equal(12))
	g.Expect(apis[3].Size).To(Equal(16))

	g.Expect(dotContent).To(Equal(`digraph G { 

"POST /books" -> "com.phodal.pholedge.book.BookController.createBook";
"com.phodal.pholedge.book.BookFactory.create" -> "com.phodal.pholedge.core.IdGenerator.generate";
"com.phodal.pholedge.book.model.Book.create" -> "com.phodal.pholedge.book.model.Book.builder";
"com.phodal.pholedge.book.model.Book.create" -> "com.phodal.pholedge.book.model.Book.id";
"com.phodal.pholedge.book.model.Book.create" -> "com.phodal.pholedge.book.model.Book.isbn";
"com.phodal.pholedge.book.model.Book.create" -> "com.phodal.pholedge.book.model.Book.name";
"com.phodal.pholedge.book.model.Book.create" -> "com.phodal.pholedge.book.model.Book.createdAt";
"com.phodal.pholedge.book.model.Book.create" -> "com.phodal.pholedge.book.model.Book.build";
"com.phodal.pholedge.book.BookFactory.create" -> "com.phodal.pholedge.book.model.Book.create";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookFactory.create";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getIsbn";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getName";
"com.phodal.pholedge.book.BookRepository.save" -> "com.phodal.pholedge.book.model.this.bookMapper.doSave";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookRepository.save";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.Book.getId";
"com.phodal.pholedge.book.BookController.createBook" -> "com.phodal.pholedge.book.BookService.createBook";

"PUT /books/{id}" -> "com.phodal.pholedge.book.BookController.updateBook";
"com.phodal.pholedge.book.BookRepository.byId" -> "com.phodal.pholedge.book.model.this.bookMapper.byId";
"com.phodal.pholedge.book.BookRepository.byId" -> "com.phodal.pholedge.core.exception.NotFoundException";
"com.phodal.pholedge.book.BookService.updateBook" -> "com.phodal.pholedge.book.BookRepository.byId";
"com.phodal.pholedge.book.BookService.updateBook" -> "com.phodal.pholedge.book.model.Book.save";
"com.phodal.pholedge.book.BookService.updateBook" -> "com.phodal.pholedge.book.model.command.UpdateBookCommand.getIsbn";
"com.phodal.pholedge.book.BookService.updateBook" -> "com.phodal.pholedge.book.model.command.UpdateBookCommand.getName";
"com.phodal.pholedge.book.BookRepository.save" -> "com.phodal.pholedge.book.model.this.bookMapper.doSave";
"com.phodal.pholedge.book.BookService.updateBook" -> "com.phodal.pholedge.book.BookRepository.save";
"com.phodal.pholedge.book.model.Book.toRepresentation" -> "com.phodal.pholedge.book.model.BookRepresentaion";
"com.phodal.pholedge.book.BookService.updateBook" -> "com.phodal.pholedge.book.model.Book.toRepresentation";
"com.phodal.pholedge.book.BookController.updateBook" -> "com.phodal.pholedge.book.BookService.updateBook";

"GET /books/" -> "com.phodal.pholedge.book.BookController.getBookList";
"com.phodal.pholedge.book.BookRepository.list" -> "com.phodal.pholedge.book.model.this.bookMapper.list";
"com.phodal.pholedge.book.BookService.getBooksLists" -> "com.phodal.pholedge.book.BookRepository.list";
"com.phodal.pholedge.book.BookController.getBookList" -> "com.phodal.pholedge.book.BookService.getBooksLists";

"GET /books/{id}" -> "com.phodal.pholedge.book.BookController.getBookById";
"com.phodal.pholedge.book.BookRepository.byId" -> "com.phodal.pholedge.book.model.this.bookMapper.byId";
"com.phodal.pholedge.book.BookRepository.byId" -> "com.phodal.pholedge.core.exception.NotFoundException";
"com.phodal.pholedge.book.BookService.getBookById" -> "com.phodal.pholedge.book.BookRepository.byId";
"com.phodal.pholedge.book.model.Book.toRepresentation" -> "com.phodal.pholedge.book.model.BookRepresentaion";
"com.phodal.pholedge.book.BookService.getBookById" -> "com.phodal.pholedge.book.model.Book.toRepresentation";
"com.phodal.pholedge.book.BookController.getBookById" -> "com.phodal.pholedge.book.BookService.getBookById";
}
`))

}
