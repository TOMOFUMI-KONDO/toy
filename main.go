package main

import (
	"fmt"

	"github.com/TOMOFUMI-KONDO/toy/ast"

	"github.com/TOMOFUMI-KONDO/toy/interpreter"
)

func main() {
	// ( 1 - 2 * 3 ) + 4
	exp := ast.NewAdd(
		ast.NewSubtract(
			ast.NewInteger(1),
			ast.NewMultiply(
				ast.NewInteger(2),
				ast.NewInteger(3),
			),
		),
		ast.NewInteger(4),
	)

	i := interpreter.NewInterpreter()
	result, err := i.Interpret(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
}
