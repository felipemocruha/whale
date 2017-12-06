package main

type IRef interface {
	setValidator(vf IFn)
	getValidator() IFn
	getWatches() IPersistentMap
	addWatch(key interface{}, callback IFn) IRef
	removeWatch(key interface{}) IRef
}
