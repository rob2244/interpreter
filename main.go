package main

import (
	"bufio"
	"os"

	"github.com/rob2244/interpreter/ast"
	"github.com/rob2244/interpreter/interpreter"
	"github.com/rob2244/interpreter/lexer"
)

// func main() {
// 	expr := "12 + 2 - 10 + 22 + 10 - 12 * 12 / 2"

// 	i := interpreter.NewInterpreter(expr)

// 	fmt.Println(i.Expr())
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		exp := scanner.Text()

// 		l := interpreter.NewLexer(exp)
// 		i := interpreter.NewMultInterpreter(l)

// 		fmt.Println(i.Expr())
// 	}
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		exp := scanner.Text()

		l := lexer.NewLexer(exp)
		p := ast.NewParser(l)
		i := interpreter.NewLISPPrinter(p)

		i.Print()
	}
}
