package aautomata

import (
	"fmt"
)

type transition struct {
	from   string
	input  string
	to     string
	before func(*AdaptiveAutomata)
	after  func(*AdaptiveAutomata)
}

func newTransition(from, input, to string, b func(*AdaptiveAutomata), a func(*AdaptiveAutomata)) *transition {
	return &transition{
		from,
		input,
		to,
		b,
		a,
	}
}

func (t *transition) ExecBefore(aa *AdaptiveAutomata) {
	if t.before != nil {
		t.before(aa)
	}
}

func (t *transition) ExecAfter(aa *AdaptiveAutomata) {
	if t.after != nil {
		t.after(aa)
	}
}

type transitionCollection struct {
	table map[string]map[string]*transition
}

func newTransitionCollection() *transitionCollection {
	return &transitionCollection{
		make(map[string]map[string]*transition),
	}
}

func (c *transitionCollection) Add(t *transition) {
	if _, exists := c.table[t.from]; !exists {
		c.table[t.from] = make(map[string]*transition)
	}

	c.table[t.from][t.input] = t
}

func (c *transitionCollection) Remove(from, input string) {
	if _, exists := c.table[from][input]; !exists {
		return
	}

	delete(c.table[from], input)

	if len(c.table[from]) == 0 {
		delete(c.table, from)
	}
}

func (c *transitionCollection) Find(state, input string) *transition {
	if trans, exists := c.table[state][input]; exists {
		return trans
	}

	return nil
}

func (c *transitionCollection) Print() {
	for _, from := range c.table {
		for _, trans := range from {
			fmt.Println(trans.from, trans.input, trans.to, trans.before, trans.after)
		}
	}
}
