package drawer

import (
	"bmstu/cc2024/lab1/dfa"
	"bmstu/cc2024/lab1/nfa"

	"fmt"
	// "github.com/dominikbraun/graph"
	// "github.com/dominikbraun/graph/draw"
	"os"
	"os/exec"
)


func DrawDFA(_dfa *dfa.DFA, fa_name string) {
  
  var buf []byte = []byte("digraph {\n")

  nodes := _dfa.GetAllStates()

  for _, node := range nodes.GetAll(){
    options := ""
    if _dfa.InitialState.N == node.N{
      options = ", color=blue"
    }
    if _dfa.AcceptStates.Contains(node){
      options = ", color=red"
    }

    buf = append(buf, []byte(fmt.Sprintf("\"%s\" [ weight=0%s ];\n", node.String(), options))...)    
  }

  for rule, dst := range _dfa.Rules{
    buf = append(buf, []byte(fmt.Sprintf("\"%s\" -> \"%s\" [ label=%s, weight=0 ];\n", rule.Src.String(), dst.String(), string(rule.Val) ) )...)    
  }


  buf = append(buf, []byte("\n}")...)

  name := fmt.Sprintf("./%s.gv", fa_name)

  os.WriteFile(name, buf, 0666)

	cmd := exec.Command("dot", "-Tsvg", name)
	stdout, err := cmd.Output()
	if err == nil {
    file, _ := os.Create(name + ".svg")
		file.Write(stdout)
	}
	// fmt.Println(stdout, err)
	// os.Remove(name)

	cmd = exec.Command("firefox", name+".svg")
	_, _ = cmd.Output()


}

func DrawNFA(_nfa *nfa.NFA, fa_name string) {
  var buf []byte = []byte("strict digraph {\n")

  nodes := _nfa.GetAllStates()

  for _, node := range nodes.GetAll(){
    options := ""
    if _nfa.InitialState.N == node.N{
      options = ", color=blue"
    }
    if _nfa.AcceptStates.Contains(node){
      options = ", color=red"
    }
    buf = append(buf, []byte(fmt.Sprintf("\"%s\" [ weight=0%s ];\n", node.String(), options))...)    
  }

  for rule, dst := range _nfa.Rules{
    for _, st := range dst.GetAll(){
      buf = append(buf, []byte(fmt.Sprintf("\"%s\" -> \"%s\" [ label=%s, weight=0 ];\n", rule.Src.String(), st.String(), string(rule.Val) ) )...)    
    }
  }


  buf = append(buf, []byte("\n}")...)

  name := fmt.Sprintf("./%s.gv", fa_name)

  os.WriteFile(name, buf, 0666)

	cmd := exec.Command("dot", "-Tsvg", name)
	stdout, err := cmd.Output()
	if err == nil {
    file, _ := os.Create(name + ".svg")
		file.Write(stdout)
	}
	// fmt.Println(stdout, err)
	os.Remove(name)

	cmd = exec.Command("firefox", name+".svg")
	_, _ = cmd.Output()


}
