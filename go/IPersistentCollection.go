package main

type IPersistentCollection interface {
	cons(o interface{}) IPersistentCollection
	empty() IPersistentCollection
	equiv(o interface{}) IPersistentCollection

	seq() ISeq
}
