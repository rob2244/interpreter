package interpreter

import (
	"fmt"
)

// TokenType defines the available tokens
type TokenType string

const (
	EOF      TokenType = "EOF"
	INTEGER  TokenType = "INTEGER"
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	DIVIDE   TokenType = "DIVIDE"
	MULTIPLY TokenType = "MULTIPLY"
)

// Token represents a syntax token
type Token struct {
	Type  TokenType
	Value interface{}
}

// NewToken creates a new Token
func NewToken(typeof TokenType, value interface{}) *Token {
	return &Token{typeof, value}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s, %s", t.Type, t.Value)
}
