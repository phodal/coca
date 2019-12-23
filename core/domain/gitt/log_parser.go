package gitt

import (
	"regexp"
	"strings"
)

var (
	rev               = `\[([\d|a-f]{5,12})\]`
	author            = `(.*?)\s\d{4}-\d{2}-\d{2}`
	date              = `\d{4}-\d{2}-\d{2}`
	changes           = `([\d-]+)[\t\s]+([\d-]+)[\t\s]+(.*)`
	complexMoveRegStr = `(.*)\{(.*)\s=>\s(.*)\}(.*)`
	basicMoveRegStr   = `(.*)\s=>\s(.*)`

	revReg         = regexp.MustCompile(rev)
	authorReg      = regexp.MustCompile(author)
	dateReg        = regexp.MustCompile(date)
	changesReg     = regexp.MustCompile(changes)
	complexMoveReg = regexp.MustCompile(complexMoveRegStr)
	basicMvReg = regexp.MustCompile(basicMoveRegStr)
)

func UpdateMessageForChange(changedFile string) (string, string, string) {
	oldFileName := changedFile
	newFileName := changedFile
	changed := complexMoveReg.FindStringSubmatch(changedFile)
	// examples: cmd/{call_graph.go => call.go}
	SUCCESS_MATCH_LENGTH := 5
	if len(changed) == SUCCESS_MATCH_LENGTH {
		var oldLastChanged = changed[4]
		// TODO: support for Windows rename
		if changed[2] == "" {
			if strings.HasPrefix(oldLastChanged, "/") {
				oldLastChanged = oldLastChanged[1:]
			}
		}

		oldFileName = changed[1] + changed[2] + oldLastChanged
		newFileName = changed[1] + changed[3] + changed[4]

		changedFile = newFileName
	}
	return changedFile, oldFileName, newFileName
}


