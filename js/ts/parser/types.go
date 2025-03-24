package parser

import (
	"fmt"

	"github.com/mrinalgaur2005/act-parser/js/ts/ast"
	"github.com/mrinalgaur2005/act-parser/js/ts/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenType]type_nud_handler
type type_led_lookup map[lexer.TokenType]type_led_handler

var type_bp_lu = bp_lookup{}
var type_led_lu = type_led_lookup{}
var type_nud_lu = type_nud_lookup{}

func parse_type(p *parser, bp binding_power) ast.Type {
	//first parse Nud
	tokenKind := p.currentTokenKind()
	nud_fn, exists := type_nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("Type_NUD Handler expected for %s\n", lexer.TokenTypeToStr(tokenKind)))
	}
	left := nud_fn(p)

	//While we have Led and current is > , parse lhs
	for type_bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := type_led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("Type_led Handler expected for %s\n", lexer.TokenTypeToStr(tokenKind)))
		}
		left = led_fn(p, left, type_bp_lu[p.currentTokenKind()])
	}
	return left
}

func type_led(kind lexer.TokenType, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud(kind lexer.TokenType, nud_fn type_nud_handler) {
	type_nud_lu[kind] = nud_fn
}

func createTokenTypeLookups() {
	type_nud(lexer.IDENTIFIER, parse_simple_type)
	type_led(lexer.OPEN_BRACK, call, parse_arr_type)
}

func parse_simple_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_arr_type(p *parser, left ast.Type, bp binding_power) ast.Type {
	p.expect(lexer.OPEN_BRACK)
	p.expect(lexer.CLOSED_BRACK)

	return ast.ArrayType{
		Underlying: left,
	}
}
