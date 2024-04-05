package dfa

import (
	"bmstu/cc2024/lab1/fsm"
	"bmstu/cc2024/lab1/nfa"
)

func NewDFAFromNFA(nfa *nfa.NFA) *DFA {
	nfa.ToWithoutEpsilon()
	I, F, Delta := SubsetConstruction(nfa)
	return NewDFA(I, F, Delta)
}

func SubsetConstruction(_nfa *nfa.NFA) (dI fsm.State, dF fsm.StateSet, dRules DFARulesMap) {
	I := _nfa.InitialState
	F := _nfa.AcceptStates
	Rules := _nfa.Rules

	dI = fsm.NewState(0)
	dF = *fsm.NewStateSet()
	dRules = DFARulesMap{}

	dStates := DFAStatesMap{}
	start := fsm.NewStateSet()
	start.Add(I)
	dStates[start] = fsm.NewState(0)

	queue := newQueue()
	queue.Push(start)
	for len(*queue) != 0 {
		dstate := queue.Pop()

		if F.Intersect(*dstate).Len() > 0 {
			dF.Add(dStates.getState(*dstate))
		}

		Sigma := _nfa.GetAllSymbols()
		for c := range Sigma {
			dnext := fsm.NewStateSet()
			for _, q := range dstate.GetAll() {
				d, ok := Rules[nfa.NewNFARule(q, c)]
				if ok {
					dnext = dnext.Union(*d)
				}
			}

			if dnext.Len() == 0 {
				continue
			}

			if !dStates.haveKey(*dnext) {
				queue.Push(dnext)
				dStates[dnext] = fsm.NewState(len(dStates))
			}

			for k := range dStates {
				if k.Equals(*dnext) {
					dnext = k // Swap to avoid problems with pointers
				}
			}
			dRules[NewDFARule(dStates[dstate], c)] = dStates[dnext]
		}
	}

	return
}

type DFAStatesMap map[*fsm.StateSet]fsm.State

// getState returns the state associated with key.
// If there is no corresponding state, it returns empty state struct.
func (dm DFAStatesMap) getState(key fsm.StateSet) fsm.State {
	if dm.haveKey(key) {
		for k := range dm {
			if k.Equals(key) {
				return dm[k]
			}
		}
	}
	return fsm.State{}
}

// haveKey returns whether DFAStatesMap has the set given as the argument "key".
// If it has, returns true.
func (dm DFAStatesMap) haveKey(key fsm.StateSet) bool {
	for k := range dm {
		if k.Equals(key) {
			return true
		}
	}
	return false
}

type queue []*fsm.StateSet

func (q *queue) Push(set *fsm.StateSet) {
	*q = append(*q, set)
}

func (q *queue) Pop() *fsm.StateSet {
	if len(*q) == 0 {
		return nil
	}
	ret := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return ret
}

func newQueue() *queue {
	return &queue{}
}
