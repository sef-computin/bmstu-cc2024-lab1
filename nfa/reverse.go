package nfa

import "bmstu/cc2024/lab1/fsm"

func (nfa *NFA) ReverseNFA() *NFA {
	ret := NewNFA(nfa.InitialState, *fsm.NewStateSet(), nil)

  ret.AcceptStates.Add(nfa.InitialState)
	// ret.AcceptStates[nfa.InitialState] = true
	for _, key := range nfa.AcceptStates.GetAll() {
		ret.InitialState = key
		break
	}

	ret.Rules = make(NFARulesMap)

	for rule, dst := range nfa.Rules {
		for _, st := range dst.GetAll() {
			ret.Rules.AddDst(NewNFARule(st, rule.Val), rule.Src)
		}
	}

	return ret
}
