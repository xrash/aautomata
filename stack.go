package aautomata

type stack []string

func newStack() stack {
	return stack{}
}

func (s *stack) push(i string) {
	*s = append(*s, i)
}

func (s *stack) pop() string {
	l := len(*s) - 1
	st := (*s)[l]
	*s = (*s)[:l]
	return st
}

func (s *stack) empty() bool {
	return len(*s) == 0
}
