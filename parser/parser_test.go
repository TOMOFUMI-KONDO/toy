package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	exp :=
		`define factorial(n) {
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
}
