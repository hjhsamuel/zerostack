package parser

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Lexer struct {
	input []rune
	pos   int
	line  int
}

func NewLexer(path string) (*Lexer, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	if !strings.HasSuffix(path, ".api") {
		return nil, errors.New("not a .api file")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var runeList []rune
	for _, r := range string(content) {
		runeList = append(runeList, r)
	}

	return &Lexer{
		input: runeList,
		line:  1,
	}, nil
}

func (l *Lexer) NextToken() (Token, error) {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]

		switch {
		case ch == '\n':
			l.pos++
			l.line++
			return Token{Type: TokenNewLine, Line: l.line}, nil

		case isWhitespace(ch):
			l.pos++
			continue

		case ch == '=':
			l.pos++
			return Token{Type: TokenAssign, Line: l.line}, nil

		case ch == '@':
			l.pos++
			return Token{Type: TokenAt, Line: l.line}, nil

		case ch == ':':
			l.pos++
			return Token{Type: TokenColon, Line: l.line}, nil

		case ch == '{':
			l.pos++
			return Token{Type: TokenLBrace, Line: l.line}, nil

		case ch == '}':
			l.pos++
			return Token{Type: TokenRBrace, Line: l.line}, nil

		case ch == '(':
			l.pos++
			return Token{Type: TokenLParen, Line: l.line}, nil

		case ch == ')':
			l.pos++
			return Token{Type: TokenRParen, Line: l.line}, nil

		case ch == '`':
			return l.readBacktick()

		case ch == '"':
			return l.readString()

		case ch == '/' && l.peek() == '/':
			return l.readComment()

		default:
			if isLetter(ch) {
				return l.readIdent()
			}
			return Token{}, fmt.Errorf("line %d: unexpected character '%c'", l.line, ch)
		}
	}

	return Token{Type: TokenEOF, Line: l.line}, nil
}

func (l *Lexer) readIdent() (Token, error) {
	start := l.pos
	for l.pos < len(l.input) && (isLetter(l.input[l.pos]) || isDigit(l.input[l.pos])) {
		l.pos++
	}
	val := string(l.input[start:l.pos])

	if val == "true" || val == "false" {
		return Token{Type: TokenBool, Value: val, Line: l.line}, nil
	}

	return Token{Type: TokenIdent, Value: val, Line: l.line}, nil
}

func (l *Lexer) readString() (Token, error) {
	l.pos++
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '"' {
		l.pos++
	}
	if l.pos >= len(l.input) {
		return Token{}, fmt.Errorf("line %d: unterminated string", l.line)
	}
	val := string(l.input[start:l.pos])
	l.pos++
	return Token{Type: TokenString, Value: val, Line: l.line}, nil
}

func (l *Lexer) readBacktick() (Token, error) {
	l.pos++
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '`' {
		l.pos++
	}
	if l.pos >= len(l.input) {
		return Token{}, fmt.Errorf("line %d: unterminated backtick", l.line)
	}
	val := string(l.input[start:l.pos])
	l.pos++
	return Token{Type: TokenBacktick, Value: val, Line: l.line}, nil
}

func (l *Lexer) readComment() (Token, error) {
	l.pos += 2
	start := l.pos
	for l.pos < len(l.input) && l.input[l.pos] != '\n' {
		l.pos++
	}
	val := string(l.input[start:l.pos])
	return Token{Type: TokenComment, Value: strings.TrimSpace(val), Line: l.line}, nil
}

func (l *Lexer) peek() rune {
	if l.pos+1 >= len(l.input) {
		return 0
	}
	return l.input[l.pos+1]
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
