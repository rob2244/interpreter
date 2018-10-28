package main

import (
	"github.com/rob2244/interpreter/ast"
	"github.com/rob2244/interpreter/interpreter"
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
	mulToken := interpreter.NewToken(interpreter.MULTIPLY, '*')
	plusToken := interpreter.NewToken(interpreter.PLUS, '+')
	mulNode := ast.NewBinOp(
		mulToken,
		ast.NewNum(interpreter.NewToken(interpreter.INTEGER, 2)),
		ast.NewNum(interpreter.NewToken(interpreter.INTEGER, 7)))

	addNode := ast.NewBinOp(plusToken, mulNode,
		ast.NewNum(interpreter.NewToken(interpreter.INTEGER, 3)))
}
