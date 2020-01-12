package cocago

import (
	. "github.com/onsi/gomega"
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
	assertPanic(t, errorReadFile)
}

func errorReadFile() {
	abs, _ := filepath.Abs("../../pkg/domain/code_file.go2")
	ProcessFile(abs)
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}