package parser

import "fmt"

type TokenStream struct {
	tokens []Token
	pos    int
}

func NewTokenStream(tokens []Token) *TokenStream {
	return &TokenStream{tokens: tokens}
}

func (ts *TokenStream) Peek() Token {
	if ts.pos >= len(ts.tokens) {
		return Token{Type: TokenEOF}
	}
	return ts.tokens[ts.pos]
}

func (ts *TokenStream) Next() Token {
	t := ts.Peek()
	ts.pos++
	return t
}

func (ts *TokenStream) Expect(tt TokenType) (Token, error) {
	t := ts.Next()
	if t.Type != tt {
		return t, fmt.Errorf(
			"line %d: expect %v, got %v (%s)",
			t.Line, tt, t.Type, t.Value,
		)
	}
	return t, nil
}

func (ts *TokenStream) SkipIgnorable() {
	for {
		switch ts.Peek().Type {
		case TokenNewLine, TokenComment:
			ts.Next()
		default:
			return
		}
	}
}
