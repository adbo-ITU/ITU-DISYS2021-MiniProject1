package main

import (
	"fmt"
	"sync"
)

const (
	N = 5
)

// Our forks and philosophers are saved in these two global arrays
var forks [N]fork
var philosophers [N]philosopher

// Only one goroutine can access forks and philosophers at once - this arbiter
// controls that
var arbiter sync.Mutex

func main() {
	// initialize the forks
	for i := range forks {
		forks[i] = NewFork()
		go forks[i].forkLife()
	}

	// initialize the philosophers
	for i := range philosophers {
		philosophers[i] = NewPhilosopher()
		go philosophers[i].philosopherLife(forks[i], forks[(i+1)%N])
	}

	// Prepare to recieve current states of the philosophers and forks
	var pState [N]PhilosopherStatus
	var fState [N]ForkStatus

	for {
		// Query all of the philosophers for their current state
		for i, p := range philosophers {
			p.input <- 1
			reply := <-p.output
			pState[i] = reply
		}

		// Query all of the forks for their current state
		for i, f := range forks {
			reply := <-f.output
			f.input <- doNothing
			fState[i] = reply
		}

		// this is a magic print that should "clear" the terminal window
		fmt.Print("\033[H\033[2J")

		var eatsSum uint64
		var thinksSum uint64
		// Print all philosopher info
		for i, philosopher := range pState {
			fmt.Printf("[P #%d]: Has eaten %d times. Has thought %d times. Is eating: %t\n",
				i,
				philosopher.numEats,
				philosopher.numThinks,
				philosopher.isEating)

			eatsSum += uint64(philosopher.numEats)
			thinksSum += uint64(philosopher.numThinks)
		}

		fmt.Printf("Total eats: %d, total thinks: %d.\n\n", eatsSum, thinksSum)

		var pickupSum uint64
		// Print all fork info
		for i, fork := range fState {
			fmt.Printf("[F #%d]: Has been picked up %d times, Is picked up: %t.\n",
				i,
				fork.numPickUps,
				fork.isPickedUp)

			pickupSum += fork.numPickUps
		}

		fmt.Printf("Total number of fork pickups: %d.\n\n", pickupSum)

		// Print additional info to check if the numbers are as expected
		fmt.Printf("pickups/eats ratio: %f (expected: 2).\n", float64(pickupSum)/(float64(eatsSum)))
		fmt.Printf("eats/thinks ratio:  %f (expected: 2/3 = 0.666..).\n", float64(eatsSum)/float64(thinksSum))
	}
}
