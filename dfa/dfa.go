// Package dfa implements Deterministic Finite Automaton(DFA).
package dfa

import "bmstu/cc2024/lab1/fsm"

type DFA struct {
	InitialState fsm.State
	AcceptStates fsm.StateSet
	Rules        DFARulesMap
}

func NewDFA(init fsm.State, accepts fsm.StateSet, rules DFARulesMap) *DFA {
	return &DFA{
		InitialState: init,
		AcceptStates: accepts,
		Rules:        rules,
	}
}

// // Minimize minimizes the DFA.
// func (dfa *DFA) Minimize() {
// 	states := .NewSet(dfa.I)
// 	for _, v := range dfa.Rules {
// 		states.Add(v)
// 	}
// 	n := states.N()
//
// 	eqMap := map[utils.State]utils.State{}
// 	for i := 0; i < n; i++ {
// 		q1 := utils.NewState(i)
// 		for j := i + 1; j < n; j++ {
// 			q2 := utils.NewState(j)
// 			if !dfa.isEquivalent(q1, q2) {
// 				continue
// 			}
// 			if _, ok := eqMap[q2]; ok {
// 				continue
// 			}
// 			eqMap[q2] = q1
// 			dfa.mergeState(q1, q2)
// 		}
// 	}
// }

func (dfa *DFA) replaceState(to, from fsm.State) {
	rules := dfa.Rules
	for arg, dst := range rules {
		if dst == from {
			rules[arg] = to
		}
	}
}

func (dfa *DFA) deleteState(q fsm.State) {
	rules := dfa.Rules
	for arg := range rules {
		if arg.Src == q {
			delete(rules, arg)
		}
	}
}

func (dfa *DFA) mergeState(to, from fsm.State) {
	dfa.replaceState(to, from)
	dfa.deleteState(from)
}

func (dfa *DFA) isEquivalent(q1, q2 fsm.State) bool {
	if !((dfa.AcceptStates.Contains(q1) && dfa.AcceptStates.Contains(q2)) ||
		(!dfa.AcceptStates.Contains(q1) && !dfa.AcceptStates.Contains(q2))) {
		return false
	}

	rules := dfa.Rules
	for k := range rules {
		if k.Src != q1 {
			continue
		}
		if rules[NewDFARule(q1, k.Val)] != rules[NewDFARule(q2, k.Val)] {
			return false
		}
	}
	return true
}

type Runtime struct {
	d   *DFA
	cur fsm.State
}

func (dfa *DFA) GetRuntime() *Runtime {
	return NewRuntime(dfa)
}

func NewRuntime(d *DFA) (r *Runtime) {
	r = &Runtime{
		d: d,
	}
	r.cur = d.InitialState
	return
}

func (r *Runtime) transit(c rune) bool {
	key := NewDFARule(r.cur, c)
	_, ok := r.d.Rules[key]
	if ok {
		r.cur = r.d.Rules[key]
		return true
	}
	return false
}

func (r *Runtime) isAccept() bool {
	accepts := r.d.AcceptStates
	if accepts.Contains(r.cur) {
		return true
	}
	return false
}

func (r *Runtime) Matching(str string) bool {
	r.cur = r.d.InitialState
	for _, c := range []rune(str) {
		if !r.transit(c) {
			return false
		}
	}
	return r.isAccept()
}
