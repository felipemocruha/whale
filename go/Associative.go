package main

type Associative interface {
	containsKey(key interface{}) bool
	entryAt(key interface{}) IMapEntry
	assoc(key, val interface{}) Assciative
	valAt(key interface{}) interface{}
	valAt(key, notFound interface{}) interface{}
}
