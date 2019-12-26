package count

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"testing"
)

func TestBuildCallMap(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	file := support.ReadFile("../../../_fixtures/count/call.json")
	_ = json.Unmarshal(file, &parsedDeps)

	callMap := BuildCallMap(parsedDeps)

	g.Expect(len(callMap)).To(Equal(4))
	g.Expect(callMap["com.phodal.pholedge.book.BookService.createBook"]).To(Equal(1))
	g.Expect(callMap["com.phodal.pholedge.book.BookService.getBookById"]).To(Equal(1))
	g.Expect(callMap["com.phodal.pholedge.book.BookService.getBooksLists"]).To(Equal(1))
}
