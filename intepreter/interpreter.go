package interpreter

import (
	"fmt"

	"github.com/TOMOFUMI-KONDO/toy/ast"
)

type Interpreter struct {
	environment map[string]int
}

func (i *Interpreter) Interpret(e ast.Expression) (int, error) {
	binaryExp, ok := e.(ast.BinaryExpression)
	if ok {
		lhs, err := i.Interpret(binaryExp.Lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := i.Interpret(binaryExp.Rhs)
		if err != nil {
			return 0, err
		}

		switch binaryExp.Operator {
		case ast.ADD:
			return lhs + rhs, nil
		case ast.SUBTRACT:
			return lhs - rhs, nil
		case ast.MULTIPLY:
			return lhs * rhs, nil
		case ast.DIVIDE:
			return lhs / rhs, nil
		default:
			return 0, fmt.Errorf("invalid operator: %v", binaryExp.Operator)
		}
	}

	intLiteralExp, ok := e.(ast.IntegerLiteral)
	if ok {
		return intLiteralExp.Value, nil
	}

	identifier, ok := e.(ast.Identifier)
	if ok {
		return i.environment[identifier.Name], nil
	}

	assignment, ok := e.(ast.Assignment)
	if ok {
		v, err := i.Interpret(assignment.Expression)
		if err != nil {
			return 0, fmt.Errorf("failed to Interpert expression of assignment: %v", err)
		}
		i.environment[assignment.Name] = v
		return v, nil
	}

	return 0, fmt.Errorf("unexpected expression: %v", e)
}
