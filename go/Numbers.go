package main

type Ops interface {
	combine(y Ops) Ops
	opsWith(x IntegerOps) Ops
	opsWith(x LongOps) Ops
	opsWith(x FloatOps) Ops
	opsWith(x DoubleOps) Ops
	opsWith(x RatioOps) Ops
	isZero(x Number) bool
	isPos(x Number) bool
	isNeg(x Number) bool
	add(x, y Number) Number
	multiply(x, y Number) Number
	divide(x, y Number) Number
	quotient(x, y Number) Number
	remainder(x, y Number) Number
	equiv(x, y Number) Number
	lt(x, y Number) Number
	negate(x, y Number) Number
	inc(x Number) Number
	dec(x Number) Number
}

type BitOps interface {
	combine(y BitOps) BitOps
	bitOpsWith(x IntegerBitOps) BitOps
	bitOpsWith(x LongBitOps) BitOps
	not(x Number) Number
	and(x, y Number) Number
	or(x, y Number) Number
	xor(x, y Number) Number
	andNot(x, y Number) Number
	clearBit(x Number, n int) Number
	setBit(x Number, n int) Number
	flipBit(x Number, n int) Number
	testBit(x Number, n int) bool
	shiftLeft(x Number, n int) Number
	shiftRight(x Number, n int) Number
}
