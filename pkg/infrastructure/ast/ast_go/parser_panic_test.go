package ast_go

import (
	"github.com/phodal/coca/cocatest"
	"path/filepath"
	"testing"
)

func Test_ShouldPainWhenReadFileError(t *testing.T) {
	t.Parallel()
	cocatest.AssertPanic(t, errorReadFile)
}

func errorReadFile() {
	abs, _ := filepath.Abs("../../pkg/domain/code_file.go2")
	testParser.ProcessFile(abs)
}
