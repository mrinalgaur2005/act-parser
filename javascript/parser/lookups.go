package parser

import (
	"github.com/mrinalgaur2005/act-parser/javascript/ast"
	"github.com/mrinalgaur2005/act-parser/javascript/lexer"
)

type binding_power int

const (
	// Lowest precedence
	default_bp binding_power = iota
	comma
	assignment
	// Ternary `?:`
	ternary
	// Logical Nullish Coalescing `??`
	nullish_coalescing
	logical
	bitwise
	relational
	shift
	additive
	multiplicative
	exponentiation
	// Unary `!`, `~`, `-`, `+`, `typeof`, `void`, `delete`
	unary
	// Function calls `()`
	call
	member
	// Optional chaining `?.`
	optional_chaining
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type smt_lookup map[lexer.TokenType]stmt_handler
type nud_lookup map[lexer.TokenType]nud_handler
type led_lookup map[lexer.TokenType]led_handler
type bp_lookup map[lexer.TokenType]binding_power

var bp_lu = bp_lookup{}
var led_lu = led_lookup{}
var nud_lu = nud_lookup{}
var stmt_lu = smt_lookup{}

func led(kind lexer.TokenType, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenType, nud_fn nud_handler) {
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenType, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {

	//assignment
	led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MUL_EQUALS, assignment, parse_assignment_expr)
	led(lexer.DIV_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MOD_EQUALS, assignment, parse_assignment_expr)
	led(lexer.EXPONENT_EQUALS, assignment, parse_assignment_expr)

	//Logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)

	//Relational
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)
	led(lexer.STRICT_EQUALS, relational, parse_binary_expr)
	led(lexer.STRICT_NOT_EQUALS, relational, parse_binary_expr)

	//Additve + MUltiplicative
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	//Literals + Symbols
	nud(lexer.NUMBER, parse_primary_expr)
	nud(lexer.STRING, parse_primary_expr)
	nud(lexer.IDENTIFIER, parse_primary_expr)
	nud(lexer.DASH, parse_prefix_expr)

	//Statements
	stmt(lexer.CONST, parse_var_declare_stmt)
	stmt(lexer.LET, parse_var_declare_stmt)
	stmt(lexer.VAR, parse_var_declare_stmt)
}
