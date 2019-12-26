package models


type JImport struct {
	Name   string
	StartLine         int
	StopLine          int
}

type JMoveStruct struct {
	JFullIdentifier

	Path string
	Deps []JImport
}
