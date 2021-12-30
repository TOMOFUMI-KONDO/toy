package ast

type Ast struct{}

func (Ast) Add(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: ADD,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) Subtract(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: SUBTRACT,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) Multiply(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: MULTIPLY,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) Divide(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: DIVIDE,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) Integer(value int) IntegerLiteral {
	return IntegerLiteral{Value: value}
}

func (Ast) Identifier(name string) Identifier {
	return Identifier{Name: name}
}

func (Ast) Assignment(name string, exp Expression) Assignment {
	return Assignment{
		Name:       name,
		Expression: exp,
	}
}

func (Ast) BlockExpression(exps []Expression) BlockExpression {
	return BlockExpression{Expressions: exps}
}

func (Ast) WhileExpression(cond Expression, body Expression) WhileExpression {
	return WhileExpression{
		Condition: cond,
		Body:      body,
	}
}

func (Ast) IfExpression(cond Expression, thenClause Expression, elseClause Expression) IfExpression {
	return IfExpression{
		Condition:  cond,
		ThenClause: thenClause,
		ElseClause: elseClause,
	}
}
