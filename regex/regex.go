package regex

import (
	"bmstu/cc2024/lab1/fsm"
	"bmstu/cc2024/lab1/nfa"
)

func ParseNFA(regexp string) *nfa.NFA {
	psr := NewParser(regexp)
	ast := psr.GetAST()
	// fmt.Println(ast.SubtreeString())
	frg := ast.Assemble(fsm.NewContext())
	// fmt.Println(frg)
	nfa := frg.Build()

	return nfa
}
