package git

import (
	"fmt"
	"github.com/phodal/coca/pkg/infrastructure/string_helper"
	"io"
	"regexp"
)

var (
	changeLogRegex = `^(\w*)(?:\((.*)\))?: (.*)$`
)

// high fix
// high features
func ShowChangeLogSummary(commits []CommitMessage, output io.Writer) {
	changeMap := BuildChangeMap(commits)
	for key, value := range changeMap {
		sortValue := string_helper.SortWord(value)
		maxSize := len(sortValue)
		if maxSize > 10 {
			maxSize = 10
		}

		fmt.Fprintf(output, "%s :\n", key)
		fmt.Fprintln(output, "---------------------")
		for _, val := range sortValue[:maxSize] {
			fmt.Fprintf(output, "%s, %d\n", val.Key, val.Value)
		}
		fmt.Fprintln(output, "=====================")
	}
}

func BuildChangeMap(commits []CommitMessage) map[string]map[string]int {
	logRegexp := regexp.MustCompile(changeLogRegex)

	var czMap = make(map[string]map[string]int)
	for _, commit := range commits {
		message := commit.Message
		if logRegexp.MatchString(message) {
			matches := logRegexp.FindStringSubmatch(message)
			if len(matches) > 3 {
				keyword := matches[1]
				//message := matches[3]

				if _, ok := czMap[keyword];!ok {
					czMap[keyword] = make(map[string]int)
				}

				for _, change := range commit.Changes {
					file := change.File
					file, oldFile, newFile := UpdateMessageForChange(file)
					if file != oldFile {
						file = newFile
					}

					czMap[keyword][file]++
				}
			}
		}
	}

	return czMap
}
