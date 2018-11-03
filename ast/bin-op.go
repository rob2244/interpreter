package ast

import (
	"robinseitz/interpreter/lexer"
)

type BinOp struct {
	Left  interface{}
	Token *lexer.Token
	Op    rune
	Right interface{}
}

func NewBinOp(token *lexer.Token, left, right interface{}) *BinOp {
	return &BinOp{left, token, token.Value.(rune), right}
}
