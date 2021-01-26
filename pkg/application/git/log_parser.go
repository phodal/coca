package git

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	rev               = `\[([\d|a-f]{5,12})\]`
	author            = `(.*?)\s\d{4}-\d{2}-\d{2}`
	date              = `\d{4}-\d{2}-\d{2}`
	changes           = `([\d-]+)[\t\s]+([\d-]+)[\t\s]+(.*)`
	complexMoveRegStr = `(.*)\{(.*)\s=>\s(.*)\}(.*)`
	basicMoveRegStr   = `(.*)\s=>\s(.*)`
	changeModel       = `\s(\w{1,6})\s(mode 100(\d){3})?\s?(.*)(\s\(\d{2}%\))?`

	revReg         = regexp.MustCompile(rev)
	authorReg      = regexp.MustCompile(author)
	dateReg        = regexp.MustCompile(date)
	changesReg     = regexp.MustCompile(changes)
	complexMoveReg = regexp.MustCompile(complexMoveRegStr)
	basicMvReg     = regexp.MustCompile(basicMoveRegStr)
	changeModeReg  = regexp.MustCompile(changeModel)
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
			oldLastChanged = strings.TrimPrefix(oldLastChanged, "/")
		}

		oldFileName = changed[1] + changed[2] + oldLastChanged
		newFileName = changed[1] + changed[3] + changed[4]

		changedFile = newFileName
	}
	return changedFile, oldFileName, newFileName
}

func ParseLog(text string) {
	allString := revReg.FindAllString(text, -1)
	if len(allString) == 1 {
		str := ""
		id := revReg.FindStringSubmatch(text)
		str = strings.Split(text, id[0])[1]
		auth := authorReg.FindStringSubmatch(str)
		str = strings.Split(str, auth[1])[1]
		dat := dateReg.FindStringSubmatch(str)
		msg := strings.Split(str, dat[0])[1]
		if len(msg) > 1 {
			msg = msg[1:]
		}

		currentCommit = CommitMessage{id[1], auth[1][1:], dat[0], msg, nil}
	} else if changesReg.MatchString(text) {
		changes := changesReg.FindStringSubmatch(text)
		deleted, _ := strconv.Atoi(changes[2])
		added, _ := strconv.Atoi(changes[1])
		fileFieldName := changes[3]
		change := FileChange{added, deleted, fileFieldName, ""}

		currentFileChangeMap[fileFieldName] = change
	} else if changeModeReg.MatchString(text) {
		buildChangeMode(text)
	} else if currentCommit.Rev != "" {
		for _, value := range currentFileChangeMap {
			currentFileChanges = append(currentFileChanges, value)
		}

		currentFileChangeMap = make(map[string]FileChange)
		currentCommit.Changes = currentFileChanges
		commits = append(commits, currentCommit)

		currentCommit = CommitMessage{"", "", "", "", nil}
		currentFileChanges = nil
	}
}

func buildChangeMode(text string) {
	matches := changeModeReg.FindStringSubmatch(text)

	CHANGE_MODE_INDEX := 4
	if len(matches) > CHANGE_MODE_INDEX {
		mode := matches[1]

		if _, ok := currentFileChangeMap[matches[CHANGE_MODE_INDEX]]; ok {
			change := currentFileChangeMap[matches[CHANGE_MODE_INDEX]]
			change.Mode = mode
			currentFileChangeMap[matches[CHANGE_MODE_INDEX]] = change
		} else if mode == "delete" {
			change := FileChange{
				Added:   0,
				Deleted: 0,
				File:    matches[CHANGE_MODE_INDEX],
				Mode:    "delete",
			}

			currentFileChanges = append(currentFileChanges, change)
		}
	}
}
