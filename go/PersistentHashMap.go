package main

type INode interface {
	assoc(shift, hash int, key interface{}, val interface{}, addedLeaf Box) INode
	without(hash int, key interface{})
	find(hash int, key interface{})
	nodeSeq() ISeq
	getHash() int
}
