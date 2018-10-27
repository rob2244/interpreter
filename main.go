package main

import (
	"fmt"
	"robinseitz/interpreter/interpreter"
)

// func main() {
// 	expr := "12 + 2 - 10 + 22 + 10 - 12 * 12 / 2"

// 	i := interpreter.NewInterpreter(expr)

// 	fmt.Println(i.Expr())
// }

func main() {
	expr := "12 * 2 / 3 * 22 / 12 + 22 - 1 + 23"

	l := interpreter.NewLexer(expr)
	i := interpreter.NewMultInterpreter(l)

	fmt.Println(i.Expr())
}
