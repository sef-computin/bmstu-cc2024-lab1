package regex

import (
	"bmstu/cc2024/lab1/lexer"
	"bmstu/cc2024/lab1/token"
	"fmt"
	"log"
)


type Parser struct {
	tokens []token.Token
	look   token.Token
}

func NewParser(s string) *Parser {
	p := &Parser{
		tokens: lexer.NewLexer(s).Scan(),
	}
	p.move()
	return p
}

func (psr *Parser) GetAST() Node {
	ast := psr.expression()
	return ast
}

func (psr *Parser) move() {
	if len(psr.tokens) == 0 {
		psr.look = token.NewToken('\x00', token.EOF)
	} else {
		psr.look = psr.tokens[0]
		psr.tokens = psr.tokens[1:]
	}
}

func (psr *Parser) moveWithValidation(expect token.Type) {
	if psr.look.Ty != expect {
		err := fmt.Sprintf("[syntax error] expect:\x1b[31m%s\x1b[0m actual:\x1b[31m%s\x1b[0m", expect, psr.look.Ty)
		log.Fatal(err)
	}
	psr.move()
}

func (psr *Parser) expression() Node {
	nd := psr.subexpr()
	psr.moveWithValidation(token.EOF)
	return nd
}

func (psr *Parser) subexpr() Node {
	nd := psr.seq()
	for {
		if psr.look.Ty == token.UNION {
			psr.moveWithValidation(token.UNION)
			nd2 := psr.seq()
			nd = NewUnion(nd, nd2)
		} else {
			break
		}
	}
	return nd
}

func (psr *Parser) seq() Node {
	if psr.look.Ty == token.LPAREN || psr.look.Ty == token.CHARACTER {
		return psr.subseq()
	}
	return NewCharacter('E')
}

func (psr *Parser) subseq() Node {
	nd := psr.sufope()
	if psr.look.Ty == token.LPAREN || psr.look.Ty == token.CHARACTER {
		nd2 := psr.subseq()
		return NewConcat(nd, nd2)
	}
	return nd
}

func (psr *Parser) sufope() Node {
	nd := psr.factor()
	switch psr.look.Ty {
	case token.STAR:
		psr.move()
		return NewStar(nd)
	}
	return nd
}

func (psr *Parser) factor() Node {
	if psr.look.Ty == token.LPAREN {
		psr.moveWithValidation(token.LPAREN)
		nd := psr.subexpr()
		psr.moveWithValidation(token.RPAREN)
		return nd
	}
	nd := NewCharacter(psr.look.V)
	psr.moveWithValidation(token.CHARACTER)
	return nd
}
