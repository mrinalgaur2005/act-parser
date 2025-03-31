package parser

import (
	"fmt"

	"github.com/mrinalgaur2005/act-parser/js/ts/ast"
	"github.com/mrinalgaur2005/act-parser/js/ts/lexer"
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
	var explicitType ast.Type
	var assignedVal ast.Expr
	isConstant := p.advance().Type == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "inside variable declaration expeected to find var name").Value

	//can be any
	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		explicitType = parse_type(p, default_bp)
	}

	if p.currentTokenKind() != lexer.SEMICOLON {
		fmt.Printf("hello %s", p.currentToken().Value)
		p.expect(lexer.ASSIGNMENT)
		assignedVal = parse_expr(p, assignment)
	} else if explicitType == nil {
		panic("Misssing either rhs side in var declaration or explicit type")
	}
	p.expect(lexer.SEMICOLON)

	if isConstant && assignedVal == nil {
		panic("'const' declarations must be initialized")
	}

	return ast.VarDeclStmt{
		ExplicitType: explicitType,
		IsConstant:   isConstant,
		VariableName: varName,
		AssignedVal:  assignedVal,
	}
}

func parse_interface_decl_stmt(p *parser) ast.Stmt {
	p.expect(lexer.INTERFACE)
	var interfaceName = p.expect(lexer.IDENTIFIER).Value
	var properties = map[string]ast.InterfaceProp{}
	var methods = map[string]ast.InterfaceMethod{}

	p.expect(lexer.OPEN_CURLY)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		var isStatic bool
		var propertyName string
		if p.currentTokenKind() == lexer.STATIC {
			isStatic = true
			p.expect(lexer.STATIC)
		}
		//properties start here
		if p.currentTokenKind() == lexer.IDENTIFIER {
			propertyName = p.expect(lexer.IDENTIFIER).Value
			p.expectError(lexer.COLON, "Expected to find colon follwing property name inside the interface decln")
			interfaceType := parse_type(p, default_bp)

			_, exists := properties[propertyName]

			if exists {
				panic(fmt.Sprintf("Property %s has already been used in the interface decl", propertyName))
			}

			properties[propertyName] = ast.InterfaceProp{
				IsStatic: isStatic,
				Type:     interfaceType,
			}

			continue
		}
		panic("cant handle methods abhi")
	}

	p.expect(lexer.CLOSE_CURLY)

	return ast.InterfaceDeclStmt{
		Properties:    properties,
		Methods:       methods,
		InterfaceName: interfaceName,
	}
}
