package cocago

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"path/filepath"
	"testing"
)


func Test_ParserLog(t *testing.T) {
	g := NewGomegaWithT(t)

	abs, _ := filepath.Abs("../../pkg/domain/code_file.go")
	ProcessFile(abs)
	g.Expect(1).To(Equal(1))
}

func TestPanic(t *testing.T) {
	cocatest.AssertPanic(t, errorReadFile)
}

func errorReadFile() {
	abs, _ := filepath.Abs("../../pkg/domain/code_file.go2")
	ProcessFile(abs)
}
