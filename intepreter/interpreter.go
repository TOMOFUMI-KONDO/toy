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
		case ast.LESS_THAN:
			if lhs < rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.LESS_OR_EQUAL:
			if lhs <= rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.GREATER_THAN:
			if lhs > rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.GREATER_OR_EQUAL:
			if lhs >= rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.EQUAL:
			if lhs == rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.NOT_EQUAL:
			if lhs != rhs {
				return 1, nil
			} else {
				return 0, nil
			}
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
			return 0, fmt.Errorf("failed to Interpert expression of assignment: %w", err)
		}
		i.environment[assignment.Name] = v
		return v, nil
	}

	ifExp, ok := e.(ast.IfExpression)
	if ok {
		//cond, err := i.Interpret(ifExp.Condition)
		//if err != nil {
		//	return 0, fmt.Errorf("failed to Interpret condition of IfExpression: %w", err)
		//}
		cond, err := i.evalCondition(ifExp.Condition)
		if err != nil {
			return 0, fmt.Errorf("failed to eval condition of IfExpression")
		}

		var result int
		if cond /* NOTE: evaluate true if cond is not 0 */ {
			result, err = i.Interpret(ifExp.ThenClause)
		} else if ifExp.ElseClause != nil {
			result, err = i.Interpret(ifExp.ElseClause)
		} else {
			// NOTE: evaluate 1 if cond is false and elseClause is nil
			return 1, nil
		}

		if err != nil {
			return 0, fmt.Errorf("failed to Interptret ThenClause of IfExpression: %w", err)
		}
		return result, nil
	}

	whileExp, ok := e.(ast.WhileExpression)
	if ok {
		// loop body while cond is true, then return 1
		for {
			cond, err := i.evalCondition(whileExp.Condition)
			if err != nil {
				return 0, fmt.Errorf("failed to eval condition of WhileExpression: %w", err)
			}

			if cond {
				if _, err := i.Interpret(whileExp.Body); err != nil {
					return 0, fmt.Errorf("failed to Interpret body of WhileExp: %w", err)
				}
			} else {
				break
			}
		}

		return 1, nil
	}

	blockExp, ok := e.(ast.BlockExpression)
	if ok {
		var err error
		result := 0

		// evaluate all expressions, then return last expression.
		for _, exp := range blockExp.Expressions {
			result, err = i.Interpret(exp)
			if err != nil {
				return 0, fmt.Errorf("failed to Interpret one of Expressions of BlockExpression: %v", err)
			}
		}

		return result, nil
	}

	return 0, fmt.Errorf("unexpected expression: %v", e)
}

func (i *Interpreter) evalCondition(cond ast.Expression) (bool, error) {
	condInt, err := i.Interpret(cond)
	if err != nil {
		return false, fmt.Errorf("failed to Interpret condition: %w", err)
	}

	return condInt != 0, nil
}
