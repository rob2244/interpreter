package ast

import "github.com/rob2244/interpreter/interpreter"

type Parser struct {
	lexer        interpreter.Lexer
	currentToken *interpreter.Token
}

func NewParser(lexer interpreter.Lexer) *Parser {
	token, _ := lexer.GetNextToken()
	return &Parser{lexer, token}
}
