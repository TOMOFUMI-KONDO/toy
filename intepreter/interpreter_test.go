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

func TestInterpreterIdentifier(t *testing.T) {
	interpreter = Interpreter{
		environment: map[string]int{
			"key": 1,
		},
	}

	exp := _ast.Identifier("key")
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpretAssignment(t *testing.T) {
	var exp ast.Expression

	exp = ast.Assignment{
		Name:       "key",
		Expression: ast.IntegerLiteral{Value: 1},
	}
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
	// environment should be set
	cond := interpreter.environment["key"]
	if cond != 1 {
		t.Errorf("environment.key = %d; want 1", cond)
	}
}

func TestInterpreterIf(t *testing.T) {
	exp := ast.IfExpression{
		Condition:  ast.IntegerLiteral{Value: 1}, // true
		ThenClause: ast.IntegerLiteral{Value: 2},
		ElseClause: ast.IntegerLiteral{Value: 3},
	}
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 2 {
		t.Errorf("result = %d; want 2", result)
	}

	exp.Condition = ast.IntegerLiteral{Value: 0} //false
	result, err = interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}

	exp.ElseClause = nil
	result, err = interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	// should be evaluated 1 if ElseClause is nil
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterWhile(t *testing.T) {
	interpreter = Interpreter{
		environment: map[string]int{
			"condition": 10,
		},
	}

	identifier := ast.Identifier{Name: "condition"}
	exp := ast.WhileExpression{
		Condition: identifier,
		// decrement environment.condition
		Body: ast.Assignment{
			Name: "condition",
			Expression: ast.BinaryExpression{
				Operator: ast.SUBTRACT,
				Lhs:      identifier,
				Rhs:      ast.IntegerLiteral{Value: 1},
			},
		},
	} // loop while environment['"key"] != 0

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
	cond := interpreter.environment["condition"]
	if cond != 0 {
		t.Errorf("environment.condition = %d; want 0", cond)
	}
}

func TestInterpreterBlock(t *testing.T) {
	identifier := ast.Identifier{Name: "a"}

	/*
		a = 0
		a = a + 10
		a * 2
	*/
	exp := ast.BlockExpression{
		Expressions: []ast.Expression{
			ast.Assignment{
				Name:       "a",
				Expression: ast.IntegerLiteral{Value: 0},
			},
			ast.Assignment{
				Name: "a",
				Expression: ast.BinaryExpression{
					Operator: ast.ADD,
					Lhs:      identifier,
					Rhs:      ast.IntegerLiteral{Value: 10},
				},
			},
			ast.BinaryExpression{
				Operator: ast.MULTIPLY,
				Lhs:      identifier,
				Rhs:      ast.IntegerLiteral{Value: 2},
			},
		},
	}

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 20 {
		t.Errorf("result = %d; want 20", result)
	}
	a := interpreter.environment["a"]
	if a != 10 {
		t.Errorf("environemnt.a = %d; want 10", a)
	}
}
