package nfa

import (
	"bmstu/cc2024/lab1/fsm"
	"fmt"
)

type Fragment struct {
	InitialState     fsm.State 
	AcceptStates     fsm.StateSet  
	Rules            NFARulesMap
}

func NewFragment() *Fragment {
	return &Fragment{
		InitialState:     fsm.NewState(0),
		AcceptStates:     *fsm.NewStateSet(),
		Rules:            NFARulesMap{},
	}
}

func (frg *Fragment) AddRule(from fsm.State, c rune, next fsm.State) {
  fmt.Println()
  fmt.Println("[DEBUG] adding rule: ", from.String(), string(c), next.String())  

	r := frg.Rules

  fmt.Println(r)

	_, ok := r[NewNFARule(from, c)]
  fmt.Println("ok = ", ok)
	if !ok {
    r[NewNFARule(from, c)] = fsm.NewStateSet()
	}
	r.AddDst(NewNFARule(from, c), next)
}

func (frg *Fragment) CreateSkeleton() (Skeleton *Fragment) {
	Skeleton = NewFragment()
	Skeleton.Rules = frg.Rules
	return
}

func (frg *Fragment) MergeRule(frg2 *Fragment) (synthesizedFrg *Fragment) {
	synthesizedFrg = frg.CreateSkeleton()
  
  fmt.Println()
  fmt.Println("[DEBUG] frg1: ", synthesizedFrg) 
  fmt.Println()

	for k, v := range frg2.Rules {
		_, ok := synthesizedFrg.Rules[k]
		if !ok {
			synthesizedFrg.Rules[k] = fsm.NewStateSet()
		}
    for _, st := range v.GetAll(){
		  synthesizedFrg.Rules[k].Add(st)
    }
	}
	return
}

func (frg *Fragment) Build() *NFA {
	return NewNFA(frg.InitialState, frg.AcceptStates, frg.Rules)
}
