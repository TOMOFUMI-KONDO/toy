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
	_ast = ast.NewAst()
	interpreter = NewInterpreter()

	code := m.Run()
	os.Exit(code)
}

func TestInterpreterInteger(t *testing.T) {
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

func TestInterpreterLessThan(t *testing.T) {
	result, err := interpreter.Interpret(_ast.LessThan(
		_ast.Integer(1),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(_ast.LessThan(
		_ast.Integer(2),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(_ast.LessThan(
		_ast.Integer(3),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}
}

func TestInterpreterLessOrEqual(t *testing.T) {
	result, err := interpreter.Interpret(_ast.LessOrEqual(
		_ast.Integer(1),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(_ast.LessOrEqual(
		_ast.Integer(2),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(_ast.LessOrEqual(
		_ast.Integer(3),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}
}

func TestInterpreterGreaterThan(t *testing.T) {
	result, err := interpreter.Interpret(_ast.GreaterThan(
		_ast.Integer(1),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(_ast.GreaterThan(
		_ast.Integer(2),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(_ast.GreaterThan(
		_ast.Integer(3),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterEqual(t *testing.T) {
	result, err := interpreter.Interpret(_ast.Equal(
		_ast.Integer(1),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(_ast.Equal(
		_ast.Integer(2),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(_ast.Equal(
		_ast.Integer(3),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}
}

func TestInterpreterNotEqual(t *testing.T) {
	result, err := interpreter.Interpret(_ast.NotEqual(
		_ast.Integer(1),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(_ast.NotEqual(
		_ast.Integer(2),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(_ast.NotEqual(
		_ast.Integer(3),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterGreaterOrEqual(t *testing.T) {
	result, err := interpreter.Interpret(_ast.GreaterOrEqual(
		_ast.Integer(1),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 0 {
		t.Errorf("result = %d; want 0", result)
	}

	result, err = interpreter.Interpret(_ast.GreaterOrEqual(
		_ast.Integer(2),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}

	result, err = interpreter.Interpret(_ast.GreaterOrEqual(
		_ast.Integer(3),
		_ast.Integer(2),
	))
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 1 {
		t.Errorf("result = %d; want 1", result)
	}
}

func TestInterpreterIdentifier(t *testing.T) {
	interpreter = Interpreter{
		varEnv:  &ast.Environment{Bindings: map[string]int{"key": 1}},
		funcEnv: map[string]ast.FunctionDefinition{},
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
	exp := _ast.Assignment("key", _ast.Integer(1))
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
	exp := _ast.If(
		_ast.Equal(_ast.Integer(1), _ast.Integer(1)), // true
		_ast.Block([]ast.Expression{_ast.Integer(2)}),
		_ast.Block([]ast.Expression{_ast.Integer(3)}),
	)
	result, err := interpreter.Interpret(exp)
	if err != nil {
		t.Errorf("failed to Interpret: %v", err)
	}
	if result != 2 {
		t.Errorf("result = %d; want 2", result)
	}

	exp.Condition = _ast.NotEqual(_ast.Integer(1), _ast.Integer(1)) //false
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
	interpreter = Interpreter{
		varEnv:  &ast.Environment{Bindings: map[string]int{"condition": 10}},
		funcEnv: map[string]ast.FunctionDefinition{},
	}

	identifier := _ast.Identifier("condition")

	/*
		while condition != 0 {
			condition = condition - 1
		}
	*/
	exp := _ast.While(
		identifier,
		_ast.Block([]ast.Expression{
			_ast.Assignment(
				"condition",
				_ast.Subtract(identifier, _ast.Integer(1)),
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
	identifier := _ast.Identifier("a")

	/*
		a = 0
		a = a + 10
		a * 2
	*/
	exp := _ast.Block(
		[]ast.Expression{
			_ast.Assignment("a", _ast.Integer(0)),
			_ast.Assignment("a", _ast.Add(identifier, _ast.Integer(10))),
			_ast.Multiply(identifier, _ast.Integer(2)),
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
	result, err := interpreter.Interpret(_ast.Println(_ast.Integer(2)))
	if err != nil {
		t.Errorf("failed to Interpret Println: %v", err)
	}
	if result != 2 {
		t.Errorf("result = %d; want 2", result)
	}
}

func TestInterpreterDefineAndCallFunction(t *testing.T) {
	n := _ast.Identifier("n")
	topLevels := []ast.TopLevel{
		/*
			define main() {
				n = 0
				fact(5);
			}
		*/
		_ast.DefineFunction("main", nil, _ast.Block(
			[]ast.Expression{
				_ast.Assignment("n", _ast.Integer(0)), // This will be overwritten by argument in fact().
				_ast.Call("fact", []ast.Expression{_ast.Integer(5)}),
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
		_ast.DefineFunction("fact", []string{"n"}, _ast.Block([]ast.Expression{
			_ast.If(
				_ast.LessThan(n, _ast.Integer(2)),
				_ast.Block([]ast.Expression{_ast.Integer(1)}),
				_ast.Block([]ast.Expression{
					_ast.Multiply(n, _ast.Call("fact", []ast.Expression{
						_ast.Subtract(n, _ast.Integer(1)),
					})),
				}),
			),
		})),
	}

	result, err := interpreter.CallMain(_ast.Program(topLevels))
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
		_ast.GlobalAssignment("n", _ast.Integer(1)),
		_ast.GlobalAssignment("m", _ast.Integer(3)), // This will be overwritten in main().
		_ast.DefineFunction("main", nil, _ast.Block([]ast.Expression{
			_ast.Assignment("m", _ast.Integer(2)),
			_ast.Add(_ast.Identifier("n"), _ast.Identifier("m")),
		})),
	}

	result, err := interpreter.CallMain(_ast.Program(topLevels))
	if err != nil {
		t.Errorf("failed to CallMain: %v", err)
	}
	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}
}
