// Package dfa implements Deterministic Finite Automaton(DFA).
package dfa

import (
	"bmstu/cc2024/lab1/fsm"
)

type DFA struct {
	InitialState fsm.State
	AcceptStates fsm.StateSet
	Rules        DFARulesMap
}

func (this *DFA) GetAllStates() fsm.StateSet {
	ret := fsm.StateSet{}

	for key, dst := range this.Rules {
		ret.Add(key.Src)
    ret.Add(dst)
	}
	return ret
}


func NewDFA(init fsm.State, accepts fsm.StateSet, rules DFARulesMap) *DFA {
	return &DFA{
		InitialState: init,
		AcceptStates: accepts,
		Rules:        rules,
	}
}

func (dfa *DFA) Minimize() {
	states := fsm.NewStateSet()
  states.Add(dfa.InitialState)
	for _, v := range dfa.Rules {
		states.Add(v)
	}
	n := states.Len()

	eqMap := map[fsm.State]fsm.State{}
	for i := 0; i < n; i++ {
		q1 := fsm.NewState(i)
		for j := i + 1; j < n; j++ {
			q2 := fsm.NewState(j)
			if !dfa.isEquivalent(q1, q2) {
				continue
			}
			if _, ok := eqMap[q2]; ok {
				continue
			}
			eqMap[q2] = q1
      // fmt.Println("Merging ", q1.String(), "and", q2.String())
			dfa.mergeState(q1, q2)
		}
	}
  // fmt.Println("eqMap: ", eqMap)
}

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
