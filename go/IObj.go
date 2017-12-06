package main

type IObj interface {
	withMeta(meta IPersistentMap) IObj

	meta() IPersistentMap
}
