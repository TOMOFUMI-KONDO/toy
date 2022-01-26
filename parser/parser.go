package parser

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TOMOFUMI-KONDO/toy/ast"
)

type Logger struct {
	*log.Logger
}

func NewLogger() Logger {
	return Logger{
		log.New(os.Stderr, "[INFO]", log.LstdFlags),
	}
}

func (l *Logger) info(format string, v ...interface{}) {
	if os.Getenv("DEBUG") == "1" {
		l.Printf(format, v...)
	}
}

var infoLog Logger

func init() {
	infoLog = NewLogger()
}

func (p *Toy) ConvertAst() error {
	return p.program(p.AST())
}

func (p *Toy) program(node *node32) error {
	infoLog.info("program\n%s\n", p.tokenStr(node))

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruletopLevel:
			topLevel, err := p.topLevel(node)
			if err != nil {
				return err
			}
			p.Program.PushTopLevel(topLevel)
		}

		node = node.next
	}

	return nil
}

func (p *Toy) topLevel(node *node32) (ast.TopLevel, error) {
	infoLog.info("topLevel\n%s\n", p.tokenStr(node))

	node = node.up
	for node != nil {
		switch node.pegRule {
		case rulefunctionDefinition:
			funcDef, err := p.functionDefinition(node)
			if err != nil {
				return nil, err
			}
			return *funcDef, nil

		case ruleglobalVariableDefinition:
			globalVarDef, err := p.globalVariableDefinition(node)
			if err != nil {
				return nil, err
			}
			return *globalVarDef, nil
		}

		node = node.next
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) functionDefinition(node *node32) (*ast.FunctionDefinition, error) {
	infoLog.info("functionDefinition\n%s\n", p.tokenStr(node))

	var name string
	var args []string
	var body *ast.BlockExpression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			arg := p.tokenStr(node)
			if name == "" {
				name = arg
			} else {
				args = append(args, arg)
			}

		case ruleblockExpression:
			var err error
			body, err = p.block(node)
			if err != nil {
				return nil, err
			}
		}

		node = node.next
	}

	funcDef := ast.NewFuncDef(name, args, *body)
	return &funcDef, nil
}

func (p *Toy) globalVariableDefinition(node *node32) (*ast.GlobalVariableDefinition, error) {
	infoLog.info("topLevel\n%s\n", p.tokenStr(node))

	var name string
	var exp ast.Expression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			name = p.tokenStr(node)

		case ruleexpression:
			var err error
			exp, err = p.expression(node)
			if err != nil {
				return nil, err
			}
		}

		node = node.next
	}

	globalVarDef := ast.NewGlobalVarDef(name, exp)
	return &globalVarDef, nil
}

func (p *Toy) expression(node *node32) (ast.Expression, error) {
	infoLog.info("expression\n%s\n", p.tokenStr(node))

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleifExpression:
			exp, err := p.ifExp(node)
			return *exp, err

		case rulewhileExpression:
			exp, err := p.while(node)
			return *exp, err

		case ruleblockExpression:
			exp, err := p.block(node)
			return *exp, err

		case ruleassignment:
			exp, err := p.assignment(node)
			return *exp, err

		case rulecomparative:
			return p.comparative(node)
		}

		node = node.next
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) ifExp(node *node32) (*ast.IfExpression, error) {
	infoLog.info("ifExp\n%s\n", p.tokenStr(node))

	var cond ast.Expression
	var thenClause *ast.BlockExpression
	var elseClause *ast.BlockExpression

	node = node.up
	for node != nil {
		var err error

		switch node.pegRule {
		case rulecomparative:
			cond, err = p.comparative(node)
			if err != nil {
				return nil, err
			}

		case ruleblockExpression:
			if thenClause == nil {
				thenClause, err = p.block(node)
			} else {
				elseClause, err = p.block(node)
			}
			if err != nil {
				return nil, err
			}
		}

		node = node.next
	}

	var ifExp ast.IfExpression
	if elseClause == nil {
		ifExp = ast.NewIfWithoutElse(cond, *thenClause)
	} else {
		ifExp = ast.NewIf(cond, *thenClause, *elseClause)
	}
	return &ifExp, nil
}

func (p *Toy) while(node *node32) (*ast.WhileExpression, error) {
	infoLog.info("assignment\n%s\n", p.tokenStr(node))

	var cond ast.Expression
	var body *ast.BlockExpression

	node = node.up
	for node != nil {
		var err error

		switch node.pegRule {
		case rulecomparative:
			cond, err = p.comparative(node)
			if err != nil {
				return nil, err
			}

		case ruleblockExpression:
			body, err = p.block(node)
		}

		node = node.next
	}

	while := ast.NewWhile(cond, *body)
	return &while, nil
}

func (p *Toy) block(node *node32) (*ast.BlockExpression, error) {
	infoLog.info("block\n%s\n", p.tokenStr(node))

	var expressions []ast.Expression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleexpression:
			exp, err := p.expression(node)
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, exp)
		}

		node = node.next
	}

	block := ast.NewBlock(expressions)
	return &block, nil
}

func (p *Toy) assignment(node *node32) (*ast.Assignment, error) {
	infoLog.info("assignment\n%s\n", p.tokenStr(node))

	var name string
	var exp ast.Expression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			name = p.tokenStr(node)

		case ruleexpression:
			var err error
			exp, err = p.expression(node)
			if err != nil {
				return nil, err
			}
		}

		node = node.next
	}

	assignment := ast.NewAssignment(name, exp)
	return &assignment, nil
}

