package fsm

type Edge struct {
	val rune
	src State
	dst State
}

func NewEdge(x rune, src, dst State) Edge {
	return Edge{x, src, dst}
}

func (this *Edge) SetValue(x rune) {
	this.val = x
}

func (this *Edge) SetSrc(src State) {
	this.src = src
}

func (this *Edge) SetDst(dst State) {
	this.dst = dst
}

func (this *Edge) GetValue() rune {
	return this.val
}

func (this *Edge) GetSrc() State {
	return this.src
}

func (this *Edge) GetDst() State {
	return this.dst
}


