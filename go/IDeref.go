package main

type IDeref interface {
	deref() (interface{}, error)
}
