package nfa

import "bmstu/cc2024/lab1/fsm"

func (nfa *NFA) ReverseNFA() *NFA {
	ret := NewNFA(nfa.InitialState, *fsm.NewStateSet(), nil)

	ret.AcceptStates.Add(nfa.InitialState)

	// ret.AcceptStates[nfa.InitialState] = true
	ret.InitialState = nfa.AcceptStates.GetAll()[0]

	ret.Rules = make(NFARulesMap)

	for rule, dst := range nfa.Rules {
		for _, st := range dst.GetAll() {
			ret.Rules.AddDst(NewNFARule(st, rule.Val), rule.Src)
		}
	}

	rule := NewNFARule(ret.InitialState, 'E')
	for idx, st := range nfa.AcceptStates.GetAll() {
		if idx < 1 {
			continue
		}
		ret.Rules.AddDst(rule, st)
	}

	return ret
}
