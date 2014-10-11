package tests

import (
	"github.com/xrash/aautomata"
	"strconv"
	"testing"
)

func TestThird(t *testing.T) {
	creator := func() *aautomata.AdaptiveAutomata {

		// Ignore blanks
		// Extract integers
		// Extract special characters
		first := func(aa *aautomata.AdaptiveAutomata) {
			aa.AddTransition("1", " ", "1", nil, nil)

			for _, s := range(D) {
				aa.AddTransition("1", s, "1a", nil, nil)
				aa.AddTransition("1a", s, "1a", nil, nil)
			}

			for _, s := range(LUS) {
				aa.AddTransition("1a", s, "1b", nil, nil)
			}

			for _, s := range(S) {
				aa.AddTransition("1", s, "z", nil, nil)
			}

			for _, s := range(L) {
				func(s string) {
					t := func(aa *aautomata.AdaptiveAutomata) {
						aa.Push(s)
					}

					aa.AddTransition("1", s, "2", nil, t)
				}(s)
			}
		}

		// Extract token 'int'
		// Extract token 'real'
		second := func(aa *aautomata.AdaptiveAutomata) {
			for _, s := range(L) {
				if (s == "i") || (s == "r") {
					continue
				}

				aa.AddTransition("2", s, "3", nil, nil)
			}

			aa.AddTransition("2", "i", "2b", nil, nil)
			aa.AddTransition("2b", "n", "2e", nil, nil)
			aa.AddTransition("2e", "t", "2h", nil, nil)
			for _, s := range(S) {
				aa.AddTransition("2h", s, "2k", nil, nil)
			}

			for _, s := range(A) {
				if s == "n" {
					continue
				}

				func(s string) {
					t := func(aa *aautomata.AdaptiveAutomata) {
						aa.Push(s)
						aa.Push("i")
					}
					aa.AddTransition("2b", s, "3", nil, t)
				}(s)
			}

			for _, s := range(A) {
				if s == "t" {
					continue
				}

				func(s string) {
					t := func(aa *aautomata.AdaptiveAutomata) {
						aa.Push(s)
						aa.Push("n")
						aa.Push("i")
					}
					aa.AddTransition("2e", s, "3", nil, t)
				}(s)
			}

			for _, s := range(LUD) {
				func(s string) {
					t := func(aa *aautomata.AdaptiveAutomata) {
						aa.Push(s)
						aa.Push("t")
						aa.Push("n")
						aa.Push("i")
					}
					aa.AddTransition("2h", s, "3", nil, t)
				}(s)
			}
		}

		third := func(aa *aautomata.AdaptiveAutomata) {
			J := 0

			var D func(*aautomata.AdaptiveAutomata, string)
			var B func(*aautomata.AdaptiveAutomata, string, string)

			D = func(aa *aautomata.AdaptiveAutomata, i string) {
				for _, s := range(S) {
					func(s string) {
						aa.RemoveTransition(i, s)
						aa.AddTransition(i, s, "9", nil, nil)
					}(s)
				}
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

				for _, s := range(S) {
					func(s string) {
						aa.AddTransition(j, s, "8", nil, d)
					}(s)
				}

				J++
			}

			for _, s := range(L) {
				func(s string) {
					b := func(aa *aautomata.AdaptiveAutomata) {
						B(aa, "3", s)
					}

					aa.AddTransition("3", s, "3a", b, nil)
				}(s)
			}

			for _, s := range(S) {
				func(s string) {
					d := func(aa *aautomata.AdaptiveAutomata) {
						D(aa, "3a")
					}

					aa.AddTransition("3a", s, "8", nil, d)
				}(s)
			}
		}

		aa := aautomata.NewAdaptiveAutomata()
		first(aa)
		second(aa)
		third(aa)

		aa.SetFinalStates("8", "9", "z", "1b", "2k")

		return aa
	}

	cases := [][]testcase{
		{{"baa}", "1", "8", true}},
		{{"baa{", "1", "8", true}},
		{
			{"baa|", "1", "8", true},
			{"baa|", "1", "9", true},
		},
		{{"int-", "1", "2k", true}},
		{{"inv-", "1", "8", true}},
		{{"&", "1", "z", true}},
		{{":", "1", "z", true}},
		{
			{"&", "1", "z", true},
			{":", "1", "z", true},
		},
		{
			{"123|", "1", "1b", true},
			{"123|", "1", "1b", true},
			{"baa|", "1", "8", true},
			{"1a|", "1", "1b", false}, // identifier starting with a number
			{"|", "1", "z", true},
		},
	}

	testcases(t, creator, cases)
}
