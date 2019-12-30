package count

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/infrastructure"
	"path/filepath"
	"testing"
)

func TestBuildCallMap(t *testing.T) {
	g := NewGomegaWithT(t)
	var parsedDeps []models.JClassNode
	codePath := "../../../_fixtures/count/call.json"
	codePath = filepath.FromSlash(codePath)
	file := infrastructure.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	callMap := BuildCallMap(parsedDeps)

	g.Expect(len(callMap)).To(Equal(4))
	g.Expect(callMap["com.phodal.pholedge.book.BookService.createBook"]).To(Equal(1))
	g.Expect(callMap["com.phodal.pholedge.book.BookService.getBookById"]).To(Equal(1))
	g.Expect(callMap["com.phodal.pholedge.book.BookService.getBooksLists"]).To(Equal(1))
}
