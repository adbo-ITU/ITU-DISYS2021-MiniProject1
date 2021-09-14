package main

import (
	"time"
)

// A type representing a philosopher entity
type philosopher struct {
	isEating  bool
	numEats   int
	numThinks int
	input     chan int
	output    chan PhilosopherStatus
}

// A type containing the information about a philosopher that is relevant for
// the outside world
type PhilosopherStatus struct {
	numEats   int
	numThinks int
	isEating  bool
}

// A sort of enum to tell the fork what to do
const doNothing = 1
const pickUp = 2
const putDown = 3

// The main philosopher life cycle - each "round" at the table, this is done
func (p philosopher) philosopherLife(leftFork fork, rightFork fork) {
	for {
		select {
		// if a status request is received, write a response to output
		case <-p.input:
			p.output <- PhilosopherStatus{numEats: p.numEats, numThinks: p.numThinks, isEating: p.isEating}
		// No status request, move on
		default:
		}

		// Only we can access the forks, so we lock the mutex
		arbiter.Lock()

		p.isEating = false

		// We have exclusive access to forks now, so we read from their channels
		leftStatus := <-leftFork.output
		rightStatus := <-rightFork.output

		willPickUp := !leftStatus.isPickedUp && !rightStatus.isPickedUp

		action := doNothing
		if willPickUp {
			action = pickUp
		}

		// Tell the fork what should happen to it
		leftFork.input <- action
		rightFork.input <- action

		if willPickUp {
			p.isEating = true
			p.numEats++
		} else {
			p.numThinks++
		}

		// We are done for now, so we release the mutex, letting others lock it
		arbiter.Unlock()

		// Everything takes "time" - think for 2ms or eat for 2ms
		time.Sleep(2 * time.Millisecond)

		// If we were eating, put down the forks again before we move on
		if p.isEating {
			arbiter.Lock()
			<-leftFork.output  // prepare left fork to receive message
			<-rightFork.output // prepare right fork to receive message
			leftFork.input <- putDown
			rightFork.input <- putDown
			arbiter.Unlock()
		}
	}
}

func NewPhilosopher() philosopher {
	inputChannel := make(chan int)
	outputChannel := make(chan PhilosopherStatus)
	return philosopher{isEating: false, numEats: 0, numThinks: 0, input: inputChannel, output: outputChannel}
}
