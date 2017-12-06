package main

type Sorted interface {
	comparator() Comparator
	entryKey(entry interface{}) interface{}
	seq(ascending bool) ISeq
	seqForm(key interface{}, ascending bool) ISeq
}
