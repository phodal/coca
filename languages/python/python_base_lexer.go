package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/infrastructure/container"
)

var TabSize = 8
var indents *container.Stack
var buffer []antlr.Token

func init() {
	indents = container.NewStack()
	buffer = make([]antlr.Token, 32)
}

type PythonBaseLexer struct {
	*antlr.BaseLexer

	firstTokenIndex int
	lastTokenIndex  int
	opened          int
	lastToken       antlr.Token
}

func (l *PythonBaseLexer) BuildDefaultToken(tokenType int) antlr.Token {
	return l.BuildTokenByType(tokenType, antlr.LexerDefaultTokenChannel, "")
}

func (l *PythonBaseLexer) BuildTokenByType(tokenType int, channel int, text string) antlr.Token {
	cpos := l.GetCharPositionInLine()
	lpos := l.GetLine()
	commonToken := antlr.NewCommonTokenFactory(false).Create(
		l.GetTokenSourceCharStreamPair(),
		tokenType,
		text,
		channel,
		l.GetCharIndex()-len(text),
		l.GetCharIndex(),
		lpos,
		cpos)

	return commonToken
}

func (l *PythonBaseLexer) EmitToken(token antlr.Token) {
	l.BaseLexer.EmitToken(token)

	if buffer[l.firstTokenIndex] != nil {
		l.lastTokenIndex = l.IncTokenInd(l.lastTokenIndex)

		if l.firstTokenIndex == l.lastTokenIndex {
			var newArray = make([]antlr.Token, len(buffer)*2)
			destIndex := len(newArray) - (len(buffer) - l.firstTokenIndex)
			copy(newArray, buffer)
			copy(newArray, buffer[:len(buffer)-l.firstTokenIndex])

			l.firstTokenIndex = destIndex
			buffer = newArray
		}
	}

	buffer[l.lastTokenIndex] = token
	l.lastToken = token
}

func (l *PythonBaseLexer) IncIndentLevel() {
	l.opened++
}

func (l *PythonBaseLexer) DecIndentLevel() {
	if l.opened > 0 {
		l.opened--
	}
}

func (l *PythonBaseLexer) NextToken() antlr.Token {
	if l.GetInputStream().LA(1) == antlr.TokenEOF && indents.Len() > 0 {
		if buffer[l.lastTokenIndex] == nil || buffer[l.lastTokenIndex].GetTokenType() != PythonLexerLINE_BREAK {
			l.EmitToken(l.BuildDefaultToken(PythonLexerLINE_BREAK))
		}

		for indents.Len() != 0 {
			l.EmitToken(l.BuildDefaultToken(PythonLexerDEDENT))
			indents.Pop()
		}
	}

	next := l.BaseLexer.NextToken()

	if buffer[l.firstTokenIndex] == nil {
		return next
	}

	var result = buffer[l.firstTokenIndex]
	buffer[l.firstTokenIndex] = nil

	if l.firstTokenIndex != l.lastTokenIndex {
		l.firstTokenIndex = l.IncTokenInd(l.firstTokenIndex)
	}

	return result
}

func (l *PythonBaseLexer) HandleNewLine() {
	l.EmitToken(l.BuildTokenByType(PythonLexerNEWLINE, antlr.LexerHidden, l.GetText()))

	next := string(rune(l.GetInputStream().LA(1)))
	if next != " " && next != "\t" && l.IsNotNewLineOrComment(next) {
		l.ProcessNewLine(0)
	}
}

func (l *PythonBaseLexer) HandleSpaces() {
	next := string(rune(l.GetInputStream().LA(1)))

	// class lost space here
	if (l.lastToken == nil || l.lastToken.GetTokenType() == PythonLexerNEWLINE) && l.IsNotNewLineOrComment(next) {
		indent := 0
		text := l.GetText()

		for _, char := range text {
			if char == '\t' {
				indent = indent + TabSize - indent%TabSize
			} else {
				indent = indent + 1
			}
		}

		l.ProcessNewLine(0)
	}

	l.EmitToken(l.BuildTokenByType(PythonLexerWS, antlr.LexerHidden, l.GetText()))
}

func (l *PythonBaseLexer) IsNotNewLineOrComment(next string) bool {
	return l.opened == 0 && next != "\r" && next != "\n" && next != "\f" && next != "#"
}

func (l *PythonBaseLexer) IncTokenInd(index int) int {
	return (index + 1) % len(buffer)
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
