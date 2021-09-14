package main

import (
	"fmt"
	"sync"
)

const (
	N = 5
)

var forks [N]fork
var philosophers [N]philosopher
var arbiter sync.Mutex

func main() {
	// initialize the forks
	for i := range forks {
		inputChannel := make(chan int)
		outputChannel := make(chan ForkStatus)
		forkBuffer := fork{isPickedUp: false, numPickUps: 0, input: inputChannel, output: outputChannel}
		forks[i] = forkBuffer
		go forkBuffer.forkLife()
	}

	// initialize the philosophers
	for i := range philosophers {
		inputChannel := make(chan int)
		outputChannel := make(chan PhilosopherStatus)
		philosopherBuffer := philosopher{id: i, isEating: false, numEats: 0, numThinks: 0, input: inputChannel, output: outputChannel}
		philosophers[i] = philosopherBuffer
		go philosophers[i].philosopherLife(forks[i], forks[(i+1)%N])
	}

	// prepare for the job of the main thread: printing the pState of the philosophers
	var pState [N]PhilosopherStatus
	var fState [N]ForkStatus

	for {
		for i, p := range philosophers {
			p.input <- 1
			reply := <-p.output
			pState[i] = reply
		}

		for i, f := range forks {
			reply := <-f.output
			f.input <- doNothing
			fState[i] = reply
		}

		// this is a magic print that should "clear" the terminal window
		fmt.Print("\033[H\033[2J")

		var eatsSum uint64
		var thinksSum uint64
		for i, philosopher := range pState {
			fmt.Printf("[P #%d]: Has eaten %d times. Has thought %d times.\n",
				i,
				philosopher.numEats,
				philosopher.numThinks)

			eatsSum += uint64(philosopher.numEats)
			thinksSum += uint64(philosopher.numThinks)
		}

		fmt.Printf("Total eats: %d, total thinks: %d.\n\n", eatsSum, thinksSum)

		var pickupSum uint64
		for i, fork := range fState {
			fmt.Printf("[F #%d]: Has been picked up %d times, Is picked up: %t.\n",
				i,
				fork.numPickUps,
				fork.isPickedUp)

			pickupSum += fork.numPickUps
		}

		fmt.Printf("Total number of fork pickups: %d.\n\n", pickupSum)

		fmt.Printf("pickups/eats ratio: %f (expected: 2).\n", float64(pickupSum)/(float64(eatsSum)))
		fmt.Printf("eats/thinks ratio:  %f (expected: 2/3 = 0.666..).\n", float64(eatsSum)/float64(thinksSum))
	}
}
