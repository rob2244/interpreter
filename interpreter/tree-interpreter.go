package interpreter

import (
	"fmt"
	"reflect"
	"robinseitz/interpreter/ast"
	"robinseitz/interpreter/lexer"
	"strings"
)

type TreeInterpreter struct {
	parser *ast.Parser
}

func NewTreeInterpreter(parser *ast.Parser) *TreeInterpreter {
	return &TreeInterpreter{parser: parser}
}

func (i *TreeInterpreter) visit(node interface{}) int64 {
	nodeStringType := reflect.TypeOf(node).String()
	p := strings.Index(nodeStringType, ".")
	nodeStringType = nodeStringType[p+1:]
	nodeStringType = "Visit" + nodeStringType

	met := reflect.ValueOf(i).MethodByName(nodeStringType)
	value := met.Call([]reflect.Value{reflect.ValueOf(node)})

	return value[0].Int()
}

func (i *TreeInterpreter) VisitBinOp(node *ast.BinOp) (int64, error) {
	switch node.Token.Type {
	case lexer.PLUS:
		return i.visit(node.Left) + i.visit(node.Right), nil
	case lexer.MINUS:
		return i.visit(node.Left) - i.visit(node.Right), nil
	case lexer.MULTIPLY:
		return i.visit(node.Left) * i.visit(node.Right), nil
	case lexer.DIVIDE:
		return i.visit(node.Left) / i.visit(node.Right), nil
	default:
		return 0, fmt.Errorf("Invalid Token")
	}
}

func (i *TreeInterpreter) VisitNum(node *ast.Num) int64 {
	return node.Value
}

func (i *TreeInterpreter) Interpret() int64 {
	tree := i.parser.Parse()
	return i.visit(tree)
}
