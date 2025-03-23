package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExpressionStmt struct {
	Exppression Expr
}

func (n ExpressionStmt) stmt() {}

type VarDeclStmt struct {
	VariableName string
	IsConstant   bool
	AssignedVal  Expr
	// ExplicitType Type
}

func (n VarDeclStmt) stmt() {}
