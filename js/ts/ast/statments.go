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
	ExplicitType Type
}

func (n VarDeclStmt) stmt() {}

type InterfaceProp struct {
	IsStatic bool
	Type     Type
}

type InterfaceMethod struct {
	IsStatic bool
	//Type Type
}
type InterfaceDeclStmt struct {
	InterfaceName string
	Properties    map[string]InterfaceProp
	Methods       map[string]InterfaceMethod
}

func (n InterfaceDeclStmt) stmt() {}

type TypeProp struct {
	IsStatic bool
	Type     Type
}

type TypeMethod struct {
	IsStatic bool
	//Type Type
}
type TypeDeclStmt struct {
	InterfaceName string
	Properties    map[string]TypeProp
	Methods       map[string]TypeMethod
}

func (n TypeDeclStmt) stmt() {}
