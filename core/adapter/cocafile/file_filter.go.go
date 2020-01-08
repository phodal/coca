package cocafile

import "strings"

var JavaTestFileFilter = func(path string) bool {
	return strings.Contains(path, "Test.java") || strings.Contains(path, "Tests.java")
}

var JavaCodeFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".java") && !JavaTestFileFilter(path)
}
