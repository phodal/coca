package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type PythonBaseParser struct {
	*antlr.BaseParser
}

var (
	Autodetect = 0
	Python2    = 2
	Python3    = 3
	Version    int
)

func (p *PythonBaseParser) CheckVersion(ver int) bool {
	return Version == Autodetect || ver == Version
}

func (p *PythonBaseParser) SetVersion(ver int) {
	if ver == 2 {
		Version = Python2
	} else if ver == 3 {
		Version = Python3
	}
}
