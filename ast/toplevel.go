package ast

type TopLevel interface {
	topLevel()
}

type FunctionDefinition struct {
	Name string
	Args []string
	Body BlockExpression
}

func (FunctionDefinition) topLevel() {}

type GlobalVariableDefinition struct {
	Name string
	Expression
}

func (GlobalVariableDefinition) topLevel() {}

type Program struct {
	Definitions []TopLevel
}
