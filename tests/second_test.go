package tests

import (
	"github.com/xrash/aautomata"
	"strconv"
	"testing"
)

func TestSecond(t *testing.T) {
	creator := func() *aautomata.AdaptiveAutomata {
		J := 0

		var B func(*aautomata.AdaptiveAutomata, string, string)
		var D func(*aautomata.AdaptiveAutomata, string)

		D = func(aa *aautomata.AdaptiveAutomata, i string) {
			aa.RemoveTransition(i, "|")
			aa.AddTransition(i, "|", "9", nil, nil)
		}

		B = func(aa *aautomata.AdaptiveAutomata, i, o string) {
			j := "j" + strconv.Itoa(J)

			d := func(aa *aautomata.AdaptiveAutomata) {
				D(aa, j)
			}

			aa.RemoveTransition(i, o)
			aa.AddTransition(i, o, j, nil, nil)

			for _, s := range(LUD) {
				func(s string) {
					b := func(aa *aautomata.AdaptiveAutomata) {
						B(aa, j, s)
					}
					aa.AddTransition(j, s, "3a", b, nil)
				}(s)
			}

			aa.AddTransition(j, "|", "8", nil, d)

			J++
		}

		aa := aautomata.NewAdaptiveAutomata()

		d := func(aa *aautomata.AdaptiveAutomata) {
			D(aa, "3a")
		}

		for _, s := range L {
			func(s string) {
				b := func(aa *aautomata.AdaptiveAutomata) {
					B(aa, "3", s)
				}

				aa.AddTransition("3", s, "3a", b, nil)
			}(s)
		}

		for _, s := range LUD {
			aa.AddTransition("3a", s, "3a", nil, nil)
		}

		aa.AddTransition("3a", "|", "8", nil, d)

		aa.SetState("3")
		aa.SetFinalStates("8", "9")

		return aa
	}

	cases := [][]testcase{
		{{"baa|", "3", "8", true}},
		{
			{"baa|", "3", "8", true},
			{"baa|", "3", "9", true},
			{"caa|", "3", "8", true},
			{"ccc|", "3", "8", true},
			{"caa|", "3", "9", true},
		},
		{{"aa2aac|", "3", "8", true}},
		{
			{"11|", "3", "3", false},
			{"11|", "3", "3", false},
			{"1a|", "3", "3", false},
			{"a1|", "3", "8", true},
			{"a1|", "3", "9", true},
		},
	}

	testcases(t, creator, cases)
}
