package ast

func NewInteger(value int) IntegerLiteral {
	return IntegerLiteral{Value: value}
}

func NewAdd(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: ADD,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewSubtract(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: SUBTRACT,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewMultiply(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: MULTIPLY,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewDivide(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: DIVIDE,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewLessThan(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: LESS_THAN,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewLessOrEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: LESS_OR_EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewGreaterThan(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: GREATER_THAN,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewGreaterOrEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: GREATER_OR_EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewNotEqual(lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: NOT_EQUAL,
		Lhs:      lhs,
		Rhs:      rhs,
	}
}

func NewIdentifier(name string) Identifier {
	return Identifier{Name: name}
}

func NewAssignment(name string, exp Expression) Assignment {
	return Assignment{
		Name:       name,
		Expression: exp,
	}
}

func NewBlock(exps []Expression) BlockExpression {
	return BlockExpression{Expressions: exps}
}

func NewWhile(cond Expression, body BlockExpression) WhileExpression {
	return WhileExpression{
		Condition: cond,
		Body:      body,
	}
}

func NewIf(cond Expression, thenClause BlockExpression, elseClause BlockExpression) IfExpression {
	return IfExpression{
		Condition:  cond,
		ThenClause: thenClause,
		ElseClause: elseClause,
	}
}

func NewPrintln(arg Expression) Println {
	return Println{
		Arg: arg,
	}
}

func NewFuncCall(name string, args []Expression) FunctionCall {
	return FunctionCall{
		Name: name,
		Args: args,
	}
}

func NewFuncDef(name string, args []string, body BlockExpression) FunctionDefinition {
	return FunctionDefinition{
		Name: name,
		Args: args,
		Body: body,
	}
}

func NewGlobalVarDef(name string, exp Expression) GlobalVariableDefinition {
	return GlobalVariableDefinition{
		Name:       name,
		Expression: exp,
	}
}

func NewProgram(topLevels []TopLevel) Program {
	return Program{Definitions: topLevels}
}
