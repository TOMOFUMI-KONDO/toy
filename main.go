package main

import (
	"fmt"

	"github.com/TOMOFUMI-KONDO/toy/ast"
	interpreter "github.com/TOMOFUMI-KONDO/toy/intepreter"
)

func main() {
	a := ast.Ast{}
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

	i := interpreter.Interpreter{}
	result, err := i.Interpret(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
}
