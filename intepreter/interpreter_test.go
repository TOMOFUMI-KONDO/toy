package interpreter

import (
	"os"
	"testing"

	"github.com/TOMOFUMI-KONDO/toy/ast"
)

var (
	_ast        ast.Ast
	interpreter Interpreter
)

func TestMain(m *testing.M) {
	_ast = ast.Ast{}
	interpreter = Interpreter{}

	code := m.Run()
	os.Exit(code)
}

func TestInterpreterIntLiteral(t *testing.T) {
	exp := _ast.Integer(1)

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpretAdd(t *testing.T) {
	exp := _ast.Add(_ast.Integer(1), _ast.Integer(2))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}
}

func TestInterpretSubtract(t *testing.T) {
	exp := _ast.Subtract(_ast.Integer(10), _ast.Integer(3))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 7 {
		t.Errorf("result = %d; want 7", result)
	}
}

func TestInterpretMultiply(t *testing.T) {
	exp := _ast.Multiply(_ast.Integer(2), _ast.Integer(5))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 10 {
		t.Errorf("result = %d; want 10", result)
	}
}

func TestInterpretDivide(t *testing.T) {
	exp := _ast.Divide(_ast.Integer(10), _ast.Integer(2))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 5 {
		t.Errorf("result = %d; want 5", result)
	}
}
