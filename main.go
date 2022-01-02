package main

import (
	"fmt"

	"github.com/TOMOFUMI-KONDO/toy/ast"
	"github.com/TOMOFUMI-KONDO/toy/interpreter"
)

func main() {
	a := ast.NewAst()
	// ( 1 - 2 * 3 ) + 4
	exp := a.Add(
		a.Subtract(
			a.Integer(1),
			a.Multiply(
				a.Integer(2),
				a.Integer(3),
			),
		),
		a.Integer(4),
	)

	i := interpreter.NewInterpreter()
	result, err := i.Interpret(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
}
