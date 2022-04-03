package ast

type Operator int

const (
	Add Operator = iota
	Subtract
	Multiply
	Divide
	LessThan
	LessOrEqual
	GreaterThan
	GreaterOrEqual
	Equal
	NotEqual
)

func (o Operator) Name() string {
	return [...]string{
		"Add",
		"Subtract",
		"Multiply",
		"Divide",
		"LessThan",
		"LessOrEqual",
		"GreaterThan",
		"GreaterOrEqual",
		"Equal",
		"NotEqual",
	}[o]
}
