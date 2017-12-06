package main

type IReduce interface {
	reduce(f IFn) (interface{}, error)
	reduce(f IFn, start interface{}) (interface{}, error)
}
