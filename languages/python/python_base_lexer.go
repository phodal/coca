package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/infrastructure/container"
)

var TabSize = 8
var indents *container.Stack

func init() {
	indents = container.NewStack()
}

type PythonBaseLexer struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken        antlr.Token
	useStrictDefault bool
	useStrictCurrent bool

	_opened         int
	buffer          []antlr.Token
	lastTokenIndex  int
	firstTokenIndex int
}

func (l *PythonBaseLexer) EmitToken(token antlr.Token) {
	l.BaseLexer.EmitToken(token)

	if l.buffer[l.firstTokenIndex] != nil {
		l.lastTokenIndex = l.IncTokenInd(l.lastTokenIndex)

		if l.firstTokenIndex != l.lastTokenIndex {
			var newArray = make([]antlr.Token, len(l.buffer)*2)
			destIndex := len(newArray) - (len(l.buffer) - l.firstTokenIndex)
			copy(newArray, l.buffer)
			copy(newArray, l.buffer[:len(l.buffer)-l.firstTokenIndex])

			l.firstTokenIndex = destIndex
			l.buffer = newArray
		}
	}

	l.buffer[l.lastTokenIndex] = token
	l.lastToken = token
}

func (l *PythonBaseLexer) IncIndentLevel() {
	l._opened++
}

func (l *PythonBaseLexer) DecIndentLevel() {
	if l._opened > 0 {
		l._opened--
	}
}

func (l *PythonBaseLexer) NextToken() antlr.Token {
	if l.GetInputStream().LA(1) == antlr.TokenEOF && indents.Len() > 0 {
		if l.buffer[l.lastTokenIndex] == nil || l.buffer[l.lastTokenIndex].GetTokenType() != PythonLexerLINE_BREAK {
			l.BaseLexer.EmitToken(l.BuildDefaultToken(PythonLexerLINE_BREAK))
		}

		for indents.Len() != 0 {
			l.BaseLexer.EmitToken(l.BuildDefaultToken(PythonLexerDEDENT))
			indents.Pop()
		}
	}

	next := l.BaseLexer.NextToken() // Get next token

	if l.buffer == nil {
		return next
	}

	if l.buffer[l.firstTokenIndex] == nil {
		return next
	}

	var result = l.buffer[l.firstTokenIndex]
	l.buffer[l.firstTokenIndex] = nil

	if l.firstTokenIndex != l.lastTokenIndex {
		l.firstTokenIndex = l.IncTokenInd(l.firstTokenIndex)
	}

	return result
}

func (l *PythonBaseLexer) BuildDefaultToken(tokenIndex int) antlr.Token {
	return l.BuildTokenByType(tokenIndex, antlr.LexerDefaultTokenChannel, "")
}

func (l *PythonBaseLexer) BuildTokenByType(tokenIndex int, channel int, text string) antlr.Token {
	cpos := l.GetCharPositionInLine()
	lpos := l.GetLine()
	token := l.GetTokenFactory().Create(l.GetTokenSourceCharStreamPair(), tokenIndex, text, channel, l.GetInputStream().Index(), l.GetInputStream().Index()-1, lpos, cpos)
	return token
}

func (l *PythonBaseLexer) HandleNewLine() {
	l.EmitToken(l.BuildTokenByType(PythonLexerNEWLINE, antlr.LexerHidden, l.GetText()))

	next := string(l.GetInputStream().LA(1))
	if next != " " && next != "\t" && l.IsNotNewLineOrComment(next) {
		l.ProcessNewLine(0)
	}
}

func (l *PythonBaseLexer) IsNotNewLineOrComment(next string) bool {
	return l._opened == 0 && next != "\r" && next != "\n" && next != "\f" && next != "#"
}

func (l *PythonBaseLexer) HandleSpaces() {
	next := string(l.GetInputStream().LA(1))

	if (l.lastToken == nil || l.lastToken.GetTokenType() == PythonLexerNEWLINE) && l.IsNotNewLineOrComment(next) {
		indent := 0
		text := l.GetText()

		for _, char := range text {
			if char == '\t' {
				indent = indent + TabSize - indent&TabSize
			} else {
				indent = 1
			}
		}

		l.ProcessNewLine(indent)
	}

	l.EmitToken(l.BuildTokenByType(PythonLexerWS, antlr.LexerHidden, l.GetText()))
}

func (l *PythonBaseLexer) IncTokenInd(index int) int {
	return (index + 1) % len(l.buffer)
}

func (l *PythonBaseLexer) ProcessNewLine(indent int) {
	l.EmitToken(l.BuildDefaultToken(PythonLexerLINE_BREAK))

	var previous = 0
	if indents.Len() != 0 {
		previous = indents.Peak().(int)
	}

	if indent > previous {
		indents.Push(indent)
		l.EmitToken(l.BuildDefaultToken(PythonLexerINDENT))
	} else {
		for indents.Len() != 0 && indents.Peak().(int) > indent {
			l.EmitToken(l.BuildDefaultToken(PythonLexerDEDENT))
			indents.Pop()
		}
	}

}
