package parser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

var (
	Autodetect = 0
	Python2    = 2
	Python3    = 3
)

type PythonBaseParser struct {
	*antlr.BaseParser
	Version int
}

func (p *PythonBaseParser) CheckVersion(ver int) bool {
	return p.Version == Autodetect || p.Version == ver
}

func (p *PythonBaseParser) SetVersion(ver int) {
	if ver == 2 {
		p.Version = Python2
	} else if ver == 3 {
		p.Version = Python3
	}
}
