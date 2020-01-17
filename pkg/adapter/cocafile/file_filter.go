package cocafile

import "strings"

var JavaTestFileFilter = func(path string) bool {
	return strings.Contains(path, "Test.java") || strings.Contains(path, "Tests.java")
}

var JavaCodeFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".java") && !JavaTestFileFilter(path)
}

var JavaFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".java")
}

var TypeScriptFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".ts")
}

var PythonFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".py")
}

var GoFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".go")
}

var PomXmlFilter = func(path string) bool {
	return strings.HasSuffix(path, "pom.xml")
}

var BuildGradleFilter = func(path string) bool {
	return strings.HasSuffix(path, "build.gradle")
}
