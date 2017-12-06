package main

type Streamable interface {
	stream() (Stream, error)
}
