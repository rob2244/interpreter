package interpreter

import (
	"fmt"
	"strconv"
	"unicode"
)

type Lexer interface {
	GetNextToken() (*Token, error)
}

type MultLexer struct {
	text string
	pos  int
}

func NewLexer(txt string) *MultLexer {
	return &MultLexer{pos: 0, text: txt}
}

func (l *MultLexer) GetNextToken() (*Token, error) {
	l.skipWhitespace()

	if l.pos > len(l.text)-1 {
		return NewToken(EOF, nil), nil
	}

	currentChar := l.currentChar()

	switch {
	case currentChar == '*':
		l.pos++
		return NewToken(MULTIPLY, currentChar), nil
	case currentChar == '/':
		l.pos++
		return NewToken(DIVIDE, currentChar), nil
	case currentChar == '+':
		l.pos++
		return NewToken(PLUS, currentChar), nil
	case currentChar == '-':
		l.pos++
		return NewToken(MINUS, currentChar), nil
	case currentChar == '(':
		l.pos++
		return NewToken(LPAREN, currentChar), nil
	case currentChar == ')':
		l.pos++
		return NewToken(RPAREN, currentChar), nil
	case unicode.IsDigit(currentChar):
		return NewToken(INTEGER, l.integer()), nil
	default:
		return nil, fmt.Errorf("Unrecognized character %v", currentChar)
	}
}

func (l *MultLexer) currentChar() rune {
	return rune(l.text[l.pos])
}

func (l *MultLexer) integer() int64 {
	var numString string

	for ; l.pos < len(l.text) && unicode.IsDigit(l.currentChar()); l.pos++ {
		numString += string(l.currentChar())
	}

	result, _ := strconv.ParseInt(numString, 10, 64)

	return result
}

func (l *MultLexer) skipWhitespace() {
	if l.pos > len(l.text)-1 {
		return
	}

	for ; l.currentChar() == ' '; l.pos++ {

	}
}
