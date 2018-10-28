package ast

import (
	"github.com/rob2244/interpreter/interpreter"
)

type Num struct {
	token *interpreter.Token
	value int64
}

func NewNum(token *interpreter.Token) *Num {
	return &Num{token, token.Value.(int64)}
}
