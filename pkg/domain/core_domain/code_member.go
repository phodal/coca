package core_domain

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
)

type CodeMember struct {
	ID            string
	Name          string
	Type          string
	ClassNodes    []jdomain.JClassNode
	FunctionNodes []CodeFunction
	Namespace     []string
	FileID        string
	DataStructID  string
	Position      CodePosition
}
