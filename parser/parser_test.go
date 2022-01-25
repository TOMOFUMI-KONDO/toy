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
	if err := setUp(toy); err != nil {
		t.Fatal(err)
	}

	i := interpreter.NewInterpreter()
	result, err := i.CallMain(toy.Program)
	if err != nil {
		t.Fatal(err)
	}

	if result != 120 {
		t.Errorf("result = %d; want %d", result, 120)
	}
}

func setUp(t *Toy) error {
	if err := t.Init(); err != nil {
		return err
	}
	if err := t.Parse(); err != nil {
		return err
	}
	if err := t.ConvertAst(); err != nil {
		return err
	}

	return nil
}
