package writer

import (
	"bmstu/cc2024/lab1/dfa"
	"bmstu/cc2024/lab1/drawer"
	"bmstu/cc2024/lab1/regex"
	"fmt"
)

func Process(_args ...string) {

	if len(_args) < 1 {
		return
	}
	_regexp := _args[0]

	fmt.Println(_regexp)

	_nfa, _ast := regex.ParseNFA(_regexp)

	fmt.Println("Tokens:")
	fmt.Println(_ast.SubtreeString())

	fmt.Println("\nInitial E-NFA")
	fmt.Println("Initial State: ", _nfa.InitialState)
	fmt.Println("Accept States", _nfa.AcceptStates)
	fmt.Println(_nfa.Rules.ToString())

	drawer.DrawNFA(_nfa, "E-nfa")

	// // fmt.Println("Deleting E")
	// // _nfa.ToWithoutEpsilon()
	// // fmt.Println("New NFA:")
	// // fmt.Println("InitialState: ", _nfa.InitialState)
	// // fmt.Println("AcceptStates: ", _nfa.AcceptStates)
	// // fmt.Println(_nfa.Rules.ToString())
	//

	// fmt.Println("DFA")
	// _dfa := dfa.NewDFAFromNFA(_nfa)
	// fmt.Println("InitialState: ", _dfa.InitialState)
	// fmt.Println("AcceptStates: ", _dfa.AcceptStates)
	// fmt.Println(_dfa.Rules.ToString())


	fmt.Println("Minimized DFA")
	min_dfa, r_fa, dr_fa, rdr_fa := dfa.MinimizeDFA(_nfa)
  drawer.DrawNFA(r_fa, "r_fa")
  drawer.DrawNFA(dr_fa, "dr_fa")
  drawer.DrawNFA(rdr_fa, "rdr_fa")
	drawer.DrawDFA(min_dfa, "min_dfa")

	fmt.Println("InitialState: ", min_dfa.InitialState)
	fmt.Println("AcceptStates: ", min_dfa.AcceptStates)
	fmt.Println(min_dfa.Rules.ToString())


	if len(_args) > 1 {
		_exprs := _args[1:]
		r := dfa.NewRuntime(min_dfa)
		for _, _exp := range _exprs {
			if r.Matching(_exp) {
				fmt.Printf("\"%s\" matches \"%s\"\n", _exp, _regexp)
			} else {
				fmt.Printf("\"%s\" does not match \"%s\"\n", _exp, _regexp)
			}
		}
	}

}
