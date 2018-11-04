package ast

import (
	"github.com/rob2244/interpreter/lexer"
)

type UnaryOp struct {
	token *lexer.Token
	expr  interface{}
}
