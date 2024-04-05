package lexer

import "bmstu/cc2024/lab1/token"

type Lexer struct {
  s []rune // string to be analyzed
}

func NewLexer(s string) *Lexer {
  return &Lexer{
    s: []rune(s),
  }
}

func (l *Lexer) Scan() (tokenList []token.Token) {
  for i := 0; i < len(l.s); i++ {
    switch l.s[i] {
    case '\x00':
      tokenList = append(tokenList, token.NewToken(l.s[i], token.EOF))
    case '|':
      tokenList = append(tokenList, token.NewToken(l.s[i], token.UNION))
    case '(':
      tokenList = append(tokenList, token.NewToken(l.s[i], token.LPAREN))
    case ')':
      tokenList = append(tokenList, token.NewToken(l.s[i], token.RPAREN))
    case '*':
      tokenList = append(tokenList, token.NewToken(l.s[i], token.STAR))
    case '\\':
      tokenList = append(tokenList, token.NewToken(l.s[i+1], token.CHARACTER))
      i++
    default:
      tokenList = append(tokenList, token.NewToken(l.s[i], token.CHARACTER))
    }
  }
  return
}
