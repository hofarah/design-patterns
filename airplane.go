package main

import "fmt"

type airplane interface {
	getName() string
	landingReq()
	taxi()
	permitLanding()
}

type airbus struct {
	name string
	mediator
}

func (a *airbus) landingReq() {
	fmt.Println("i am " + a.getName() + " i want to land")
	if a.mediator.canLand(a) {
		fmt.Println(a.getName() + ": ok, i am landing")
	} else {
		fmt.Println(a.getName() + ": ok, i am waiting")
	}
}
func (a *airbus) taxi() {
	fmt.Println(a.getName() + ": i am leaving airport")
	a.mediator.notifyFree()
}
func (a *airbus) getName() string {
	return a.name
}
func (a *airbus) permitLanding() {
	fmt.Println(a.getName() + ": Arrival Permitted. Landing")
}
