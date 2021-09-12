package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	N = 5
)

var forks [N]fork
var philosophers [N]philosopher
var arbiter sync.Mutex

func main() {

	eventChannel := make(chan philosopher)

	// initialize the forks
	for i, _ := range forks {
		inputBuffer := make(chan int)
		outputBuffer := make(chan ForkStatus)
		forkBuffer := fork{isPickedUp: false, numPickUps: 0, input: inputBuffer, output: outputBuffer}
		forks[i] = forkBuffer
		go forkBuffer.forkLife()
	}

	// initialize the philosophers
	for i, _ := range philosophers {
		philosopherBuffer := philosopher{id: i, isEating: false, numEats: 0, numThinks: 0, events: eventChannel}
		philosophers[i] = philosopherBuffer
		go philosophers[i].philosopherLife(forks[i], forks[(i+1)%N])
	}

	// prepare for the job of the main thread: printing the state of the philosophers
	clock := time.Now()
	timer := time.Now()
	var state [N]philosopher

	for {
		select {
		case msg := <-eventChannel:
			state[msg.id] = msg
		}

		// measure time
		timer = time.Now()

		// if 500 milliseconds have passed since last check
		if timer.Sub(clock) >= 2000*1000 {

			// this is a magic print that should "clear" the terminal window
			fmt.Print("\033[H\033[2J")

			for _, philosopher := range state {
				fmt.Printf("[P #%d]: Has eaten %d times. Has thought %d times.\n",
					philosopher.id,
					philosopher.numEats,
					philosopher.numThinks)
			}

			// reset the clock
			clock = time.Now()
		}

	}
}
