package ast

type Environment struct {
	Bindings map[string]int
	next     *Environment
}

func NewEnvironment(next *Environment) *Environment {
	return &Environment{
		Bindings: map[string]int{},
		next:     next,
	}
}

func (e *Environment) FindBinding(name string) map[string]int {
	if _, ok := e.Bindings[name]; ok {
		return e.Bindings
	}

	if e.next != nil {
		return e.next.FindBinding(name)
	}

	return nil
}
