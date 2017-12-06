package main

type Reversible interface {
	rseq() (ISeq, error)
}
