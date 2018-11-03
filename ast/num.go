package ast

import (
	"robinseitz/interpreter/lexer"
)

type Num struct {
	Token *lexer.Token
	Value int64
}

func NewNum(token *lexer.Token) *Num {
	return &Num{token, token.Value.(int64)}
}
