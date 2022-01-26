package interpreter

import (
	"os"
	"testing"

	"github.com/TOMOFUMI-KONDO/toy/ast"
)

var (
	interpreter Interpreter
)

func TestMain(m *testing.M) {
	interpreter = NewInterpreter()

	code := m.Run()
	os.Exit(code)
}

func TestInterpreterInteger(t *testing.T) {
	exp := ast.NewInteger(1)

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpretAdd(t *testing.T) {
	exp := ast.NewAdd(ast.NewInteger(1), ast.NewInteger(2))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}
}

func TestInterpretSubtract(t *testing.T) {
	exp := ast.NewSubtract(ast.NewInteger(10), ast.NewInteger(3))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 7 {
		t.Errorf("result = %d; want 7", result)
	}
}

func TestInterpretMultiply(t *testing.T) {
	exp := ast.NewMultiply(ast.NewInteger(2), ast.NewInteger(5))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 10 {
		t.Errorf("result = %d; want 10", result)
	}
}

func TestInterpretDivide(t *testing.T) {
	exp := ast.NewDivide(ast.NewInteger(10), ast.NewInteger(2))

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 5 {
		t.Errorf("result = %d; want 5", result)
	}
}

func TestInterpreterLessThan(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewLessThan(
		ast.NewInteger(1),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(ast.NewLessThan(
		ast.NewInteger(2),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(ast.NewLessThan(
		ast.NewInteger(3),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}
}

func TestInterpreterLessOrEqual(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewLessOrEqual(
		ast.NewInteger(1),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(ast.NewLessOrEqual(
		ast.NewInteger(2),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(ast.NewLessOrEqual(
		ast.NewInteger(3),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}
}

func TestInterpreterGreaterThan(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewGreaterThan(
		ast.NewInteger(1),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(ast.NewGreaterThan(
		ast.NewInteger(2),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(ast.NewGreaterThan(
		ast.NewInteger(3),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterEqual(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewEqual(
		ast.NewInteger(1),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(ast.NewEqual(
		ast.NewInteger(2),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(ast.NewEqual(
		ast.NewInteger(3),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}
}

func TestInterpreterNotEqual(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewNotEqual(
		ast.NewInteger(1),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(ast.NewNotEqual(
		ast.NewInteger(2),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(ast.NewNotEqual(
		ast.NewInteger(3),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterGreaterOrEqual(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewGreaterOrEqual(
		ast.NewInteger(1),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(ast.NewGreaterOrEqual(
		ast.NewInteger(2),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(ast.NewGreaterOrEqual(
		ast.NewInteger(3),
		ast.NewInteger(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterIdentifier(t *testing.T) {
	interpreter.varEnv = &ast.Environment{Bindings: map[string]int{"key": 1}}

	exp := ast.NewIdentifier("key")
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpretAssignment(t *testing.T) {
	exp := ast.NewAssignment("key", ast.NewInteger(1))
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
	// varEnv should be set
	cond := interpreter.varEnv.Bindings["key"]
	if cond != 1 {
		t.Errorf("varEnv.key = %d; want 1", cond)
	}
}

func TestInterpreterIf(t *testing.T) {
	exp := ast.NewIf(
		ast.NewEqual(ast.NewInteger(1), ast.NewInteger(1)), // true
		ast.NewBlock([]ast.Expression{ast.NewInteger(2)}),
		ast.NewBlock([]ast.Expression{ast.NewInteger(3)}),
	)
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 2 {
		t.Errorf("result = %d; want 2", result)
	}

	exp.Condition = ast.NewNotEqual(ast.NewInteger(1), ast.NewInteger(1)) //false
	result, err = interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}

	exp.ElseClause.Expressions = nil
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
	interpreter.varEnv = &ast.Environment{Bindings: map[string]int{"condition": 10}}

	identifier := ast.NewIdentifier("condition")

	/*
		while condition != 0 {
			condition = condition - 1
		}
	*/
	exp := ast.NewWhile(
		identifier,
		ast.NewBlock([]ast.Expression{
			ast.NewAssignment(
				"condition",
				ast.NewSubtract(identifier, ast.NewInteger(1)),
			),
		}),
	)

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
	cond := interpreter.varEnv.Bindings["condition"]
	if cond != 0 {
		t.Errorf("varEnv.condition = %d; want 0", cond)
	}
}

func TestInterpreterBlock(t *testing.T) {
	identifier := ast.NewIdentifier("a")

	/*
		a = 0
		a = a + 10
		a * 2
	*/
	exp := ast.NewBlock(
		[]ast.Expression{
			ast.NewAssignment("a", ast.NewInteger(0)),
			ast.NewAssignment("a", ast.NewAdd(identifier, ast.NewInteger(10))),
			ast.NewMultiply(identifier, ast.NewInteger(2)),
		},
	)

	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 20 {
		t.Errorf("result = %d; want 20", result)
	}
	a := interpreter.varEnv.Bindings["a"]
	if a != 10 {
		t.Errorf("environemnt.a = %d; want 10", a)
	}
}

func TestInterpreterPrintln(t *testing.T) {
	result, err := interpreter.Interpret(ast.NewPrintln(ast.NewInteger(2)))
	if err != nil {
		t.Errorf("failed to Interpret Println: %v", err)
	}
	if result != 2 {
		t.Errorf("result = %d; want 2", result)
	}
}

func TestInterpreterDefineAndCallFunction(t *testing.T) {
	n := ast.NewIdentifier("n")
	topLevels := []ast.TopLevel{
		/*
			define main() {
				n = 0
				fact(5);
			}
		*/
		ast.NewFuncDef("main", nil, ast.NewBlock(
			[]ast.Expression{
				ast.NewAssignment("n", ast.NewInteger(0)), // This will be overwritten by argument in fact().
				ast.NewFuncCall("fact", []ast.Expression{ast.NewInteger(5)}),
			},
		)),
		/*
			define fact(n) {
				if(n < 2)  {
					1;
				} else {
					n  * fact(n - 1);
				}
			}
		*/
		ast.NewFuncDef("fact", []string{"n"}, ast.NewBlock([]ast.Expression{
			ast.NewIf(
				ast.NewLessThan(n, ast.NewInteger(2)),
				ast.NewBlock([]ast.Expression{ast.NewInteger(1)}),
				ast.NewBlock([]ast.Expression{
					ast.NewMultiply(n, ast.NewFuncCall("fact", []ast.Expression{
						ast.NewSubtract(n, ast.NewInteger(1)),
					})),
				}),
			),
		})),
	}

	result, err := interpreter.CallMain(ast.NewProgram(topLevels))
	if err != nil {
		t.Errorf("failed to CallMain: %v", err)
	}
	// 5! = 120
	if result != 120 {
		t.Errorf("result = %d; want 120", result)
	}
}

func TestInterpreterGlobalVarDef(t *testing.T) {
	topLevels := []ast.TopLevel{
		/*
			n = 1
			m = 2

			define main() {
				n + m
			}
		*/
		ast.NewGlobalVarDef("n", ast.NewInteger(1)),
		ast.NewGlobalVarDef("m", ast.NewInteger(3)), // This will be overwritten in main().
		ast.NewFuncDef("main", nil, ast.NewBlock([]ast.Expression{
			ast.NewAssignment("m", ast.NewInteger(2)),
			ast.NewAdd(ast.NewIdentifier("n"), ast.NewIdentifier("m")),
		})),
	}

	result, err := interpreter.CallMain(ast.NewProgram(topLevels))
	if err != nil {
		t.Errorf("failed to CallMain: %v", err)
	}
	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}
}
