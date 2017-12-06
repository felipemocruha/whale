package main

type IPersistentStack interface {
	peek() interface{}
	pop() IPersistentStack

	cons(o interface{}) IPersistentCollection
	empty() IPersistentCollection
	equiv(o interface{}) IPersistentCollection
	seq() ISeq
}
