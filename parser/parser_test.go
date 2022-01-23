package parser

import (
	"testing"

	"github.com/TOMOFUMI-KONDO/toy/interpreter"
)

func TestParser(t *testing.T) {
	exp := `define factorial(n) {
	if n<2 {
		1
	} else {
		n*factorial(n-1)
	}
}
define main() {
	factorial(5)
}`
	toy := &Toy{Buffer: exp}
	if err := toy.Init(); err != nil {
		t.Fatal(err)
	}
	if err := toy.Parse(); err != nil {
		t.Fatal(err)
	}

	if err := toy.ConvertAst(); err != nil {
		t.Fatal(err)
	}

	_interpreter := interpreter.NewInterpreter()
	result, err := _interpreter.CallMain(toy.Program)
	if err != nil {
		t.Fatal(err)
	}

	expected := 120
	if result != expected {
		t.Errorf("result = %d; want %d", result, expected)
	}
}
