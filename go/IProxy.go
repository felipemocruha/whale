package main

type IProxy interface {
	__initClojureFnMappings(m IPersistentMap)
	__updateClojureFnMappings(m IPersistentMap)
	__getClojureFnMappings() IPersistentMap
}
