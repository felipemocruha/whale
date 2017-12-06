package main

type Settable interface {
	doSet(val interface{}) (interface{}, error)
	doReset(val interface{}) (interface{}, error)
}
