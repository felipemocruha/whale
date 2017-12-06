package main

type IMapEntry interface {
	key() interface{}
	val() interface{}
}
