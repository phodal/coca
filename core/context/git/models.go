package git

import "time"

type CodeAgeDisplay struct {
	EntityName string
	Month      string
}

type TeamSummary struct {
	EntityName  string
	AuthorCount int
	RevsCount   int
}

type ProjectInfo struct {
	EntityName string
	Authors    map[string]string
	Revs       map[string]string
	Age        time.Time
}

type GitSummary struct {
	Commits  int
	Entities int
	Changes  int
	Authors  int
}

type CommitMessage struct {
	Rev     string
	Author  string
	Date    string
	Message string
	Changes []FileChange
}

type FileChange struct {
	Added   int
	Deleted int
	File    string
	Mode    string
}
