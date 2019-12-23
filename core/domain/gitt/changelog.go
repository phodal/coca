package gitt

import (
	"fmt"
	"regexp"
)

var (
	changeLogRegex = `^(\w*)(?:\((.*)\))?: (.*)$`
)

func ShowChangeLogSummary(commits []CommitMessage) {
	changeMap := BuildChangeMap(commits)
	fmt.Println(changeMap)
}

func BuildChangeMap(commits []CommitMessage) map[string]map[string]int {
	logRegexp := regexp.MustCompile(changeLogRegex)

	var czMap = make(map[string]map[string]int)
	for _, commit := range commits {
		message := commit.Message
		if logRegexp.MatchString(message) {
			matchs := logRegexp.FindStringSubmatch(message)
			if len(matchs) > 3 {
				keyword := matchs[1]
				//message := matchs[3]

				if _, ok := czMap[keyword];!ok {
					czMap[keyword] = make(map[string]int)
				}

				for _, change := range commit.Changes {
					czMap[keyword][change.File]++
				}
			}
		}
	}

	return czMap
}
