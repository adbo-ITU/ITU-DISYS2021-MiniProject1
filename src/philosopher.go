package main

import "time"

type philosopher struct {
	id        int
	isEating  bool
	numEats   int
	numThinks int
	events    chan<- philosopher
}

const doNothing = 1
const pickUp = 2
const putDown = 3

func (p philosopher) philosopherLife(leftFork fork, rightFork fork) {
	for {
		arbiter.Lock()

		leftStatus := <-leftFork.output
		rightStatus := <-rightFork.output

		willPickUp := !leftStatus.isPickedUp && !rightStatus.isPickedUp

		action := doNothing
		if willPickUp {
			action = pickUp
		}

		leftFork.input <- action
		rightFork.input <- action

		if willPickUp {
			p.isEating = true
			p.numEats++
			p.events <- p
		} else {
			p.numThinks++
			p.events <- p
		}

		arbiter.Unlock()

		time.Sleep(2 * time.Millisecond)

		if p.isEating {
			arbiter.Lock()
			<-leftFork.output  // prepare left fork to receive message
			<-rightFork.output // prepare right fork to receive message
			leftFork.input <- putDown
			rightFork.input <- putDown
			arbiter.Unlock()
			time.Sleep(2 * time.Millisecond)
		}
	}
}
