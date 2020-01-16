package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/infrastructure/container"
)

var scopeStrictModes *container.Stack

func init() {
	scopeStrictModes = container.NewStack()
}

// TypeScriptBaseLexer state
type TypeScriptBaseLexer struct {
	*antlr.BaseLexer

	lastToken        antlr.Token
	useStrictDefault bool
	useStrictCurrent bool
}

func (l *TypeScriptBaseLexer) IsStartOfFile() bool {
	return l.lastToken == nil
}

// IsStrictMode is self explanatory.
func (l *TypeScriptBaseLexer) IsStrictMode() bool {
	return l.useStrictCurrent
}

func (l *TypeScriptBaseLexer) SetUseStrictDefault(value bool) {
	l.useStrictCurrent = value
	l.useStrictDefault = value
}

func (l *TypeScriptBaseLexer) GetStrictDefault() bool {
	return l.useStrictDefault
}

// NextToken from the character stream.
func (l *TypeScriptBaseLexer) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on default channel
		l.lastToken = next
	}
	return next
}

// ProcessOpenBrace is called when a { is encountered during
// lexing, we push a new scope everytime.
func (l *TypeScriptBaseLexer) ProcessOpenBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if scopeStrictModes.Len() > 0 && scopeStrictModes.Peak().(bool) {
		l.useStrictCurrent = true
	}
	scopeStrictModes.Push(l.useStrictCurrent)
}

// ProcessCloseBrace is called when a } is encountered during
// lexing, we pop a scope unless we're inside global scope.
func (l *TypeScriptBaseLexer) ProcessCloseBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if scopeStrictModes.Len() > 0 {
		l.useStrictCurrent = scopeStrictModes.Pop().(bool)
	}
}

// ProcessStringLiteral is called when lexing a string literal.
func (l *TypeScriptBaseLexer) ProcessStringLiteral() {
	if l.lastToken == nil || l.lastToken.GetTokenType() == TypeScriptLexerOpenBrace {
		if l.GetText() == `"use strict"` || l.GetText() == "'use strict'" {
			if scopeStrictModes.Len() > 0 {
				scopeStrictModes.Pop()
			}
			l.useStrictCurrent = true
			scopeStrictModes.Push(l.useStrictCurrent)
		}
	}
}

// IsRegexPossible returns true if the lexer can match a
// regex literal.
func (l *TypeScriptBaseLexer) IsRegexPossible() bool {
	if l.lastToken == nil {
		return true
	}
	switch l.lastToken.GetTokenType() {
	case
		TypeScriptLexerIdentifier,
		TypeScriptLexerNullLiteral,
		TypeScriptLexerBooleanLiteral,
		TypeScriptLexerThis,
		TypeScriptLexerCloseBracket,
		TypeScriptLexerCloseParen,
		TypeScriptLexerOctalIntegerLiteral,
		TypeScriptLexerDecimalLiteral,
		TypeScriptLexerHexIntegerLiteral,
		TypeScriptLexerStringLiteral,
		TypeScriptLexerPlusPlus,
		TypeScriptLexerMinusMinus:
		return false
	default:
		return true
	}
}
