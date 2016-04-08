package tests

import (
	"github.com/xrash/aautomata"
	"strconv"
	"testing"
)

func Test0n1n(t *testing.T) {
	creator := func() *aautomata.AdaptiveAutomata {
		J := 0
		M := 0

		var B func(*aautomata.AdaptiveAutomata, string, string)

		B = func(aa *aautomata.AdaptiveAutomata, a, b string) {
			j := "j" + strconv.Itoa(J)
			m := "m" + strconv.Itoa(M)

			_B := func(aa *aautomata.AdaptiveAutomata) {
				B(aa, j, m)
			}

			aa.AddTransition(a, "0", j, nil, _B)
			aa.AddTransition(j, "1", m, nil, nil)
			aa.AddTransition(m, "1", b, nil, nil)

			J++
			M++
		}

		aa := aautomata.NewAdaptiveAutomata()

		b := func(aa *aautomata.AdaptiveAutomata) {
			B(aa, "K", "F")
		}

		aa.AddTransition("S", "0", "K", nil, b)
		aa.AddTransition("K", "1", "F", nil, nil)
		aa.SetFinalStates("F")

		return aa
	}

	cases := [][]testcase{
		{{"0011", "S", "F", true}},
		{{"01", "S", "F", true}},
		{{"000111", "S", "F", true}},
		{{"001111", "S", "F", false}},
		{{"00011", "S", "m0", false}},
	}

	testcases(t, creator, cases)
}
