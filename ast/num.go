package ast

import (
	"github.com/rob2244/interpreter/lexer"
)

type Num struct {
	Token *lexer.Token
	Value int64
}

func NewNum(token *lexer.Token) *Num {
	return &Num{token, token.Value.(int64)}
}
