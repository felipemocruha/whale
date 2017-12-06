package main

type IPersistentMap interface {
	assoc(key, val interface{}) IPersistentMap
	assocEx(key, val interface{}) (IPersistentMap, error)
	without(key interface{}) (IPersistentMap, error)
}
