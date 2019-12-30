package git

import (
	"fmt"
	"github.com/phodal/coca/core/infrastructure/str_helper"
	"regexp"
)

var (
	changeLogRegex = `^(\w*)(?:\((.*)\))?: (.*)$`
)

// high fix
// high features
//
func ShowChangeLogSummary(commits []CommitMessage) {
	changeMap := BuildChangeMap(commits)
	for key, value := range changeMap {
		sortValue := str_helper.RankByWordCount(value)
		maxSize := len(sortValue)
		if maxSize > 10 {
			maxSize = 10
		}

		fmt.Println(key  + ":")
		fmt.Println("---------------------")
		for _, val := range sortValue[:maxSize] {
			fmt.Println(val.Key, val.Value)
		}
		fmt.Println("=====================")
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
