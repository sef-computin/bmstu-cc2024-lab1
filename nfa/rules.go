package nfa

import (
	"bmstu/cc2024/lab1/fsm"
	"fmt"
)

type NFARulesMap map[NFARule]*fsm.StateSet

func (this *NFARulesMap) AddDst(rule NFARule, st fsm.State) {
	// fmt.Println("adding new dst:", rule, st)
	// fmt.Println((*this)[rule][st])
  if _, ok := (*this)[rule]; !ok {
		(*this)[rule] = fsm.NewStateSet()
	}
	(*this)[rule].Add(st)
}

func (this NFARulesMap) ToString() string {
	s := ""

	for key, val := range this {
		for _, st := range val.GetAll() {
			s += fmt.Sprintf("%s --%c--> %s\n", key.Src.String(), key.Val, st.String())
		}
	}

	return s
}

type NFARule struct {
	Src fsm.State
	Val rune
}

func NewNFARule(from fsm.State, val rune) NFARule {
	return NFARule{Src: from, Val: val}
}
