package gitt

import (
	"bytes"
	"coca/core/domain/gitt/apriori"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var currentCommitMessage CommitMessage
var currentFileChanges []FileChange
var commitMessages []CommitMessage


var (
	rev       = `\[([\d|a-f]{5,12})\]`
	author    = `(.*?)\s\d{4}-\d{2}-\d{2}`
	date      = `\d{4}-\d{2}-\d{2}`
	changes   = `([\d-]+)[\t\s]+([\d-]+)[\t\s]+(.*)`
	moveRegSt = `(.*)\{(.*)\s=>\s(.*)\}(.*)`

	revReg     = regexp.MustCompile(rev)
	authorReg  = regexp.MustCompile(author)
	dateReg    = regexp.MustCompile(date)
	changesReg = regexp.MustCompile(changes)
	moveReg    = regexp.MustCompile(moveRegSt)
)

func BuildCommitMessage() []CommitMessage {
	historyArgs := []string{"log", "--pretty=format:[%h] %aN %ad %s", "--date=short", "--numstat"}
	cmd := exec.Command("git", historyArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return BuildMessageByInput(string(out))
}

func BuildMessageByInput(inputStr string) []CommitMessage {
	commitMessages = nil
	splitStr := strings.Split(inputStr, "\n")
	for _, str := range splitStr {
		ParseLog(str)
	}

	return commitMessages
}

func CalculateCodeAge(messages []CommitMessage) []CodeAgeDisplay {
	timeFormat := "2006-01-02"

	ages := make(map[string]CodeAge)
	for _, commitMessage := range messages {
		for _, change := range commitMessage.Changes {
			if ages[change.File].File == "" {
				date, _ := time.Parse(timeFormat, commitMessage.Date)
				ages[change.File] = *&CodeAge{change.File, date}
			}
		}
	}

	var agesArray []CodeAge
	for _, info := range ages {
		agesArray = append(agesArray, *&CodeAge{info.File, info.Age})
	}

	sort.Slice(agesArray, func(i, j int) bool {
		return agesArray[i].Age.Before(agesArray[j].Age)
	})

	var agesDisplay []CodeAgeDisplay
	for _, info := range agesArray {
		const secondsOfOneMonth = 2600640
		month := time.Now().Sub(info.Age).Seconds() / secondsOfOneMonth
		displayMonth := strconv.FormatFloat(month, 'f', 2, 64)
		agesDisplay = append(agesDisplay, *&CodeAgeDisplay{info.File, displayMonth})
	}

	return agesDisplay
}

func GetTeamSummary(messages []CommitMessage) []TeamSummary {
	infos := make(map[string]TeamInformation)
	for _, commitMessage := range messages {
		for _, change := range commitMessage.Changes {
			if moveReg.MatchString(change.File) {
				infos = switchFile(infos, change.File)
			}

			if infos[change.File].EntityName == "" {
				authors := make(map[string]string)
				authors[commitMessage.Author] = commitMessage.Author
				revs := make(map[string]string)
				revs[commitMessage.Rev] = commitMessage.Rev

				infos[change.File] = *&TeamInformation{change.File, authors, revs}
			} else {
				infos[change.File].Authors[commitMessage.Author] = commitMessage.Author
				infos[change.File].Revs[commitMessage.Rev] = commitMessage.Rev
			}
		}
	}

	var sortInfos []TeamSummary
	for _, info := range infos {
		sortInfos = append(sortInfos, *&TeamSummary{info.EntityName, len(info.Authors), len(info.Revs)})
	}

	sort.Slice(sortInfos, func(i, j int) bool {
		return sortInfos[i].RevsCount > sortInfos[j].RevsCount
	})

	return sortInfos
}

// 反向查询
func switchFile(infos map[string]TeamInformation, changedFile string) map[string]TeamInformation {
	changed := moveReg.FindStringSubmatch(changedFile)
	// examples: cmd/{call_graph.go => call.go}
	if len(changed) >= 5 {
		oldFileName := changed[1] + changed[2] + changed[4]
		newFileName := changed[1] + changed[3] + changed[4]

		fmt.Println(infos, oldFileName, newFileName)
		if _, ok := infos[oldFileName]; ok {
			oldInfo := infos[oldFileName]
			delete(infos, oldFileName)
			infos[newFileName] = oldInfo
		}
	}

	return infos
}

type TopAuthor struct {
	Name        string
	CommitCount int
	LineCount   int
}

func GetRelatedFiles(commitMessages []CommitMessage, relatedConfig []byte) []apriori.RelationRecord {
	var dataset [][]string
	for _, commitMessage := range commitMessages {
		var set []string
		for _, change := range commitMessage.Changes {
			if strings.HasSuffix(change.File, ".java") && !strings.HasSuffix(change.File, "Test.java") {
				if strings.Contains(change.File, "core/main/java/") {
					split := strings.Split(change.File, "core/main/java/")
					change.File = strings.ReplaceAll(split[1], "/", ".")
				}

				set = append(set, change.File)
			}
		}

		if len(set) > 2 {
			dataset = append(dataset, set)
		}
	}

	var newOptions apriori.Options = apriori.NewOptions(0.1, 0.9, 0, 0)

	decoder := json.NewDecoder(bytes.NewReader(relatedConfig))
	decoder.UseNumber()
	error := decoder.Decode(&newOptions)
	if error != nil {
		log.Fatal(error)
		return nil
	}

	fmt.Println(newOptions)
	apriori := apriori.NewApriori(dataset)
	result := apriori.Calculate(newOptions)

	for _, res := range result {
		items := res.GetSupportRecord().GetItems()
		if len(items) > 2 {
			fmt.Println(items)
			fmt.Println(res.GetSupportRecord().GetSupport())
		}
	}

	return result
}

func GetTopAuthors(commitMessages []CommitMessage) []TopAuthor {
	authors := make(map[string]*TopAuthor)
	for _, commitMessage := range commitMessages {
		if authors[commitMessage.Author] == nil {
			authors[commitMessage.Author] = &TopAuthor{commitMessage.Author, 0, 0}
		}
		authors[commitMessage.Author].CommitCount++
		for _, change := range commitMessage.Changes {
			authors[commitMessage.Author].LineCount = authors[commitMessage.Author].LineCount + change.Added
			authors[commitMessage.Author].LineCount -= change.Deleted
		}
	}

	var topAuthors []TopAuthor
	for _, info := range authors {
		topAuthors = append(topAuthors, *&TopAuthor{info.Name, info.CommitCount, info.LineCount})
	}

	sort.Slice(topAuthors, func(i, j int) bool {
		return topAuthors[i].CommitCount > topAuthors[j].CommitCount
	})

	return topAuthors
}

func BasicSummary(commitMessages []CommitMessage) *GitSummary {
	authors := make(map[string]string)
	entities := make(map[string]string)
	commits := len(commitMessages)
	changes := 0

	for _, commitMessage := range commitMessages {
		authors[commitMessage.Author] = commitMessage.Author
		for _, change := range commitMessage.Changes {
			entities[change.File] = change.File
			if change.Added > 0 {
				changes++
			}
			if change.Deleted > 0 {
				changes--
			}
		}
	}

	authorSummary := len(authors)
	entitySummary := len(entities)

	gitSummary := &GitSummary{commits, entitySummary, changes, authorSummary}
	return gitSummary
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
		msg = msg[1:]

		currentCommitMessage = *&CommitMessage{id[1], auth[1][1:], dat[0], msg, nil}
	} else if changesReg.MatchString(text) {
		changes := changesReg.FindStringSubmatch(text)
		deleted, _ := strconv.Atoi(changes[2])
		added, _ := strconv.Atoi(changes[1])
		change := &FileChange{added, deleted, changes[3]}

		currentFileChanges = append(currentFileChanges, *change)
	} else {
		if currentCommitMessage.Rev != "" {
			currentCommitMessage.Changes = currentFileChanges
			commitMessages = append(commitMessages, currentCommitMessage)

			currentCommitMessage = *&CommitMessage{"", "", "", "", nil}
			currentFileChanges = nil
		}
	}
}
