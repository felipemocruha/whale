package main

type ISeq interface {
	first() interface{}
	next() ISeq
	more() ISeq
	cons(o interface{}) ISeq
}
