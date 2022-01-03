package ast

type Ast struct{}

func NewAst() Ast {
	return Ast{}
}

func (Ast) Integer(value int) IntegerLiteral {
	return IntegerLiteral{Value: value}
}

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

func (Ast) LessThan(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: LESS_THAN,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) LessOrEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: LESS_OR_EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) GreaterThan(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: GREATER_THAN,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) GreaterOrEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: GREATER_OR_EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) Equal(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func (Ast) NotEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: NOT_EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
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

func (Ast) Block(exps []Expression) BlockExpression {
	return BlockExpression{Expressions: exps}
}

func (Ast) While(cond Expression, body Expression) WhileExpression {
	return WhileExpression{
		Condition: cond,
		Body:      body,
	}
}

func (Ast) If(cond Expression, thenClause Expression, elseClause Expression) IfExpression {
	return IfExpression{
		Condition:  cond,
		ThenClause: thenClause,
		ElseClause: elseClause,
	}
}

func (Ast) Println(arg Expression) Println {
	return Println{
		Arg: arg,
	}
}

func (Ast) Call(name string, args []Expression) FunctionCall {
	return FunctionCall{
		Name: name,
		Args: args,
	}
}

func (Ast) DefineFunction(name string, args []string, body Expression) FunctionDefinition {
	return FunctionDefinition{
		Name: name,
		Args: args,
		Body: body,
	}
}

func (Ast) GlobalAssignment(name string, exp Expression) GlobalVariableDefinition {
	return GlobalVariableDefinition{
		Name:       name,
		Expression: exp,
	}
}

func (Ast) Program(topLevels []TopLevel) Program {
	return Program{Definitions: topLevels}
}
