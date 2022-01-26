package ast

type Expression interface {
	expression()
}

type IntegerLiteral struct {
	Value int
}

func (IntegerLiteral) expression() {}

func NewInteger(value int) IntegerLiteral {
	return IntegerLiteral{Value: value}
}

type BinaryExpression struct {
	Operator Operator
	Lhs      Expression
	Rhs      Expression
}

func (BinaryExpression) expression() {}

func NewBinary(op Operator, lhs, rhs Expression) BinaryExpression {
	return BinaryExpression{
		Operator: op,
		Lhs:      lhs,
		Rhs:      rhs,
	}
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

type Assignment struct {
	Name       string
	Expression Expression
}

func (Assignment) expression() {}

func NewAssignment(name string, exp Expression) Assignment {
	return Assignment{
		Name:       name,
		Expression: exp,
	}
}

type Identifier struct {
	Name string
}

func (Identifier) expression() {}

func NewIdentifier(name string) Identifier {
	return Identifier{Name: name}
}

type BlockExpression struct {
	Expressions []Expression
}

func (BlockExpression) expression() {}

func NewBlock(exps []Expression) BlockExpression {
	return BlockExpression{Expressions: exps}
}

type WhileExpression struct {
	Condition Expression
	Body      BlockExpression
}

func (WhileExpression) expression() {}

func NewWhile(cond Expression, body BlockExpression) WhileExpression {
	return WhileExpression{
		Condition: cond,
		Body:      body,
	}
}

type IfExpression struct {
	Condition  Expression
	ThenClause BlockExpression
	ElseClause BlockExpression
}

func (IfExpression) expression() {}

func NewIf(cond Expression, thenClause BlockExpression, elseClause BlockExpression) IfExpression {
	return IfExpression{
		Condition:  cond,
		ThenClause: thenClause,
		ElseClause: elseClause,
	}
}

func NewIfWithoutElse(cond Expression, thenClause BlockExpression) IfExpression {
	return IfExpression{
		Condition:  cond,
		ThenClause: thenClause,
	}
}

type Println struct {
	Arg Expression
}

func (Println) expression() {}

func NewPrintln(arg Expression) Println {
	return Println{
		Arg: arg,
	}
}

type FunctionCall struct {
	Name string
	Args []Expression
}

func (FunctionCall) expression() {}

func NewFuncCall(name string, args []Expression) FunctionCall {
	return FunctionCall{
		Name: name,
		Args: args,
	}
}
