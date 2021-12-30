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
