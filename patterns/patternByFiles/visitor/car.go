package main

type carV interface {
	getType() string
	accept(visitor)
}
