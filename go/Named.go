package main

type Named interface {
	getNamespace() string
	getName() string
}
