package ast

type TopLevel interface {
	topLevel()
}

type FunctionDefinition struct {
	Name string
	Args []string
	Body Expression
}

func (FunctionDefinition) topLevel() {}

type GlobalVariableDefinition struct {
	Name       string
	Expression Expression
}

func (GlobalVariableDefinition) topLevel() {}

type Program struct {
	Definitions []TopLevel
}
