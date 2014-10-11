package tests

import (
	"testing"
	"github.com/xrash/aautomata"
)

type testcase struct {
	input        string
	initialState string
	finalState   string
	accept       bool
}

func testcases(t *testing.T, creator func() *aautomata.AdaptiveAutomata, cases [][]testcase) {
	for _, batch := range cases {
		aa := creator()

		for _, c := range batch {
			aa.SetState(c.initialState)
			state, accept := aa.Run(c.input, true)

			if state != c.finalState || accept != c.accept {
				t.Fail()
			}
		}
	}
}
