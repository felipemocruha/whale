package main

type IReference interface {
	alterMeta(alter IFn, args ISeq) (IPersistentMap, error)
	resetMeta(m IPersistentMap) IPersistentMap
}
