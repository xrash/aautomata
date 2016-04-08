package aautomata

import (
	"fmt"
	"strings"
)

type debug bool

func (d debug) print(i ...interface{}) {
	if d {
		fmt.Print(i...)
	}
}

func (d debug) println(i ...interface{}) {
	if d {
		fmt.Println(i...)
	}
}

func (d debug) printf(s string, i ...interface{}) {
	if d {
		fmt.Printf(s, i...)
	}
}

type AdaptiveAutomata struct {
	scanner     *scanner
	transitions *transitionCollection
	state       string
	final       map[string]bool
}

func NewAdaptiveAutomata() *AdaptiveAutomata {
	return &AdaptiveAutomata{
		newScanner(),
		newTransitionCollection(),
		"",
		nil,
	}
}

func (aa *AdaptiveAutomata) AddTransition(from, input, to string, b func(*AdaptiveAutomata), a func(*AdaptiveAutomata)) {
	aa.transitions.Add(newTransition(from, input, to, b, a))
}

func (aa *AdaptiveAutomata) RemoveTransition(from, input string) {
	aa.transitions.Remove(from, input)
}

func (aa *AdaptiveAutomata) SetState(s string) {
	aa.state = s
}

func (aa *AdaptiveAutomata) SetFinalStates(s ...string) {
	aa.final = make(map[string]bool)

	for _, state := range s {
		aa.final[state] = true
	}
}

func (aa *AdaptiveAutomata) Run(input string, d ...bool) (string, bool) {
	var debug debug

	if len(d) > 0 && d[0] {
		debug = true
	}

	debug.println("testing input ", input)

	aa.scanner.Init(strings.NewReader(input))

	var trans *transition

	for aa.scanner.Scan() {
		debug.printf("%v\t%v\t", aa.state, aa.scanner.Text())

		state := aa.state
		input := aa.scanner.Text()

		trans = aa.transitions.Find(state, input)
		if trans == nil {
			debug.println("transition not found")
			debug.println("")
			return state, false
		}

		trans.ExecBefore(aa)

		trans = aa.transitions.Find(state, input)
		if trans == nil {
			debug.println("transition not found")
			debug.println("")
			return state, false
		}

		trans.ExecAfter(aa)

		aa.state = trans.to

		debug.printf("%v\n", trans.to)
	}

	accept := aa.final[aa.state]

	debug.println(aa.state, accept)
	debug.println("")

	return aa.state, accept
}

func (aa *AdaptiveAutomata) Push(s string) {
	aa.scanner.Push(s)
}

func (aa *AdaptiveAutomata) Print() {
	aa.transitions.Print()
}
