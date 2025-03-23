package parser

import (
	"github.com/mrinalgaur2005/act-parser/javascript/ast"
	"github.com/mrinalgaur2005/act-parser/javascript/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}
	expression := parse_expr(p, default_bp)

	p.expect(lexer.SEMICOLON)

	return ast.ExpressionStmt{
		Exppression: expression,
	}
}

func parse_var_declare_stmt(p *parser) ast.Stmt {
	isConstant := p.advance().Type == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "inside variable declaration expeected to find var name").Value

	p.expect(lexer.ASSIGNMENT)
	assignedVal := parse_expr(p, assignment)
	p.expect(lexer.SEMICOLON)

	return ast.VarDeclStmt{
		IsConstant:   isConstant,
		VariableName: varName,
		AssignedVal:  assignedVal,
	}
}
