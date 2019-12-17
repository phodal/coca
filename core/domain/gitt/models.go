package gitt

import "time"

type CodeAge struct {
	File string
	Age  time.Time
}

type CodeAgeDisplay struct {
	File  string
	Month string
}

type TeamSummary struct {
	EntityName  string
	AuthorCount int
	RevsCount   int
}

type TeamInformation struct {
	EntityName string
	Authors    map[string]string
	Revs       map[string]string
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
}
