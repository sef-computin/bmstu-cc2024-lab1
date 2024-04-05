package main

import (
	"bmstu/cc2024/lab1/dfa"
	"bmstu/cc2024/lab1/regex"
	"fmt"
)

func main() {
	nfa := regex.ParseNFA("(a|b)abb")

	fmt.Println("Initial E-NFA")
	fmt.Println("Initial State: ", nfa.InitialState)
	fmt.Println("Accept States", nfa.AcceptStates)
	fmt.Println(nfa.Rules.ToString())

	fmt.Println("Deleting E")
	nfa.ToWithoutEpsilon()
	fmt.Println("New NFA:")
	fmt.Println("InitialState: ", nfa.InitialState)
	fmt.Println("AcceptStates: ", nfa.AcceptStates)
	fmt.Println(nfa.Rules.ToString())

	fmt.Println("DFA")
  dfa := dfa.NewDFAFromNFA(nfa)
	fmt.Println("InitialState: ", dfa.InitialState)
	fmt.Println("AcceptStates: ", dfa.AcceptStates)
	fmt.Println(dfa.Rules.ToString())

	// fmt.Println("NFA reverse")
	// nfa = nfa.ReverseNFA()
	//
	// fmt.Println("Initial State: ", nfa.InitialState)
	// fmt.Println("Accept States", nfa.AcceptStates)
	// fmt.println(nfa.rules.tostring())
}
