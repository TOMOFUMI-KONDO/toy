package ast

type Operator int

const (
	ADD Operator = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
)

func (o Operator) Name() string {
	return [...]string{"ADD", "SUBTRACT", "MULTIPLY", "DIVIDE"}[o]
}
