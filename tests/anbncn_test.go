package tests

import (
	"github.com/xrash/aautomata"
	"strconv"
	"testing"
)

func TestAnbncn(t *testing.T) {
	creator := func() *aautomata.AdaptiveAutomata {
		I := 0
		J := 0
		K := 0

		var F func(*aautomata.AdaptiveAutomata, string, string, string)
		var G func(*aautomata.AdaptiveAutomata, string)

		F = func(aa *aautomata.AdaptiveAutomata, x, y, z string) {
			i := "i" + strconv.Itoa(I)
			j := "j" + strconv.Itoa(J)
			k := "k" + strconv.Itoa(K)

			f := func(aa *aautomata.AdaptiveAutomata) {
				F(aa, i, j, k)
			}

			g := func(aa *aautomata.AdaptiveAutomata) {
				G(aa, k)
			}

			aa.AddTransition(x, "a", i, nil, f)
			aa.AddTransition(i, "b", j, nil, g)
			aa.AddTransition(j, "b", y, nil, nil)
			aa.AddTransition(k, "c", z, nil, nil)

			I++
			J++
			K++
		}

		G = func(aa *aautomata.AdaptiveAutomata, x string) {
			aa.AddTransition("B", "c", x, nil, nil)
		}

		aa := aautomata.NewAdaptiveAutomata()

		f := func(aa *aautomata.AdaptiveAutomata) {
			F(aa, "A", "B", "C")
		}

		g := func(aa *aautomata.AdaptiveAutomata) {
			G(aa, "C")
		}

		aa.AddTransition("S", "a", "A", nil, f)
		aa.AddTransition("A", "b", "B", nil, g)
		aa.SetFinalStates("C")

		return aa
	}

	cases := [][]testcase{
		{{"aaabbbccc", "S", "C", true}},
		{{"abc", "S", "C", true}},
		{{"aaa", "S", "i1", false}},
	}

	testcases(t, creator, cases)
}
