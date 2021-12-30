package ast

type Operator int

const (
	ADD Operator = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
	LESS_THAN
	LESS_OR_EQUAL
	GREATER_THAN
	GREATER_OR_EQUAL
	EQUAL
	NOT_EQUAL
)

func (o Operator) Name() string {
	return [...]string{
		"ADD",
		"SUBTRACT",
		"MULTIPLY",
		"DIVIDE",
		"LESS_THAN",
		"LESS_OR_EQUAL",
		"GREATER_THAN",
		"GREATER_OR_EQUAL",
		"EQUAL",
		"NOT_EQUAL",
	}[o]
}
