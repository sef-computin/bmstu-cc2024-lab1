package regex

import (
	"bmstu/cc2024/lab1/fsm"
	"bmstu/cc2024/lab1/nfa"
	"fmt"
)

const (
	TypeCharacter = "Character"
	TypeUnion     = "Union"
	TypeConcat    = "Concat"
	TypeStar      = "Star"
)

type Node interface {
	SubtreeString() string

	Assemble(*fsm.Context) *nfa.Fragment
}

type Character struct {
	Ty string
	V  rune
}

func (c *Character) String() string {
	return c.SubtreeString()
}

func NewCharacter(r rune) *Character {
	return &Character{
		Ty: TypeCharacter,
		V:  r,
	}
}

func (c *Character) Assemble(ctx *fsm.Context) *nfa.Fragment {
	newFrg := nfa.NewFragment()

	q1 := fsm.NewState(ctx.Increment())
	q2 := fsm.NewState(ctx.Increment())

	newFrg.AddRule(q1, c.V, q2)

	newFrg.InitialState = q1
	newFrg.AcceptStates.Add(q2)

	return newFrg
}

func (c *Character) SubtreeString() string {
	return fmt.Sprintf("\x1b[32m%s('%s')\x1b[32m", c.Ty, string(c.V))
}

type Union struct {
	Ty   string
	Ope1 Node
	Ope2 Node
}

func (u *Union) String() string {
	return u.SubtreeString()
}

func NewUnion(ope1, ope2 Node) *Union {
	return &Union{
		Ty:   TypeUnion,
		Ope1: ope1,
		Ope2: ope2,
	}
}

func (u *Union) Assemble(ctx *fsm.Context) *nfa.Fragment {
	newFrg := nfa.NewFragment()
	frg1 := u.Ope1.Assemble(ctx)
	frg2 := u.Ope2.Assemble(ctx)

	newState := fsm.NewState(ctx.Increment())

	newFrg = frg1.MergeRule(frg2)
	newFrg.AddRule(newState, 'E', frg1.InitialState)
	newFrg.AddRule(newState, 'E', frg2.InitialState)

	newFrg.InitialState = newState
  for _, st := range frg1.AcceptStates.GetAll(){
    newFrg.AcceptStates.Add(st)
  }
  for _, st := range frg2.AcceptStates.GetAll(){
    newFrg.AcceptStates.Add(st)
  }

	return newFrg
}

func (u *Union) SubtreeString() string {
	return fmt.Sprintf("\x1b[36m%s(%s, %s\x1b[36m)\x1b[0m", u.Ty, u.Ope1.SubtreeString(), u.Ope2.SubtreeString())
}

type Concat struct {
	Ty   string
	Ope1 Node
	Ope2 Node
}

func (c *Concat) String() string {
	return c.SubtreeString()
}

func NewConcat(ope1, ope2 Node) *Concat {
	return &Concat{
		Ty:   TypeConcat,
		Ope1: ope1,
		Ope2: ope2,
	}
}

func (c *Concat) Assemble(ctx *fsm.Context) *nfa.Fragment {
	newFrg := nfa.NewFragment()
	frg1 := c.Ope1.Assemble(ctx)
	frg2 := c.Ope2.Assemble(ctx)


  

	newFrg = frg1.MergeRule(frg2)
	for _, q := range frg1.AcceptStates.GetAll() {
		newFrg.AddRule(q, 'E', frg2.InitialState)
	}

	newFrg.InitialState = frg1.InitialState
  for _, state := range frg2.AcceptStates.GetAll(){
    newFrg.AcceptStates.Add(state)
  }
	
  return newFrg
}

func (c *Concat) SubtreeString() string {
	return fmt.Sprintf("\x1b[31m%s(%s, %s\x1b[31m)\x1b[0m", c.Ty, c.Ope1.SubtreeString(), c.Ope2.SubtreeString())
}

type Star struct {
	Ty  string
	Ope Node
}

func (s *Star) String() string {
	return s.SubtreeString()
}

func NewStar(ope Node) *Star {
	return &Star{
		Ty:  TypeStar,
		Ope: ope,
	}
}

func (s *Star) Assemble(ctx *fsm.Context) *nfa.Fragment {
	orgFrg := s.Ope.Assemble(ctx)
	newFrg := orgFrg.CreateSkeleton()

	newState1 := fsm.NewState(ctx.Increment())
	newState2 := fsm.NewState(ctx.Increment())

	newFrg.AddRule(newState1, 'E', newState2)
	newFrg.AddRule(newState1, 'E', orgFrg.InitialState)
	for _, q := range orgFrg.AcceptStates.GetAll() {
		newFrg.AddRule(q, 'E', newState2)
		newFrg.AddRule(q, 'E', orgFrg.InitialState)
	}

	newFrg.InitialState = newState1
	newFrg.AcceptStates.Add(orgFrg.InitialState)
	newFrg.AcceptStates.Add(newState2)

	return newFrg
}

func (s *Star) SubtreeString() string {
	return fmt.Sprintf("\x1b[33m%s(%s\x1b[33m)\x1b[0m", s.Ty, s.Ope.SubtreeString())
}

