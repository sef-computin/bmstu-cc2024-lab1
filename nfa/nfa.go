package nfa

import (
	"bmstu/cc2024/lab1/fsm"
	// "fmt"
)

type NFA struct {
	InitialState fsm.State
	AcceptStates fsm.StateSet
	Rules        NFARulesMap
}

func NewNFA(init fsm.State, accepts fsm.StateSet, rules NFARulesMap) *NFA {
	return &NFA{
		InitialState: init,
		AcceptStates: accepts,
		Rules:        rules,
	}
}

func (this *NFA) GetAllStates() fsm.StateSet {
	ret := fsm.StateSet{}

	for key, dst := range this.Rules {
		ret.Add(key.Src)
    for _, st := range dst.GetAll(){
      ret.Add(st)
    }
	}
	return ret
}

func (this *NFA) GetAllSymbols() map[rune]interface{} {
	ret := map[rune]interface{}{}

	for key := range this.Rules {
		ret[key.Val] = true
	}
	return ret
}

func (this *NFA) CalcDst(q fsm.State, c rune) (*fsm.StateSet, bool) {
	s, ok := this.Rules[NewNFARule(q, c)]

	if ok {
		return s, true
	}
	return nil, false
}

func (this *NFA) ToWithoutEpsilon() {
	if isSubset(&this.AcceptStates, this.epsilonClosure(this.InitialState)) {
		this.AcceptStates.Add(this.InitialState)
	}
	this.Rules = this.removeEpsilonRule()
}

func (this *NFA) removeEpsilonRule() (newRule NFARulesMap) {
	newRule = NFARulesMap{}

	states, sym := this.GetAllStates(), this.GetAllSymbols()

	delete(sym, 'E')

  // fmt.Println(states)
  // fmt.Println(sym)

	for _, q := range states.GetAll() {
		for c := range sym {
			for _, mid := range this.epsilonClosure(q).GetAll() {
				dst := this.epsilonExpand(mid, c)

				for _, st := range dst.GetAll() {
					newRule.AddDst(NewNFARule(q, c), st)
				}
			}
		}
	}

	for k := range newRule {
		if newRule[k].Len() == 0 {
			delete(newRule, k)
		}
	}

	return
}

func (this *NFA) epsilonClosure(state fsm.State) (reachable *fsm.StateSet) {
	reachable = fsm.NewStateSet()
  reachable.Add(state)

	modified := true
	for modified {
		modified = false
		for _, q := range reachable.GetAll() {
			dst, ok := this.CalcDst(q, 'E')
			if !ok || isSubset(dst, reachable) {
				continue
			}
			reachable = reachable.Union(*dst)
			modified = true
		}
	}
	return}

func (this *NFA) epsilonExpand(state fsm.State, symbol rune) (final *fsm.StateSet) {
	first := this.epsilonClosure(state)

	second := fsm.NewStateSet()

	for _, q := range first.GetAll() {
		if dst, ok := this.CalcDst(q, symbol); ok {
			for _, st := range dst.GetAll() {
				second.Add(st)
			}
		}
	}

	final = fsm.NewStateSet()
	for _, q := range second.GetAll() {
		dst := this.epsilonClosure(q)
		for _, st := range dst.GetAll() {
			final.Add(st)
		}
	}

	return
}

func isSubset(subset, superset *fsm.StateSet) bool {
	for _, key := range subset.GetAll() {
		if !superset.Contains(key) {
			return false
		}
	}
	return true
}


