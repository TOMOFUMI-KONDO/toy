package ast

type Expression interface {
	express()
}

type BinaryExpression struct {
	Operator Operator
	Lhs      Expression
	Rhs      Expression
}

func (BinaryExpression) express() {}

type IntegerLiteral struct {
	Value int
}

func (IntegerLiteral) express() {}

type Assignment struct {
	Name       string
	Expression Expression
}

func (Assignment) express() {}

type Identifier struct {
	Name string
}

func (Identifier) express() {}
