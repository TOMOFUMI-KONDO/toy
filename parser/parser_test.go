package parser

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/TOMOFUMI-KONDO/toy/interpreter"
)

type testCase struct {
	expression string
	expected   int
	printed    string
}

func (t *testCase) String() string {
	return fmt.Sprintf("expression: %s\nexpectred: %d\nprinted: %s\n", t.expression, t.expected, t.printed)
}

var tests = []testCase{
	// test integer
	{
		`define main() {
			1
		}`,
		1,
		"",
	},
	// test multiply
	{
		`define main() {
			2*3
		}`,
		6,
		"",
	},
	// test divide
	{
		`define main() {
			10/2
		}`,
		5,
		"",
	},
	// test add
	{
		`define main() {
			1+2
		}`,
		3,
		"",
	},
	// test subtract
	{
		`define main() {
			3-1
		}`,
		2,
		"",
	},
	// test lessThan
	{
		`define main() {
			1<2
		}`,
		1,
		"",
	},
	{
		`define main() {
			2<2
		}`,
		0,
		"",
	},
	{
		`define main() {
			3<2
		}`,
		0,
		"",
	},
	// test lessOrEqual
	{
		`define main() {
			1<=2
		}`,
		1,
		"",
	},
	{
		`define main() {
			2<=2
		}`,
		1,
		"",
	},
	{
		`define main() {
			3<=2
		}`,
		0,
		"",
	},
	// test greaterThan
	{
		`define main() {
			3>2
		}`,
		1,
		"",
	},
	{
		`define main() {
			2>2
		}`,
		0,
		"",
	},
	{
		`define main() {
			1>2
		}`,
		0,
		"",
	},
	// test greaterOrEqual
	{
		`define main() {
			3>=2
		}`,
		1,
		"",
	},
	{
		`define main() {
			2>=2
		}`,
		1,
		"",
	},
	{
		`define main() {
			1>=2
		}`,
		0,
		"",
	},
	// test Equal
	{
		`define main() {
			1==2
		}`,
		0,
		"",
	},
	{
		`define main() {
			2==2
		}`,
		1,
		"",
	},
	{
		`define main() {
			3==2
		}`,
		0,
		"",
	},
	// test NotEqual
	{
		`define main() {
			1!=2
		}`,
		1,
		"",
	},
	{
		`define main() {
			2!=2
		}`,
		0,
		"",
	},
	{
		`define main() {
			3!=2
		}`,
		1,
		"",
	},
	// test println
	{
		`define main() {
			println(2)
		}`,
		2,
		"2",
	},
	// test functionCall
	{
		`define two() {
			2
		}
		define main() {
			two()
		}`,
		2,
		"",
	},
	{
		`define oneArg(n) {
			n*2
		}
		define main() {
			oneArg(2)
		}`,
		4,
		"",
	},
	{
		`define twoArgs(a,b) {
			a+b
		}
		define main() {
			twoArgs(2,3)
		}`,
		5,
		"",
	},
	// test assignment
	{
		`define main() {
			n=2
		}`,
		2,
		"",
	},
	{
		`define shadow(n) {
			n=n*2
			n
		}
		define main() {
			n=2
			shadow(n)
		}`,
		4,
		"",
	},
	// test block
	{
		`define main() {
			{
				2
				3
				4
			}
		}`,
		4,
		"",
	},
	// test while
	{
		`define main() {
			n=5
			while n>0 {
				n=n-1
			}
			n
		}`,
		0,
		"",
	},
	// test if
	{
		`define main() {
			if 1 {
				2
			}
		}`,
		2,
		"",
	},
	{
		`define main() {
			if 0 {
				2
			} else {
				3
			}
		}`,
		3,
		"",
	},
	// test globalVariableDefinition
	{
		`global n=2
		define main() {
			n
		}`,
		2,
		"",
	},
	// test complex program
	{
		`define factorial(n) {
			if n<2 {
				1
			} else {
				n*factorial(n-1)
			}
		}
		define main() {
			result=factorial(5)
			println(result)
		}`,
		120,
		"120",
	},
}

func TestParser(t *testing.T) {
	for _, test := range tests {
		toy := &Toy{Buffer: test.expression}
		if err := setUp(toy); err != nil {
			t.Fatalf("%v\ntestCase = \n%s", err, test.String())
		}

		var buf bytes.Buffer
		i := interpreter.NewInterpreterWithWriter(&buf)

		result, err := i.CallMain(toy.Program)
		if err != nil {
			t.Fatalf("%v\ntestCase = \n%s", err, test.String())
		}

		if result != test.expected {
			t.Errorf("result = %d\ntestCase = \n%s", result, test.String())
		}

		printed := string(buf.Bytes())
		if printed != test.printed {
			t.Errorf("printed = %s\ntestCase = \n%s", printed, test.String())
		}
	}
}

func setUp(t *Toy) error {
	if err := t.Init(); err != nil {
		return err
	}
	if err := t.Parse(); err != nil {
		return err
	}
	return t.ConvertAst()
}
