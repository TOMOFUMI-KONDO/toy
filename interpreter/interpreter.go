package interpreter

import (
	"fmt"
	"io"
	"os"

	"github.com/TOMOFUMI-KONDO/toy/ast"
)

const MainFuncName = "main"

type Interpreter struct {
	varEnv  *ast.Environment
	funcEnv map[string]ast.FunctionDefinition
	writer  io.Writer
}

func NewInterpreter() Interpreter {
	return Interpreter{
		varEnv:  ast.NewEnvironment(nil),
		funcEnv: map[string]ast.FunctionDefinition{},
		writer:  os.Stdout,
	}
}

func NewInterpreterWithWriter(w io.Writer) Interpreter {
	i := NewInterpreter()
	i.writer = w
	return i
}

func (i *Interpreter) Interpret(intf ast.Expression) (int, error) {
	switch exp := intf.(type) {
	case ast.BinaryExpression:
		lhs, err := i.Interpret(exp.Lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := i.Interpret(exp.Rhs)
		if err != nil {
			return 0, err
		}

		switch exp.Operator {
		case ast.Add:
			return lhs + rhs, nil
		case ast.Subtract:
			return lhs - rhs, nil
		case ast.Multiply:
			return lhs * rhs, nil
		case ast.Divide:
			return lhs / rhs, nil
		case ast.LessThan:
			if lhs < rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.LessOrEqual:
			if lhs <= rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.GreaterThan:
			if lhs > rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.GreaterOrEqual:
			if lhs >= rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.Equal:
			if lhs == rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		case ast.NotEqual:
			if lhs != rhs {
				return 1, nil
			} else {
				return 0, nil
			}
		default:
			return 0, fmt.Errorf("invalid operator: %v", exp.Operator)
		}

	case ast.IntegerLiteral:
		return exp.Value, nil

	case ast.Identifier:
		b := i.varEnv.FindBinding(exp.Name)
		return b[exp.Name], nil

	case ast.Assignment:
		v, err := i.Interpret(exp.Expression)
		if err != nil {
			return 0, fmt.Errorf("failed to Interpert expression of assignment: %w", err)
		}

		b := i.varEnv.FindBinding(exp.Name)
		if b == nil {
			// assign new environment
			i.varEnv.Bindings[exp.Name] = v
		} else {
			// reassignment
			b[exp.Name] = v
		}

		return v, nil

	case ast.IfExpression:
		cond, err := i.evalCondition(exp.Condition)
		if err != nil {
			return 0, fmt.Errorf("failed to eval condition of IfExpression")
		}

		var result int
		if cond /* NOTE: evaluate true if cond is not 0 */ {
			result, err = i.Interpret(exp.ThenClause)
		} else if exp.ElseClause.Expressions != nil {
			result, err = i.Interpret(exp.ElseClause)
		} else {
			// NOTE: evaluate 1 if cond is false and elseClause is nil
			return 1, nil
		}

		if err != nil {
			return 0, fmt.Errorf("failed to Interptret ThenClause of IfExpression: %w", err)
		}
		return result, nil

	case ast.WhileExpression:
		// loop body while cond is true, then return 1
		for {
			cond, err := i.evalCondition(exp.Condition)
			if err != nil {
				return 0, fmt.Errorf("failed to eval condition of WhileExpression: %w", err)
			}

			if cond {
				if _, err := i.Interpret(exp.Body); err != nil {
					return 0, fmt.Errorf("failed to Interpret body of WhileExp: %w", err)
				}
			} else {
				break
			}
		}

		return 1, nil

	case ast.BlockExpression:
		var err error
		result := 0

		// evaluate all expressions, then return last expression.
		for _, exp := range exp.Expressions {
			result, err = i.Interpret(exp)
			if err != nil {
				return 0, fmt.Errorf("failed to Interpret one of Expressions of BlockExpression: %v", err)
			}
		}

		return result, nil

	case ast.Println:
		result, err := i.Interpret(exp.Arg)
		if err != nil {
			return 0, fmt.Errorf("failed to Interpret Println: %v", err)
		}

		if _, err := fmt.Fprint(i.writer, result); err != nil {
			return 0, err
		}

		return result, nil

	case ast.FunctionCall:
		funcDef, ok := i.funcEnv[exp.Name]
		if !ok {
			return 0, fmt.Errorf("function %s is not found", exp.Name)
		}

		var actualArgs []int
		for _, param := range exp.Args {
			result, err := i.Interpret(param)
			if err != nil {
				return 0, fmt.Errorf("failed to Interpret one of FunctionCall Args: %v", err)
			}
			actualArgs = append(actualArgs, result)
		}

		// make backup of variable definitions and restore later
		varEnvBackup := i.varEnv
		defer func() { i.varEnv = varEnvBackup }()

		// map function args to interpreter's variable definitions
		i.varEnv = ast.NewEnvironment(i.varEnv)
		for j, argName := range funcDef.Args {
			i.varEnv.Bindings[argName] = actualArgs[j]
		}

		// interpret with function scoped variable definitions
		result, err := i.Interpret(funcDef.Body)
		if err != nil {
			return 0, fmt.Errorf("failed to Interpret body of FunctionDefinition: %v", err)
		}

		return result, nil

	default:
		return 0, fmt.Errorf("unexpected expression: %v", exp)
	}
}

func (i *Interpreter) CallMain(program ast.Program) (int, error) {
	topLevels := program.Definitions
	for _, topLevel := range topLevels {
		funcDef, ok := topLevel.(ast.FunctionDefinition)
		if ok {
			i.funcEnv[funcDef.Name] = funcDef
			continue
		}

		globalVarDef, ok := topLevel.(ast.GlobalVariableDefinition)
		if ok {
			result, err := i.Interpret(globalVarDef.Expression)
			if err != nil {
				return 0, fmt.Errorf("failed to Interpret Expression of GlobalVariable Definition: %v", err)
			}
			i.varEnv.Bindings[globalVarDef.Name] = result
			continue
		}
	}

	if mainFunc, ok := i.funcEnv[MainFuncName]; ok {
		result, err := i.Interpret(mainFunc.Body)
		if err != nil {
			return 0, fmt.Errorf("failed to Interpret body of mainFunction: %v", err)
		}
		return result, nil
	} else {
		return 0, fmt.Errorf("this program doesn't have %s() function", MainFuncName)
	}
}

func (i *Interpreter) evalCondition(cond ast.Expression) (bool, error) {
	condInt, err := i.Interpret(cond)
	if err != nil {
		return false, fmt.Errorf("failed to Interpret condition: %w", err)
	}

	// eval as false if and only if condInt is 0
	return condInt != 0, nil
}
