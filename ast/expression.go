package ast

type Expression interface {
	expression()
}

type IntegerLiteral struct {
	Value int
}

func (IntegerLiteral) expression() {}

type BinaryExpression struct {
	Operator Operator
	Lhs      Expression
	Rhs      Expression
}

func (BinaryExpression) expression() {}

type Assignment struct {
	Name       string
	Expression Expression
}

func (Assignment) expression() {}

type Identifier struct {
	Name string
}

func (Identifier) expression() {}

type BlockExpression struct {
	Expressions []Expression
}

func (BlockExpression) expression() {}

type WhileExpression struct {
	Condition Expression
	Body      Expression
}

func (WhileExpression) expression() {}

type IfExpression struct {
	Condition  Expression
	ThenClause Expression
	ElseClause Expression
}

func (IfExpression) expression() {}

type FunctionCall struct {
	Name string
	Args []Expression
}

func (FunctionCall) expression() {}
