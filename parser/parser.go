package parser

import (
	"fmt"
	"strconv"

	"github.com/TOMOFUMI-KONDO/toy/ast"
)

func (p *Toy) ConvertAst() error {
	return p.program(p.AST())
}

func (p *Toy) program(node *node32) error {
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
	node = node.up
	for node != nil {
		switch node.pegRule {
		case rulefunctionDefinition:
			funcDef, err := p.functionDefinition(node)
			if err != nil {
				return nil, err
			}
			return funcDef, nil

		case ruleglobalVariableDefinition:
			// TODO
		}
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) functionDefinition(node *node32) (*ast.FunctionDefinition, error) {
	var name string
	var args []string
	var body *ast.BlockExpression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			arg := string(p.token(node))
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

func (p *Toy) expression(node *node32) (ast.Expression, error) {
	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleifExpression:
			return p.ifExp(node)
		}
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) block(node *node32) (*ast.BlockExpression, error) {
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

func (p *Toy) ifExp(node *node32) (*ast.IfExpression, error) {
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
			if thenClause.Expressions == nil {
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

	ifExp := ast.NewIf(cond, *thenClause, *elseClause)
	return &ifExp, nil
}

func (p *Toy) funcCall(node *node32) (*ast.FunctionCall, error) {
	var name string
	var args []ast.Expression

	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			identifier, err := p.identifier(node)
			if err != nil {
				return nil, err
			}
			name = identifier.Name

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
	var lhs, rhs ast.Expression
	var operator ast.Operator = -1

	node = node.up
	for node != nil {
		switch node.pegRule {
		case rulemultitive:
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
	node = node.up
	for node != nil {
		switch node.pegRule {
		case rulecomparative:
			// TODO
		case ruleprintln:
			// TODO
		case rulefunctionCall:
			return p.funcCall(node)

		case ruleidentifier:
			return p.identifier(node)

		case ruleinteger:
			return p.integer(node)
		}

		node = node.next
	}

	return nil, fmt.Errorf("not reach here")
}

func (p *Toy) comparativeOperator(node *node32) (ast.Operator, error) {
	node = node.up
	for node != nil {
		switch string(p.token(node)) {
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

		node = node.next
	}

	return -1, fmt.Errorf("not reach here")
}

func (p *Toy) additiveOperator(node *node32) (ast.Operator, error) {
	node = node.up
	for node != nil {
		switch string(p.token(node)) {
		case "+":
			return ast.ADD, nil
		case "-":
			return ast.SUBTRACT, nil
		}

		node = node.next
	}

	return -1, fmt.Errorf("not reach here")
}

func (p *Toy) multitiveOperator(node *node32) (ast.Operator, error) {
	node = node.up
	for node != nil {
		switch string(p.token(node)) {
		case "*":
			return ast.MULTIPLY, nil
		case "/":
			return ast.DIVIDE, nil
		}

		node = node.next
	}

	return -1, fmt.Errorf("not reach here")
}

func (p *Toy) identifier(node *node32) (ast.Identifier, error) {
	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleidentifier:
			s := string(p.token(node))
			return ast.NewIdentifier(s), nil
		}

		node = node.next
	}

	return ast.Identifier{}, fmt.Errorf("not reach here")
}

func (p *Toy) integer(node *node32) (ast.IntegerLiteral, error) {
	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleinteger:
			n, err := runesToInt(p.token(node))
			if err != nil {
				return ast.IntegerLiteral{}, err
			}
			return ast.NewInteger(n), nil
		}

		node = node.next
	}

	return ast.IntegerLiteral{}, fmt.Errorf("not reach here")
}

func (p *Toy) token(node *node32) []rune {
	return p.buffer[node.begin:node.end]
}

func runesToInt(r []rune) (int, error) {
	return strconv.Atoi(string(r))
}
