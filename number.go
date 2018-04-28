package whale

type adder interface {
	Add(n interface{}) interface{}
}

type subtracter interface {
	Sub(n interface{}) interface{}
}

type multiplier interface {
	Mul(n interface{})
}

type divider interface {
	Div(n interface{}) interface{}
}

type greater interface {
	Gt(n interface{}) interface{}
}

type lesser interface {
	Lt(n interface{}) interface{}
}

type greatEqualer interface {
	Ge(n interface{}) interface{}
}

type lessEqualer interface {
	Le(n interface{}) interface{}
}

type equaler interface {
	Eq(n interface{}) interface{}
}

type Number interface {
	adder
	subtracter
	multiplier
	divider
	greater
	lesser
	greatEqualer
	lessEqualer
	equaler
	Value() interface{}
}

type Int int
type Float float64

func Add(n1, n2 Number) Number {
	return n1.Value() + n2.Value()
}

func (i Int) Value() interface{} {
	return i
}

func (f Float) Value() interface{} {
	return f
}
