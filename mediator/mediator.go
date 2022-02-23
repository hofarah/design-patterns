package main

type mediator interface {
	canLand(airplane) bool
	notifyFree()
}
