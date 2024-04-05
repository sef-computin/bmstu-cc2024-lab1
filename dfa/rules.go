package dfa

import (
	"bmstu/cc2024/lab1/fsm"
	"fmt"
)

type DFARulesMap map[DFARule]fsm.State

func (this DFARulesMap) ToString() string {
	s := ""

	for key, st := range this {
		s += fmt.Sprintf("%s --%c--> %s\n", key.Src.String(), key.Val, st.String())
	}

	return s
}

type DFARule struct {
	Src fsm.State
	Val rune
}

func NewDFARule(from fsm.State, val rune) DFARule {
	return DFARule{Src: from, Val: val}
}
