package main

type IPersistentSet interface {
	disjoin(key interface{}) (IPersistentSet, error)
	contains(key interface{}) bool
	get(key interface{}) interface{}

	count() int
	cons(o interface{}) IPersistentCollection
	empty() IPersistentCollection
	equiv(o interface{}) IPersistentCollection
	seq() ISeq
}
