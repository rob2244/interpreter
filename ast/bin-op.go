package ast

import "github.com/rob2244/interpreter/interpreter"

type BinOp struct {
	left  interface{}
	token *interpreter.Token
	op    rune
	right interface{}
}

func NewBinOp(token *interpreter.Token, left, right interface{}) *BinOp {
	return &BinOp{left, token, token.Value.(rune), right}
}
