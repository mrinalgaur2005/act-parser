package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExpressionStmt struct {
	Exppression Expr
}

func (n ExpressionStmt) stmt() {}
