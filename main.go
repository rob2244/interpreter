package main

import (
	"fmt"
	"robinseitz/interpreter/interpreter"
)

func main() {
	expr := "3+4"

	i := interpreter.NewInterpreter(expr)

	fmt.Println(i.Expr())
}
