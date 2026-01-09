package api

import (
	"fmt"
)

type Parser struct {
	ts *TokenStream
}

func NewParser(path string) (*Parser, error) {
	lexer, err := NewLexer(path)
	if err != nil {
		return nil, err
	}
	var tokens []Token
	for {
		t, err := lexer.NextToken()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
		if t.Type == TokenEOF {
			break
		}
	}

	p := &Parser{ts: NewTokenStream(tokens)}

	return p, nil
}

func (p *Parser) Parse() (*APIDefinition, error) {
	api := &APIDefinition{}

	for p.ts.Peek().Type != TokenEOF {
		p.ts.SkipIgnorable()

		tok := p.ts.Peek()

		if tok.Type == TokenIdent {
			switch tok.Value {
			case "syntax":
				t, err := p.parseSyntax()
				if err != nil {
					return nil, err
				}
				api.Syntax = t
				continue
			case "type":
				t, err := p.parseType()
				if err != nil {
					return nil, err
				}
				api.Types = append(api.Types, t)
				continue
			}
		}

		if tok.Type == TokenAt {
			g, err := p.parseGroup()
			if err != nil {
				return nil, err
			}
			api.Groups = append(api.Groups, g)
			continue
		}

		p.ts.Next()
	}

	return api, nil
}

func (p *Parser) parseSyntax() (string, error) {
	if _, err := p.ts.Expect(TokenIdent); err != nil {
		return "", err
	}

	if _, err := p.ts.Expect(TokenAssign); err != nil {
		return "", err
	}

	name, err := p.ts.Expect(TokenString)
	if err != nil {
		return "", err
	}
	return name.Value, nil
}

