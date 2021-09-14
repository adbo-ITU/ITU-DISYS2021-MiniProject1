package main

import (
	"time"
)

type philosopher struct {
	id        int
	isEating  bool
	numEats   int
	numThinks int
	input     chan int
	output    chan PhilosopherStatus
}

type PhilosopherStatus struct {
	numEats   int
	numThinks int
}

const doNothing = 1
const pickUp = 2
const putDown = 3

func (p philosopher) philosopherLife(leftFork fork, rightFork fork) {
	for {
		select {
		// if an input request is received, write a response to output
		case <-p.input:
			p.output <- PhilosopherStatus{numEats: p.numEats, numThinks: p.numThinks}
		// No input request, move on
		default:
		}

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
		} else {
			p.numThinks++
		}

		arbiter.Unlock()

		time.Sleep(2 * time.Millisecond)

		if p.isEating {
			arbiter.Lock()
			<-leftFork.output  // prepare left fork to receive message
			<-rightFork.output // prepare right fork to receive message
			leftFork.input <- putDown
			rightFork.input <- putDown
			p.isEating = false
			arbiter.Unlock()
		}
	}
}
