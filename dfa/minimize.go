package dfa

import (
	"bmstu/cc2024/lab1/nfa"
	// "fmt"
)

func MinimizeDFA(_nfa *nfa.NFA) (_dfa *DFA, _r_fa *nfa.NFA, _dr_fa *nfa.NFA, _rdr_fa *nfa.NFA) {

	r_fa := _nfa.ReverseNFA()

  _r_fa = nfa.NewNFA(r_fa.InitialState, r_fa.AcceptStates, r_fa.Rules)

	dr_temp := NewDFAFromNFA(r_fa)

	dr_fa := nfa.NewNFA(dr_temp.InitialState, dr_temp.AcceptStates, make(nfa.NFARulesMap))
  
	for key, st := range dr_temp.Rules {
		rule := nfa.NewNFARule(key.Src, key.Val)
		dr_fa.Rules.AddDst(rule, st)
	}

  _dr_fa = nfa.NewNFA(dr_fa.InitialState, dr_fa.AcceptStates, dr_fa.Rules)

	rdr_fa := dr_fa.ReverseNFA()
  _rdr_fa = nfa.NewNFA(rdr_fa.InitialState, rdr_fa.AcceptStates, rdr_fa.Rules)
  
	_dfa = NewDFAFromNFA(rdr_fa)

	return
}