func (p *Parser) parseType() (*TypeDef, error) {
	if _, err := p.ts.Expect(TokenIdent); err != nil {
		return nil, err
	}

	name, err := p.ts.Expect(TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.ts.Expect(TokenLBrace); err != nil {
		return nil, err
	}

	t := &TypeDef{Name: name.Value}

	for {
		p.ts.SkipIgnorable()

		if p.ts.Peek().Type == TokenRBrace {
			break
		}

		if p.ts.Peek().Type != TokenIdent {
			p.ts.Next()
			continue
		}

		fieldName := p.ts.Peek().Value

		if p.ts.Peek().Type == TokenIdent {
			p.ts.Next()
			f := &Field{
				Name: fieldName,
				Type: p.ts.Peek().Value,
			}

			p.ts.Next()

			if p.ts.Peek().Type == TokenBacktick {
				f.Tag = p.ts.Next().Value
			}
			if p.ts.Peek().Type == TokenComment {
				f.Comment = p.ts.Next().Value
			}

			t.Fields = append(t.Fields, f)
		} else {
			t.Embed = append(t.Embed, fieldName)
		}
	}

	p.ts.Next()
	return t, nil
}

func (p *Parser) parseGroup() (*Group, error) {
	if _, err := p.ts.Expect(TokenAt); err != nil {
		return nil, err
	}

	if _, err := p.ts.Expect(TokenIdent); err != nil {
		return nil, err
	}

	route, err := p.parseRoute()
	if err != nil {
		return nil, err
	}
	p.ts.SkipIgnorable()
	if group, err := p.ts.Expect(TokenIdent); err != nil || group.Value != "group" {
		if err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf(`expected "group" got "%s"`, group.Value)
		}
	}

	name, err := p.ts.Expect(TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.ts.Expect(TokenLBrace); err != nil {
		return nil, err
	}

	g := &Group{Name: name.Value, RouteMeta: route}

	for {
		p.ts.SkipIgnorable()
		if p.ts.Peek().Type == TokenRBrace {
			break
		}
		if p.ts.Peek().Type == TokenAt {
			h, err := p.parseHandler()
			if err != nil {
				return nil, err
			}
			g.Handlers = append(g.Handlers, h)
		} else {
			p.ts.Next()
		}
	}
	p.ts.Next()
	return g, nil
}

func (p *Parser) parseRoute() (*RouteAnnotation, error) {
	if _, err := p.ts.Expect(TokenLParen); err != nil {
		return nil, err
	}

	meta := &RouteAnnotation{}

	for {
		p.ts.SkipIgnorable()

		if p.ts.Peek().Type == TokenRParen {
			p.ts.Next()
			break
		}

		key, err := p.ts.Expect(TokenIdent)
		if err != nil {
			return nil, err
		}

		if _, err := p.ts.Expect(TokenColon); err != nil {
			return nil, err
		}

		p.ts.SkipIgnorable()

		switch key.Value {
		case "tag":
			val, err := p.ts.Expect(TokenIdent)
			if err != nil {
				return nil, err
			}
			meta.Tag = val.Value

		case "auth":
			val, err := p.ts.Expect(TokenBool)
			if err != nil {
				return nil, err
			}
			meta.Auth = val.Value == "true"

		default:
			return nil, fmt.Errorf("line %d: unknown route attribute %s", key.Line, key.Value)
		}
	}

	return meta, nil
}

func (p *Parser) parseHandler() (*Handler, error) {
	if _, err := p.ts.Expect(TokenAt); err != nil {
		return nil, err
	}

	if p.ts.Peek().Type != TokenIdent {
		return nil, fmt.Errorf(`expected "ident" got "%d"`, p.ts.Peek().Type)
	}

	var (
		doc *DocAnnotation
		err error
	)
	switch p.ts.Next().Value {
	case "doc":
		doc, err = p.parseDoc()
		if err != nil {
			return nil, err
		}
	case "handler":
	default:
		return nil, fmt.Errorf(`unknown handler "%s"`, p.ts.Peek().Value)
	}

	p.ts.SkipIgnorable()

	if _, err := p.ts.Expect(TokenAt); err != nil {
		return nil, err
	}
	if handler, err := p.ts.Expect(TokenIdent); err != nil || handler.Value != "handler" {
		if err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf(`expected "handler" got "%s"`, handler.Value)
		}
	}

	funcName, err := p.ts.Expect(TokenIdent)
	if err != nil {
		return nil, err
	}

	p.ts.SkipIgnorable()

	method, err := p.ts.Expect(TokenIdent)
	if err != nil {
		return nil, err
	}

	path, err := p.ts.Expect(TokenString)
	if err != nil {
		return nil, err
	}

	if _, err := p.ts.Expect(TokenLParen); err != nil {
		return nil, err
	}
	req, err := p.ts.Expect(TokenIdent)
	if err != nil {
		return nil, err
	}
	if _, err := p.ts.Expect(TokenRParen); err != nil {
		return nil, err
	}

	if returns, err := p.ts.Expect(TokenIdent); err != nil || returns.Value != "returns" {
		if err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf(`expected "returns" got "%s"`, returns.Value)
		}
	}
	if _, err := p.ts.Expect(TokenLParen); err != nil {
		return nil, err
	}
	rsp, err := p.ts.Expect(TokenIdent)
	if err != nil {
		return nil, err
	}
	if _, err := p.ts.Expect(TokenRParen); err != nil {
		return nil, err
	}

	return &Handler{
		Method:   method.Value,
		Path:     path.Value,
		ReqType:  req.Value,
		RspType:  rsp.Value,
		FuncName: funcName.Value,
		Doc:      doc,
	}, nil
}

func (p *Parser) parseDoc() (*DocAnnotation, error) {
	if _, err := p.ts.Expect(TokenLParen); err != nil {
		return nil, err
	}

	doc := &DocAnnotation{}

	for {
		p.ts.SkipIgnorable()

		if p.ts.Peek().Type == TokenRParen {
			p.ts.Next()
			break
		}

		key, err := p.ts.Expect(TokenIdent)
		if err != nil {
			return nil, err
		}

		if key.Value == "summary" {
			val, err := p.ts.Expect(TokenString)
			if err != nil {
				return nil, err
			}
			doc.Summary = val.Value
		}
	}

	return doc, nil
}
