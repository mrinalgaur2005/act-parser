package parser

import (
	"fmt"

	"github.com/mrinalgaur2005/act-parser/js/ts/ast"
	"github.com/mrinalgaur2005/act-parser/js/ts/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTokenTypeLookups()
	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	p := createParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: Body,
	}
}

// helper
func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenType {
	return p.currentToken().Type
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) expectError(expectedKind lexer.TokenType, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Type

	if kind != expectedKind {
		if err != nil {
			err = fmt.Sprintf("Expected %s but received %s insted \n", lexer.TokenTypeToStr(expectedKind), lexer.TokenTypeToStr(kind))
		}
		fmt.Printf("inside nill error %s and %s", kind, expectedKind)
		panic(err)
	}
	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenType) lexer.Token {
	return p.expectError(expectedKind, nil)
}
