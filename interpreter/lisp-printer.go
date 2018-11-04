package interpreter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rob2244/interpreter/ast"
	"github.com/rob2244/interpreter/lexer"
)

type LISPPrinter struct {
	parser *ast.Parser
}

func NewLISPPrinter(parser *ast.Parser) *LISPPrinter {
	return &LISPPrinter{parser}
}

func (l *LISPPrinter) Print() {
	ast := l.parser.Parse()
	builder := &strings.Builder{}
	fmt.Println(l.visit(ast, builder).String())
}

func (l *LISPPrinter) visit(node interface{}, builder *strings.Builder) *strings.Builder {
	nodeStringType := reflect.TypeOf(node).String()
	p := strings.Index(nodeStringType, ".")
	nodeStringType = nodeStringType[p+1:]
	nodeStringType = "Visit" + nodeStringType

	met := reflect.ValueOf(l).MethodByName(nodeStringType)
	value := met.Call([]reflect.Value{reflect.ValueOf(node), reflect.ValueOf(builder)})

	return value[0].Interface().(*strings.Builder)
}

func (l *LISPPrinter) VisitBinOp(node *ast.BinOp, builder *strings.Builder) (*strings.Builder, error) {
	switch node.Token.Type {
	case lexer.PLUS:

		builder.WriteString("( + ")
		l.visit(node.Left, builder)
		l.visit(node.Right, builder)
		builder.WriteString(")")
		return builder, nil

	case lexer.MINUS:
		builder.WriteString("(- ")
		l.visit(node.Left, builder)
		l.visit(node.Right, builder)
		builder.WriteString(")")
		return builder, nil

	case lexer.MULTIPLY:
		builder.WriteString("(* ")
		l.visit(node.Left, builder)
		l.visit(node.Right, builder)
		builder.WriteString(")")
		return builder, nil

	case lexer.DIVIDE:
		builder.WriteString("(/ ")
		l.visit(node.Left, builder)
		l.visit(node.Right, builder)
		builder.WriteString(")")
		return builder, nil

	default:
		return builder, fmt.Errorf("Invalid Token")
	}
}

func (l *LISPPrinter) VisitNum(node *ast.Num, builder *strings.Builder) *strings.Builder {
	builder.WriteString(strconv.Itoa(int(node.Value)) + " ")
	return builder
}
