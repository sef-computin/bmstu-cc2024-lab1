package fsm

import "fmt"

type State struct {
	N int
}

func NewState(n int) State {
	return State{
		N: n,
	}
}

func (s State) String() string {
	return fmt.Sprintf("q%d", s.N)
}

func (s State) Equals(st State) bool {
	return st.N == s.N
}

type Context struct {
	N int
}

func NewContext() *Context {
	return &Context{
		N: -1,
	}
}

func (ctx *Context) Increment() int {
	ctx.N++
	return ctx.N
}