func (p *Toy) println(node *node32) (*ast.Println, error) {
	infoLog.info("println\n%s\n", p.tokenStr(node))

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleexpression:
			exp, err := p.expression(node)
			if err != nil {
				return nil, err
			}
			printlnExp := ast.NewPrintln(exp)
			return &printlnExp, nil
		}

		node = node.next
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) funcCall(node *node32) (*ast.FunctionCall, error) {
	infoLog.info("funcCall\n%s\n", p.tokenStr(node))

	var name string
	var args []ast.Expression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			name = p.identifier(node).Name

		case ruleexpression:
			exp, err := p.expression(node)
			if err != nil {
				return nil, err
			}
			args = append(args, exp)
		}

		node = node.next
	}

	funcCall := ast.NewFuncCall(name, args)
	return &funcCall, nil
}

func (p *Toy) comparative(node *node32) (ast.Expression, error) {
	infoLog.info("comparative\n%s\n", p.tokenStr(node))

	var lhs, rhs ast.Expression
	var operator ast.Operator = -1

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleadditive:
			additive, err := p.additive(node)
			if err != nil {
				return nil, err
			}

			if lhs == nil {
				lhs = additive
			} else {
				rhs = additive
			}

		case rulecomparativeOperator:
			op, err := p.comparativeOperator(node)
			if err != nil {
				return nil, err
			}
			operator = op
		}

		node = node.next
	}

	if operator == -1 {
		return lhs, nil
	} else {
		return ast.NewBinary(operator, lhs, rhs), nil
	}
}

func (p *Toy) additive(node *node32) (ast.Expression, error) {
	infoLog.info("additive\n%s\n", p.tokenStr(node))

	var lhs, rhs ast.Expression
	var operator ast.Operator = -1

	node = node.up
	for node != nil {
		switch node.pegRule {
		case rulemultitive:
			multitive, err := p.multitive(node)
			if err != nil {
				return nil, err
			}

			if lhs == nil {
				lhs = multitive
			} else {
				rhs = multitive
			}

		case ruleadditiveOperator:
			op, err := p.additiveOperator(node)
			if err != nil {
				return nil, err
			}
			operator = op
		}

		node = node.next
	}

	if operator == -1 {
		return lhs, nil
	} else {
		return ast.NewBinary(operator, lhs, rhs), nil
	}
}

func (p *Toy) multitive(node *node32) (ast.Expression, error) {
	infoLog.info("multitive\n%s\n", p.tokenStr(node))

	var lhs, rhs ast.Expression
	var operator ast.Operator = -1

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleprimary:
			primary, err := p.primary(node)
			if err != nil {
				return nil, err
			}

			if lhs == nil {
				lhs = primary
			} else {
				rhs = primary
			}

		case rulemultitiveOperator:
			op, err := p.multitiveOperator(node)
			if err != nil {
				return nil, err
			}
			operator = op
		}

		node = node.next
	}

	if operator == -1 {
		return lhs, nil
	} else {
		return ast.NewBinary(operator, lhs, rhs), nil
	}
}

func (p *Toy) primary(node *node32) (ast.Expression, error) {
	infoLog.info("primary\n%s\n", p.tokenStr(node))

	node = node.up
	for node != nil {
		switch node.pegRule {
		case rulecomparative:
			return p.comparative(node)
		case ruleprintln:
			exp, err := p.println(node)
			return *exp, err
		case rulefunctionCall:
			exp, err := p.funcCall(node)
			return *exp, err

		case ruleidentifier:
			exp := p.identifier(node)
			return *exp, nil

		case ruleinteger:
			exp, err := p.integer(node)
			return *exp, err
		}

		node = node.next
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) comparativeOperator(node *node32) (ast.Operator, error) {
	infoLog.info("comparativeOperator\n%s\n", p.tokenStr(node))

	switch p.tokenStr(node) {
	case "<":
		return ast.LESS_THAN, nil
	case ">":
		return ast.GREATER_THAN, nil
	case "<=":
		return ast.LESS_OR_EQUAL, nil
	case ">=":
		return ast.GREATER_OR_EQUAL, nil
	case "==":
		return ast.EQUAL, nil
	case "!=":
		return ast.NOT_EQUAL, nil
	}

	return -1, fmt.Errorf("not reach here")
}

func (p *Toy) additiveOperator(node *node32) (ast.Operator, error) {
	infoLog.info("comparativeOperator\n%s\n", p.tokenStr(node))

	switch p.tokenStr(node) {
	case "+":
		return ast.ADD, nil
	case "-":
		return ast.SUBTRACT, nil
	}

	return -1, fmt.Errorf("not reach here")
}

func (p *Toy) multitiveOperator(node *node32) (ast.Operator, error) {
	infoLog.info("multitiveOperator\n%s\n", p.tokenStr(node))

	switch p.tokenStr(node) {
	case "*":
		return ast.MULTIPLY, nil
	case "/":
		return ast.DIVIDE, nil
	}

	return -1, fmt.Errorf("not reach here")
}

func (p *Toy) identifier(node *node32) *ast.Identifier {
	infoLog.info("identifier\n%s\n", p.tokenStr(node))

	s := p.tokenStr(node)
	identifier := ast.NewIdentifier(s)
	return &identifier
}

func (p *Toy) integer(node *node32) (*ast.IntegerLiteral, error) {
	infoLog.info("integer\n%s\n", p.tokenStr(node))

	n, err := p.tokenInt(node)
	if err != nil {
		return nil, err
	}

	integer := ast.NewInteger(n)
	return &integer, nil
}

func (p *Toy) token(node *node32) []rune {
	return p.buffer[node.begin:node.end]
}

func (p *Toy) tokenStr(node *node32) string {
	return string(p.token(node))
}

func (p *Toy) tokenInt(node *node32) (int, error) {
	return strconv.Atoi(p.tokenStr(node))
}
