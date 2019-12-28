package todo

import parser "github.com/phodal/coca/languages/java"

type EmptyListener struct {
	parser.BaseJavaParserListener
}

func NewEmptyListener() *EmptyListener {
	return &EmptyListener{}
}
