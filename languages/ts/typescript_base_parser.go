package parser

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// TypeScriptBaseParser implementation.
type TypeScriptBaseParser struct {
	*antlr.BaseParser
}

// Short for p.prev(str string)
func (p *TypeScriptBaseParser) p(str string) bool {
	return p.prev(str)
}

// Whether the previous token value equals to str.
func (p *TypeScriptBaseParser) prev(str string) bool {
	return p.GetTokenStream().LT(-1).GetText() == str
}

// Short for p.next(str string)
func (p *TypeScriptBaseParser) n(str string) bool {
	return p.next(str)
}

// Whether the next token value equals to str.
func (p *TypeScriptBaseParser) next(str string) bool {
	return p.GetTokenStream().LT(1).GetText() == str
}

func (p *TypeScriptBaseParser) notLineTerminator() bool {
	return !p.here(TypeScriptParserLineTerminator)
}

func (p *TypeScriptBaseParser) notOpenBraceAndNotFunction() bool {
	nextTokenType := p.GetTokenStream().LT(1).GetTokenType()
	return nextTokenType != TypeScriptParserOpenBrace && nextTokenType != TypeScriptParserFunction
}

func (p *TypeScriptBaseParser) closeBrace() bool {
	return p.GetTokenStream().LT(1).GetTokenType() == TypeScriptParserCloseBrace
}

// Returns true if on the current index of the parser's
// token stream a token of the given type exists on the
// Hidden channel.
func (p *TypeScriptBaseParser) here(_type int) bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	// Check if the token resides on the HIDDEN channel and if it's of the
	// provided type.
	return ahead.GetChannel() == antlr.LexerHidden && ahead.GetTokenType() == _type
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *TypeScriptBaseParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return false
	}

	if ahead.GetTokenType() == TypeScriptParserLineTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == TypeScriptParserWhiteSpaces {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == TypeScriptParserMultiLineComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == TypeScriptParserLineTerminator)
}
