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

func NewFuncDef(name string, args []string, body BlockExpression) FunctionDefinition {
	return FunctionDefinition{
		Name: name,
		Args: args,
		Body: body,
	}
}

type GlobalVariableDefinition struct {
	Name string
	Expression
}

func (GlobalVariableDefinition) topLevel() {}

func NewGlobalVarDef(name string, exp Expression) GlobalVariableDefinition {
	return GlobalVariableDefinition{
		Name:       name,
		Expression: exp,
	}
}

type Program struct {
	Definitions []TopLevel
}

func NewProgram(topLevels []TopLevel) Program {
	return Program{Definitions: topLevels}
}

func (p *Program) PushTopLevel(tp TopLevel) {
	p.Definitions = append(p.Definitions, tp)
}
