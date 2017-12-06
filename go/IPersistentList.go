package main

type IPersistentList interface {
	peek() interface{}
	pop() IPersistentStack
}
