package interpreter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rob2244/interpreter/ast"
	"github.com/rob2244/interpreter/lexer"
)

type RPNPrinter struct {
	parser *ast.Parser
}

func NewRPNPrinter(parser *ast.Parser) *RPNPrinter {
	return &RPNPrinter{parser}
}

func (r *RPNPrinter) Print() {
	ast := r.parser.Parse()
	builder := &strings.Builder{}
	fmt.Println(r.visit(ast, builder).String())
}

func (r *RPNPrinter) visit(node interface{}, builder *strings.Builder) *strings.Builder {
	nodeStringType := reflect.TypeOf(node).String()
	p := strings.Index(nodeStringType, ".")
	nodeStringType = nodeStringType[p+1:]
	nodeStringType = "Visit" + nodeStringType

	met := reflect.ValueOf(r).MethodByName(nodeStringType)
	value := met.Call([]reflect.Value{reflect.ValueOf(node), reflect.ValueOf(builder)})

	return value[0].Interface().(*strings.Builder)
}

func (r *RPNPrinter) VisitBinOp(node *ast.BinOp, builder *strings.Builder) (*strings.Builder, error) {
	switch node.Token.Type {
	case lexer.PLUS:
		r.visit(node.Left, builder)
		r.visit(node.Right, builder)
		builder.WriteString("+ ")
		return builder, nil

	case lexer.MINUS:
		r.visit(node.Left, builder)
		r.visit(node.Right, builder)
		builder.WriteString("- ")
		return builder, nil

	case lexer.MULTIPLY:
		r.visit(node.Left, builder)
		r.visit(node.Right, builder)
		builder.WriteString("* ")
		return builder, nil

	case lexer.DIVIDE:
		r.visit(node.Left, builder)
		r.visit(node.Right, builder)
		builder.WriteString("/ ")
		return builder, nil

	default:
		return builder, fmt.Errorf("Invalid Token")
	}
}

func (r *RPNPrinter) VisitNum(node *ast.Num, builder *strings.Builder) *strings.Builder {
	builder.WriteString(strconv.Itoa(int(node.Value)) + " ")
	return builder
}
