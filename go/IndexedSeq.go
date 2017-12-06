package main

type IndexedSeq interface {
	index() int

	first() interface{}
	next() ISeq
	more() ISeq
	cons(o interface{}) ISeq
	count() int
}
