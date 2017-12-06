package main

type C int

const (
	STATEMENT = iota
	EXPRESSION
	RETURN
	EVAL
)

type Expr interface {
	eval() (interface{}, error)
	emit(context C, fn FnExpr, gen GeneratorAdapter)
}

type IParser interface {
	parse(context C, form interface{}) (Expr, error)
}

type AssignableExpr interface {
	evalAssign(val Expr) (interface{}, error)
	emitAssign(context C, fn FnExpr, gen GeneratorAdapter, val Expr)
}

type MaybePrimitiveExpr interface {
	emitUnboxed(context C, fn FnExpr, gen GeneratorAdapter)
}
