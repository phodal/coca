package cocafile

import (
	"fmt"
	ignore "github.com/sabhiram/go-gitignore"
	"os"
	"path/filepath"
	"strings"
)

func GetJavaFiles(codeDir string) []string {
	return GetFilesWithFilter(codeDir, JavaCodeFileFilter)
}

func GetFilesWithFilter(codeDir string, filter func(path string) bool) []string {
	files := make([]string, 0)
	gitIgnore, err := ignore.CompileIgnoreFile(filepath.FromSlash(codeDir + "/.gitignore"))
	if err != nil {
		fmt.Println(err)
	}

	fi, err := os.Stat(codeDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fi.Mode().IsRegular() {
		files = append(files, codeDir)
		return files
	}

	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if gitIgnore != nil {
			if gitIgnore.MatchesPath(path) {
				return nil
			}
		}

		if strings.Contains(path, "testData") {
			return nil
		}

		if filter(path) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func GetJavaTestFiles(codeDir string) []string {
	return GetFilesWithFilter(codeDir, JavaTestFileFilter)
}
