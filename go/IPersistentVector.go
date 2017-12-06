package main

type IPersistentVector interface {
	nth(i int) interface{}
	assocN(i int, val interface{}) IPersistentVector
	cons(o interface{}) IPersistentVector

	containsKey(key interface{}) bool
	entryAt(key interface{}) IMapEntry
	assoc(key, val interface{}) Assciative
	valAt(key interface{}) interface{}
	valAt(key, notFound interface{}) interface{}

	peek() interface{}
	pop() IPersistentStack
	cons(o interface{}) IPersistentCollection
	empty() IPersistentCollection
	equiv(o interface{}) IPersistentCollection
	seq() ISeq
}
