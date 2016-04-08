package tests

import (
	"github.com/xrash/aautomata"
	"strconv"
	"testing"
)

func TestFirst(t *testing.T) {
	creator := func() *aautomata.AdaptiveAutomata {
		K := 0
		M := 0

		var A func(*aautomata.AdaptiveAutomata, string, string, string)

		A = func(aa *aautomata.AdaptiveAutomata, i, j, n string) {
			k := "k" + strconv.Itoa(K)
			m := "m" + strconv.Itoa(M)

			a := func(aa *aautomata.AdaptiveAutomata) {
				A(aa, k, m, i)
			}

			aa.AddTransition(k, "B", m, nil, nil)
			aa.AddTransition(m, ")", j, nil, nil)
			aa.AddTransition(i, "(", k, nil, a)

			K++
			M++
		}

		aa := aautomata.NewAdaptiveAutomata()

		a := func(aa *aautomata.AdaptiveAutomata) {
			A(aa, "2", "3", "1")
		}

		aa.AddTransition("1", "B", "4", nil, nil)
		aa.AddTransition("1", "(", "2", nil, a)
		aa.AddTransition("2", "B", "3", nil, nil)
		aa.AddTransition("3", ")", "4", nil, nil)

		aa.SetFinalStates("4")

		return aa
	}

	cases := [][]testcase{
		{{"((B))", "1", "4", true}},
		{{"((((B))))", "1", "4", true}},
		{{"B", "1", "4", true}},
		{{"((B)", "1", "3", false}},
	}

	testcases(t, creator, cases)
}
