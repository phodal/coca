package cocago

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func Test_ParserLog(t *testing.T) {
	g := NewGomegaWithT(t)

	abs, _ := filepath.Abs("../../pkg/domain/code_file.go")
	results := ProcessFile(abs)
	g.Expect(len(results.Members)).To(Equal(2))
}
