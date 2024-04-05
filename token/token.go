package token

import "fmt"

type Type int

const (
	CHARACTER Type = iota
	UNION
	STAR
	LPAREN
	RPAREN
	EOF
)

func (k Type) String() string {
	switch k {
	case CHARACTER:
		return "CHARACTER"
	case UNION:
		return "UNION"
	case STAR:
		return "STAR"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case EOF:
		return "EOF"
	default:
		return ""
	}
}

type Token struct {
	V  rune 
	Ty Type
}

func (t Token) String() string {
	return fmt.Sprintf("V -> \x1b[32m%v\x1b[0m\tKind -> \x1b[32m%v\x1b[0m", string(t.V), t.Ty)
}

// NewToken returns a new Token.
func NewToken(value rune, k Type) Token {
	return Token{
		V:  value,
		Ty: k,
	}
}
