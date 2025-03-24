package parser

import (
	"fmt"
	"strconv"

	"github.com/mrinalgaur2005/act-parser/javascript/ast"
	"github.com/mrinalgaur2005/act-parser/javascript/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	//first parse NUd
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD Handler expected for %s\n", lexer.TokenTypeToStr(tokenKind)))
	}
	left := nud_fn(p)

	//While we have Led and current is > , parse lhs
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("NUD Handler expected for %s\n", lexer.TokenTypeToStr(tokenKind)))
		}
		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}
	return left
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Cant crate a primary_expr in %s\n", lexer.TokenTypeToStr(p.currentTokenKind())))
	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parse_assignment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()

	rhs := parse_expr(p, bp)

	return ast.AssignmentExpr{
		Operator: operatorToken,
		Value:    rhs,
		Assigne:  left,
	}
}

func parse_prefix_expr(p *parser) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, default_bp)

	return ast.PrefixExpr{
		Operator:  operatorToken,
		RightExpr: rhs,
	}
}
